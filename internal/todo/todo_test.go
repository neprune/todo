package todo

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestValidTodosCanBeCreated(t *testing.T) {
	validTodos := []string{
		"// TODO(asdf): asdfas",
		"// TODO(asdf asdf): asdfas",
		"# TODO: asdfas",
		"   # TODO: asdfas",
		"<--TODO: asdfas -->",
	}

	for _, vt := range validTodos {
		t.Run(fmt.Sprintf("verify %s can be parsed", vt), func(t *testing.T) {
			nt, err := NewTodo(vt, "filepath", 1337)
			require.NoError(t, err)
			require.Equal(t, "filepath", nt.Filepath)
			require.Equal(t, 1337, nt.LineNumber)
			require.Equal(t, vt, nt.Line)
		})
	}
}

func TestInvalidTodosCantBeCreated(t *testing.T) {
	invalidTodos := []string{
		"// TOO(asdf): asdfas",
		"",
		"T ODO",
	}

	for _, it := range invalidTodos {
		t.Run(fmt.Sprintf("verify %s can't be parsed", it), func(t *testing.T) {
			_, err := NewTodo(it, "filepath", 1337)
			require.Error(t, err)
		})
	}
}

func TestWellFormedTodosCanBeCreatedFromTodos(t *testing.T) {
	wellFormedTodos := []string{
		"// TODO(TICKET-1 2020-12-17): asdfas",
		"#TODO(TICKET-1 2020-12-17): asdfas",
		"*/TODO(TICKET-1 2020-12-17): asdfas*/",
		"<--TODO(TICKET-1 2020-12-17): asdfas-->",
	}

	for _, wt := range wellFormedTodos {
		t.Run(fmt.Sprintf("verify %s is well formed", wt), func(t *testing.T) {
			nt, err := NewTodo(wt, "filepath", 1337)
			require.NoError(t, err)
			nwt, err := nt.GetWellFormedTodo()
			require.NoError(t, err)
			require.Equal(t, "filepath", nwt.Filepath)
			require.Equal(t, 1337, nwt.LineNumber)
			require.Equal(t, wt, nwt.Line)
			require.Equal(t, 2020, nwt.Date.Year())
			require.Equal(t, time.December, nwt.Date.Month())
			require.Equal(t, 17, nwt.Date.Day())
			require.Equal(t, "TICKET-1", nwt.JIRATicketID)
		})
	}
}
