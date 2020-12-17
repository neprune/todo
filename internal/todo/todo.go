package todo

import "time"

// Todo represents a well-formed todo.
type Todo struct {
	// JIRATicketID is the ID of the JIRA ticket associated with the todo.
	JIRATicketID string
	// Date is the date at which the todo was first registered.
	Date time.Time
	// Filepath is the path of the file the todo lives in.
	Filepath string
	// LineNumber is the line number the todo can be found at.
	LineNumber int
}
