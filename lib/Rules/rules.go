package rules

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
	"context"
	"log"

	github_service "github.com/0xSOLIDarnost/dao-bot/lib/github"
	github_utils "github.com/0xSOLIDarnost/dao-bot/lib/utils" // use this to fetch repo by link
	"github.com/joho/godotenv"
	//passport "IKY-telegram-bot/artifacts/TGPassport"
)

var myenv map[string]string

// file with settings for enviroment
const envLoc = ".env"

// attach rules github readme link to a chat
func SetRules() {

}

// get content of rules by repo link. should return contains of README.md file
// TODO: test it!
func GetRules(repo_url string, access_token string) (string, error) {
	ctx := context.Background()
	owner, repo, err := github_utils.ParseGithubRepoURL(repo_url) //repo also should called 'README'
	if err != nil {
		//return nil, err
	}

	//res := r.ContentsURL
	service, err := github_service.NewGithubService(access_token)
	if err != nil {
		log.Println("can't create gh service")
		//return nil, err
	}

	client := service.NextClient()

	path := "/README.md" // TODO: check it

	file, _, _, err := client.HTTP.Repositories.GetContents(ctx, owner, repo, path, nil)
	if err != nil {
		return "", err
	}
	f_content, err := file.GetContent()
	if err != nil {
		return "", err
	}
	return f_content, err

}

// load enviroment variables from .env file
func loadEnv() {
	var err error
	if myenv, err = godotenv.Read(envLoc); err != nil {
		log.Printf("could not load env from %s: %v", envLoc, err)
	}
}
