# TODO

```
A command-line tool for monitoring TODOs.

Flags:
      --help              Show context-sensitive help (also try --help-long and --help-man).
  -c, --config=todo.yaml  The path to the config file.

Commands:
  help [<command>...]
    Show help.

  check config-format
    Check the config format is well-formed.

  report terminal
    Returns a report in terminal output.

  report webpage [<flags>]
    Generates a static web page for the report.

  assert no-bad-todos
    Fails if there are TODOs that can't be parsed.
```
