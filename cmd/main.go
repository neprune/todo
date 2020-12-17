package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	app    = kingpin.New("todo", "A command-line tool for monitoring TODOs.")
	config = app.Flag("config", "The path to the config file.").Default("todo.yaml").Envar("TODO_CONFIG").Short('c').File()

	check = app.Command("check", "Check set-up.")
	checkConfigFormat = check.Command("config-format", "Check the config format is well-formed.")

	report = app.Command("report", "Generate a report.")
	reportTerminal = report.Command("terminal", "Returns a report in terminal output.")
	reportWeb = report.Command("webpage", "Generates a static web page for the report.")
	webOutput = reportWeb.Flag("out", "The path to write the web report to.").Default("index.html").Envar("TODO_WEB_OUT").Short('o').File()

	assert = app.Command("assert", "Make an assertion.")
	assertNoBadTodos = assert.Command("no-bad-todos", "Fails if there are TODOs that can't be parsed.")
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case checkConfigFormat.FullCommand():
		break

	case reportTerminal.FullCommand():
		break

	case reportWeb.FullCommand():
		break

	case assertNoBadTodos.FullCommand():
		break
	}
}

