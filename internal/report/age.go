package report

import (
	"fmt"
	"github.com/neprune/todo/internal/table"
	"github.com/neprune/todo/internal/todo"
	"github.com/olekukonko/tablewriter"
	"os"
	"sort"
	"time"
)

// Age reports on old TODOs.
type Age struct {
	// NumberOfTodosExceedingWarningAge is the number of TODOs that are older than the configured number of days to
	// warn after.
	NumberOfTodosExceedingWarningAge int
	// TodosExceedingWarningAgeSortedByOldestFirst are the todos in descending order of age.
	TodosExceedingWarningAgeSortedByOldestFirst []*todo.WellFormedTodo
	// WarningDays is the number of days after which a TODO will trigger a warning.
	WarningDays int
}

// GenerateAgeReport generates a Age report.
func GenerateAgeReport(wfts []*todo.WellFormedTodo, warningDays int) *Age {
	var old []*todo.WellFormedTodo
	for _, wft := range wfts {
		if time.Since(wft.Date).Hours() > float64(warningDays*24) {
			old = append(old, wft)
		}
	}
	sort.Slice(old, func(i, j int) bool {
		return old[i].Date.After(old[j].Date)
	})
	return &Age{
		NumberOfTodosExceedingWarningAge:            len(old),
		TodosExceedingWarningAgeSortedByOldestFirst: old,
		WarningDays: warningDays,
	}
}

// OutputToTerminal prints the Age report to terminal.
func (a *Age) OutputToTerminal() {
	t := tablewriter.NewWriter(os.Stdout)
	table.WriteWellFormedTodoTable(a.TodosExceedingWarningAgeSortedByOldestFirst, t)
	fmt.Println()
	fmt.Println("Age Report:")
	fmt.Println("===============")
	fmt.Println()
	fmt.Printf("There are %d TODOs older than %d days.\n", a.NumberOfTodosExceedingWarningAge, a.WarningDays)
	fmt.Println()
	t.Render()
	fmt.Println()
}
