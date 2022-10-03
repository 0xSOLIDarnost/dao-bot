package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"

	//"math"
	"os"

	union "github.com/MoonSHRD/IKY-telegram-bot/artifacts"
	passport "github.com/MoonSHRD/IKY-telegram-bot/artifacts/TGPassport"

	//union "github.com/daseinsucks/MultisigBot/artifacts"

	//passport "IKY-telegram-bot/artifacts/TGPassport"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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

var mainKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Verify personal wallet")),
)
var nullAddress common.Address = common.HexToAddress("0x0000000000000000000000000000000000000000")

//to operate the bot, put a text file containing key for your bot acquired from telegram "botfather" to the same directory with this file
var tgApiKey, err = os.ReadFile(".secret")
var bot, _ = tgbotapi.NewBotAPI(string(tgApiKey))

type user struct {
	tgid          int64
	tg_username   string
	dialog_status int64
	setup_status  int64
	repo          string
	dao           string
	votingtype    int64
	vtt           string
}

var baseURL = "http://localhost:3000/dao"

var user_id_query = "?user_id="
var chat_query = "&chat_id="
var address_query = "&address="
var type_query = "&votingtype="
var contract_query = "&votingtokencontract="
var name_query = "&daoname="

//localhost:3000/dao?user_id=1337&chat_id=1337&address=23746624386&votingtype=1&votingtokencontract=3278465ASDW23&daoname=lol

//main database for dialogs, key (int64) is telegram user id
var userDatabase = make(map[int64]user) // consider to change in persistend data storage?

func main() {

	_ = godotenv.Load()
	ctx := context.Background()
	pk := os.Getenv("PK") // load private key from env
	gateway := os.Getenv("GATEWAY_RINKEBY_WS")

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
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(4))
	//auth2:= bind.NewKeyedTransactorWithChainID(privateKey,big.NewInt(4))
	//NewKeyedTransactorWithChainID

	accountAddress := common.HexToAddress("0xc905803BbC804fECDc36850281fEd6520A346AC5")
	balance, _ := client.BalanceAt(ctx, accountAddress, nil) //our balance
	fmt.Printf("Balance of the validator bot: %d\n", balance)

	// Setting up Union
	union, err := union.NewUnionCaller(common.HexToAddress("0x9c6C6CBDA53E72A6e25C5F9AcE5b1Ef87Ac8635b"), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Union contract: %v", err)
	}

	log.Printf("session with union initialized")

	passport, err := passport.NewPassportCaller(common.HexToAddress("0xd8d32BB03ED024757Ad6f8585ee5973910328Cc6"), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Union contract: %v", err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	//TODO: add check tgid == daoaddress(tgid)
	//whenever bot gets a new message, check for user id in the database happens, if it's a new user, the entry in the database is created.
	for update := range updates {

		if update.Message != nil {
			if _, ok := userDatabase[update.Message.From.ID]; !ok {
				userDatabase[update.Message.From.ID] = user{update.Message.Chat.ID, update.Message.Chat.UserName, 0, 0, "0", "0", 0, "0"}

				isRegistered := checkDao(auth, union, update.Message.Chat.ID)
				if isRegistered {
					if updateDb, ok := userDatabase[update.Message.From.ID]; ok {
						updateDb.dialog_status = 1
						userDatabase[update.Message.From.ID] = updateDb
					}
				} else {
					msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, "Your Union is not registered yet! \nLet's register it! \n First, send me the link to your repo.")
					bot.Send(msg)
					if updateDb, ok := userDatabase[update.Message.From.ID]; ok {
						updateDb.dialog_status = 0
						updateDb.setup_status = 1
						userDatabase[update.Message.From.ID] = updateDb
					}
				}

			} else {
				switch userDatabase[update.Message.From.ID].setup_status {

				case 1:
					if updateDb, ok := userDatabase[update.Message.From.ID]; ok {
						updateDb.repo = update.Message.Text
						updateDb.setup_status = 2
						userDatabase[update.Message.From.ID] = updateDb
						chatvar := update.Message.Chat
						uservar := update.Message.From
						isAdmin := checkAdmin(chatvar, uservar)
						if !isAdmin {
							msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, "Sorry, but only admin of the chat may connect it to the DAO!")
							bot.Send(msg)
							delete(userDatabase, update.Message.From.ID)
						}
						isUserRegistered := checkUser(auth, passport, update.Message.From.ID)
						if !isUserRegistered {
							msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, "Sorry, but before attaching DAO you should apply for passport here:")
							bot.Send(msg)
							delete(userDatabase, update.Message.From.ID)
						}
						if isUserRegistered && isAdmin {
							msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, "Okay, tell me your Multisig address!")
							bot.Send(msg)
						}
					}
				case 2:
					if updateDb, ok := userDatabase[update.Message.From.ID]; ok {
						updateDb.dao = update.Message.Text
						updateDb.setup_status = 3
						userDatabase[update.Message.From.ID] = updateDb
						msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, "Cool, now I need to know your voting token's type")
						msg.ReplyMarkup = numericKeyboard
						bot.Send(msg)
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

						if updateDb, ok := userDatabase[update.Message.From.ID]; ok {
							updateDb.votingtype = tokenType
							updateDb.setup_status = 4
							userDatabase[update.Message.From.ID] = updateDb
							msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, "Okay, last question: what's the address of your voting token?")
							msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
							bot.Send(msg)
						}

					} else {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "That's not the type!")
						bot.Send(msg)
					}

				case 4:
					if updateDb, ok := userDatabase[update.Message.From.ID]; ok {
						updateDb.vtt = update.Message.Text
						updateDb.setup_status = 4
						userDatabase[update.Message.From.ID] = updateDb

						msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, "Alright, now apply for union here:")
						bot.Send(msg)

						userIDint := userDatabase[update.Message.From.ID].tgid
						userID := strconv.FormatInt(userIDint, 10)

						chatIDint := update.Message.Chat.ID
						chatID := strconv.FormatInt(chatIDint, 10)

						address := userDatabase[update.Message.From.ID].dao

						votingTypeint := userDatabase[update.Message.From.ID].votingtype
						votingType := strconv.FormatInt(votingTypeint, 10)

						vtt := userDatabase[update.Message.From.ID].vtt
						chatName := update.Message.Chat.UserName

						link := baseURL + user_id_query + userID + chat_query + chatID + address_query + address + type_query + votingType + contract_query + vtt + name_query + chatName

						msg2 := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, link)
						bot.Send(msg2)
					}
				}
			}

		} else if userDatabase[update.Message.From.ID].setup_status == 0 {

			switch userDatabase[update.Message.From.ID].dialog_status {

			//first check for user status, (for a new user status 0 is set automatically), then user reply for the first bot message is logged to a database as name AND user status is updated
			case 0:
				isRegistered := checkDao(auth, union, update.Message.Chat.ID)
				if isRegistered {
					if updateDb, ok := userDatabase[update.Message.From.ID]; ok {
						updateDb.dialog_status = 1
						updateDb.setup_status = 0
						userDatabase[update.Message.From.ID] = updateDb
					}
				} else {
					msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, "Your Union is not registered yet! \n Please register it at <unionbot>")
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
