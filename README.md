# TODO

A CLI tool that lets you keep on top of your TODOs.

Set up a `todo.yaml` config:
```
# Glob patterns that will match with files containing your source code.
src_glob_patterns:
- src/*

# The number of days after which a warning is generated for a TODO.
warning_age_days: 30
```

Update your `TODO`s to be in the form `TODO(<JIRA Ticket ID> <DATE YYYY-MM-DD>): <DETAIL>`.

Run `todo report`:
```
> todo report

Hygiene Report:
===============

There are 9 TODOs in total.
44.44% TODOs are well formed.
The 5 badly formed TODOs are:

+--------------------------------+--------------------+--------------------------------+
|            LOCATION            |        LINE        |          PARSE ERROR           |
+--------------------------------+--------------------+--------------------------------+
| badly_formed_todos/todos.txt:1 | // TODO(one): one  | failed to parse as TODO(<JIRA  |
|                                |                    | TICKET ID> <YYYY-MM-DD>):      |
|                                |                    | <TICKET DETAIL>                |
| badly_formed_todos/todos.txt:2 | // TODO(two ): two | failed to parse as TODO(<JIRA  |
|                                |                    | TICKET ID> <YYYY-MM-DD>):      |
|                                |                    | <TICKET DETAIL>                |
| badly_formed_todos/todos.txt:3 | # TODO: three      | failed to parse as TODO(<JIRA  |
|                                |                    | TICKET ID> <YYYY-MM-DD>):      |
|                                |                    | <TICKET DETAIL>                |
| badly_formed_todos/todos.txt:4 |    # TODO: four    | failed to parse as TODO(<JIRA  |
|                                |                    | TICKET ID> <YYYY-MM-DD>):      |
|                                |                    | <TICKET DETAIL>                |
| badly_formed_todos/todos.txt:5 | <--TODO: five -->  | failed to parse as TODO(<JIRA  |
|                                |                    | TICKET ID> <YYYY-MM-DD>):      |
|                                |                    | <TICKET DETAIL>                |
+--------------------------------+--------------------+--------------------------------+


Age Report:
===============

There are 3 TODOs older than 2 days.

+-------------------------------+--------+-------------+----------+
|           LOCATION            |  AGE   | JIRA TICKET |  DETAIL  |
+-------------------------------+--------+-------------+----------+
| well_formed_todos/todos.txt:3 | 2 days | TICKET-3    | three */ |
| well_formed_todos/todos.txt:2 | 3 days | TICKET-2    | two      |
| well_formed_todos/todos.txt:1 | 4 days | TICKET-1    | one      |
+-------------------------------+--------+-------------+----------+

```

You can also run `todo assert only-well-formed-todos` to assert that there are no badly formed TODOs in the repo - useful for local / CI checks.

There's also `todo assert no-old-todos` which could be useful if you want to trigger some event when there is at least one very old TODO still in the repo.

What's next:
* JIRA integration - see which TODOs are in progress in JIRA, which ones have tickets that are clsoed etc.
* Generate a static webpage for the report- with links to the LOCs in GH and links to the JIRA tickets
* Turn into a plug and play GH action

## Full Usage

```
> todo --help
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
