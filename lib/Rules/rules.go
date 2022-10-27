package main

/**
*		TODO:
*		1. installRules (constructor?) should store(?) link to github repo with readme
*		2. GetRules should get readme file from associated repo and pass it to main routine (as string?)
*		3. AddRule should accept string and open new PR to repo with rules (add strings to README.MD)
*		If user want to add rule he should do that through main goroutine(?) which should start voting from Voter package
*		If vote was succesfull, *THEN* main call AddRule command
*
*
 */

import (
	"log"

	"github.com/joho/godotenv"
	//g_utils "github.com/SporkHubr/Spork-go/tree/master/lib/utils"		// use this to fetch repo by link

	passport "github.com/MoonSHRD/IKY-telegram-bot/artifacts/TGPassport"

	//passport "IKY-telegram-bot/artifacts/TGPassport"

	"github.com/ethereum/go-ethereum/common"
)







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
