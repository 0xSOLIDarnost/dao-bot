package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"os"

	"github.com/joho/godotenv"

	union "github.com/MoonSHRD/IKY-telegram-bot/artifacts"
	passport "github.com/MoonSHRD/IKY-telegram-bot/artifacts/TGPassport"

	//passport "IKY-telegram-bot/artifacts/TGPassport"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var yesNoKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Yes"),
		tgbotapi.NewKeyboardButton("No")),
)

var optionKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("WhoIs"),
		tgbotapi.NewKeyboardButton("KARMA")),
)

var trustKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Trust/Untrust user"),
		tgbotapi.NewKeyboardButton("See who trust/untrust user")),
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

type VoteEntityConfig struct {
	chat_id		int64
	voting_token_address common.Address
	voting_type	uint8
	poll_config tgbotapi.SendPollConfig
}


type event_iterator = *passport.PassportPassportAppliedIterator // For filter  @TODO: consider removing

// event we got from blockchain
type event_bc = *passport.PassportPassportApplied

// channel to get this event from blockchain
var ch = make(chan *passport.PassportPassportApplied)
var ch_index = make(chan *passport.PassportPassportAppliedIndexed)

var ch_approved = make(chan *passport.PassportPassportApproved)

//main database for dialogs, key (int64) is telegram user id
var userDatabase = make(map[int64]user) // consider to change in persistend data storage?

var msgTemplates = make(map[string]string)


var tg_id_query = "?user_tg_id="
var tg_username_query = "&user_tg_name="

var myenv map[string]string

// file with settings for enviroment
const envLoc = ".env"


/** 
*		TODO:
*		1. ask what do u want (input transaction ID, possible get this from main thread)
*		2. create poll with %subject% and vote yes/no, send it to chat id
*		3. create go routine sleeping until dealine
*		4. get results
*		5. get arrays of tgid[] who voted yes and no
*		6. for each tgid ask wallet from passport contract
*		7. for each wallet ask if they have balanceOf(snapshot) of tokens, which associated with this chat id from Union contract
*		8. only those, who have voting token is counted
*		9. filter results, get final count yes or no
*		10. return result in main thread through channel
*				
*/







func main() {

	loadEnv()
	ctx := context.Background()
	pk := myenv["PK"] // load private key from env

	msgTemplates["hello"] = "Hey, this bot is attaching personal wallets to telegram user & collective wallets to chat id"
	msgTemplates["case0"] = "Go to link and attach your tg_id to your metamask wallet"
	msgTemplates["await"] = "Awaiting for verification"
	msgTemplates["case1"] = "You have successfully authorized your wallet to your account. Now you can use additional functions"
	msgTemplates["who_is"] = "Input wallet address to know it's associated telegram nickname"
	msgTemplates["karma"] = "Karma system allow users to express trust/untrust to specific tg user or see who is trust/untrust to this user. Data is immutable and store in blockchain"
	msgTemplates["trust_link"] = "Send telegram nickname of person who you are willing to trust/untrust"
	msgTemplates["who_trust"] = "Send telegram nickname of person to see who trust/untrust it"


	//var baseURL = "http://localhost:3000/"
	//var baseURL = "https://ikytest-gw0gy01is-s0lidarnost.vercel.app/"
	//	var baseURL = myenv["BASEURL"];



	bot, err = tgbotapi.NewBotAPI(string(tgApiKey))
	if err != nil {
		log.Panic(err)
	}

	// Connecting to blockchain network
	//  client, err := ethclient.Dial(os.Getenv("GATEWAY"))	// for global env config
	client, err := ethclient.Dial(myenv["GATEWAY_GOERLI_WS"]) // load from local .env file
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

	// check calls
	// check balance
	accountAddress := common.HexToAddress(myenv["ACCOUNT_ADDRESS"])
	balance, _ := client.BalanceAt(ctx, accountAddress, nil) //our balance
	fmt.Printf("Balance of the validator bot: %d\n", balance)

	// Setting up Passport Contract
	//passportCenter, err := passport.NewPassport(common.HexToAddress(myenv["PASSPORT_ADDRESS"]), client)
	

	unionContract, err := union.NewUnion(common.HexToAddress(myenv["UNION_ADDRESS"]), client)
	if err != nil {
		log.Fatalln("can't estible connection with Union contract: %v",err)
	}

	sessionUnion := &union.UnionSession{
		Contract: unionContract,
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


	log.Printf("session with passport center & union initialized")

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	//whenever bot gets a new message, check for user id in the database happens, if it's a new user, the entry in the database is created.
	for update := range updates {

		if update.Message != nil  {
			//if update.Message.Poll.
			if _, ok := userDatabase[update.Message.From.ID]; !ok {



				userDatabase[update.Message.From.ID] = user{update.Message.Chat.ID, update.Message.Chat.UserName, 0}
				msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, msgTemplates["hello"])

		//		msg.ReplyMarkup = mainKeyboard
				bot.Send(msg)
			} else {

				switch userDatabase[update.Message.From.ID].dialog_status {

				//first check for user status, (for a new user status 0 is set automatically), then user reply for the first bot message is logged to a database as name AND user status is updated
				case 0:
					if updateDb, ok := userDatabase[update.Message.From.ID]; ok {



						msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, msgTemplates["case0"])
						bot.Send(msg)

						subject := update.Message.Text
						vote := ConstructVote(sessionUnion,update.Message.From.ID,subject)
						// does we need to cast vote?

						tgid := userDatabase[update.Message.From.ID].tgid
						user_name := userDatabase[update.Message.From.ID].tg_username
						fmt.Println(user_name)
						tgid_array := make([]int64, 1)
						tgid_array[0] = tgid

			
						updateDb.dialog_status = 1
						userDatabase[update.Message.From.ID] = updateDb
						
					}
					fallthrough // МЫ ЛЕД ПОД НОГАМИ МАЙОРА!
				case 1:
					if updateDb, ok := userDatabase[update.Message.From.ID]; ok {
						msg := tgbotapi.NewMessage(userDatabase[update.Message.From.ID].tgid, msgTemplates["case1"])
						msg.ReplyMarkup = optionKeyboard
						bot.Send(msg)
						updateDb.dialog_status = 2
						userDatabase[update.Message.From.ID] = updateDb
						
					}
			}   }
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



// allow bot to get tg nickname associated with this eth wallet
func WhoIsAddress(session *passport.PassportSession,address_to_check common.Address) (string,error){
	passport, err := session.GetPassportByAddress(address_to_check)
	if err != nil {
		log.Println("cant get passport associated with this address, possible it's not registred yet: ")
		log.Print(err)
		 return "error",err
	}
	nickname := passport.UserName
	return nickname,nil

}

// construct simple off-chain vote
func ConstructPoll(subject string, chat_id int64) (tgbotapi.SendPollConfig) {
	
	var poll = tgbotapi.NewPoll(chat_id,subject,"yes","no")
	return poll
}

// Get voting meta and construct simple off-chain poll (SendConfig)
func ConstructVote(session *union.UnionSession,chat_id int64, subject string) (*VoteEntityConfig) {
	dao_address, err := session.GetDaoAddressbyChatId(chat_id)
	if err != nil {
		log.Println("can't find dao registred with this chat id, possible not registred yet: ")
		log.Println(err)
	//	return "error", err
	}
	dao, err := session.Daos(dao_address)
	v_token_address := dao.VotingToken
	// erc20, erc20Snapshot, erc721
	v_token_type := dao.VotingType
	poll := ConstructPoll(subject,chat_id)
	var v *VoteEntityConfig
	v.chat_id = chat_id
	v.voting_token_address = v_token_address
	v.voting_type = v_token_type
	v.poll_config = poll
	return v
}

func StarteVoteSession(session *union.UnionSession, chat_id int64, subject string) {
	vote_start_config := ConstructVote(session,chat_id,subject)
	close_date := vote_start_config.poll_config.CloseDate
	open_period := vote_start_config.poll_config.OpenPeriod
	fmt.Println("close date:", close_date)
	fmt.Println("open period: ", open_period)
	//current_date :=    TODO: ask current day
}