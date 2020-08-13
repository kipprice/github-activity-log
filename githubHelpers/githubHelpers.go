package githubHelpers

import (
	"github.com/google/go-github/github"
	html "github.com/kipprice/github-activity-log/htmlHelpers"
)

func PrintActivityLogForUser(username string, lookbackDays int) {
	client := createClient()
	
	ctx := createContext()

	// get the user themselves
	// TODO: is this needed?
	userService := client.Users
	user, _, _ := userService.Get(ctx, username)

	html.Header(user.GetName(), 2)
	
	var opts *github.SearchOptions
	
	searchService := client.Search
	open, _, _ := searchService.Issues(ctx, generatePRSearchString(username, 0, false), opts)
	html.Header("As Author", 3)
	loopThroughPRs(open, client, "")

	closed, _, _ := searchService.Issues(ctx, generatePRSearchString(username, lookbackDays, false), opts)
	loopThroughPRs(closed, client, "")


	activeReview, _, _ := searchService.Issues(ctx, generatePRSearchString(username, 0, true), opts)
	html.Header("As Reviewer", 3)
	loopThroughPRs(activeReview, client, username)

	reviewed, _, _ := searchService.Issues(ctx, generatePRSearchString(username, lookbackDays, true), opts)
	loopThroughPRs(reviewed, client, username)
}

func loopThroughPRs (prList *github.IssuesSearchResult, client *github.Client, notAuthor string) {

	columns := 3
	if notAuthor != "" { columns = 4 }

	html.StartGrid(columns);
	for i := 0; i < len(prList.Issues); i++ {
		issue := prList.Issues[i]
		if issue.GetUser().GetLogin() == notAuthor { continue }

		printIssue(issue, client, notAuthor)
	}
	html.EndGrid();
}
