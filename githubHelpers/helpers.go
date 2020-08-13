package githubHelpers

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var GithubToken string;

func loadToken() {
	// TODO: load in the token
}

func createClient() *github.Client {
	
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