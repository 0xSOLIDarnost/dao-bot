package voter

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

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/joho/godotenv"

	union "github.com/MoonSHRD/IKY-telegram-bot/artifacts"
	token_erc20 "github.com/MoonSHRD/IKY-telegram-bot/artifacts/ERC20"
	passport "github.com/MoonSHRD/IKY-telegram-bot/artifacts/TGPassport"

	//passport "IKY-telegram-bot/artifacts/TGPassport"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var yesNoKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Yes"),
		tgbotapi.NewKeyboardButton("No")),
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
	chat_id              int64
	voting_token_address common.Address
	voting_type          uint8
	poll_config          tgbotapi.SendPollConfig
}

type ActiveVote struct {
	yes    int64
	no     int64
	voters []Voter
}

type Voter struct {
	tgid    int64
	accepts bool
}

var poll = make(map[string]ActiveVote) // map pollId => ActiveVote

//main database for dialogs, key (int64) is telegram user id
var userDatabase = make(map[int64]user) // consider to change in persistend data storage?

var msgTemplates = make(map[string]string)

var myenv map[string]string

// file with settings for enviroment
const envLoc = ".env"

// load enviroment variables from .env file
func loadEnv() {
	var err error
	if myenv, err = godotenv.Read(envLoc); err != nil {
		log.Printf("could not load env from %s: %v", envLoc, err)
	}
}

// allow bot to get tg nickname associated with this eth wallet
func WhoIsAddress(session *passport.PassportSession, address_to_check common.Address) (string, error) {
	passport, err := session.GetPassportByAddress(address_to_check)
	if err != nil {
		log.Println("cant get passport associated with this address, possible it's not registred yet: ")
		log.Print(err)
		return "error", err
	}
	nickname := passport.UserName
	return nickname, nil

}

// construct simple off-chain vote
func ConstructPoll(subject string, chat_id int64) tgbotapi.SendPollConfig {

	var poll = tgbotapi.NewPoll(chat_id, subject, "Yes", "No")
	return poll
}

// Get voting meta and construct simple off-chain poll (SendConfig)
func ConstructVote(session *union.UnionSession, chat_id int64, subject string) (*VoteEntityConfig, error) {
	dao_address, err := session.GetDaoAddressbyChatId(chat_id)
	if err != nil {
		log.Println("can't find dao registred with this chat id, possible not registred yet: ")
		log.Println(err)
		return nil, err
	}
	dao, err := session.Daos(dao_address)
	v_token_address := dao.VotingToken
	// erc20, erc20Snapshot, erc721
	v_token_type := dao.VotingType
	poll := ConstructPoll(subject, chat_id)
	var v *VoteEntityConfig
	v.chat_id = chat_id
	v.voting_token_address = v_token_address
	v.voting_type = v_token_type
	v.poll_config = poll
	return v, err
}

// TODO
func StarteVoteSession(session *union.UnionSession, chat_id int64, subject string, duration int64) {
	vote_start_config, err := ConstructVote(session, chat_id, subject)
	if err != nil {
		//return nil, err
		log.Println("error in ConstructVote")
		log.Println(err)
	}

	current_date := time.Now().Unix()
	close_date := current_date + duration
	vote_start_config.poll_config.CloseDate = close_date
	open_period := vote_start_config.poll_config.OpenPeriod
	fmt.Println("close date:", close_date)
	fmt.Println("open period: ", open_period)
	timeToSleep := time.Duration(duration)
	time.Sleep(timeToSleep)
	CalculatePollResult()
}

// Getter for user wallet
func GetAddressByTgID(passportSession *passport.PassportSession, tg_id int64) (common.Address, error) {
	user_address, err := passportSession.GetPassportWalletByID(tg_id)
	if err != nil {
		return common.HexToAddress("0x0"), err
	} else {
		return user_address, err
	}
}

// Create instanse of token contract with corresponding type and calculate balanceOf for []address
func CalculatePower(client_bc *ethclient.Client, auth *bind.TransactOpts, token_address common.Address, token_type uint8, user_addresses []common.Address) (*big.Int, error) {

	var sum *big.Int
	// Enum token_type, 0 = erc 20
	if token_type == uint8(0) {
		tokenContract, err := token_erc20.NewTokenERC20(token_address, client_bc)
		if err != nil {
			//return nil, err  //TODO return err
			log.Println("can't estiblish connection with VoteToken contract")
			log.Println(err)
		}
		sessionToken := &token_erc20.TokenERC20Session{
			Contract: tokenContract,
			CallOpts: bind.CallOpts{
				Pending: false,
				From:    auth.From,
				Context: context.Background(),
			},
			TransactOpts: bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasLimit: 0,
				GasPrice: nil,
				Context:  context.Background(),
			},
		}

		for _, user_address := range user_addresses {
			votePower, err := sessionToken.BalanceOf(user_address)
			if err != nil {
				log.Println("can't get balanceOf user, possible wrong address for token contract")
				log.Println(err)
				return nil, err
			}
			sum = sum.Add(sum, votePower) // not sure it should work this way
		}
		//return sum,err

	} // TODO: add calculating for other token type (ERC20Snapshot), NFT ..
	// doing this require add this sample contracts in IKY and make new release
	return sum, err
}

// This function calculete final vote power to specific []tg_ids related to chat_id group. Assuming tg_ids is already sorted yes/no array
func CalculateVotePower(client_bc *ethclient.Client, auth *bind.TransactOpts, chat_id int64, tg_ids []int64) (*big.Int, error) {

	// Setting up Passport Session
	passportContract, err := passport.NewPassport(common.HexToAddress(myenv["PASSPORT_ADDRESS"]), client_bc)
	if err != nil {
		log.Fatalln("can't estible connection with Passport contract: %v", err)
	}
	sessionPassport := &passport.PassportSession{
		Contract: passportContract,
		CallOpts: bind.CallOpts{
			Pending: false,
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

	// Setting up Union Contract
	unionContract, err := union.NewUnion(common.HexToAddress(myenv["UNION_ADDRESS"]), client_bc)
	if err != nil {
		log.Fatalln("can't establish connection with Union contract: %v", err)
	}
	sessionUnion := &union.UnionSession{
		Contract: unionContract,
		CallOpts: bind.CallOpts{
			Pending: false,
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

	// create of user addresses and get them by tgid
	var user_addresses = make([]common.Address, 0)
	for _, tg_id := range tg_ids {
		user_address, err := GetAddressByTgID(sessionPassport, tg_id)
		if err != nil {
			log.Println("user with %d tgid was not found in passport")
			log.Println(err)
			return nil, err
		} else {
			user_addresses = append(user_addresses, user_address)
		}
	}

	// get vote_type
	dao_address, err := sessionUnion.GetDaoAddressbyChatId(chat_id)
	if err != nil {
		log.Println("chat with this chat_id has not registred in Union", chat_id)
		log.Println(err)
		return nil, err
	}
	dao, err := sessionUnion.Daos(dao_address)
	if err != nil {
		log.Println("can't find DAO struct with associated address %v")
		log.Println("this shit can't happen")
		return nil, err
	}
	voting_type := dao.VotingType
	voting_token_address := dao.VotingToken

	summary_vote_power, err := CalculatePower(client_bc, auth, voting_token_address, voting_type, user_addresses)
	if err != nil {
		log.Println("error in counting tokens:")
		log.Println(err)
		return nil, err
	}
	return summary_vote_power, err
}

// Accepts []int64 tg_ids 'yes' and 'now' and generates final flag
func CalculatePollResult(client_bc *ethclient.Client, auth *bind.TransactOpts, chat_id int64, yes_ids []int64, no_ids []int64) (bool, error) {
	var accepted bool
	yes_power, err := CalculateVotePower(client_bc, auth, chat_id, yes_ids)
	if err != nil {
		return false, err
	}
	no_power, err := CalculateVotePower(client_bc, auth, chat_id, no_ids)
	if err != nil {
		return false, err
	}

	//if yes_power > no_power true, otherwise false
	switch yes_power.Cmp(no_power) {
	case -1:
		accepted = false

	case 0:
		// parity
		accepted = false
	case 1:
		accepted = true
	}

	return accepted, err
}

func CalculatePersonPower(client_bc *ethclient.Client, auth *bind.TransactOpts, token_address common.Address, passportSession *passport.PassportSession, token_type uint8, tgid int64) (*big.Int, error) {
	defer handlePanic()

	var sum *big.Int
	// Enum token_type, 0 = erc 20
	if token_type == uint8(0) {
		tokenContract, err := token_erc20.NewTokenERC20(token_address, client_bc)
		if err != nil { //TODO return err
			log.Println("can't estiblish connection with VoteToken contract")
			log.Println(err)
		}
		sessionToken := &token_erc20.TokenERC20Session{
			Contract: tokenContract,
			CallOpts: bind.CallOpts{
				Pending: false,
				From:    auth.From,
				Context: context.Background(),
			},
			TransactOpts: bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasLimit: 0,
				GasPrice: nil,
				Context:  context.Background(),
			},
		}

		user_address, err := GetAddressByTgID(passportSession, tgid)
		if err != nil {
			log.Println("can't get user address, possible wrong address for token contract")
			log.Println(err)
			return nil, err
		}

		sum, err := sessionToken.BalanceOf(user_address)
		if err != nil {
			log.Println("can't get balanceOf user, possible wrong address for token contract")
			log.Println(err)
			return nil, err
		}
	}

	return sum, err
}

func handlePanic() {
	// detect if panic occurs or not
	a := recover()

	if a != nil {
		fmt.Println("RECOVER", a)
	}

}

//this function manages incoming votes in polls and returns 2 bool values
//so we keep using it until finished variable == true
func VoteInProgress(update *tgbotapi.Update, client_bc *ethclient.Client, auth *bind.TransactOpts, token_address common.Address, passportSession *passport.PassportSession, token_type uint8) (bool, bool) {
	var accepted bool
	var finished bool
	if update.PollAnswer != nil {
		if _, ok := poll[update.PollAnswer.PollID]; !ok {

			accepts := false
			if update.PollAnswer.OptionIDs == 0 { //"Yes" should be first option in poll
				accepts = true
			}

			var y, n int64 = 0, 0 //beginning of votepower calculation

			votepower, err := CalculatePersonPower(client_bc, auth, token_address, passportSession, token_type, update.PollAnswer.User.ID)
			if err != nil {
				votepower = big.NewInt(0)
			}

			finalVotePower := votepower.Int64()
			if accepts {
				y = finalVotePower
			} else {
				n = finalVotePower
			}
			firstVoter := Voter{update.PollAnswer.User.ID, accepts}
			poll[update.PollAnswer.PollID] = ActiveVote{y, n, []Voter{firstVoter}}

		} else {

			if update.Poll.IsClosed == false {

				existingVoters := poll[update.PollAnswer.PollID].voters
				alreadyVoted := false

				var previouslyAccepted bool

				for _, votersList := range existingVoters {
					if votersList.tgid == update.PollAnswer.User.ID {
						alreadyVoted = true
						previouslyAccepted = votersList.accepts
					}
				}

				votepower, err := CalculatePersonPower(client_bc, auth, token_address, passportSession, token_type, update.PollAnswer.User.ID)
				if err != nil {
					votepower = big.NewInt(0)
				}
				finalVotePower := votepower.Int64()

				updateVoteStatus := poll[update.PollAnswer.PollID]

				if alreadyVoted { //if user already voted, we just change the uint value of yes/no

					if previouslyAccepted {
						updateVoteStatus.yes = updateVoteStatus.yes - finalVotePower
						updateVoteStatus.no = updateVoteStatus.no + finalVotePower

						poll[update.PollAnswer.PollID] = updateVoteStatus
					} else {
						updateVoteStatus.yes = updateVoteStatus.yes + finalVotePower
						updateVoteStatus.no = updateVoteStatus.no - finalVotePower

						poll[update.PollAnswer.PollID] = updateVoteStatus
					}

				} else { //if user did not vote, we add him to an array with his opinion

					accepts := false
					if update.PollAnswer.OptionIDs == 0 {
						accepts = true //"Yes" should be first option in poll
					}
					if accepts {
						updateVoteStatus.yes = updateVoteStatus.yes + finalVotePower
					} else {
						updateVoteStatus.no = updateVoteStatus.no + finalVotePower
					}

					newVoter := Voter{update.PollAnswer.User.ID, accepts}
					i := updateVoteStatus.voters
					updateArray := append(i, newVoter)
					updateVoteStatus.voters = updateArray
					poll[update.PollAnswer.PollID] = updateVoteStatus
				}

			} else { //if poll is closed, then next incoming update triggers this block, which returns the result
				finalYes := poll[update.PollAnswer.PollID].yes
				finalNo := poll[update.PollAnswer.PollID].no

				finished = true

				if finalYes > finalNo {
					accepted = true
				}

			}

		}

	}

	return accepted, finished
}
