package todo

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var (
	wellFormedTodoPattern = regexp.MustCompile(`TODO\((?P<ticket>[\w-]+)\s(?P<date>[\w-]+)\):(.*)`)
)

// Todo represents any line of code containing "TODO".
type Todo struct {
	// Filepath is the path of the file the todo lives in.
	Filepath string
	// LineNumber is the line number the todo can be found at.
	LineNumber int
	// Line is the contents of the line of code.
	Line string
}

// WellFormedTodo represents a well-formed todo.
type WellFormedTodo struct {
	*Todo

	// JIRATicketID is the ID of the JIRA ticket associated with the todo.
	JIRATicketID string
	// Date is the date at which the todo was first registered.
	Date time.Time
	// Detail is a summary of the TODO.
	Detail string
}

func NewTodo(line string, filepath string, number int) (*Todo, error) {
	if !strings.Contains(line, "TODO") {
		return nil, fmt.Errorf("invalid construction: line does not contain TODO: %s", line)
	}
	return &Todo{
		Line: line,
		Filepath: filepath,
		LineNumber: number,
	}, nil
}

func (t *Todo) GetWellFormedTodo() (*WellFormedTodo, error) {
	ss := wellFormedTodoPattern.FindAllStringSubmatch(t.Line, -1)
	if len(ss) != 1 || len(ss[0]) != 4 {
		return nil, fmt.Errorf("failed to parse well-formed todo from %s", t.Line)
	}

	ticket := ss[0][1]
	date := ss[0][2]
	detail := ss[0][3]

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, fmt.Errorf("failed to parse date %s into yyyy-mm-dd", date)
	}

	return &WellFormedTodo{
		t,
		ticket,
		parsedDate,
		detail,
	}, nil
}