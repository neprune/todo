package report

import (
	"fmt"
	"github.com/neprune/todo/internal/table"
	"github.com/neprune/todo/internal/todo"
	"github.com/olekukonko/tablewriter"
	"os"
)

type Hygiene struct {
	NumberOfTodos            int
	NumberOfBadlyFormedTodos int
	PercentageWellFormed     float32
	BadlyFormedTodos         []*todo.BadlyFormedTodo
}

func GenerateHygieneReport(wfts []*todo.WellFormedTodo, bfts []*todo.BadlyFormedTodo) *Hygiene {
	return &Hygiene{
		NumberOfTodos:            len(wfts) + len(bfts),
		NumberOfBadlyFormedTodos: len(bfts),
		PercentageWellFormed:     (float32(len(wfts)) / float32(len(bfts)+len(wfts))) * 100,
		BadlyFormedTodos:         bfts,
	}
}

func (h *Hygiene) OutputToTerminal() {
	t := tablewriter.NewWriter(os.Stdout)
	table.WriteBadlyFormedTodoTable(h.BadlyFormedTodos, t)
	fmt.Println()
	fmt.Printf("There are %d TODOs in total.\n", h.NumberOfTodos)
	fmt.Printf("%.2f%% TODOs are well formed.\n", h.PercentageWellFormed)
	fmt.Printf("The %d badly formed TODOs are:\n", h.NumberOfBadlyFormedTodos)
	fmt.Println()
	t.Render()
	fmt.Println()
}
