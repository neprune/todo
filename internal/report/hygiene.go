package report

import (
	"fmt"
	"github.com/neprune/todo/internal/table"
	"github.com/neprune/todo/internal/todo"
	"github.com/olekukonko/tablewriter"
	"os"
)

// Hygiene summarises the distribution of well-formed vs badly formed TODOs.
type Hygiene struct {
	// NumberOfTodos is the total number of TODOs.
	NumberOfTodos int
	// PercentageWellFormed is the percentage of TODOs that are well formed.
	PercentageWellFormed float32
	// BadlyFormedTodos are the badly formed todos.
	BadlyFormedTodos []todo.BadlyFormedTodo
}

// GenerateHygeineReport generates a Hygiene report.
func GenerateHygieneReport(wfts []todo.WellFormedTodo, bfts []todo.BadlyFormedTodo) *Hygiene {
	return &Hygiene{
		NumberOfTodos:        len(wfts) + len(bfts),
		PercentageWellFormed: (float32(len(wfts)) / float32(len(bfts)+len(wfts))) * 100,
		BadlyFormedTodos:     bfts,
	}
}

// OutputToTerminal prints the Hygiene report to terminal.
func (h *Hygiene) OutputToTerminal() {
	t := tablewriter.NewWriter(os.Stdout)
	table.WriteBadlyFormedTodoTable(h.BadlyFormedTodos, t)
	fmt.Println()
	fmt.Println("Hygiene Report:")
	fmt.Println("===============")
	fmt.Println()
	fmt.Printf("There are %d TODOs in total.\n", h.NumberOfTodos)
	fmt.Printf("%.2f%% TODOs are well formed.\n", h.PercentageWellFormed)
	fmt.Printf("The %d badly formed TODOs are:\n", len(h.BadlyFormedTodos))
	fmt.Println()
	t.Render()
	fmt.Println()
}
