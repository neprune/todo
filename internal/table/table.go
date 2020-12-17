package table

import (
	"fmt"
	"github.com/neprune/todo/internal/todo"
	"github.com/olekukonko/tablewriter"
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
