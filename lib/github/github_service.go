package lib

/**
*		Github service is a code provided by GitJob project (https://github.com/SporkHubr)
*		Authors: JackBekket, ChronosX88, Terisback
*		Original repo is 'private' so I can't just import original source
*		Also probably most part of it wouldn't really need here, so I consider refactor this lib and swipe all we do not need after finishing rules modeule
 */

import (
	"context"
	"errors"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ChronosX88/go-throttle"
	"github.com/reactivex/rxgo/v2"
	"github.com/shurcooL/githubv4"

	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

const (
	PerPage          = 30 // Max 100, but it triggers anti abuse system
	DefaultRateLimit = 200 * time.Millisecond
)

type GithubService struct {
	// Current github client for this instance of service
	//
	// Even after CycleProxy(), Client will contain current GithubClient instance used for api calls
	//
	// If you created github service with token, it will be the only client in that service (nothing will change on CycleProxy())
	client *GithubClient
	// Used when no token provided to NewGithubService
	clients []*GithubClient
	idx     uint64
}

type GithubClient struct {
	HTTP         *github.Client
	httpThrottle *throttle.Throttle
	GraphQL      *githubv4.Client
}

func NewGithubService(ctx context.Context, token string) (*GithubService, error) {
	if token == "" {
		return nil, errors.New("provided github token is empty")
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	graphClient := githubv4.NewClient(tc)

	return &GithubService{
		client: &GithubClient{
			HTTP:         client,
			httpThrottle: throttle.NewThrottle(ctx, nil),
			GraphQL:      graphClient,
		},
	}, nil
}

func (gs *GithubService) NextClient() *GithubClient {
	// If got token authorized client
	if gs.client != nil {
		return gs.client
	}

	return gs.clients[atomic.AddUint64(&gs.idx, 1)%uint64(len(gs.clients))]
}

// Creates Observable that you can wait via `<-gs.Throttle()`
//
// Follows the default rate limit of 2 seconds for the github
func (gc *GithubClient) Throttle(rateLimits ...time.Duration) <-chan rxgo.Item {
	return gc.httpThrottle.Take(sumDurations(rateLimits...), 1.0).Observe()
}

func (gc *GithubClient) DefaultThrottle() *GithubClient {
	<-gc.httpThrottle.Take(DefaultRateLimit, 1.0).Observe()
	return gc
}

// Returns true if github service was created with token
func (gc *GithubService) Authorized() bool {
	return gc.client != nil
}

func (gs *GithubService) ListRepositories(user string, opts *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error) {
	return gs.
		NextClient().
		DefaultThrottle().
		HTTP.Repositories.List(context.Background(), user, opts)
}

func (gs *GithubService) SearchCode(pageNumber int, perPageSum int, query string) (*github.CodeSearchResult, bool, error) {
	var isLastPage bool

	opts := &github.SearchOptions{ListOptions: github.ListOptions{Page: 1, PerPage: 10}}

	if pageNumber != -1 && perPageSum != -1 {
		opts = &github.SearchOptions{ListOptions: github.ListOptions{Page: pageNumber, PerPage: perPageSum}}
	}

	result, resp, err := gs.
		NextClient().DefaultThrottle().
		HTTP.Search.Code(context.Background(), query, opts)
	if err != nil {
		return nil, false, err
	}

	if resp.NextPage == 0 {
		isLastPage = true
	}

	return result, isLastPage, nil
}

func (gs *GithubService) ListAllRepositoriesAsync(user string) rxgo.Observable {
	resultChannel := make(chan rxgo.Item)
	go func(resultChannel chan rxgo.Item) {
		defer close(resultChannel)
		opt := &github.RepositoryListOptions{
			Type:      "username",
			Sort:      "updated",
			Direction: "desc",
			ListOptions: github.ListOptions{
				Page:    1,
				PerPage: PerPage,
			},
		}

		// get repos of user
		repos, res, err := gs.ListRepositories(user, opt)
		if err != nil {
			resultChannel <- rxgo.Error(err)
			return
		}
		var rRepos []*github.Repository
		for res.NextPage != 0 {
			rRepos, res, err = gs.ListRepositories(user, opt)
			if err != nil {
				resultChannel <- rxgo.Error(err)
				return
			}
			repos = append(repos, rRepos...)
			opt.Page++
		}
		resultChannel <- rxgo.Of(repos)
	}(resultChannel)
	return rxgo.FromChannel(resultChannel)
}

func (gs *GithubService) ListAllOrganizationsAsync(user string) rxgo.Observable {
	res := make(chan rxgo.Item)
	go func(resultChannel chan rxgo.Item) {
		defer close(resultChannel)
		orgListOpts := &github.ListOptions{
			Page:    1,
			PerPage: PerPage,
		}
		orgs, res, err := gs.NextClient().DefaultThrottle().HTTP.Organizations.List(context.Background(), user, orgListOpts)
		if err != nil {
			resultChannel <- rxgo.Error(err)
			return
		}
		var rOrgs []*github.Organization
		for res.NextPage != 0 {
			rOrgs, res, err = gs.NextClient().DefaultThrottle().HTTP.Organizations.List(context.Background(), user, orgListOpts)
			if err != nil {
				resultChannel <- rxgo.Error(err)
				return
			}
			orgs = append(orgs, rOrgs...)
			orgListOpts.Page++
		}
		resultChannel <- rxgo.Of(orgs)
	}(res)
	return rxgo.FromChannel(res)
}

func (gs *GithubService) ListAllReposInOrgAsync(org string) rxgo.Observable {
	res := make(chan rxgo.Item)
	go func(resultChannel chan rxgo.Item) {
		defer close(resultChannel)
		orgRepoListOpts := &github.RepositoryListByOrgOptions{
			ListOptions: github.ListOptions{
				Page:    1,
				PerPage: PerPage,
			},
		}
		orgRepos, res, err := gs.NextClient().DefaultThrottle().HTTP.Repositories.ListByOrg(context.Background(), org, orgRepoListOpts)
		if err != nil {
			resultChannel <- rxgo.Error(err)
			return
		}
		for res.NextPage != 0 {
			var nOrgRepos []*github.Repository
			nOrgRepos, res, err = gs.NextClient().DefaultThrottle().HTTP.Repositories.ListByOrg(context.Background(), org, orgRepoListOpts)
			if err != nil {
				resultChannel <- rxgo.Error(err)
				return
			}
			orgRepos = append(orgRepos, nOrgRepos...)
			orgRepoListOpts.Page++
		}
		resultChannel <- rxgo.Of(orgRepos)
	}(res)
	return rxgo.FromChannel(res)
}

func (gs *GithubService) GetLanguagesForRepoAsync(repo *github.Repository) rxgo.Observable {
	res := make(chan rxgo.Item)
	go func(resultChannel chan rxgo.Item) {
		defer close(resultChannel)

		if *repo.Size == 0 {
			resultChannel <- rxgo.Of(nil)
			return
		}

		languages, _, err := gs.NextClient().DefaultThrottle().HTTP.Repositories.ListLanguages(context.Background(), repo.GetOwner().GetLogin(), repo.GetName())
		if err != nil {
			resultChannel <- rxgo.Error(err)
			return
		}
		res := map[string]interface{}{}
		res["languages"] = languages
		res["repo"] = repo
		resultChannel <- rxgo.Of(res)
	}(res)
	return rxgo.FromChannel(res)
}

func (gs *GithubService) GetTotalCountOfCommitsByUsername(owner, repo, username string) (int, error) {
	contributors, res, err := gs.NextClient().DefaultThrottle().HTTP.Repositories.ListContributorsStats(context.Background(), owner, repo)
	for res.StatusCode == 202 {
		contributors, res, err = gs.NextClient().DefaultThrottle().HTTP.Repositories.ListContributorsStats(context.Background(), owner, repo)
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		return 0, err
	}
	var totalCountOfCommitsByUsername int
	for _, contributor := range contributors {
		if *contributor.Author.Login == username {
			totalCountOfCommitsByUsername = *contributor.Total
		}
	}
	return totalCountOfCommitsByUsername, nil
}

func (gs *GithubService) GetRepoStarsCount(owner, repoName string) (int, error) {
	repo, _, err := gs.NextClient().DefaultThrottle().HTTP.Repositories.Get(context.Background(), owner, repoName)
	if err != nil {
		return 0, err
	}
	return repo.GetStargazersCount(), nil
}

func (gs *GithubService) GetRepoForksCount(owner, repoName string) (int, error) {
	repo, _, err := gs.NextClient().DefaultThrottle().HTTP.Repositories.Get(context.Background(), owner, repoName)
	if err != nil {
		return 0, err
	}
	return repo.GetForksCount(), nil
}

func (gs *GithubService) GetUser(username string) (*github.User, error) {
	user, _, err := gs.NextClient().DefaultThrottle().HTTP.Users.Get(context.Background(), username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (gs *GithubService) GetRepo(owner, repoName string) (*github.Repository, error) {
	repo, _, err := gs.NextClient().DefaultThrottle().HTTP.Repositories.Get(context.Background(), owner, repoName)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (gs *GithubService) ListRepoLanguages(owner, repo string) (map[string]int, error) {
	languages, res, err := gs.NextClient().DefaultThrottle().HTTP.Repositories.ListLanguages(context.Background(), owner, repo)
	for res.StatusCode == 202 {
		languages, res, err = gs.NextClient().DefaultThrottle().HTTP.Repositories.ListLanguages(context.Background(), owner, repo)
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		return nil, err
	}
	return languages, nil
}

func (gs *GithubService) GetOutsideContributionsAsync(user string) rxgo.Observable {
	res := make(chan rxgo.Item)
	go func(resultChannel chan rxgo.Item) {
		defer close(resultChannel)

		repos, err := gs.GetOutsideContributions(user)
		if err != nil {
			resultChannel <- rxgo.Error(err)
			return
		}

		resultChannel <- rxgo.Of(repos)
	}(res)

	return rxgo.FromChannel(res)
}

func (gs *GithubService) GetOutsideContributions(user string) ([]*github.Repository, error) {
	type reposNames struct {
		NameWithOwner string
	}

	var query struct {
		User struct {
			RepositoriesContributedTo struct {
				Nodes    []reposNames
				PageInfo struct {
					EndCursor   githubv4.String
					HasNextPage bool
				}
			} `graphql:"repositoriesContributedTo(first: 100, contributionTypes: [COMMIT, PULL_REQUEST, REPOSITORY], after: $contribCursor)"`
		} `graphql:"user(login: $user)"`
	}

	variables := map[string]interface{}{
		"user":          githubv4.String(user),
		"contribCursor": (*githubv4.String)(nil),
	}

	var reposNamesArray []reposNames
	for {
		err := gs.NextClient().DefaultThrottle().GraphQL.Query(context.Background(), &query, variables)
		if err != nil {
			return nil, err
		}
		reposNamesArray = append(reposNamesArray, query.User.RepositoriesContributedTo.Nodes...)
		if !query.User.RepositoriesContributedTo.PageInfo.HasNextPage {
			break
		}
		variables["contribCursor"] = githubv4.NewString(query.User.RepositoriesContributedTo.PageInfo.EndCursor)
	}

	var repos []*github.Repository
	for _, v := range reposNamesArray {
		splittedRepoName := strings.Split(v.NameWithOwner, "/")
		if len(splittedRepoName) != 2 {
			return nil, errors.New("invalid repo name")
		}
		repo, err := gs.GetRepo(splittedRepoName[0], splittedRepoName[1])
		if err != nil {
			return nil, err
		}
		repos = append(repos, repo)
	}

	return repos, nil
}

func sumDurations(durations ...time.Duration) time.Duration {
	var duration time.Duration

	for _, d := range durations {
		duration += d
	}

	return duration
}
