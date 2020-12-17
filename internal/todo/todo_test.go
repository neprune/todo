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

func TestWellFormedTodosCanBeParsedFromTodos(t *testing.T) {
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
			wft, bft := nt.Parse()
			require.Nil(t, bft)
			require.Equal(t, "filepath", wft.Filepath)
			require.Equal(t, 1337, wft.LineNumber)
			require.Equal(t, wt, wft.Line)
			require.Equal(t, 2020, wft.Date.Year())
			require.Equal(t, time.December, wft.Date.Month())
			require.Equal(t, 17, wft.Date.Day())
			require.Equal(t, "TICKET-1", wft.JIRATicketID)
		})
	}
}
