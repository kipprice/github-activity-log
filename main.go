package main

import (
	"strings"
	"strconv"
	"os"
	"github.com/kipprice/github-activity-log/githubHelpers"
	"github.com/kipprice/github-activity-log/htmlHelpers"
)

func main() {
	printPage();
}

func printPage() {
	rawUsers := os.Getenv("GITHUB_TEAM")

	users := strings.Split(string(rawUsers), ",")


	lookbackDays, err := strconv.Atoi(os.Getenv("LOOKBACK_DAYS"))
	if err != nil { lookbackDays = 7 }

	// Start the rendering
	htmlHelpers.StartPageHtml()
	for _, user := range users {
		if len(user) == 0 { continue }
		githubHelpers.PrintActivityLogForUser(user, lookbackDays)
	}
	htmlHelpers.EndPageHtml()
}