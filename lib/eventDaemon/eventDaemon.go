package eventDaemon

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/joho/godotenv"

	union "github.com/MoonSHRD/IKY-telegram-bot/artifacts"

	multisig "github.com/0xSOLIDarnost/MultisigLegacy/artifacts/multisig"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
)

// channel to get this event from blockchain
//var ch = make(chan *passport.PassportPassportApplied)
//var ch_index = make(chan *passport.PassportPassportAppliedIndexed)




var GlobalClient *ethclient.Client
var GlobalAuth *bind.TransactOpts

var chat_ids = make([]int64,0) 

var chat_wallets = make(map[int64]common.Address)

type SubmissionMsg struct {
	chat_id int64
	SubmissionEvent *multisig.MultiSigWalletSubmission
}
// var globalChan chan *SubmissionMsg

var myenv map[string]string

// file with settings for enviroment
const envLoc = ".env"

func Start(masterChannel chan *SubmissionMsg) {


	loadEnv()
	ctx := context.Background()
	pk := myenv["PK"] // load private key from env

	// Connecting to blockchain network
	//  client, err := ethclient.Dial(os.Getenv("GATEWAY"))	// for global env config
	client, err := ethclient.Dial(myenv["GATEWAY_GOERLI_WS"]) // load from local .env file
	if err != nil {
		log.Fatalf("could not connect to Ethereum gateway: %v\n", err)
	}
	defer client.Close()
	GlobalClient = client

	// setting up private key in proper format
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
	}

	// Creating an auth transactor
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(4))
	GlobalAuth = auth

	// check calls
	// check balance
	accountAddress := common.HexToAddress("0x16d97A46030C5D3D705bca45439e48529997D8b2")
	balance, _ := client.BalanceAt(ctx, accountAddress, nil) //our balance
	fmt.Printf("Balance of the validator bot: %d\n", balance)


	// setting up union contract
	UnionCenter, err := union.NewUnion(common.HexToAddress(myenv["UNION_ADDRESS"]), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Union contract: %v", err)
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


	// Get current counter (how much chat_id is registred and make them in array)
	counter, err := GetUnionsCounter(sessionUnion)
	if err != nil {
		log.Printf(err.Error())
	}
	log.Println(counter)
	
	counter_int := counter.Int64()
	log.Println(counter_int)

	i := int64(0)
	for i = 0;  i < counter_int; i++ {

		// Get chatID
		chat_id, err := GetChatID(sessionUnion,big.NewInt(i))
		if err != nil {
			log.Printf(err.Error())
		}
		//chat_ids[i] = chat_id
		chat_ids = append(chat_ids,chat_id)
		fmt.Println("found new chat id: ",chat_ids[i])
	}

	// we also subscribe to any new registred dao
	var new_daos_ch = make(chan *union.UnionApprovedJoin)
	subscriptionNewDaos,err := SubscribeForApprovedUnions(sessionUnion,new_daos_ch)

	
	for _,id := range chat_ids {
		// get wallet addresses
		chat_wallets[id], err = GetAddressDao(sessionUnion,id)
		if err != nil {
			log.Printf(err.Error())
		}
		fmt.Printf("for chat_id: %d  ",id)
		fmt.Println("dao wallet address is: ", chat_wallets[id])
	}
	
	//var submission_ch = make(chan *multisig.MultiSigWalletSubmission)
	var submission_msg = make(chan *SubmissionMsg)


	
	for chatID,address := range chat_wallets {
		//go SubscribeForSubmittedTransactions()
		go InitiateMultisigSession(ctx,address,submission_msg,chatID)
		fmt.Printf("for chat_id: %d  ",chatID)
		fmt.Println("subscribed for multisig address: ", address)
	}
	


	EventLoop:
	for {
		select {
		case <-ctx.Done():
			{
				subscriptionNewDaos.Unsubscribe();
			break EventLoop
			}
	case NewDao:= <-new_daos_ch:
		{
			fmt.Println("/n")
			fmt.Println("found new chat id::", NewDao.ChatId)
			fmt.Println("with associated wallet address:", NewDao.MultyWalletAddress)
			go InitiateMultisigSession(ctx,NewDao.MultyWalletAddress,submission_msg,NewDao.ChatId.Int64())

		}
	case NewSubmission := <-submission_msg:
		{
			fmt.Println("new submission has found in MAIN thread, channels work!")
			fmt.Println("/n")
			fmt.Println("transaction id:",NewSubmission.SubmissionEvent.TransactionId)
			fmt.Println("chat id to sent msg:",NewSubmission.chat_id)
			fmt.Println("data:",NewSubmission.SubmissionEvent.Raw)
			masterChannel<- NewSubmission
		}
		}
		}



//	log.Printf("Authorized on account %s", bot.Self.UserName)

	/** TODO:
	*	0. Make initialize (main) function
	*	1. subscribe for Union event "Approved"
	*	2. each time someone got approve -- InitiateMultisigSession(ctx.Background, Multisig_address)
	*	3. each time we get some "Submit" event from Multisig we need to forward it to main module with tethered chat_id
	*
	*/




} // end of main func



// Initiate event listener session (go routine) for SUBMIT Events
func InitiateMultisigSession(ctx context.Context,dao_wallet_address common.Address, listenchan chan *SubmissionMsg, chat_id int64)  {

//	ctx := context.Background()
	MultisigInstance,err := multisig.NewMultiSigWallet(dao_wallet_address,GlobalClient)
	if err != nil {
		log.Fatalf("Failed to instantiate a Multisig_wallet contract: %v", err)
	}

	//Wrap a session
	sessionMultisig := &multisig.MultiSigWalletSession{
		Contract: MultisigInstance,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    GlobalAuth.From,
			Context: context.Background(),
		},
		TransactOpts: bind.TransactOpts{
			From:     GlobalAuth.From,
			Signer:   GlobalAuth.Signer,
			GasLimit: 0,   // 0 automatically estimates gas limit
			GasPrice: nil, // nil automatically suggests gas price
			Context:  context.Background(),
		},
	}

	var ch = make(chan *multisig.MultiSigWalletSubmission)
	subscription,err := SubscribeForSubmittedTransactions(sessionMultisig,ch)

	// Infinite loop for specific Multisig submission
	EventLoop:
	for {
		select {
		case <-ctx.Done():
			{
			subscription.Unsubscribe();
			break EventLoop
			}
	case eventResult:= <-ch:
		{
			fmt.Println("/n")
			//fmt.Println("Somebody want to submit new tx, his address:")
			fmt.Println("Destination for outcoming tx:", eventResult.Raw.Address)
			fmt.Println("Data for outcoming tx:", eventResult.Raw.Data)
			var msg *SubmissionMsg
			msg.chat_id = chat_id
			msg.SubmissionEvent = eventResult
			listenchan <-msg
		}
		}
		}
}


// load enviroment variables from .env file
func loadEnv() {
	var err error
	if myenv, err = godotenv.Read(envLoc); err != nil {
		log.Printf("could not load env from %s: %v", envLoc, err)
	}
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
		//log.Println("GetChatID: ")
		//log.Println(chatId)
		return chatId, err
	}
}


func GetAddressDao(session *union.UnionSession, chat_id int64) (common.Address, error) {
	dao_address,err := session.GetDaoAddressbyChatId(chat_id)
	if err != nil {
		return common.BigToAddress(nil), err
	} else {
		return dao_address, err
	}
}


// subscribing for ApprovedJoin events. This event represent that a new dao has been registred. We use watchers without fast-forwarding past events
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

/*
// go routine for forwarding new found 
func GetNewDaos(session *union.UnionSession) {

	var new_daos = make(chan *union.UnionApprovedJoin)
	subscriptionNewDaos,err := SubscribeForApprovedUnions(session,new_daos)

}
*/

// TODO: add Anonymouse event for each time of event in multisig (without indexed values)
func SubscribeForSubmittedTransactions(session *multisig.MultiSigWalletSession, listenChannel chan<- *multisig.MultiSigWalletSubmission) (event.Subscription, error) {
	subscription, err := session.Contract.WatchSubmission(&bind.WatchOpts{
		Start: nil,
		Context: nil,
	}, listenChannel,
	)
	if err != nil {
		return nil, err
	}
	log.Println("subscribed to Multisig submission transactions")
	return subscription,err
}

//func SubscribeTo

func IsTokenTransfer(session *multisig.MultiSigWalletSession, data []byte) (bool,error) {
	transfer, err := session.Contract.IsTransfer(&session.CallOpts,data)
	if err != nil {
		return false, err
	} else {
		return transfer,err
	}
}

