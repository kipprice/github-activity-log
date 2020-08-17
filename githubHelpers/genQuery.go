package githubHelpers

import (
	"time"
)

func generatePRSearchString (username string, lookbackDays int, asCommenter bool) string {
	
	
	query := "archived:false is:PR user:codecademy-engineering user:Codecademy base:master base:main base:next"

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