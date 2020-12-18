package web

import (
	"fmt"
	conf "github.com/neprune/todo/internal/config"
	"github.com/neprune/todo/internal/report"
	"html/template"
	"os"
)

type data struct {
	age     report.Age
	hygiene report.Hygiene
	jira    report.JIRA
	config  conf.Config
}

func GenerateWebPage(hygiene report.Hygiene, age report.Age, jira report.JIRA, config conf.Config, outPath string) error {
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
		age:     age,
		hygiene: hygiene,
		jira:    jira,
		config:  config,
	})
	if err != nil {
		return fmt.Errorf("failed to render web page template: %w", err)
	}
	return nil
}
