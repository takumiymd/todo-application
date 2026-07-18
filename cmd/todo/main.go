package main

import (
	"os"

	"takumiymd/todo-application/internal"
	"takumiymd/todo-application/takumifmt"
	"takumiymd/todo-application/takumios"
)

func main() {
	args := os.Args

	// base case
	if len(args) < 2 {
		takumios.Print("Must follow: todo [add|list|complete|delete] [arguments]")
		return
	}

	command := args[1]
	dbFile := "todo.db"

	var myTodos internal.List

	_ = myTodos.Load(dbFile)

	// Route the correct cmd to the handler functions
	switch command {
	case "add":
		if len(args) < 3 {
			takumios.Print("Error: Please provide the task to add")
			return
		}
		task := args[2]

		myTodos.Add(task)
		myTodos.Save(dbFile)

		takumios.Print("Task added: " + task)

	case "list":
		if len(myTodos) == 0 {
			takumios.Print("No Task found:")
			return
		}

		takumios.Print("Todo lists:")

		for _, task := range myTodos {
			status := "pending"
			if task.Completed {
				status = "done"
			}

			output := takumifmt.IntToString(task.ID) + ":" + task.Task + " " + status
			takumios.Print(output)
		}

	case "complete":
		if len(args) < 3 {
			takumios.Print("Error: Task ID required to complete")
			return
		}

		id := takumifmt.StringToInt(args[2])
		err := myTodos.Complete(id)
		if err != nil {
			takumios.Print("Error: " + err.Error())
			return
		}

		myTodos.Save(dbFile)
		takumios.Print("Task " + args[2] + " changed to Complete")

	case "delete":
		if len(args) < 3 {
			takumios.Print("Error: Task ID required to delete")
			return
		}

		id := takumifmt.StringToInt(args[2])
		err := myTodos.Delete(id)
		if err != nil {
			takumios.Print("Error: " + err.Error())
			return
		}

		myTodos.Save(dbFile)
		takumios.Print("Task " + args[2] + " deleted")

	default:
		takumios.Print("Invalid command: " + command)
		takumios.Print("Available command following: [add|list|complete|delete]")
	}
}
