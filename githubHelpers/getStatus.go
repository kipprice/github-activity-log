package githubHelpers

import (
	"github.com/google/go-github/github"
)

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