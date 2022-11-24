package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	//"math"
	"os"

	union "github.com/MoonSHRD/IKY-telegram-bot/artifacts"
	passport "github.com/MoonSHRD/IKY-telegram-bot/artifacts/TGPassport"

	multisig "github.com/0xSOLIDarnost/MultisigLegacy/artifacts/multisig"
	pogreb "github.com/akrylysov/pogreb"

	//union "github.com/daseinsucks/MultisigBot/artifacts"

	rules "github.com/0xSOLIDarnost/dao-bot/lib/rules"
	voter "github.com/0xSOLIDarnost/dao-bot/lib/voter"

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
		tgbotapi.NewKeyboardButton("ERC721"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ERC20Votes")),
)

var mainKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Start a vote"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Start a vote for rules"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Help message"),
		tgbotapi.NewKeyboardButton("Get rules")),
)

var nullAddress common.Address = common.HexToAddress("0x0000000000000000000000000000000000000000")

//to operate the bot, put a text file containing key for your bot acquired from telegram "botfather" to the same directory with this file
var tgApiKey, err = os.ReadFile(".secret")
var bot, _ = tgbotapi.NewBotAPI(string(tgApiKey))

type user struct {
	Chat         *tgbotapi.Chat
	Usertype     *tgbotapi.User
	Tgid         int64
	ChatID       int64
	TgChatName   string
	DialogStatus int64
	SetupStatus  int64
	Repo         string
	Dao          string
	VotingType   uint8
	VTC          string
	PollTopic    string
	PollDuration int64
	VoteType     int64
}

var pollToChat = make(map[string]int64)

var pollToBeginning = make(map[string]int64)

var chatToPoll = make(map[int64]string)

type event_bc = *union.UnionApplicationForJoinIndexed

var baseURL = "http://localhost:3000/dao"

var user_id_query = "?user_id="
var chat_query = "&chat_id="
var address_query = "&address="
var type_query = "&votingtype="
var contract_query = "&votingtokencontract="
var name_query = "&daoname="

var ch_index = make(chan *union.UnionApplicationForJoinIndexed)

//localhost:3000/dao?user_id=1337&chat_id=1337&address=23746624386&VotingType=1&votingtokencontract=3278465ASDW23&daoname=lol

//main database for dialogs, key (int64) is telegram chat id
var userDatabase = make(map[int64]user) // consider to change in persistend data storage?

func main() {

	tgdb, _ := pogreb.Open("telegramdb", nil)
	_ = godotenv.Load()
	ctx := context.Background()
	pk := os.Getenv("PK") // load private key from env
	gateway := os.Getenv("GATEWAY_GOERLI_WS")

	gitToken := os.Getenv("GITHUB_TOKEN")
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

	passportCaller, err := passport.NewPassportCaller(common.HexToAddress(passportAddress), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Passport contract: %v", err)
	}

	passportCenter, err := passport.NewPassport(common.HexToAddress(passportAddress), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a TGPassport contract: %v", err)
	}

	passportSession := &passport.PassportSession{
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

	log.Printf("session with passport initialized")

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	helptext := "Hi! This bot is designed to operate DAO. If you're not registered, we'll walk you through it.\nOtherwise, use the provided keyboard to operate it"

	updates := bot.GetUpdatesChan(u)
	//TODO: add check Tgid == daoaddress(Tgid)
	//whenever bot gets a new message, check for user id in the database happens, if it's a new user, the entry in the database is created.
	for update := range updates {
		if update.Message != nil {
			fmt.Println("got message! Dialog status:", userDatabase[update.Message.Chat.ID].DialogStatus)
		}

		if update.Message != nil {
			if _, ok := userDatabase[update.Message.Chat.ID]; !ok {

				userDatabase[update.Message.Chat.ID] = user{update.Message.Chat, update.Message.From, update.Message.From.ID, update.Message.Chat.ID, update.Message.Chat.Title, 0, 0, "0", "0", 0, "0", "0", 0, 0}

				isRegistered := checkDao(auth, UnionCaller, update.Message.Chat.ID)
				if isRegistered {
					updateDb := userDatabase[update.Message.Chat.ID]

					updateDb = restoreUserViaJson(tgdb, update.Message.Chat.ID) //if the chat is registered, it is present in our db
					updateDb.SetupStatus = 0
					updateDb.DialogStatus = 1

					fmt.Println(updateDb)

					msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, helptext)
					msg.ReplyMarkup = mainKeyboard
					bot.Send(msg)
					userDatabase[update.Message.Chat.ID] = updateDb

				} else {
					himsg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, helptext)
					bot.Send(himsg)
					msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, "Your Union is not registered yet!\nLet's register it!\nFirst, send me the link to your Repo.")
					bot.Send(msg)
					if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {
						updateDb.DialogStatus = 0
						updateDb.SetupStatus = 1
						userDatabase[update.Message.Chat.ID] = updateDb
					}
				}

			} else if userDatabase[update.Message.Chat.ID].SetupStatus != 0 {
				switch userDatabase[update.Message.Chat.ID].SetupStatus {

				case 1:
					if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {
						updateDb.Repo = update.Message.Text
						updateDb.SetupStatus = 2
						userDatabase[update.Message.Chat.ID] = updateDb
						chatvar := userDatabase[update.Message.Chat.ID].Chat
						uservar := userDatabase[update.Message.Chat.ID].Usertype
						isAdmin := checkAdmin(chatvar, uservar)
						if !isAdmin {
							msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, "Sorry, but only admin of the Chat may connect it to the DAO!")
							bot.Send(msg)
							delete(userDatabase, update.Message.Chat.ID)
						}
						isUserRegistered := checkUser(auth, passportCaller, userDatabase[update.Message.Chat.ID].Tgid)
						if !isUserRegistered {
							msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, "Sorry, but before attaching DAO you should apply for passport here:")
							bot.Send(msg)
							delete(userDatabase, update.Message.Chat.ID)
						}
						if isUserRegistered && isAdmin {
							msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, "Okay, now please add solidaobot as a collaborator here: "+userDatabase[update.Message.Chat.ID].Repo+"/settings/access"+"\n"+"We need it to edit the rules based on polls.")
							bot.Send(msg)
							msg1 := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, "After you're done, please send me your Multisignature wallet address.")
							bot.Send(msg1)
						}
					}

				case 2:
					if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {
						updateDb.Dao = update.Message.Text
						updateDb.SetupStatus = 3
						userDatabase[update.Message.Chat.ID] = updateDb

						daoaddress := userDatabase[update.Message.Chat.ID].Dao
						wallet, _ := multisig.NewMultiSigWalletCaller(common.HexToAddress(daoaddress), client)

						botIsOwner := checkBotAsOwner(auth, wallet, accountAddress)

						if botIsOwner {
							msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, "Cool, now I need to know your voting token's type")
							msg.ReplyMarkup = numericKeyboard
							bot.Send(msg)
						} else {
							msg1 := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, "Sorry, but this bot is not the owner of the multisig wallet.")
							bot.Send(msg1)
							delete(userDatabase, update.Message.Chat.ID)
						}
					}
				case 3:
					if update.Message.Text == "ERC20Snapshot" || update.Message.Text == "ERC20" || update.Message.Text == "ERC721" {

						var tokenType uint8
						if update.Message.Text == "ERC20" {
							tokenType = 0
						} else if update.Message.Text == "ERC20Votes" {
							tokenType = 1
						} else if update.Message.Text == "ERC721" {
							tokenType = 2
						}

						if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {
							updateDb.VotingType = tokenType
							updateDb.SetupStatus = 4
							userDatabase[update.Message.Chat.ID] = updateDb
							msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, "Okay, last question: what's the address of your voting token?")
							msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
							bot.Send(msg)
						}

					} else {
						msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, "That's not the type!")
						bot.Send(msg)
					}

				case 4:
					if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {
						updateDb.VTC = update.Message.Text
						updateDb.SetupStatus = 4
						userDatabase[update.Message.Chat.ID] = updateDb

						msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, "Alright, now apply for union here:")
						bot.Send(msg)

						userIDint := userDatabase[update.Message.Chat.ID].Tgid
						userID := fmt.Sprint(userIDint)

						chatIDint := userDatabase[update.Message.Chat.ID].ChatID
						chatID := fmt.Sprint(chatIDint)

						address := userDatabase[update.Message.Chat.ID].Dao

						votingTypeint := userDatabase[update.Message.Chat.ID].VotingType
						votingType := fmt.Sprint(votingTypeint)

						VTC := userDatabase[update.Message.Chat.ID].VTC
						chatName := userDatabase[update.Message.Chat.ID].TgChatName

						link := baseURL + user_id_query + userID + chat_query + chatID + address_query + address + type_query + votingType + contract_query + VTC + name_query + chatName

						msg2 := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, link)
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
									msg = tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, "Your application for chat "+applyer_tg_string+" was received!")
									msg.ReplyMarkup = mainKeyboard
									bot.Send(msg)
									ApproveDAO(auth, UnionSession, eventResult.MultyWalletAddress)
									if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {
										updateDb.SetupStatus = 0
										updateDb.DialogStatus = 1
										userDatabase[update.Message.Chat.ID] = updateDb

										userToSave, _ := json.Marshal(userDatabase[update.Message.Chat.ID])
										dbKey := []byte(strconv.FormatInt(update.Message.Chat.ID, 10))

										tgdb.Put(dbKey, userToSave)

										subscription.Unsubscribe()
										break EventLoop
									}

								}
							}
						}
					}
				}
			} else {

				switch userDatabase[update.Message.Chat.ID].DialogStatus {

				//first check for user status, (for a new user status 0 is set automatically), then user reply for the first bot message is logged to a database as name AND user status is updated

				case 1: //main standby status, awaiting for commands (they should be entered in this switch statement)

					switch update.Message.Text {

					case "Start a vote":
						if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {
							updateDb.DialogStatus = 2
							updateDb.VoteType = 0
							updateDb.Tgid = update.Message.From.ID
							updateDb.PollDuration = time.Now().Unix() + 60

							userDatabase[update.Message.Chat.ID] = updateDb
							msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, "Okay, let's start a vote! Enter the topic.")
							msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
							bot.Send(msg)

						}
					case "Start a vote for rules":
						if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {
							updateDb.DialogStatus = 2
							updateDb.VoteType = 1

							updateDb.Tgid = update.Message.From.ID
							updateDb.PollDuration = time.Now().Unix() + 60

							userDatabase[update.Message.Chat.ID] = updateDb
							msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, "Okay, let's start a vote for rules! Enter the suggested rule.")
							msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
							bot.Send(msg)

						}

					case "Help message":
						msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, helptext)
						bot.Send(msg)

					case "Get rules":
						rules, err := rules.GetRules(ctx, userDatabase[update.Message.Chat.ID].Repo, gitToken)
						text := "Failed to get rules :( \n Are you sure you added bot as a collaborator in your repo?"
						if err == nil {
							text = rules
						}
						msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, text)
						bot.Send(msg)
					}

				case 2:

					if update.Message.From.ID == userDatabase[update.Message.Chat.ID].Tgid && time.Now().Unix() < userDatabase[update.Message.Chat.ID].PollDuration {
						if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {
							updateDb.PollTopic = update.Message.Text
							updateDb.DialogStatus = 3
							userDatabase[update.Message.Chat.ID] = updateDb
							msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, "Okay, for how long in hours you want to be active?")
							bot.Send(msg)

						}
					} else {
						if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {
							updateDb.PollTopic = update.Message.Text
							updateDb.DialogStatus = 1
							userDatabase[update.Message.Chat.ID] = updateDb
						}
					}

				case 3:

					if update.Message.From.ID == userDatabase[update.Message.Chat.ID].Tgid && time.Now().Unix() < userDatabase[update.Message.Chat.ID].PollDuration {

						duration, err := strconv.ParseInt(update.Message.Text, 10, 64)
						if err == nil && duration > 0 {

							if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {

								updateDb.PollDuration = duration
								updateDb.DialogStatus = 4

								userDatabase[update.Message.Chat.ID] = updateDb

								timeOfEnd := time.Now().Unix() + (duration * 60) //TODO: change to 3600!!!

								timeUTCstring := time.Unix(timeOfEnd, 0).Format("15:04:05 02-01-2006")

								msgWithTime := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, "Please, send any message to this chat after "+timeUTCstring+" to summarize the results.")
								bot.Send(msgWithTime)

								poll := voter.StartPoll(userDatabase[update.Message.Chat.ID].ChatID, userDatabase[update.Message.Chat.ID].PollDuration, userDatabase[update.Message.Chat.ID].PollTopic)

								if userDatabase[update.Message.Chat.ID].VoteType == 1 {
									msgWithRule := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, "@"+update.Message.From.UserName+" wants to add the following rule: "+userDatabase[update.Message.Chat.ID].PollTopic)

									bot.Send(msgWithRule)
									poll = voter.StartPoll(userDatabase[update.Message.Chat.ID].ChatID, userDatabase[update.Message.Chat.ID].PollDuration, "Do we accept to add this rule?")
								}

								sentMessage, _ := bot.Send(poll)

								pollToChat[sentMessage.Poll.ID] = userDatabase[update.Message.Chat.ID].ChatID
								pollToBeginning[sentMessage.Poll.ID] = time.Now().Unix()
								chatToPoll[update.Message.Chat.ID] = sentMessage.Poll.ID
							}
						} else {
							msg := tgbotapi.NewMessage(userDatabase[update.Message.Chat.ID].ChatID, "Please, send the poll desired duration in hours")
							bot.Send(msg)
						}
					} else {
						if updateDb, ok := userDatabase[update.Message.Chat.ID]; ok {
							updateDb.PollTopic = update.Message.Text
							updateDb.DialogStatus = 1
							userDatabase[update.Message.Chat.ID] = updateDb
						}
					}
				}
			}
		}
		if (update.PollAnswer != nil && userDatabase[pollToChat[update.PollAnswer.PollID]].DialogStatus == 4) || (update.Message != nil && userDatabase[update.Message.Chat.ID].DialogStatus == 4) {

			fmt.Println("got update! status is still 4")
			var pollkey string
			var ChatID int64

			if update.PollAnswer != nil {
				pollkey = update.PollAnswer.PollID
				ChatID = pollToChat[pollkey]
			} else if update.Message != nil {
				ChatID = update.Message.Chat.ID
				pollkey = chatToPoll[ChatID]
			}

			tokenAddress := common.HexToAddress(userDatabase[ChatID].VTC)
			tokenType := userDatabase[ChatID].VotingType
			duration := userDatabase[ChatID].PollDuration * 60
			beginning := pollToBeginning[pollkey]
			accepted, finished := voter.VoteInProgress(duration, beginning, update, client, auth, tokenAddress, passportSession, tokenType, pollkey)
			fmt.Println("Finished is:", finished)
			fmt.Println("Accepted is:", accepted)

			if finished {
				if updateDb, ok := userDatabase[ChatID]; ok {

					updateDb.DialogStatus = 1
					text := "Was declined!"
					if accepted {
						text = "Was accepted!"
						if updateDb.VoteType == 1 {
							err := rules.AddRule(ctx, updateDb.Repo, gitToken, updateDb.PollTopic)
							if err == nil {
								text = "Pull request with rule was created! Please, merge it at " + updateDb.Repo + "/pulls"
							} else {
								text = "Error occured in pull request opening :(\nDid you add this bot as a colaborator in your repo?"
								fmt.Println(err)
							}
						}
					}

					userDatabase[ChatID] = updateDb
					msg := tgbotapi.NewMessage(userDatabase[ChatID].ChatID, text)
					msg.ReplyMarkup = mainKeyboard
					bot.Send(msg)

				}
			} else if time.Now().Unix() > (beginning + duration + 3600) {
				if updateDb, ok := userDatabase[ChatID]; ok {
					updateDb.DialogStatus = 1
					userDatabase[ChatID] = updateDb
				}
				msg := tgbotapi.NewMessage(userDatabase[ChatID].ChatID, "Poll was ignored :(")
				msg.ReplyMarkup = mainKeyboard
				bot.Send(msg)

			}

		}
	}
}

func checkDao(auth *bind.TransactOpts, pc *union.UnionCaller, Tgid int64) bool {
	registration, err := pc.DaoAddresses(&bind.CallOpts{
		From:    auth.From,
		Context: context.Background(),
	}, Tgid)

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

func checkUser(auth *bind.TransactOpts, pc *passport.PassportCaller, Tgid int64) bool {
	registration, err := pc.TgIdToAddress(&bind.CallOpts{
		From:    auth.From,
		Context: context.Background(),
	}, Tgid)

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

func checkAdmin(Chat *tgbotapi.Chat, user *tgbotapi.User) bool {
	thing := tgbotapi.ChatConfigWithUser{
		ChatID:             Chat.ID,
		SuperGroupUsername: Chat.ChatConfig().SuperGroupUsername,
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

func checkBotAsOwner(auth *bind.TransactOpts, pc *multisig.MultiSigWalletCaller, botAddress common.Address) bool {
	isOwner, _ := pc.IsOwner(&bind.CallOpts{
		From:    auth.From,
		Context: context.Background(),
	}, botAddress)

	return isOwner
}

func restoreUserViaJson(database *pogreb.DB, chatid int64) user {
	defer handlePanic()

	thing, _ := database.Get([]byte(strconv.FormatInt(chatid, 10)))

	var userstruct user
	json.Unmarshal(thing, &userstruct)
	return userstruct
}

func handlePanic() {

	// detect if panic occurs or not
	a := recover()

	if a != nil {
		fmt.Println("RECOVER", a)
	}

}
