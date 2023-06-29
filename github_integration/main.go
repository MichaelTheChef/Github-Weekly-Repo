package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	token           = "YOUR_GITHUB_ACCESS_TOKEN"
	topReposPerPage = 10
	emailRecipient  = "your-email@example.com"
	emailSender     = "sender-email@example.com"
	emailSMTPServer = "smtp.example.com"
	emailSMTPPort   = 587
	emailSMTPUser   = "smtp-user"
	emailSMTPPass   = "smtp-password"
)

func main() {
	ctx := context.Background()

	// Authenticate with GitHub using your access token
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Get the current date and the date of one week ago
	now := time.Now()
	oneWeekAgo := now.AddDate(0, 0, -7)

	// Set up the search options for repositories created in the past week
	opts := &github.SearchOptions{
		Sort:        "stars",
		Order:       "desc",
		ListOptions: github.ListOptions{PerPage: topReposPerPage},
	}

	query := fmt.Sprintf("created:%s..%s", oneWeekAgo.Format("2006-01-02"), now.Format("2006-01-02"))

	// Perform the search query to fetch the top repositories
	result, _, err := client.Search.Repositories(ctx, query, opts)
	if err != nil {
		log.Fatalf("Failed to fetch top repositories: %v", err)
	}

	// Prepare the email body with the top repositories
	emailBody := "Weekly Top GitHub Repositories:\n\n"
	for _, repo := range result.Repositories {
		emailBody += fmt.Sprintf("- %s (%d stars)\n", *repo.FullName, *repo.StargazersCount)
	}

	// Send the email
	sendEmail(emailRecipient, "Weekly Top GitHub Repositories", emailBody)
}
