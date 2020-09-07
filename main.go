package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

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
	r.LoadHTMLGlob("./templates/*")

	// index page: just send usernames, which will be fetched async
	r.GET("/", func(ctx *gin.Context) {
		setupPage()

		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{"usernames": usernames})
	})

	// fetch user data async
	r.GET("/activity/:username", func(ctx *gin.Context) {
		out := printUser(ctx.Param("username"))
		ctx.String(http.StatusOK, out)
	})

	// info message
	port := os.Getenv("PORT")
	fmt.Printf("Go to http://localhost:%v to view GitHub activity", port)

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

func printUser(username string) string {
	return githubHelpers.PrintActivityLogForUser(username, lookbackDays)
}
