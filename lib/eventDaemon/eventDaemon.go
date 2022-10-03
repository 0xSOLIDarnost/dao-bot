package event_demon

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"os"

	"github.com/joho/godotenv"

	union "github.com/MoonSHRD/IKY-telegram-bot/artifacts"
	passport "github.com/MoonSHRD/IKY-telegram-bot/artifacts/TGPassport"

	//GroupWallet

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

	// setting up union contract
	UnionCenter, err := union.NewUnion(common.HexToAddress("0x2658da2258849ad6a2104704F4f085644aD45d0D"), client)
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

	//Wrap union session
	sessionUnion := &union.UnionSession{
		Contract: UnionCenter,
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

	log.Printf("session with union center initialized")

//	log.Printf("Authorized on account %s", bot.Self.UserName)





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



func GetUnionsCounter(session *union.UnionSession) (*big.Int, error) {

	// sessionUnion.Contract.
	dao_counter,err :=session.GetDaoCount()
	if err != nil {
		return nil, err
	} else {
		return dao_counter,err
	}
}

func GetChatID(session *union.UnionSession, counter *big.Int) (int64, error) {
	chatId, err := session.ChatIdArray(counter)
	if err != nil {
		return 0, err
	} else {
		return chatId, err
	}
}


// subscribing for ApprovedJoin events. We use watchers without fast-forwarding past events
func SubscribeForApprovedUnions(session *union.UnionSession, listenChannel chan<- *union.UnionApprovedJoin) (event.Subscription, error) {
	subscription, err := session.Contract.WatchApprovedJoin(&bind.WatchOpts{
		Start:   nil, //last block
		Context: nil, // nil = no timeout
	}, listenChannel,
	)
	if err != nil {
		return nil, err
	}
	return subscription, err
}

