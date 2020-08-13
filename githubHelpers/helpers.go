package githubHelpers

import (
	"os"
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var GithubToken string;

func loadToken() {
	GithubToken = os.Getenv("GITHUB_TOKEN")
}

func createClient() *github.Client {
	if GithubToken == "" { loadToken() }

	ctx := createContext();
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client
}

func createContext() context.Context {
	return context.Background();
}