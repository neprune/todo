package main

import (
	"fmt"
	conf "github.com/neprune/todo/internal/config"
	"github.com/neprune/todo/internal/harvest"
	table "github.com/neprune/todo/internal/table"
	"github.com/neprune/todo/internal/todo"
	"github.com/olekukonko/tablewriter"
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

	_, bfts := todo.ParseAllTodos(ts...)

	switch cmd {
	case report.FullCommand():
		break

	case assertWellFormedTodosOnly.FullCommand():
		if len(bfts) == 0 {
			fmt.Println("No badly formed todos found!")
		}
		t := tablewriter.NewWriter(os.Stdout)
		table.WriteBadlyFormedTodoTable(bfts, t)
		fmt.Println()
		fmt.Printf("%d badly formed todos found:\n", len(bfts))
		t.Render()
		fmt.Println()
		kingpin.Fatalf("Assertion failed")
	}
}
