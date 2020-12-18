package main

import (
	conf "github.com/neprune/todo/internal/config"
	"github.com/neprune/todo/internal/git"
	"github.com/neprune/todo/internal/harvest"
	rep "github.com/neprune/todo/internal/report"
	"github.com/neprune/todo/internal/todo"
	"github.com/neprune/todo/internal/web"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"os"
)

var (
	app          = kingpin.New("todo", "A command-line tool for monitoring TODOs.")
	config       = app.Flag("config", "The path to the config file.").Default("todo.yaml").Envar("TODO_CONFIG").Short('c').File()
	jiraUsername = app.Flag("jira-username", "The username to use to login to JIRA.").Envar("JIRA_USERNAME").String()
	jiraToken    = app.Flag("jira-token", "The token to use to login to JIRA.").Envar("JIRA_TOKEN").String()

	assert                    = app.Command("assert", "Make an assertion.")
	assertWellFormedTodosOnly = assert.Command("well-formed-todos-only", "Fails if there are TODOs that don't conform to the expected format.")
	assertNoOldTodos          = assert.Command("no-old-todos", "Fails if there are any TODOs exceeding the warning limit..")
	assertConsistentWithJIRA  = assert.Command("consistent-with-jira", "Fails if there are TODOs with non-existent or complete tickets.")

	report                  = app.Command("report", "Generate a report.")
	reportTerminal          = report.Command("terminal", "Output the report to terminal.")
	reportWeb               = report.Command("web", "Generate a static web page for the report.")
	reportWebOutputFilepath = reportWeb.Flag("web-output-filepath", "The filepath to write the webpage to.").Default("index.html").Short('o').String()
)

func main() {
	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))
	configData, err := ioutil.ReadAll(*config)
	kingpin.FatalIfError(err, "failed to load config file %s", (*config).Name())

	c, err := conf.LoadFromYAMLData(configData)
	kingpin.FatalIfError(err, "failed to parse config file %s", (*config).Name())

	ts, err := harvest.TodosFromGlobPatterns(c.SrcGlobPatterns)
	kingpin.FatalIfError(err, "failed to harvest todos")

	wfts, bfts := todo.ParseAllTodos(ts...)

	switch cmd {
	case reportTerminal.FullCommand():
		hygiene := rep.GenerateHygieneReport(wfts, bfts)
		hygiene.OutputToTerminal()
		age := rep.GenerateAgeReport(wfts, c.WarningAgeDays)
		age.OutputToTerminal()
		jira, err := rep.GenerateJIRAReport(wfts, c.JIRAAddress, *jiraUsername, *jiraToken)
		kingpin.FatalIfError(err, "failed to generate JIRA report")
		jira.OutputToTerminal()
		break

	case reportWeb.FullCommand():
		hygiene := rep.GenerateHygieneReport(wfts, bfts)
		age := rep.GenerateAgeReport(wfts, c.WarningAgeDays)
		jira, err := rep.GenerateJIRAReport(wfts, c.JIRAAddress, *jiraUsername, *jiraToken)
		kingpin.FatalIfError(err, "failed to generate JIRA report")
		dir, err := os.Getwd()
		kingpin.FatalIfError(err, "failed to get working dir")
		commit, err := git.GetCommit(dir)
		kingpin.FatalIfError(err, "failed to get current git repo commit")
		kingpin.FatalIfError(web.GenerateWebPage(*hygiene, *age, *jira, *c, *reportWebOutputFilepath, commit), "failed to generate web page")
		break

	case assertWellFormedTodosOnly.FullCommand():
		hygiene := rep.GenerateHygieneReport(wfts, bfts)
		hygiene.OutputToTerminal()
		if len(hygiene.BadlyFormedTodos) > 0 {
			kingpin.Fatalf("Assertion failed")
		}

	case assertNoOldTodos.FullCommand():
		age := rep.GenerateAgeReport(wfts, c.WarningAgeDays)
		age.OutputToTerminal()
		if len(age.TodosExceedingWarningAgeSortedByOldestFirst) > 0 {
			kingpin.Fatalf("Assertion failed")
		}

	case assertConsistentWithJIRA.FullCommand():
		jira, err := rep.GenerateJIRAReport(wfts, c.JIRAAddress, *jiraUsername, *jiraToken)
		kingpin.FatalIfError(err, "failed to generate JIRA report")
		jira.OutputToTerminal()
		if len(jira.TodosWithMissingIssues) > 0 || len(jira.TodosWithClosedIssues) > 0 || len(jira.TodosWithDoneIssues) > 0 {
			kingpin.Fatalf("Assertion failed")
		}
	}
}
