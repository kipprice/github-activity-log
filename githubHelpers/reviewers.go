package githubHelpers

import (
	"fmt"
	"github.com/google/go-github/github"
)

func getReviewers(issue github.Issue, reviewMap map[string]string) string {
	out := ""
	for u, v := range reviewMap {
		if len(out) > 0 { out += "<br>" }
		out += fmt.Sprintf("%v %v", emojiForStatus(v), u)
	}
	return out
}

func emojiForStatus (status string) string {
	switch (status) {
		
	case "APPROVED":
		return "✅"

	case "COMMENTED":
		return "💬"

	case "CHANGES_REQUESTED":
		return "⛔️"
	}
	return ""
}