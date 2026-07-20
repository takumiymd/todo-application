package takumifmt

// IntToString converts a base 10 integer into string format
func IntToString(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""

	// Use modulo logic for getting digits one by one
	for n > 0 {
		digit := n % 10

		// '0' is the rune 48 in ASCII
		result = string(rune('0'+digit)) + result

		n = n / 10
	}

	return result
}

// StringtoInt converts numbers in string into  base 10 integers
func StringToInt(s string) int {
	result := 0

	for _, char := range s {
		// '0' is 48 in ASCII
		// char - 48 = digit
		digit := int(char - '0')

		result = (result * 10) + digit

	}
	return result
}

// BoolToString converts a boolean into string format
func BoolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// BuildCSVLine creates a CSV format
func BuildCSVLine(id int, task string, completed bool) string {
	return IntToString(id) + "," + task + "," + BoolToString(completed)
}

// Join takes a slice of strings and concatenates them with a separater in b/w
func Join(elements []string, separator string) string {
	// base case
	if len(elements) == 0 {
		return ""
	}

	result := elements[0]

	for i := 1; i < len(elements); i++ {
		result = result + separator + elements[i]
	}

	return result
}

// StringToBool converts a text string into a boolean
func StringToBool(s string) bool {
	if s == "true" {
		return true
	}
	return false
}

// Split slices a string into an array of strings based on a separator rune
func Split(s string, separator rune) []string {
	var result []string
	var currentWord string

	for _, char := range s {
		if char == separator {
			result = append(result, currentWord)
			currentWord = ""
		} else {
			currentWord = currentWord + string(char)
		}
	}

	result = append(result, currentWord)

	return result
}

// Contains checks if the substr exists inside the string s
func Contains(s string, substr string) bool {
	// base case
	if len(substr) == 0 {
		return true
	}
	if len(substr) > len(s) {
		return false
	}

	// sliding window algorithm
	for i := 0; i <= len(s)-len(substr); i++ {
		match := true

		for j := 0; j < len(substr); j++ {
			if s[i+j] != substr[j] {
				match = false
				break
			}
		}

		if match {
			return true
		}
	}

	return false
}
