package githubHelpers

import (
	"github.com/google/go-github/github"
	html "github.com/kipprice/github-activity-log/htmlHelpers"
	"time"
	"fmt"
	"net/url"
	"strings"
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


func generatePRSearchString (username string, lookbackDays int, asCommenter bool) string {
	
	query := "archived:false is:PR user:codecademy-engineering user:Codecademy"

	if (asCommenter) {
		query += " commenter:" + username
	} else {
		query += " author:" + username
	}

	if (lookbackDays > 0) {
		lastWeek := time.Now().AddDate(0, 0, -1 * lookbackDays);
		query += " is:closed updated:>" + lastWeek.Format("2006-01-02")
	} else {
		query += " is:open"
	}
	
	return query 
}

func loopThroughPRs (prList *github.IssuesSearchResult, client *github.Client, notAuthor string) {

	columns := 3
	if notAuthor != "" { columns = 4 }

	fmt.Printf("<div style='display: grid; grid-template-columns: 1fr repeat(%v, 10vw);'>", columns)
	for i := 0; i < len(prList.Issues); i++ {
		issue := prList.Issues[i]
		if issue.GetUser().GetLogin() == notAuthor { continue }

		printIssue(issue, client, notAuthor)
	}
	fmt.Println("</div>")
}

func printIssue (issue github.Issue, client *github.Client, notAuthor string) {
	ctx := createContext()

	if (issue.IsPullRequest()) {
		u, _ := url.Parse(issue.GetPullRequestLinks().GetURL())

		pathPieces := strings.Split(u.Path, "/")
		owner := pathPieces[2]
		repo := pathPieces[3]

		service := client.PullRequests
		pr, _, _ := service.Get(ctx, owner, repo, issue.GetNumber())
		var opts *github.ListOptions
		reviews, _, _ := service.ListReviews(ctx, owner, repo, issue.GetNumber(), opts)

		var reviewMap = make(map[string]string)
		for r := 0; r < len(reviews); r ++ {
			review := reviews[r]
			userId := review.User.GetLogin()
			if (userId == issue.GetUser().GetLogin()) { continue }

			reviewMap[userId] = review.GetState()
		} 
	
		link := fmt.Sprintf("https://github.com/%v/%v/pull/%v", owner, repo, pr.GetNumber())


		html.A(pr.GetTitle(), link, "title")

		// author
		if notAuthor != "" { html.Span(pr.GetUser().GetLogin(), "author") }

		// status
		html.Span(getStatus(issue, reviewMap), "status")

		// link to repo
		html.A(fmt.Sprintf("%v/%v #%v", owner, repo, pr.GetNumber()), link, "link")

		// comments
		html.Span(fmt.Sprintf("%v comments", pr.GetComments()), "comments")

	}
}

func getStatus (issue github.Issue, reviewMap map[string]string) string {
	if (issue.GetState() == "Closed") { return "closed" }

	approvalCount := 0
	commentCount := 0
	changesNeededCount := 0
	for _, v := range reviewMap {
		switch (v) {

			case "CHANGES_NEEDED":
				changesNeededCount += 1;
				break;

			case "COMMENTED":
				commentCount += 1;
				break;

			case "APPROVED":
				approvalCount += 1;
				break;
		}
	}

	if changesNeededCount > 0 { return "Changes Needed" }
	if (approvalCount == len(reviewMap)) { return "Approved" }
	if (approvalCount > 0 ) { return "Partial Approval" }
	if (commentCount > 0) { return "Has Feedback" }
	return "Not Reviewed"
}
