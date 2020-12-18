package jira

import "fmt"

// GenerateJIRATicketURL creates a link to the given ticket.
func GenerateJIRATicketURL(jiraURL string, ticketID string) string {
	return fmt.Sprintf("%sbrowse/%s", jiraURL, ticketID)
}
