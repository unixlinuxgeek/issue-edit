package issueedit

import (
	"github.com/google/go-github/v64/github"
	"testing"
	"time"
)

// create issue test
// Example1: $ go test -v -run TestCreate -u unixlinuxgeek -r gopl-issues  -t <SECRET>
// Example2: $ go test -v -run TestCreate -u=unixlinuxgeek -r=gopl-issues  -t=<SECRET>
func TestCreate(t *testing.T) {

	now := time.Now().Local()
	issueReq := &github.IssueRequest{
		Title: github.String("test_issue_" + now.Format("2006-01-02_15:04")),
	}

	err := Create(*usr, *rep, *tok, issueReq)
	if err != nil {
		t.Errorf("test_create: %s\n", err)
	}
	t.Log("TestCreate is PASSED\n")
}

// read all issues test
// Example1: $ go test -v -run TestReadAll -u unixlinuxgeek -r gopl-issues -t <SECRET>
// Example2: $ go test -v -run TestReadAll -u=unixlinuxgeek -r=gopl-issues -t=<SECRET>
func TestReadAll(t *testing.T) {
	err := ReadAll(*usr, *rep, *tok)
	if err != nil {
		t.Errorf("test_read: %s\n", err)
	}
	t.Log("TestReadAll is PASSED\n")
}

// read single issue test
// Example1: $ go test -v -run TestReadOne -u unixlinuxgeek -r gopl-issues -num 4 -t <SECRET>
// Example2: $ go test -v -run TestReadOne -u=unixlinuxgeek -r=gopl-issues -num=4 -t=<SECRET>
func TestReadOne(t *testing.T) {
	err := ReadOne(*usr, *rep, *tok, *issueNum)
	if err != nil {
		t.Errorf("test_read: %s\n", err)
	}
	t.Log("TestReadOne is PASSED\n")
}

// update issue test
// Example1: $ go test -v -run TestUpdate -u unixlinuxgeek -r gopl-issues -issueNum 6 -t <SECRET>
// Example2: $ go test -v -run TestUpdate -u=unixlinuxgeek -r=gopl-issues -issueNum=6 -t=<SECRET>
func TestUpdate(t *testing.T) {
	n := time.Now()
	issueReq := &github.IssueRequest{Title: github.String("test_issue_updated_" + n.Format("2006-01-02 15:04"))}
	err := Update(*usr, *rep, *tok, issueReq, *issueNum)
	if err != nil {
		t.Errorf("test_update: %s", err.Error())
	}
	t.Log("TestUpdate is PASSED\n")
}

// close issue test
// Example1: $ go test -v -run TestClose -u unixlinuxgeek -r gopl-issues -issueNum 6 -t <SECRET>
// Example2: $ go test -v -run TestClose -u=unixlinuxgeek -r=gopl-issues -issueNum=6 -t=<SECRET>
func TestClose(t *testing.T) {
	err := Close(*usr, *rep, *tok, *issueNum)
	if err != nil {
		t.Errorf("%s\n", err)
	}
	t.Log("TestClose is PASSED\n")
}
