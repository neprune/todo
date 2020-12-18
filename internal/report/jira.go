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

func GenerateJIRAReport(wfts []todo.WellFormedTodo, address string, username string, password string) (*JIRA, error) {
	tp := jira.BasicAuthTransport{
		Username: username,
		Password: password,
	}
	client, err := jira.NewClient(tp.Client(), address)
	if err != nil {
		return nil, fmt.Errorf("failed to establish connection with JIRA: %w", err)
	}

	var missingIssues []todo.WellFormedTodo
	var closedIssues []todo.WellFormedTodo
	var doneIssues []todo.WellFormedTodo

	for _, wft := range wfts {
		copy := wft
		issue, _, err := client.Issue.Get(wft.JIRATicketID, nil)

		if err != nil && strings.Contains(err.Error(), "does not exist") || issue == nil {
			missingIssues = append(missingIssues, copy)
			continue
		}

		if err != nil {
			return nil, fmt.Errorf("failed to query for issue: %w", err)
		}

		if issue.Fields.Status.Name == "Closed" {
			closedIssues = append(closedIssues, copy)
			continue
		}

		if issue.Fields.Status.Name == "Done" {
			doneIssues = append(doneIssues, copy)
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
	t := tablewriter.NewWriter(os.Stdout)
	fmt.Println()
	fmt.Println("JIRA Report:")
	fmt.Println("===============")
	fmt.Println()

	fmt.Println("TODOs with done issues:")
	table.WriteWellFormedTodoTable(j.TodosWithDoneIssues, t)
	t.Render()
	fmt.Println()

	fmt.Println("TODOs with closed issues:")
	table.WriteWellFormedTodoTable(j.TodosWithClosedIssues, t)
	t.Render()
	fmt.Println()

	fmt.Println("TODOs with missing issues:")
	table.WriteWellFormedTodoTable(j.TodosWithMissingIssues, t)
	t.Render()
	fmt.Println()
}
