# RFC: CLI Todo Application

## 1. Overview

A command-line Todo application built without the Go standard library. It relies entirely on custom internal packages (`takumios` for OS/file operations and `takumifmt` for string formatting and terminal output).

## 2. Command Line Interface (CLI) UX

The user will interact with the application by passing arguments to the compiled binary. `takumios` will parse the raw command-line arguments.

- `todo add "Buy groceries"` — Adds a new task.
- `todo list` — Displays all active and completed tasks.
- `todo complete 1` — Marks the task with ID 1 as completed.
- `todo delete 1` — Removes the task with ID 1 from the list.
- todo edit 1 "Buy organic groceries" — Updates the task with ID 1 to the new string.

## 3. Data Schema

The core application logic will hold tasks in memory using a slice of structs:

```go
type Todo struct {
    ID        int
    Task      string
    Completed bool
}
```

## 4. Storage Format

Since standard encoding libraries (like JSON) are banned, the state will be persisted to a local text file (`todo.db`) using a custom pipe-delimited format. `takumios` will handle reading and writing this file.

**Format:** `{ID}|{Task}|{Completed}`

Example `todo.db` file:

```
1|Set up Go module|true
2|Write Design Doc|false
3|Build CLI router|false
```

