package githubHelpers

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-github/github"
	html "github.com/kipprice/github-activity-log/htmlHelpers"
)

// printIssue generates the markup for each PR
func printIssue(issue github.Issue, client *github.Client, notAuthor string) string {
	out := ""
	ctx := createContext()

	if !issue.IsPullRequest() {
		return out
	}

	// ==> Gather all of the data
	u, _ := url.Parse(issue.GetPullRequestLinks().GetURL())

	pathPieces := strings.Split(u.Path, "/")
	owner := pathPieces[2]
	repo := pathPieces[3]

	service := client.PullRequests
	pr, _, _ := service.Get(ctx, owner, repo, issue.GetNumber())
	var opts *github.ListOptions
	reviews, _, _ := service.ListReviews(ctx, owner, repo, issue.GetNumber(), opts)

	var reviewMap = generateReviewMap(reviews, issue)

	link := fmt.Sprintf("https://github.com/%v/%v/pull/%v", owner, repo, pr.GetNumber())

	// ==> Print all of the things

	// PR Title
	out += html.StartSpan("")
	out += html.A(fmt.Sprintf("%v/%v #%v", owner, repo, pr.GetNumber()), link, "label")
	out += html.Div(pr.GetTitle(), "title", "")
	out += html.EndSpan()

	// author
	if notAuthor != "" {
		out += html.Span(pr.GetUser().GetLogin(), "author", "")
	}

	// status
	out += html.Span(getStatus(pr, reviewMap), "status", "")

	// reviewers
	out += html.Span(getReviewers(issue, reviewMap), "reviewers", "")

	// comments
	// TODO: map this back to reviewers too
	out += html.Span(fmt.Sprintf("%v comments", pr.GetComments()), "comments", "")

	return out
}

// generateReviewMap creates the mapping of reviews to reviewers
func generateReviewMap(reviews []*github.PullRequestReview, issue github.Issue) map[string]string {
	reviewMap := make(map[string]string)
	for r := 0; r < len(reviews); r++ {
		review := reviews[r]
		userId := review.User.GetLogin()
		if userId == issue.GetUser().GetLogin() {
			continue
		}

		reviewMap[userId] = review.GetState()
	}
	return reviewMap
}
