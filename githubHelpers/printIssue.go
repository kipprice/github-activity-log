package githubHelpers

import (
	"fmt"
	"strings"
	"net/url"
	"github.com/google/go-github/github"
	html "github.com/kipprice/github-activity-log/htmlHelpers"
)


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