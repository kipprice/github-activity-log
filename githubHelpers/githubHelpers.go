package githubHelpers

import (
	"github.com/google/go-github/github"
	html "github.com/kipprice/github-activity-log/htmlHelpers"
)

// PrintActivityLogForUser generates a readout of all of the activity related to
// a particular user when it comes to GitHub
func PrintActivityLogForUser(username string, lookbackDays int) string {
	out := ""
	client := createClient()

	ctx := createContext()

	// get the user themselves
	// TODO: is this needed?
	userService := client.Users
	user, _, _ := userService.Get(ctx, username)

	out += html.Header(user.GetName(), 2)

	var opts *github.SearchOptions

	searchService := client.Search

	// ==> get all PRs authored by this individual
	open, _, _ := searchService.Issues(ctx, generatePRSearchString(username, 0, false), opts)
	out += html.Header("As Author", 3)
	out += loopThroughPRs(open, client, "")

	closed, _, _ := searchService.Issues(ctx, generatePRSearchString(username, lookbackDays, false), opts)
	out += loopThroughPRs(closed, client, "")

	// ==> get all PRs reviewed bu this individual
	activeReview, _, _ := searchService.Issues(ctx, generatePRSearchString(username, 0, true), opts)
	out += html.Header("As Reviewer", 3)
	out += loopThroughPRs(activeReview, client, username)

	reviewed, _, _ := searchService.Issues(ctx, generatePRSearchString(username, lookbackDays, true), opts)
	out += loopThroughPRs(reviewed, client, username)

	return out
}

func loopThroughPRs(prList *github.IssuesSearchResult, client *github.Client, notAuthor string) string {
	out := ""

	columns := 3
	if notAuthor != "" {
		columns = 4
	}

	out += html.StartGrid(columns)

	for i := 0; i < len(prList.Issues); i++ {
		issue := prList.Issues[i]
		if issue.GetUser().GetLogin() == notAuthor {
			continue
		}

		out += printIssue(issue, client, notAuthor)
	}

	out += html.EndGrid()

	return out
}
