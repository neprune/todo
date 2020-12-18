package main

import (
	conf "github.com/neprune/todo/internal/config"
	"github.com/neprune/todo/internal/harvest"
	rep "github.com/neprune/todo/internal/report"
	"github.com/neprune/todo/internal/todo"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"os"
)

var (
	app    = kingpin.New("todo", "A command-line tool for monitoring TODOs.")
	config = app.Flag("config", "The path to the config file.").Default("todo.yaml").Envar("TODO_CONFIG").Short('c').File()

	assert                    = app.Command("assert", "Make an assertion.")
	assertWellFormedTodosOnly = assert.Command("well-formed-todos-only", "Fails if there are TODOs that don't conform to the expected format.")

	report = app.Command("report", "Generate a report.")
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
	case report.FullCommand():
		hygiene := rep.GenerateHygieneReport(wfts, bfts)
		hygiene.OutputToTerminal()
		age := rep.GenerateAgeReport(wfts, c.WarningAgeDays)
		age.OutputToTerminal()
		break

	case assertWellFormedTodosOnly.FullCommand():
		hygiene := rep.GenerateHygieneReport(wfts, bfts)
		hygiene.OutputToTerminal()
		if hygiene.NumberOfBadlyFormedTodos > 0 {
			kingpin.Fatalf("Assertion failed")
		}
	}
}
