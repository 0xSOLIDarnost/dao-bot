package event_demon

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/joho/godotenv"

	union "github.com/MoonSHRD/IKY-telegram-bot/artifacts"
	passport "github.com/MoonSHRD/IKY-telegram-bot/artifacts/TGPassport"

	// TODO: fix it
	//multisig "github.com/daseinsucks/MultisigLegacy/artifacts"
	//multisig "github.com/0xSOLIDarnost/dao-bot/artifacts/multisig"
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

var myenv map[string]string

// file with settings for enviroment
const envLoc = ".env"

func main() {

	loadEnv()
	ctx := context.Background()
	pk := myenv["PK"] // load private key from env

	// Connecting to blockchain network
	//  client, err := ethclient.Dial(os.Getenv("GATEWAY"))	// for global env config
	client, err := ethclient.Dial(myenv["GATEWAY_RINKEBY_WS"]) // load from local .env file
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

	/** TODO:
	*	1. subscribe for Union event "Approved"
	*	2. each time someone got approve -- InitiateMultisigSession(ctx.Background, Multisig_address)
	*	3. each time we get some "Submit" event from Multisig we need to forward it to main module with tethered chat_id
	*
	*/




} // end of main func

func InitiateMultisigSession(ctx context.Context,group_wallet_address string)  {

//	ctx := context.Background()
	MultisigInstance,err := multisig.NewMultisigwallet(common.HexToAddress(group_wallet_address),GlobalClient)
	if err != nil {
		log.Fatalf("Failed to instantiate a Multisig_wallet contract: %v", err)
	}

	//Wrap a session
	sessionMultisig := &multisig.MultisigwalletSession{
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

	var ch = make(chan *multisig.MultisigwalletSubmission)
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
			fmt.Println("Destination for outcoming tx:", eventResult.Raw.Address)
			fmt.Println("Data for outcoming tx:", eventResult.Raw.Data)

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


// TODO: add Anonymouse event for each time of event in multisig (without indexed values)
func SubscribeForSubmittedTransactions(session *multisig.MultisigwalletSession, listenChannel chan<- *multisig.MultisigwalletSubmission) (event.Subscription, error) {
	subscription, err := session.Contract.WatchSubmission(&bind.WatchOpts{
		Start: nil,
		Context: nil,
	}, listenChannel,
	)
	if err != nil {
		return nil, err
	}
	return subscription,err
}

