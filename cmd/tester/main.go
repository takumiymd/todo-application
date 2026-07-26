package main

import (
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

	takumios.Print("")
	takumios.Print("===> Tests has done")
}

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
}
