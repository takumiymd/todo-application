package main

import (
	"os"

	"takumiymd/todo-application/internal"
	"takumiymd/todo-application/takumifmt"
	"takumiymd/todo-application/takumios"
	"takumiymd/todo-application/takumitest"
)

func main() {
	takumios.Print("=====================")
	takumios.Print("RUNNING TEST HARNESS")
	takumios.Print("=====================")
	takumios.Print("")

	testFormatPackage()
	testInternalPackage()

	takumios.Print("")
	takumios.Print("===> Tests have done")
}

// testFormatPackage is unit tests evalutating single functions
func testFormatPackage() {
	takumios.Print("--- Testing takumifmt ---")

	takumitest.AssertString(
		"PadRight: Add correct number of spaces",
		"Hello     ",
		takumifmt.PadRight("Hello", 10),
	)

	takumitest.AssertString(
		"PadRight: Pad empty string correctly",
		"     ",
		takumifmt.PadRight("", 5),
	)

	takumitest.AssertString(
		"PadRight: Handle zero padding wihtout modifying string",
		"hello",
		takumifmt.PadRight("hello", 0),
	)

	takumitest.AssertString(
		"PadRight: Ignore padding if string is longer",
		"Takumi",
		takumifmt.PadRight("Takumi", 3),
	)

	takumitest.AssertInt(
		"StringToInt: Parse single digit",
		5,
		takumifmt.StringToInt("5"),
	)

	takumitest.AssertInt(
		"StringToInt: Parse multi digits number",
		1234,
		takumifmt.StringToInt("1234"),
	)

	takumitest.AssertInt(
		"StringToInt: Parse zero",
		0,
		takumifmt.StringToInt("0"),
	)

	takumitest.AssertString(
		"IntToString: Convert single digit",
		"7",
		takumifmt.IntToString(7),
	)

	takumitest.AssertString(
		"IntToString: Convert multi digits number",
		"9012",
		takumifmt.IntToString(9012),
	)

	takumitest.AssertString(
		"IntToString: Convert zero",
		"0",
		takumifmt.IntToString(0),
	)

	takumitest.AssertString(
		"BoolToString: Convert true",
		"true",
		takumifmt.BoolToString(true),
	)

	takumitest.AssertString(
		"BoolToString: Convert false",
		"false",
		takumifmt.BoolToString(false),
	)

	takumitest.AssertString(
		"StringToBool: Parse 'true'",
		"true",
		takumifmt.BoolToString(takumifmt.StringToBool("true")),
	)

	takumitest.AssertString(
		"StringToBool: Parse 'false'",
		"false",
		takumifmt.BoolToString(takumifmt.StringToBool("false")),
	)

	takumitest.AssertString(
		"BuildCSVLine: Create CSV record format",
		"1,Buy milk,false",
		takumifmt.BuildCSVLine(1, "Buy milk", false),
	)

	takumitest.AssertInt(
		"Split: Number of fields after splitting CSV",
		3,
		len(takumifmt.Split("1,Buy milk,false", ',')),
	)
}

// testInternalPackage is integration tests evaluating entire workflow
func testInternalPackage() {
	takumios.Print("")
	takumios.Print("--- Testing internal ---")

	var list internal.List

	takumitest.AssertInt(
		"List: New list is empty",
		0,
		len(list),
	)

	list.Add("First task")
	takumitest.AssertInt(
		"List: Length is 1 after Add",
		1,
		len(list),
	)
	takumitest.AssertInt(
		"List: First task has ID 1",
		1,
		list[0].ID,
	)
	takumitest.AssertString(
		"List: First task has correct description",
		"First task",
		list[0].Task,
	)
	takumitest.AssertString(
		"List: First task is not completed initially",
		"false",
		takumifmt.BoolToString(list[0].Completed),
	)

	list.Add("Second task")
	takumitest.AssertInt(
		"List: Length is 2 after second Add",
		2,
		len(list),
	)
	takumitest.AssertInt(
		"List: Second task has ID 2",
		2,
		list[1].ID,
	)

	err := list.Complete(2)
	if err != nil {
		takumios.Print("[FAIL] List: Completing task returned error: " + err.Error())
	} else {
		takumitest.AssertString(
			"List: Second task is completed",
			"true",
			takumifmt.BoolToString(list[1].Completed),
		)
	}

	tempDB := "test_todo.db"
	defer os.Remove(tempDB)

	err = list.Save(tempDB)
	if err != nil {
		takumios.Print("[FAIL] List: Saving list returned error: " + err.Error())
	}

	var loadedList internal.List
	err = loadedList.Load(tempDB)
	if err != nil {
		takumios.Print("[FAIL] List: Loading list returned error: " + err.Error())
	}

	takumitest.AssertInt(
		"List: Loaded list length match",
		2,
		len(loadedList),
	)
	takumitest.AssertInt(
		"List: Loaded first task ID match",
		1,
		loadedList[0].ID,
	)
	takumitest.AssertString(
		"List: Loaded first task name match",
		"First task",
		loadedList[0].Task,
	)
	takumitest.AssertString(
		"List: Loaded first task status match",
		"false",
		takumifmt.BoolToString(loadedList[0].Completed),
	)
	takumitest.AssertInt(
		"List: Loaded second task ID match",
		2,
		loadedList[1].ID,
	)
	takumitest.AssertString(
		"List: Loaded second task name match",
		"Second task",
		loadedList[1].Task,
	)
	takumitest.AssertString(
		"List: Loaded second task status match",
		"true",
		takumifmt.BoolToString(loadedList[1].Completed),
	)

	err = list.Edit(1, "Updated first task")
	if err != nil {
		takumios.Print("[FAIL] List: Editing task returned error: " + err.Error())
	} else {
		takumitest.AssertString(
			"List: First task string was successfully updated",
			"Updated first task",
			list[0].Task,
		)
	}

	err = list.Delete(1)
	if err != nil {
		takumios.Print("[FAIL] List: Deleting task returned error: " + err.Error())
	} else {
		takumitest.AssertInt(
			"List: Length is 1 after Delete",
			1,
			len(list),
		)
		takumitest.AssertInt(
			"List: Remaining task is ID 2",
			2,
			list[0].ID,
		)
	}

	list.Add("Third task")
	list.Sweep()

	takumitest.AssertInt(
		"List: Length is 1 after Sweep",
		1,
		len(list),
	)

	takumitest.AssertInt(
		"List: Remaining task is ID 3",
		3,
		list[0].ID,
	)

	takumitest.AssertString(
		"List: Remaining task description match",
		"Third task",
		list[0].Task,
	)

	takumitest.AssertString(
		"List: Remaining task is pending",
		"false",
		takumifmt.BoolToString(list[0].Completed),
	)
}
