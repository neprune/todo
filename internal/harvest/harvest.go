package harvest

import (
	"bufio"
	"github.com/neprune/todo/internal/todo"
	"os"
	"path/filepath"
	"strings"
)

// TodosFromGlobPatterns harvests all todos using the given glob patterns.
func TodosFromGlobPatterns(globPatterns []string) ([]todo.Todo, error) {
	var todos []todo.Todo

	for _, p := range globPatterns {
		ts, err := TodosFromGlobPattern(p)
		if err != nil {
			return todos, err
		}
		todos = append(todos, ts...)
	}
	return todos, nil
}

// TodosFromGlobPattern harvests all todos in files found under the given glob pattern.
func TodosFromGlobPattern(globPattern string) ([]todo.Todo, error) {
	var todos []todo.Todo

	ms, err := filepath.Glob(globPattern)
	if err != nil {
		return todos, err
	}

	for _, m := range ms {
		ts, err := TodosFromFile(m)
		if err != nil {
			return todos, err
		}
		todos = append(todos, ts...)
	}
	return todos, nil
}

// TodosFromFile harvests all todos in a given file.
func TodosFromFile(filename string) ([]todo.Todo, error) {
	var todos []todo.Todo

	file, err := os.Open(filename)
	if err != nil {
		return todos, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		if strings.Contains(scanner.Text(), "TODO") {
			t, err := todo.NewTodo(scanner.Text(), filename, lineNumber)
			if err != nil {
				return todos, err
			}
			todos = append(todos, *t)
		}
	}
	return todos, nil
}
