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
	"reflect"
	"time"

	"github.com/joho/godotenv"

	token_erc20 "github.com/MoonSHRD/IKY-telegram-bot/artifacts/ERC20"
	passport "github.com/MoonSHRD/IKY-telegram-bot/artifacts/TGPassport"

	//passport "IKY-telegram-bot/artifacts/TGPassport"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//to operate the bot, put a text file containing key for your bot acquired from telegram "botfather" to the same directory with this file
var tgApiKey, err = os.ReadFile(".secret")

//type containing all the info about user input

type ActiveVote struct {
	yes    *big.Int
	no     *big.Int
	voters []Voter
}

type Voter struct {
	tgid    int64
	accepts bool
}

var poll = make(map[string]ActiveVote) // map pollId => ActiveVote

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

// Getter for user wallet
func GetAddressByTgID(passportSession *passport.PassportSession, tg_id int64) (common.Address, error) {
	user_address, err := passportSession.GetPassportWalletByID(tg_id)
	if err != nil {
		return common.HexToAddress("0x0"), err
	} else {
		return user_address, err
	}
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

		sum, err = sessionToken.BalanceOf(user_address)
		if err != nil {
			log.Println("can't get balanceOf user, possible wrong address for token contract")
			log.Println(err)
			return nil, err
		}
	}
	fmt.Println(sum)
	return sum, err
}

func handlePanic() {
	// detect if panic occurs or not
	a := recover()

	if a != nil {
		fmt.Println("RECOVER", a)
	}

}

func StartPoll(chatId int64, durationInHours int64, topic string) tgbotapi.SendPollConfig {
	pollToSend := tgbotapi.NewPoll(chatId, topic, "Yes", "No")

	currentDate := time.Now().Unix()
	hoursInSeconds := durationInHours * 120
	timeToClose := currentDate + hoursInSeconds
	fmt.Println(hoursInSeconds)
	timeToCloseInt := int(timeToClose)
	pollToSend.CloseDate = timeToCloseInt
	pollToSend.IsAnonymous = false

	return pollToSend
}

//this function manages incoming votes in polls and returns 2 bool values
//so we keep using it until finished variable == true
func VoteInProgress(update tgbotapi.Update, client_bc *ethclient.Client, auth *bind.TransactOpts, token_address common.Address, passportSession *passport.PassportSession, token_type uint8) (bool, bool) {
	var accepted bool
	var finished bool
	var yesChosen = []int{0} //I'm not sure this'll work. It's supposed to be equal [0] as a "Yes" answer in a poll
	if update.PollAnswer != nil {
		if _, ok := poll[update.PollAnswer.PollID]; !ok {

			accepts := false
			selectedYes := reflect.DeepEqual(yesChosen, update.PollAnswer.OptionIDs)
			if selectedYes { //"Yes" should be first option in poll
				accepts = true
			}

			var y, n *big.Int = big.NewInt(0), big.NewInt(0) //beginning of votepower calculation

			firstVoter := Voter{update.PollAnswer.User.ID, accepts}
			poll[update.PollAnswer.PollID] = ActiveVote{y, n, []Voter{firstVoter}}
			fmt.Println("Voting successful!")

		} else {

			if !update.Poll.IsClosed {

				existingVoters := poll[update.PollAnswer.PollID].voters
				alreadyVoted := false

				var voterNumber int
				for n, votersList := range existingVoters {
					if votersList.tgid == update.PollAnswer.User.ID {
						alreadyVoted = true
						voterNumber = n
					}
				}

				updateVoteStatus := poll[update.PollAnswer.PollID]

				if alreadyVoted { //if user already voted, we just make another check whether he selected yes/no and store it

					accepts := false
					selectedYes := reflect.DeepEqual(yesChosen, update.PollAnswer.OptionIDs)
					if selectedYes { //"Yes" should be first option in poll
						accepts = true
					}

					updateVoteStatus.voters[voterNumber].accepts = accepts
					poll[update.PollAnswer.PollID] = updateVoteStatus

				} else { //if user did not vote, we add him to an array with his opinion

					accepts := false
					selectedYes := reflect.DeepEqual(yesChosen, update.PollAnswer.OptionIDs)
					if selectedYes {
						accepts = true //"Yes" should be first option in poll
					}

					newVoter := Voter{update.PollAnswer.User.ID, accepts}
					i := updateVoteStatus.voters
					updateArray := append(i, newVoter)
					updateVoteStatus.voters = updateArray
					poll[update.PollAnswer.PollID] = updateVoteStatus
					fmt.Println("Voting successful!")
				}

			} else { //if poll is closed, then next incoming update triggers this block, which returns the result

				fmt.Println("Calculating happens!")
				var YesPower, NoPower *big.Int

				existingVoters := poll[update.PollAnswer.PollID].voters

				for _, votersList := range existingVoters {

					votepower, err := CalculatePersonPower(client_bc, auth, token_address, passportSession, token_type, votersList.tgid)
					if err != nil {
						votepower = big.NewInt(0)
					}

					if votersList.accepts {
						var c *big.Int //I'm not sure how the Add. func works, so I've decided to add a local variable
						c.Add(YesPower, votepower)
						YesPower = c
					} else {
						var c *big.Int
						c.Add(NoPower, votepower)
						NoPower = c
					}
				}

				//final result calculation
				switch YesPower.Cmp(NoPower) {

				case -1:
					accepted = false
				case 0:
					// parity
					accepted = false
				case 1:
					accepted = true
				}

				finished = true

			}

		}

	}
	return accepted, finished
}
