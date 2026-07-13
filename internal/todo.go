package internal

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
