package utils

import (
	"context"
	"errors"
	"net/url"

	//"github.com/SporkHubr/Spork-go/lib"
	"github.com/google/go-github/v38/github"

	s "strings"

	"golang.org/x/oauth2"
)

// fetch repo by name + owner
func FetchRepoByName(owner string, reponame string, token string) (*github.Repository, error) {
	ctx := context.Background()
	var client *github.Client
	if token != "" {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	} else {
		client = github.NewClient(nil)
	}
	repo, _, err := client.Repositories.Get(ctx, owner, reponame)
	return repo, err
}

// ParseGithubRepoURL returns owner & repo by given github repo url
func ParseGithubRepoURL(source string) (string, string, error) {
	u, err := url.Parse(source)
	if err != nil {
		return "", "", err
	}

	splittedPath := s.Split(u.Path, "/")
	if len(splittedPath) != 3 {
		return "", "", errors.New("invalid github repo url")
	}
	owner := splittedPath[1]
	repo := splittedPath[2]
	return owner, repo, nil
}

func GetLanguage(repo *github.Repository) string {
	//	client := github.NewClient(nil)
	language := repo.GetLanguage()
	return language
}

/*
func GetLanguages(client *lib.GithubClient, owner string, name string) ([]string, error) {
	ctx := context.Background()
	langs, _, err := client.HTTP.Repositories.ListLanguages(ctx, owner, name)
	if err != nil {
		return nil, err
	}

	languages := make([]struct {
		lang  string
		bytes int
	}, 0, len(langs))
	for k, v := range langs {
		languages = append(languages, struct {
			lang  string
			bytes int
		}{
			lang:  k,
			bytes: v,
		})
	}

	sort.Slice(languages, func(i, j int) bool {
		return languages[i].bytes > languages[j].bytes
	})

	l := make([]string, 0, len(languages))
	for _, v := range languages {
		l = append(l, v.lang)
	}

	return l, nil
}
*/
