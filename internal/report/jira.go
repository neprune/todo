package report

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/neprune/todo/internal/table"
	"github.com/neprune/todo/internal/todo"
	"github.com/olekukonko/tablewriter"
	"os"
	"strings"
)

type JIRA struct {
	TodosWithMissingIssues []todo.WellFormedTodo
	TodosWithClosedIssues  []todo.WellFormedTodo
	TodosWithDoneIssues    []todo.WellFormedTodo
}

func GenerateJIRAReport(wfts []todo.WellFormedTodo, address string, username string, token string) (*JIRA, error) {
	tp := jira.BasicAuthTransport{
		Username: username,
		Password: token,
	}
	client, err := jira.NewClient(tp.Client(), address)
	if err != nil {
		return nil, fmt.Errorf("failed to establish connection with JIRA: %w", err)
	}

	var missingIssues []todo.WellFormedTodo
	var closedIssues []todo.WellFormedTodo
	var doneIssues []todo.WellFormedTodo

	for i, _ := range wfts {
		wft := wfts[i]
		issue, _, err := client.Issue.Get(wft.JIRATicketID, nil)

		if err != nil && strings.Contains(err.Error(), "does not exist") || issue == nil {
			missingIssues = append(missingIssues, wft)
			continue
		}

		if err != nil {
			return nil, fmt.Errorf("failed to query for issue: %w", err)
		}

		if issue.Fields.Status.Name == "Closed" {
			closedIssues = append(closedIssues, wft)
			continue
		}

		if issue.Fields.Status.Name == "Done" {
			doneIssues = append(doneIssues, wft)
			continue
		}
	}

	return &JIRA{
		TodosWithMissingIssues: missingIssues,
		TodosWithClosedIssues:  closedIssues,
		TodosWithDoneIssues:    doneIssues,
	}, nil
}

func (j *JIRA) OutputToTerminal() {
	fmt.Println()
	fmt.Println("JIRA Report:")
	fmt.Println("===============")
	fmt.Println()

	fmt.Println("TODOs with done issues:")
	t := tablewriter.NewWriter(os.Stdout)
	table.WriteWellFormedTodoTable(j.TodosWithDoneIssues, t)
	t.Render()
	fmt.Println()

	fmt.Println("TODOs with closed issues:")
	t = tablewriter.NewWriter(os.Stdout)
	table.WriteWellFormedTodoTable(j.TodosWithClosedIssues, t)
	t.Render()
	fmt.Println()

	fmt.Println("TODOs with missing issues:")
	t = tablewriter.NewWriter(os.Stdout)
	table.WriteWellFormedTodoTable(j.TodosWithMissingIssues, t)
	t.Render()
	fmt.Println()
}
