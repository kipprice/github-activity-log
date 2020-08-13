package main

import (
	"strings"
	"io/ioutil"
	"github.com/kipprice/github-activity-log/githubHelpers"
)

func main() {
	printPage();
}

func printPage() {
	rawUsers, _ := ioutil.ReadFile("/config/.github_team")
	lookback, _ := ioutil.ReadFile("/config/.")
	users := strings.Split(string(rawUsers), "\n")

	githubHelpers.StartPageHtml()
	for _, user := range users {
		githubHelpers.PrintActivityLogForUser(users, 14)
	}
	githubHelpers.EndPageHtml()
}