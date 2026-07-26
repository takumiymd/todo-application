package takumitest

import (
	"takumiymd/todo-application/takumios"

	"takumiymd/todo-application/takumifmt"
)

// AssertString checks if two stirngs match
func AssertString(testName string, expected string, current string) {
	if expected == current {
		takumios.Print("[PASS] " + testName)
		return
	}

	takumios.Print("[FAIL] '" + testName + "'")
	takumios.Print("  Expected:  '" + expected + "'")
	takumios.Print("  Current:   '" + current + "'")
}

// AssertInt checks if two integers match
func AssertInt(testName string, expected, int, current int) {
	if expected == current {
		takumios.Print("[PASS] " + testName)
		return
	}

	takumios.Print("[FAIL] " + testName)
	takumios.Print("  Expected: " + takumifmt.IntToString(expected))
	takumios.Print("  Current:  " + takumifmt.IntToString(current))
}
