package main

import (
	conf "github.com/neprune/todo/internal/config"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"os"
)

var (
	app    = kingpin.New("todo", "A command-line tool for monitoring TODOs.")
	config = app.Flag("config", "The path to the config file.").Default("todo.yaml").Envar("TODO_CONFIG").Short('c').File()

	assert = app.Command("assert", "Make an assertion.")
	assertNoBadTodos = assert.Command("no-bad-todos", "Fails if there are TODOs that can't be parsed.")

	report = app.Command("report", "Generate a report.")
	reportTerminal = report.Command("terminal", "Returns a report in terminal output.")
	reportWeb = report.Command("webpage", "Generates a static web page for the report.")
	webOutput = reportWeb.Flag("out", "The path to write the web report to.").Default("index.html").Envar("TODO_WEB_OUT").Short('o').File()
)

func main() {
	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))
	configData, err := ioutil.ReadAll(*config)
	kingpin.FatalIfError(err, "failed to load config file %s", (*config).Name())
	c, err := conf.LoadFromYAMLData(configData)
	kingpin.FatalIfError(err, "failed to parse config file %s", (*config).Name())

	switch cmd {
	case reportTerminal.FullCommand():
		break

	case reportWeb.FullCommand():
		break

	case assertNoBadTodos.FullCommand():
		break
	}
}

