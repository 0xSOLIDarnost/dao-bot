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
	"fmt"
	"log"
	"time"

	github_service "github.com/0xSOLIDarnost/dao-bot/lib/github"
	github_utils "github.com/0xSOLIDarnost/dao-bot/lib/utils" // use this to fetch repo by link
	"github.com/google/go-github/v38/github"
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
func GetRules(ctx context.Context, repo_url string, access_token string) (string, error) {
	owner, repo, err := github_utils.ParseGithubRepoURL(repo_url) //repo also should called 'README'
	if err != nil {
		return "", err
	}

	service, err := github_service.NewGithubService(ctx, access_token)
	if err != nil {
		log.Println("can't create gh service")
		return "", err
	}

	readme, _, err := service.NextClient().HTTP.Repositories.GetReadme(ctx, owner, repo, nil)
	if err != nil {
		return "", err
	}

	return readme.GetContent()
}

// addRule Creates PR in repository which adding {content} to README file
func AddRule(ctx context.Context, repo_url string, access_token string, content string) error {
	owner, repo, err := github_utils.ParseGithubRepoURL(repo_url)
	if err != nil {
		return err
	}

	service, err := github_service.NewGithubService(ctx, access_token)
	if err != nil {
		return err
	}

	repository, _, err := service.NextClient().HTTP.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return err
	}

	// ? Check case when file does not exist
	readme, resp, err := service.NextClient().HTTP.Repositories.GetReadme(ctx, owner, repo, nil)
	if err != nil && resp.StatusCode != 404 {
		fmt.Printf("Failed to get contents from github (%v): %v", resp.StatusCode, err)
		return err
	}

	readmeContent, err := readme.GetContent()
	if err != nil {
		return err
	}

	// Get default branch for reference
	master, _, err := service.NextClient().HTTP.Repositories.GetBranch(ctx, owner, repo, repository.GetDefaultBranch(), true)
	if err != nil {
		return err
	}

	// Create new branch
	branchName := fmt.Sprintf("new-rule-%v", time.Now().Unix())
	ref := "refs/heads/" + branchName
	_, _, err = service.NextClient().HTTP.Git.CreateRef(ctx, owner, repo, &github.Reference{
		Ref: &ref,
		Object: &github.GitObject{
			SHA: master.Commit.SHA,
		},
	})
	if err != nil {
		return err
	}

	// Update README file with new rule and commit
	commitName := "Add rule"
	updatedContent := readmeContent + "\n" + content
	_, _, err = service.NextClient().HTTP.Repositories.UpdateFile(ctx, owner, repo, "README.md", &github.RepositoryContentFileOptions{
		Message: &commitName,
		Content: []byte(updatedContent),
		SHA:     readme.SHA,
		Branch:  &branchName,
	})
	if err != nil {
		return err
	}

	// Open Pull Request
	_, _, err = service.NextClient().HTTP.PullRequests.Create(ctx, owner, repo, &github.NewPullRequest{
		Title: &commitName,
		Head:  &branchName,
		Base:  master.Name,
	})
	if err != nil {
		return err
	}

	return nil
}

// load enviroment variables from .env file
func loadEnv() {
	var err error
	if myenv, err = godotenv.Read(envLoc); err != nil {
		log.Printf("could not load env from %s: %v", envLoc, err)
	}
}
