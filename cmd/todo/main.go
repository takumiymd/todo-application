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
		showHelp()
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

	case "search":
		if len(args) < 3 {
			takumios.Print("Error: Keyword to search is required")
			return
		}
		keyword := args[2]
		foundCount := 0

		takumios.Print("Search result for '" + keyword + "':")

		for _, task := range myTodos {
			if takumifmt.Contains(task.Task, keyword) {
				status := "pending"
				if task.Completed {
					status = "done"
				}

				output := takumifmt.IntToString(task.ID) + ":" + task.Task + " " + status
				takumios.Print(output)
				foundCount++
			}
		}
		if foundCount == 0 {
			takumios.Print("No tasks found containing such word.")
		}

	case "list":
		if len(myTodos) == 0 {
			takumios.Print("No tasks found.")
			return
		}

		// task 4 length
		maxTaskLen := 4
		for _, task := range myTodos {
			runes := []rune(task.Task)
			if len(runes) > maxTaskLen {
				maxTaskLen = len(runes)
			}
		}

		// for the cols
		taskDashes := ""
		for i := 0; i < maxTaskLen; i++ {
			taskDashes += "-"
		}

		// corners and intersections of table
		border := "+------+-" + taskDashes + "-+----------+"

		// top border and header
		takumios.Print(border)
		header := "| ID   | " + takumifmt.PadRight("Task", maxTaskLen) + " | Status   |"
		takumios.Print(header)
		takumios.Print(border)

		// print tasks with PadRight func
		for _, task := range myTodos {
			// 8 characters length so we dont have to call PadRight just for the status
			status := "pending "
			if task.Completed {
				status = "done    "
			}

			idStr := takumifmt.IntToString(task.ID)

			paddedID := takumifmt.PadRight(idStr, 4)
			paddedTask := takumifmt.PadRight(task.Task, maxTaskLen)

			row := "| " + paddedID + " | " + paddedTask + " | " + status + " |"
			takumios.Print(row)

		}

		takumios.Print(border)

	case "help":
		showHelp()

	case "edit":
		if len(args) < 4 {
			takumios.Print("Error: Task id and new task description required")
			return
		}

		id := takumifmt.StringToInt(args[2])
		newTask := args[3]

		err := myTodos.Edit(id, newTask)
		if err != nil {
			takumios.Print("Error: " + err.Error())
			return
		}

		myTodos.Save(dbFile)
		takumios.Print("Task " + args[2] + " updated to: " + newTask)

	case "sweep":
		myTodos.Sweep()
		myTodos.Save(dbFile)
		takumios.Print("Sweep completed succeessfully.")

	default:
		takumios.Print("Invalid command: " + command)
		takumios.Print("")
		showHelp()
	}
}

// showHelp prints the app's command line usage guide and available commands
func showHelp() {
	takumios.Print("Usage: todo <command> [arguments]")
	takumios.Print("")
	takumios.Print("Commands:")

	takumios.Print("  " + takumifmt.PadRight("add", 12) + "Add a new task")
	takumios.Print("  " + takumifmt.PadRight("list", 12) + "List all tasks")
	takumios.Print("  " + takumifmt.PadRight("complete", 12) + "Mark a task as completed")
	takumios.Print("  " + takumifmt.PadRight("delete", 12) + "Delete a task by id")
	takumios.Print("  " + takumifmt.PadRight("search", 12) + "Search for tasks by keyword")
	takumios.Print("  " + takumifmt.PadRight("help", 12) + "Show this help menu")
	takumios.Print("  " + takumifmt.PadRight("edit", 12) + "Edit an existing task by id")
	takumios.Print("  " + takumifmt.PadRight("sweep", 12) + "Sweep completed tasks")
}
