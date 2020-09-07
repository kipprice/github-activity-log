package githubHelpers

import (
	"fmt"

	"github.com/google/go-github/github"
)

func getStatus(pr *github.PullRequest, reviewMap map[string]string) string {
	if pr.GetState() == "Closed" {
		return "closed"
	}

	approvalCount := 0
	commentCount := 0
	changesNeededCount := 0
	totalReviews := 0

	for _, v := range reviewMap {
		totalReviews++

		switch v {

		case "CHANGES_NEEDED":
			changesNeededCount++
			break

		case "COMMENTED":
			commentCount++
			break

		case "APPROVED":
			approvalCount++
			break
		}
	}

	if changesNeededCount > 0 {
		return fmt.Sprintf("Changes Needed")
	}
	if approvalCount == len(reviewMap) && approvalCount > 0 {
		return fmt.Sprintf("Approved")
	}
	if approvalCount > 0 {
		return fmt.Sprintf("Partial Approval")
	}
	if commentCount > 0 {
		return fmt.Sprintf("Has Feedback")
	}
	return "Not Reviewed"
}
