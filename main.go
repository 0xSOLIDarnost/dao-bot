package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	//"math"
	"os"

	union "github.com/MoonSHRD/IKY-telegram-bot/artifacts"
	passport "github.com/MoonSHRD/IKY-telegram-bot/artifacts/TGPassport"

	multisig "github.com/0xSOLIDarnost/MultisigLegacy/artifacts/multisig"
	//union "github.com/daseinsucks/MultisigBot/artifacts"

	//passport "IKY-telegram-bot/artifacts/TGPassport"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"

	"github.com/joho/godotenv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ERC20"),
		tgbotapi.NewKeyboardButton("ERC20Snapshot"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ERC721")),
)

var nullAddress common.Address = common.HexToAddress("0x0000000000000000000000000000000000000000")

//to operate the bot, put a text file containing key for your bot acquired from telegram "botfather" to the same directory with this file
var tgApiKey, err = os.ReadFile(".secret")
var bot, _ = tgbotapi.NewBotAPI(string(tgApiKey))

type user struct {
	chat          *tgbotapi.Chat
	usertype      *tgbotapi.User
	tgid          int64
	chatid        int64
	tg_chatname   string
	dialog_status int64
	setup_status  int64
	repo          string
	dao           string
	votingtype    int64
	vtt           string
}

type event_bc = *union.UnionApplicationForJoinIndexed

var baseURL = "http://localhost:3000/dao"

var user_id_query = "?user_id="
var chat_query = "&chat_id="
var address_query = "&address="
var type_query = "&votingtype="
var contract_query = "&votingtokencontract="
var name_query = "&daoname="

var ch_index = make(chan *union.UnionApplicationForJoinIndexed)

//localhost:3000/dao?user_id=1337&chat_id=1337&address=23746624386&votingtype=1&votingtokencontract=3278465ASDW23&daoname=lol

//main database for dialogs, key (int64) is telegram user id
var userDatabase = make(map[int64]user) // consider to change in persistend data storage?

func main() {

	_ = godotenv.Load()
	ctx := context.Background()
	pk := os.Getenv("PK") // load private key from env
	gateway := os.Getenv("GATEWAY_GOERLI_WS")

	accAddress := os.Getenv("ACCOUNT_ADDRESS")
	contractAddress := os.Getenv("UNION_ADDRESS")
	passportAddress := os.Getenv("PASSPORT_ADDRESS")

	bot, err = tgbotapi.NewBotAPI(string(tgApiKey))
	if err != nil {
		log.Panic(err)
	}

	// Connecting to blockchain network
	//  client, err := ethclient.Dial(os.Getenv("GATEWAY"))	// for global env config
	client, err := ethclient.Dial(gateway) // load from local .env file
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
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(5))
	if err != nil {
		log.Fatalf("could not connect to auth gateway: %v\n", err)
	}
	//auth2:= bind.NewKeyedTransactorWithChainID(privateKey,big.NewInt(4))
	//NewKeyedTransactorWithChainID

	accountAddress := common.HexToAddress(accAddress)
	balance, _ := client.BalanceAt(ctx, accountAddress, nil) //our balance
	fmt.Printf("Balance of the validator bot: %d\n", balance)

	// Setting up Union
	UnionCaller, err := union.NewUnionCaller(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Union contract: %v", err)
	}

	UnionSession, err := union.NewUnion(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Union contract: %v", err)
	}

	session := &union.UnionSession{
		Contract: UnionSession,
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

	log.Printf("session with union initialized")

	passport, err := passport.NewPassportCaller(common.HexToAddress(passportAddress), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Passport contract: %v", err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	//TODO: add check tgid == daoaddress(tgid)
	//whenever bot gets a new message, check for user id in the database happens, if it's a new user, the entry in the database is created.
	for update := range updates {

		if update.Message != nil {
			if _, ok := userDatabase[update.Message.Chat.ID]; !ok {
				userDatabase[update.Message.Chat.ID] = user{update.Message.Chat, update.Message.From, update.Message.From.ID, update.Message.Chat.ID, update.Message.Chat.Title, 0, 0, "0", "0", 0, "0"}

				isRegistered := checkDao(auth, UnionCaller, update.Message.Chat.ID)
				if isRegistered {
					if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {
						updateDb.dialog_status = 1
						userDatabase[update.Message.Chat.ID] = updateDb
					}
				} else {
					msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].chatid, "Your Union is not registered yet!\nLet's register it!.\nFirst, send me the link to your repo.")
					bot.Send(msg)
					if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {
						updateDb.dialog_status = 0
						updateDb.setup_status = 1
						userDatabase[update.Message.Chat.ID] = updateDb
					}
				}

			} else {
				switch userDatabase[update.Message.Chat.ID].setup_status {

				case 1:
					if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {
						updateDb.repo = update.Message.Text
						updateDb.setup_status = 2
						userDatabase[update.Message.Chat.ID] = updateDb
						chatvar := userDatabase[update.Message.Chat.ID].chat
						uservar := userDatabase[update.Message.Chat.ID].usertype
						isAdmin := checkAdmin(chatvar, uservar)
						if !isAdmin {
							msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].chatid, "Sorry, but only admin of the chat may connect it to the DAO!")
							bot.Send(msg)
							delete(userDatabase, update.Message.Chat.ID)
						}
						isUserRegistered := checkUser(auth, passport, userDatabase[update.Message.Chat.ID].tgid)
						if !isUserRegistered {
							msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].chatid, "Sorry, but before attaching DAO you should apply for passport here:")
							bot.Send(msg)
							delete(userDatabase, update.Message.Chat.ID)
						}
						if isUserRegistered && isAdmin {
							msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].chatid, "Okay, tell me your Multisig address!")
							bot.Send(msg)
						}
					}

				case 2:
					if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {
						updateDb.dao = update.Message.Text
						updateDb.setup_status = 3
						userDatabase[update.Message.Chat.ID] = updateDb

						daoaddress := userDatabase[update.Message.Chat.ID].dao
						wallet, _ := multisig.NewMultisigwalletCaller(common.HexToAddress(daoaddress), client)

						botIsOwner := checkBotAsOwner(auth, wallet, accountAddress)

						if botIsOwner {
							msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].chatid, "Cool, now I need to know your voting token's type")
							msg.ReplyMarkup = numericKeyboard
							bot.Send(msg)
						} else {
							msg1 := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].chatid, "Sorry, but this bot is not the owner of the multisig wallet.")
							bot.Send(msg1)
							delete(userDatabase, update.Message.Chat.ID)
						}
					}
				case 3:
					if update.Message.Text == "ERC20Snapshot" || update.Message.Text == "ERC20" || update.Message.Text == "ERC721" {

						var tokenType int64
						if update.Message.Text == "ERC20" {
							tokenType = 0
						} else if update.Message.Text == "ERC20Snapshot" {
							tokenType = 1
						} else if update.Message.Text == "ERC20Votes" {
							tokenType = 2
						}

						if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {
							updateDb.votingtype = tokenType
							updateDb.setup_status = 4
							userDatabase[update.Message.Chat.ID] = updateDb
							msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].chatid, "Okay, last question: what's the address of your voting token?")
							msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
							bot.Send(msg)
						}

					} else {
						msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].chatid, "That's not the type!")
						bot.Send(msg)
					}

				case 4:
					if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {
						updateDb.vtt = update.Message.Text
						updateDb.setup_status = 4
						userDatabase[update.Message.Chat.ID] = updateDb

						msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].chatid, "Alright, now apply for union here:")
						bot.Send(msg)

						userIDint := userDatabase[update.Message.Chat.ID].tgid
						userID := fmt.Sprint(userIDint)

						chatIDint := userDatabase[update.Message.Chat.ID].chatid
						chatID := fmt.Sprint(chatIDint)

						address := userDatabase[update.Message.Chat.ID].dao

						votingTypeint := userDatabase[update.Message.Chat.ID].votingtype
						votingType := fmt.Sprint(votingTypeint)

						vtt := userDatabase[update.Message.Chat.ID].vtt
						chatName := userDatabase[update.Message.Chat.ID].tg_chatname

						link := baseURL + user_id_query + userID + chat_query + chatID + address_query + address + type_query + votingType + contract_query + vtt + name_query + chatName

						msg2 := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].chatid, link)
						bot.Send(msg2)

						tgid_array := make([]int64, 1)
						tgid_array[0] = chatIDint

						subscription, err := SubscribeForApplicationsIndexed(session, ch_index, tgid_array) // this is subscription to INDEXED event. This mean we can pass what exactly value of argument we want to see

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
									fmt.Println("DAO tg_id:", eventResult.ChatId)
									fmt.Println("DAO wallet address:", eventResult.MultyWalletAddress)
									applyer_tg_string := fmt.Sprint(eventResult.ChatId)
									msg = tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].chatid, " your application have been recived "+applyer_tg_string)
									bot.Send(msg)
									ApproveDAO(auth, UnionSession, eventResult.MultyWalletAddress)
									if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {
										updateDb.setup_status = 0
										userDatabase[update.Message.Chat.ID] = updateDb
										subscription.Unsubscribe()
										break EventLoop
									}

								}
							}
						}
					}
				}
			}
		} else if userDatabase[update.Message.Chat.ID].setup_status == 0 {

			switch userDatabase[update.Message.Chat.ID].dialog_status {

			//first check for user status, (for a new user status 0 is set automatically), then user reply for the first bot message is logged to a database as name AND user status is updated
			case 0:
				isRegistered := checkDao(auth, UnionCaller, update.Message.Chat.ID)
				if isRegistered {
					if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {
						updateDb.dialog_status = 1
						updateDb.setup_status = 0
						userDatabase[update.Message.From.ID] = updateDb
					}
				} else {
					msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].chatid, "Your Union is not registered yet! \n Please register it at <unionbot>")
					bot.Send(msg)
				}

				//all other logic may be implemented here

			}
		}
	}
}

func checkDao(auth *bind.TransactOpts, pc *union.UnionCaller, tgid int64) bool {
	registration, err := pc.DaoAddresses(&bind.CallOpts{
		From:    auth.From,
		Context: context.Background(),
	}, tgid)

	log.Println(registration)

	if err != nil {
		log.Print(err)
	}

	if registration == nullAddress {
		return false
	} else {
		return true
	}
}

func checkUser(auth *bind.TransactOpts, pc *passport.PassportCaller, tgid int64) bool {
	registration, err := pc.TgIdToAddress(&bind.CallOpts{
		From:    auth.From,
		Context: context.Background(),
	}, tgid)

	log.Println(registration)

	if err != nil {
		log.Print(err)
	}

	if registration == nullAddress {
		return false
	} else {
		return true
	}
}

func checkAdmin(chat *tgbotapi.Chat, user *tgbotapi.User) bool {
	thing := tgbotapi.ChatConfigWithUser{
		ChatID:             chat.ID,
		SuperGroupUsername: chat.ChatConfig().SuperGroupUsername,
		UserID:             user.ID,
	}
	thing2 := tgbotapi.GetChatMemberConfig{thing}

	chatmember, _ := bot.GetChatMember(thing2)
	if chatmember.Status == "administrator" || chatmember.Status == "creator" {
		return true
	} else {
		return false
	}
}

func SubscribeForApplicationsIndexed(session *union.UnionSession, listenChannel chan<- *union.UnionApplicationForJoinIndexed, chat_id []int64) (event.Subscription, error) {
	subscription, err := session.Contract.WatchApplicationForJoinIndexed(&bind.WatchOpts{
		Start:   nil, //last block
		Context: nil, // nil = no timeout
	}, listenChannel,
		chat_id,
	)
	if err != nil {
		return nil, err
	}
	return subscription, err
}

func ApproveDAO(auth *bind.TransactOpts, pc *union.Union, dao_address common.Address) {

	tx_to_approve, err := pc.ApproveJoin(
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
		}, dao_address,
	)

	if err != nil {
		log.Println("cant send approval request to contract: ")
		log.Print(err)
	}

	fmt.Printf("transaction for APPROVAL DAO sent! Please wait for tx %s to be confirmed. \n", tx_to_approve.Hash().Hex())
}

func checkBotAsOwner(auth *bind.TransactOpts, pc *multisig.MultisigwalletCaller, botAddress common.Address) bool {
	isOwner, _ := pc.IsOwner(&bind.CallOpts{
		From:    auth.From,
		Context: context.Background(),
	}, botAddress)

	return isOwner
}
