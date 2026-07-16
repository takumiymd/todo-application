package internal

import (
	"takumiymd/todo-application/takumifmt"
	"takumiymd/todo-application/takumios"
)

// Todo represents a single task in our application
type Todo struct {
	ID        int
	Task      string
	Completed bool
}

// List represents our collection of tasks
// we make it a custom type so we can attach methods to it
type List []Todo

// TaskError represents an error that occurs during thask operations
type TaskError struct {
	message string
}

// Error returns the string of message of the error
func (e *TaskError) Error() string {
	return e.message
}

// NewTaskError creates a new TaskError and return its memory address
func NewTaskError(text string) error {
	return &TaskError{message: text}
}

// Add creates a new task and appends it to the list
// Time complexity:O(1)
// Space complexity:O(1)
func (l *List) Add(task string) {
	todo := Todo{
		// Generate a simple ID based on the current length of the list
		ID:        len(*l) + 1,
		Task:      task,
		Completed: false,
	}
	*l = append(*l, todo)
}

// Complete finds a task by its ID and marks it as finished
func (l *List) Complete(id int) error {
	ls := *l

	for i := range ls {
		if ls[i].ID == id {
			ls[i].Completed = true
			return nil
		}
	}
	return NewTaskError("Task not found")
}

// Delete finds a task by its ID and removes it from the list
func (l *List) Delete(id int) error {
	ls := *l

	for i := range ls {
		if ls[i].ID == id {
			*l = append(ls[:i], ls[i+1:]...)
			return nil
		}
	}
	return NewTaskError("task not found")
}

// Save converts the list of todos to a CSV string and writes it to disk
func (l *List) Save(filename string) error {
	var lines []string

	for _, todo := range *l {
		line := takumifmt.BuildCSVLine(todo.ID, todo.Task, todo.Completed)
		lines = append(lines, line)

	}

	data := takumifmt.Join(lines, "\n")

	err := takumios.WriteFile(filename, data)
	if err != nil {
		return NewTaskError("failed to save files: " + err.Error())
	}

	return nil
}

// Load reads a CSV file and parse the text line by line and convert those strings back into Todo structs
func (l *List) Load(filename string) error {
	data, err := takumios.ReadFile(filename)
	if err != nil {
		return nil
	}

	*l = []Todo{}

	lines := takumifmt.Split(data, '\n')

	for _, line := range lines {
		if line == "" {
			continue
		}
		fields := takumifmt.Split(line, ',')

		if len(fields) == 3 {
			id := takumifmt.StringToInt(fields[0])
			taskName := fields[1]
			isCompleted := takumifmt.StringToBool(fields[2])

			todo := Todo{
				ID:        id,
				Task:      taskName,
				Completed: isCompleted,
			}

			*l = append(*l, todo)
		}
	}

	return nil
}
