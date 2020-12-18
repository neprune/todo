package todo

import (
	"errors"
	"fmt"
	"github.com/neprune/todo/internal/github"
	"github.com/neprune/todo/internal/jira"
	"regexp"
	"strings"
	"time"
)

const (
	// HumanReadableFormat describes the required format for well-formed comments for humans.
	HumanReadableFormat = "TODO(<JIRA TICKET ID> <YYYY-MM-DD>): <TICKET DETAIL>"
)

var (
	wellFormedTodoPattern = regexp.MustCompile(`TODO\((?P<ticket>[\w-]+)\s(?P<date>[\w-]+)\):\s(.*)`)
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

func (t *Todo) String() string {
	return fmt.Sprintf("%s:%d %s", t.Filepath, t.LineNumber, t.Line)
}

// WellFormedTodo represents a well-formed todo.
type WellFormedTodo struct {
	Todo

	// JIRATicketID is the ID of the JIRA ticket associated with the todo.
	JIRATicketID string
	// Date is the date at which the todo was first registered.
	Date time.Time
	// Detail is a summary of the TODO.
	Detail string
}

// Age returns the number of days since the TODO was created.
func (w *WellFormedTodo) Age() int {
	return int(time.Since(w.Date).Hours() / 24)
}

// GithubLocURL returns the URL of the associated LOC.
func (w *WellFormedTodo) GithubLocURL(githubRepoUrl string, commitOrBranch string) string {
	return github.GenerateLOCGithubURL(githubRepoUrl, commitOrBranch, w.Filepath, w.LineNumber)
}

// JIRATicketURL returns a URL to the associated JIRA ticket.
func (w *WellFormedTodo) JIRATicketURL(jiraURL string) string {
	return jira.GenerateJIRATicketURL(jiraURL, w.JIRATicketID)
}

func (w *WellFormedTodo) String() string {
	return fmt.Sprintf("%s %s %s %s:%d", w.Date, w.JIRATicketID, w.Detail, w.Line, w.LineNumber)
}

// BadlyFormedTodo represents a badly-formed todo.
type BadlyFormedTodo struct {
	Todo
	// ParseError describes the reason the todo was not well-formed.
	ParseError error
}

func (b *BadlyFormedTodo) String() string {
	return fmt.Sprintf("%s <parse error: %s>", b.Todo, b.ParseError)
}

// GithubLocURL returns the URL of the associated LOC.
func (b *BadlyFormedTodo) GithubLocURL(githubRepoUrl string, commitOrBranch string) string {
	return github.GenerateLOCGithubURL(githubRepoUrl, commitOrBranch, b.Filepath, b.LineNumber)
}

// NewTodo creates a new todo, checking that the line is valid.
func NewTodo(line string, filepath string, number int) (*Todo, error) {
	if !strings.Contains(line, "TODO") {
		return nil, fmt.Errorf("invalid construction: line does not contain TODO: %s", line)
	}
	return &Todo{
		Line:       line,
		Filepath:   filepath,
		LineNumber: number,
	}, nil
}

// Parse either returns a WellFormedTodo or a BadlyFormedTodo.
func (t *Todo) Parse() (*WellFormedTodo, *BadlyFormedTodo) {
	ss := wellFormedTodoPattern.FindAllStringSubmatch(t.Line, -1)
	if len(ss) != 1 || len(ss[0]) != 4 {
		return nil, &BadlyFormedTodo{*t, fmt.Errorf("failed to parse as %s", HumanReadableFormat)}
	}

	ticket := ss[0][1]
	date := ss[0][2]
	detail := ss[0][3]

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, &BadlyFormedTodo{*t, errors.New("failed to parse date as yyyy-mm-dd")}
	}

	return &WellFormedTodo{
		*t,
		ticket,
		parsedDate,
		detail,
	}, nil
}

// ParseAllTodos parses all given todos and returns the BadlyFormedTodos and the WellFormedTodos.
func ParseAllTodos(todos ...Todo) ([]WellFormedTodo, []BadlyFormedTodo) {
	var bfts []BadlyFormedTodo
	var wfts []WellFormedTodo
	for _, t := range todos {
		wft, bft := t.Parse()
		if bft != nil {
			bfts = append(bfts, *bft)
		} else {
			wfts = append(wfts, *wft)
		}
	}
	return wfts, bfts
}
