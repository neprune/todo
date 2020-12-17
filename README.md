# TODO

```
+ go run ../cmd/main.go --help
usage: todo [<flags>] <command> [<args> ...]

A command-line tool for monitoring TODOs.

Flags:
      --help              Show context-sensitive help (also try --help-long and --help-man).
  -c, --config=todo.yaml  The path to the config file.

Commands:
  help [<command>...]
    Show help.

  assert well-formed-todos-only
    Fails if there are TODOs that don't conform to the expected format.

  report
    Generate a report.
```
