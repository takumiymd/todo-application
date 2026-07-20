# Todo Application API Reference

This document provides a summary of the core packages, types, and functions used in the Todo application.

## CLI Commands (`cmd/todo/main.go`)

The application is interacted with via the command line and supports the following operations:
*   **`add`**: Requires a task string as an argument. Adds the new task to the list and saves it to `todo.db`.
*   **`list`**: Displays all tasks, outputting the task ID, the task name, and its status as either "pending" or "done".
*   **`complete`**: Requires a task ID as an argument. Marks the specified task as complete and saves the updated state.
*   **`delete`**: Requires a task ID as an argument. Removes the specified task from the list and saves the updated state.
*   **`search`**: Requires a keyword string as an argument. Displays all tasks that contain the specified keyword.

---

## Package `internal` (`internal/todo.go`)

Manages the core business logic, state, and data structures for the tasks.

### Types
*   **`Todo`**: A struct representing a single task containing an `ID` (int), `Task` (string), and `Completed` (bool).
*   **`List`**: A custom type defined as a slice of `Todo` structs (`[]Todo`). 
*   **`TaskError`**: A custom error struct that holds an error message string.

### Methods
*   **`NewTaskError(text string) error`**: Creates a new `TaskError`.
*   **`(*TaskError) Error() string`**: Returns the string representation of the error message.
*   **`(*List) Add(task string)`**: Creates a new `Todo` task and appends it to the list.
*   **`(*List) Complete(id int) error`**: Finds a task by its ID and marks its `Completed` status as true.
*   **`(*List) Delete(id int) error`**: Finds a task by its ID and removes it from the list.
*   **`(*List) Save(filename string) error`**: Converts the list of todos into a CSV formatted string and writes it to disk.
*   **`(*List) Load(filename string) error`**: Reads a file, parses the text line by line, and populates the list with `Todo` structs.

---

## Package `takumifmt` (`takumifmt/format.go`)

Handles custom string formatting, type conversion, and manipulation without relying on the Go standard library.

### Functions
*   **`IntToString(n int) string`**: Converts a base 10 integer into a string format.
*   **`StringToInt(s string) int`**: Converts a string representation of a number into a base 10 integer.
*   **`BoolToString(b bool) string`**: Converts a boolean value into "true" or "false" strings.
*   **`StringToBool(s string) bool`**: Converts the text string "true" to a boolean `true`, and returns `false` otherwise.
*   **`BuildCSVLine(id int, task string, completed bool) string`**: Constructs a comma separated CSV line format.
*   **`Join(elements []string, separator string) string`**: Concatenates a slice of strings with a specified separator.
*   **`Split(s string, separator rune) []string`**: Slices a string into an array of strings divided by a separator rune.
*   **`Contains(s string, substr string) bool`**: Checks if a substring exists within a given string, returning true if a match is found.

---

## Package `takumios` (`takumios/io.go`)

Provides custom wrappers for OS interactions, specifically file I/O and terminal output.

### Functions
*   **`WriteFile(filename string, data string) error`**: Writes a string of data directly to disk with `0644` /`+rw`permissions.
*   **`ReadFile(filename string) (string, error)`**: Reads a file's contents from the disk and returns the data as a string.
*   **`Print(s string)`**: Writes a string directly to the terminal via standard output.
