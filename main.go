package main

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kipprice/github-activity-log/githubHelpers"
)

var lookbackDays int
var usernames string

func main() {
	startServer()
}

func startServer() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(ctx *gin.Context) {
		setupPage()

		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"usernames": usernames})
	})

	// TODO: make this through an AJAX request
	r.GET("/activity/:username", func(ctx *gin.Context) {
		out := printUser(ctx.Param("username"))
		ctx.String(http.StatusOK, out)
	})
	r.Run()
}

func setupPage() {
	usernames = os.Getenv("GITHUB_TEAM")
	var err error
	lookbackDays, err = strconv.Atoi(os.Getenv("LOOKBACK_DAYS"))
	if err != nil {
		lookbackDays = 7
	}
}

func printPage() string {
	out := ""
	users := strings.Split(string(usernames), ",")

	// Start the rendering
	//out += htmlHelpers.StartPageHTML()
	for _, user := range users {
		if len(user) == 0 {
			continue
		}
		out += githubHelpers.PrintActivityLogForUser(user, lookbackDays)
	}
	//out += htmlHelpers.EndPageHTML()

	return out
}

func printUser(username string) string {
	return githubHelpers.PrintActivityLogForUser(username, lookbackDays)
}
