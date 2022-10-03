package event_demon

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"os"

	"github.com/joho/godotenv"

	passport "github.com/MoonSHRD/IKY-telegram-bot/artifacts/TGPassport"
	//passport "IKY-telegram-bot/artifacts/TGPassport"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var yesNoKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Yes"),
		tgbotapi.NewKeyboardButton("No")),
)

var mainKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Verify personal wallet")),
)

//to operate the bot, put a text file containing key for your bot acquired from telegram "botfather" to the same directory with this file
var tgApiKey, err = os.ReadFile(".secret")
var bot, error1 = tgbotapi.NewBotAPI(string(tgApiKey))

//type containing all the info about user input
type user struct {
	tgid          int64
	tg_username   string
	dialog_status int64
}

type event_iterator = *passport.PassportPassportAppliedIterator // For filter  @TODO: consider removing

// event we got from blockchain
type event_bc = *passport.PassportPassportApplied

// channel to get this event from blockchain
var ch = make(chan *passport.PassportPassportApplied)
var ch_index = make(chan *passport.PassportPassportAppliedIndexed)

//main database for dialogs, key (int64) is telegram user id
var userDatabase = make(map[int64]user) // consider to change in persistend data storage?

var msgTemplates = make(map[string]string)

var baseURL = "http://localhost:3000/"
var tg_id_query = "?user_tg_id="
var tg_username_query = "&user_tg_name="

var myenv map[string]string

// file with settings for enviroment
const envLoc = ".env"

func main() {

	loadEnv()
	ctx := context.Background()
	pk := myenv["PK"] // load private key from env

	msgTemplates["hello"] = "Hey, this bot is attaching personal wallets to telegram user & collective wallets to chat id"
	msgTemplates["case0"] = "Go to link and attach your tg_id to your metamask wallet"
	msgTemplates["await"] = "Awaiting for verification"
	msgTemplates["case1"] = "You have successfully authorized your wallet to your account"

	bot, err = tgbotapi.NewBotAPI(string(tgApiKey))
	if err != nil {
		log.Panic(err)
	}

	// Connecting to blockchain network
	//  client, err := ethclient.Dial(os.Getenv("GATEWAY"))	// for global env config
	client, err := ethclient.Dial(myenv["GATEWAY_RINKEBY_WS"]) // load from local .env file
	if err != nil {
		log.Fatalf("could not connect to Ethereum gateway: %v\n", err)
	}
	defer client.Close()

	// setting up private key in proper format
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
	}

	// Creating an auth transactor
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(4))

	// check calls
	// check balance
	accountAddress := common.HexToAddress("0x16d97A46030C5D3D705bca45439e48529997D8b2")
	balance, _ := client.BalanceAt(ctx, accountAddress, nil) //our balance
	fmt.Printf("Balance of the validator bot: %d\n", balance)

	// Setting up Passport Contract
	passportCenter, err := passport.NewPassport(common.HexToAddress("0x2658da2258849ad6a2104704F4f085644aD45d0D"), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a TGPassport contract: %v", err)
	}

	// Wrap the Passport contract instance into a session
	session := &passport.PassportSession{
		Contract: passportCenter,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    auth.From,
			Context: context.Background(),
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: 0,   // 0 automatically estimates gas limit
			GasPrice: nil, // nil automatically suggests gas price
			Context:  context.Background(),
		},
	}

	log.Printf("session with passport center initialized")

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	//whenever bot gets a new message, check for user id in the database happens, if it's a new user, the entry in the database is created.
	for update := range updates {

		if update.Message != nil {
			if _, ok := userDatabase[update.Message.From.ID]; !ok {

				userDatabase[update.Message.From.ID] = user{update.Message.Chat.ID, update.Message.Chat.UserName, 0}
				msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, msgTemplates["hello"])
				msg.ReplyMarkup = mainKeyboard
				bot.Send(msg)
			} else {

				switch userDatabase[update.Message.From.ID].dialog_status {

				//first check for user status, (for a new user status 0 is set automatically), then user reply for the first bot message is logged to a database as name AND user status is updated
				case 0:
					if updateDb, ok := userDatabase[update.Message.From.ID]; ok {

						msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, msgTemplates["case0"])
						bot.Send(msg)

						tgid := userDatabase[update.Message.From.ID].tgid
						user_name := userDatabase[update.Message.From.ID].tg_username
						fmt.Println(user_name)
						tgid_string := fmt.Sprint(tgid)
						tgid_array := make([]int64,1)
						tgid_array[0] = tgid
						link := baseURL + tg_id_query + tgid_string + tg_username_query + "@" + user_name
						msg = tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, link)
						bot.Send(msg)

						//subscription, err := SubscribeForApplications(session, ch)   //  this is ordinary subscription to NORMAL event
						subscription, err := SubscribeForApplicationsIndexed(session, ch_index,tgid_array)	// this is subscription to INDEXED event. This mean we can pass what exactly value of argument we want to see

						if err != nil {
							log.Fatal(err)
						}
					EventLoop:
						for {
							select {
							case <-ctx.Done():
								{
									subscription.Unsubscribe()
									break EventLoop
								}
							case eventResult := <-ch_index:
								{
									fmt.Println("User tg_id:", eventResult.ApplyerTg)
									fmt.Println("User wallet address:", eventResult.WalletAddress)
										applyer_tg_string := strconv.FormatInt(eventResult.ApplyerTg,10)
										msg = tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, " your application have been recived "+applyer_tg_string)
										bot.Send(msg)
										ApprovePassport(auth, passportCenter, eventResult.WalletAddress) 
										subscription.Unsubscribe()
										break EventLoop
								}
							/*  Use next snippet to work with regular events (when args are NOT INDEXED)
							*	In this approach we parsing results from event and awaiting for values to match
							case eventResult := <-ch:
								{
									//fmt.Println("\n")
									fmt.Println("User tg_id:", eventResult.ApplyerTg)
									event_tgid := eventResult.ApplyerTg
									fmt.Println("User wallet address:", eventResult.WalletAddress)
									if event_tgid == tgid {
										applyer_tg_string := strconv.FormatInt(eventResult.ApplyerTg,10)
										msg = tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, " your application have been recived "+applyer_tg_string)
										bot.Send(msg)
										ApprovePassport(auth, passportCenter, eventResult.WalletAddress)
										subscription.Unsubscribe()
										break EventLoop
									}
								} */
							}
							
						}
						updateDb.dialog_status = 1
						userDatabase[update.Message.From.ID] = updateDb
					}

				case 1:
					if updateDb, ok := userDatabase[update.Message.From.ID]; ok {
						//updateDb.dialog_status = 2
						msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, msgTemplates["case1"])
						bot.Send(msg)
						userDatabase[update.Message.From.ID] = updateDb
					}

				}
			}
		}
	}

} // end of main func



// load enviroment variables from .env file
func loadEnv() {
	var err error
	if myenv, err = godotenv.Read(envLoc); err != nil {
		log.Printf("could not load env from %s: %v", envLoc, err)
	}
}

// subscribing for Applications events. We use watchers without fast-forwarding past events
func SubscribeForApplications(session *passport.PassportSession, listenChannel chan<- *passport.PassportPassportApplied) (event.Subscription, error) {
	subscription, err := session.Contract.WatchPassportApplied(&bind.WatchOpts{
		Start:   nil, //last block
		Context: nil, // nil = no timeout
	}, listenChannel,
	)
	if err != nil {
		return nil, err
	}
	return subscription, err
}

// subscribing for Applications events. We use watchers without fast-forwarding past events
func SubscribeForApplicationsIndexed(session *passport.PassportSession, listenChannel chan<- *passport.PassportPassportAppliedIndexed, applierTGID []int64) (event.Subscription, error) {
	subscription, err := session.Contract.WatchPassportAppliedIndexed(&bind.WatchOpts{
		Start:   nil, //last block
		Context: nil, // nil = no timeout
	}, listenChannel,
	   applierTGID,
	)
	if err != nil {
		return nil, err
	}
	return subscription, err
}

func ApprovePassport(auth *bind.TransactOpts, pc *passport.Passport, user_address common.Address) {

	tx_to_approve, err := pc.ApprovePassport(
		&bind.TransactOpts{
			From:      auth.From,
			Nonce:     nil,
			Signer:    auth.Signer,
			Value:     big.NewInt(0),
			GasPrice:  nil,
			GasFeeCap: nil,
			GasTipCap: nil,
			GasLimit:  0,
			Context:   context.Background(),
		}, user_address,
	)

	if err != nil {
		log.Println("cant send approval reques to contract: ")
		log.Print(err)
	}

	fmt.Printf("transaction for APPROVAL passport sent! Please wait for tx %s to be confirmed. \n", tx_to_approve.Hash().Hex())

}

func DeclinePassport(auth *bind.TransactOpts, pc *passport.Passport, user_address common.Address) {

	tx_to_approve, err := pc.DeclinePassport(
		&bind.TransactOpts{
			From:      auth.From,
			Nonce:     nil,
			Signer:    auth.Signer,
			Value:     big.NewInt(0),
			GasPrice:  nil,
			GasFeeCap: nil,
			GasTipCap: nil,
			GasLimit:  0,
			Context:   context.Background(),
		}, user_address,
	)

	if err != nil {
		log.Println("cant send DECLINING reques to contract: ")
		log.Print(err)
	}

	fmt.Printf("transaction for DECLINING passport sent! Please wait for tx %s to be confirmed. \n", tx_to_approve.Hash().Hex())

}
