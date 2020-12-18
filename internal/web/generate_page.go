package web

import (
	"fmt"
	conf "github.com/neprune/todo/internal/config"
	"github.com/neprune/todo/internal/report"
	"html/template"
	"os"
)

type data struct {
	Age     report.Age
	Hygiene report.Hygiene
	JIRA    report.JIRA
	Config  conf.Config
	Commit  string
}

func GenerateWebPage(hygiene report.Hygiene, age report.Age, jira report.JIRA, config conf.Config, outPath string, commit string) error {
	tmpl, err := template.New("report").Parse(reportTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse report template: %w", err)
	}
	f, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close()

	err = tmpl.Execute(f, data{
		Age:     age,
		Hygiene: hygiene,
		JIRA:    jira,
		Config:  config,
		Commit:  commit,
	})
	if err != nil {
		return fmt.Errorf("failed to render web page template: %w", err)
	}
	return nil
}
