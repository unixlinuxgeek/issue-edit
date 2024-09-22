package issueedit

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/go-github/v64/github"
)

var tok = flag.String("t", "", "gitHub user token")
var usr = flag.String("u", "", "repository owner")
var rep = flag.String("r", "", "repository name")
var issueNum = flag.Int("num", 1, "issue number")

// ReadAll read all issues
// Usage:
// $ go test -v -run TestReadAll -u unixlinuxgeek -r gopl-issues -t <GITHUB_TOKEN>
// or
// $ go test -v -run TestReadAll -u=unixlinuxgeek -r=gopl-issues -t=<GITHUB_TOKEN>
func ReadAll(user, repo, tk string) error {
	client := github.NewClient(nil).WithAuthToken(tk)

	issues, _, err := client.Issues.ListByRepo(context.Background(), user, repo, nil)
	if err != nil {
		return fmt.Errorf("client.Issues.ListByRepo: %s", err)
	}

	repository, _, err := client.Repositories.Get(context.TODO(), user, repo)
	if err != nil {
		return fmt.Errorf("client.Repositories.Get:  %s", err)
	}
	if *repository.HasIssues {

		fmt.Print("------------------------------------------------------------------\n")
		fmt.Printf("repository name: %s\n", *repository.Name)
		fmt.Printf("language: %s\n", *repository.Language)
		fmt.Printf("open issue count: %d\n", *repository.OpenIssuesCount)

		for i, issue := range issues {
			fmt.Print("------------------------------------------------------------------\n")
			fmt.Printf("issue number: %d\n", *issue.Number)
			fmt.Printf("issue state: %s\n", *issue.State)
			fmt.Printf("issue title: %s\n", *issue.Title)
			if i == len(issues)-1 {
				fmt.Print("------------------------------------------------------------------\n")
			}
		}

	} else {
		fmt.Printf("repository %s does not contain any issues\n", *rep)
	}

	return nil
}

// ReadOne Read only one issue
func ReadOne(user, repo, tk string, issueNum int) error {
	client := github.NewClient(nil).WithAuthToken(tk)

	issues, _, err := client.Issues.ListByRepo(context.Background(), user, repo, nil)
	if err != nil {
		return fmt.Errorf("client.Issues.ListByRepo: %s", err)
	}

	repository, _, err := client.Repositories.Get(context.TODO(), user, repo)
	if err != nil {
		return fmt.Errorf("client.Repositories.Get:  %s", err)
	}
	if *repository.HasIssues {

		for _, issue := range issues {
			if *issue.Number == issueNum {
				fmt.Print("----------------------------------------------------------------------------\n")
				fmt.Printf("issue title: %s\n", *issue.Title)
				fmt.Printf("issue number: %d\n", *issue.Number)
				fmt.Printf("issue id: %d\n", *issue.ID)
				fmt.Printf("issue state: %s\n", *issue.State)
				fmt.Printf("issue labels: %s\n", issue.Labels)
				fmt.Printf("issue url: %s\n", *issue.URL)
				fmt.Print("----------------------------------------------------------------------------\n")
			}
		}
	} else {
		fmt.Printf("repository %s does not contain any issues\n", *rep)
	}

	return nil
}

// Create creating a new issue
// Example1: $ go test -v -run TestCreate -u unixlinuxgeek -r gopl-issues  -t <SECRET>
// Example2: $ go test -v -run TestCreate -u=unixlinuxgeek -r=gopl-issues  -t=<SECRET>
func Create(user, repo, tk string, issueReq *github.IssueRequest) error {
	client := github.NewClient(nil).WithAuthToken(tk)

	// Example issueReq:
	//is := &github.IssueRequest{
	//	Title: github.String("test_issue_" + now.Format("2006-01-02_15:04")),
	//}
	iss, _, err := client.Issues.Create(context.TODO(), user, repo, issueReq)

	if err != nil {
		return err
	}
	fmt.Printf("issue state: %s, owner: %s, repo: %s\n", *iss.State, *usr, *rep)
	return nil
}

// Update updating an issue
// Example1: $ go test -v -run TestUpdate -u unixlinuxgeek -r gopl-issues -num 6 -t <SECRET>
// Example2: $ go test -v -run TestUpdate -u=unixlinuxgeek -r=gopl-issues -num=6 -t=<SECRET>
func Update(user, repo, tk string, issueReq *github.IssueRequest, issueNum int) error {
	client := github.NewClient(nil).WithAuthToken(tk)

	// Example: issueReq:
	// issueReq := &github.IssueRequest{
	//    Title: github.String("ok"),
	//    ...
	// }
	issue, _, err := client.Issues.Edit(context.Background(), user, repo, issueNum, issueReq)

	if issue != nil {
		fmt.Printf("issue number: %d\n", issue.Number)
		fmt.Printf("issue id: %d\n", *issue.ID)
		fmt.Printf("issue title: %d\n", issue.Title)
		if issue.User.Name != nil {
			fmt.Printf("issue user: %s\n", *issue.User.Name)
		}
		fmt.Printf("issue url: %s\n", *issue.URL)
		fmt.Printf("issue state: %s\n", *issue.State)
		fmt.Printf("issue labels: %v\n", issue.Labels)
	}

	if err != nil {
		return err
	}

	return nil
}

// Close change issue status to closed
// Example1: $ go test -v -run TestClose -u unixlinuxgeek -r gopl-issues -num 6 -t <SECRET>
// Example2: $ go test -v -run TestClose -u=unixlinuxgeek -r=gopl-issues -num=6 -t=<SECRET>
func Close(user, repo, tk string, issueNum int) error {
	client := github.NewClient(nil).WithAuthToken(tk)

	is := &github.IssueRequest{
		State: github.String("closed"),
	}
	iss, _, err := client.Issues.Edit(context.TODO(), user, repo, issueNum, is)

	if err != nil {
		return err
	}

	fmt.Printf("Issue state: %s, owner: %s, repo: %s. issue number: %d\n", *iss.State, *usr, *rep, issueNum)
	return nil
}
