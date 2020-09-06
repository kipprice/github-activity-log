package githubHelpers

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func generatePRSearchString(username string, lookbackDays int, asCommenter bool) string {

	orgsQuery := generateOrgsQuery()
	branchesQuery := generateBranchesQuery()

	query := fmt.Sprintf("archived:false is:PR %v %v", orgsQuery, branchesQuery)

	if asCommenter {
		query += " commenter:" + username
	} else {
		query += " author:" + username
	}

	if lookbackDays > 0 {
		lastWeek := time.Now().AddDate(0, 0, -1*lookbackDays)
		query += " is:closed updated:>" + lastWeek.Format("2006-01-02")
	} else {
		query += " is:open"
	}

	return query
}

func generateOrgsQuery() string {
	return generateQueryFromEnvVar("GITHUB_ORGS", "user")
}

func generateBranchesQuery() string {
	return generateQueryFromEnvVar("GITHUB_BRANCHES", "base")
}

func generateQueryFromEnvVar(envVarName string, prefix string) string {
	rawValues := os.Getenv(envVarName)
	values := strings.Split(string(rawValues), ",")
	query := ""
	for _, val := range values {
		if len(val) == 0 {
			continue
		}
		query += fmt.Sprintf(" %v:%v", prefix, val)
	}
	return query
}
