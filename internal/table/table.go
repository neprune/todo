package table

import (
	"fmt"
	"github.com/neprune/todo/internal/todo"
	"github.com/olekukonko/tablewriter"
	"math"
	"time"
)

func location(t *todo.Todo) string {
	return fmt.Sprintf("%s:%d", t.Filepath, t.LineNumber)
}

func WriteBadlyFormedTodoTable(bfts []*todo.BadlyFormedTodo, t *tablewriter.Table) {
	t.SetHeader([]string{"Location", "Line", "Parse Error"})
	for _, bft := range bfts {
		t.Append([]string{location(bft.Todo), bft.Line, bft.ParseError.Error()})
	}
}

func WriteWellFormedTodoTable(bfts []*todo.WellFormedTodo, t *tablewriter.Table) {
	t.SetHeader([]string{"Location", "Age", "JIRA Ticket", "Detail"})
	for _, bft := range bfts {
		t.Append([]string{location(bft.Todo), fmt.Sprintf("%d days", math.Floor(time.Since(bft.Date).Hours())/24), bft.JIRATicketID, bft.Detail})
	}
}
