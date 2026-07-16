package takumios

import (
	"os"
)

// WriteFile takes a filename and a string of data, and write it to disk
func WriteFile(filename string, data string) error {
	// 0644 is +rw
	err := os.WriteFile(filename, []byte(data), 0o644)
	if err != nil {
		return err
	}
	return nil
}

// ReadFile takes a filename, reads its contents, and returns it as a string
func ReadFile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// Print writes a string directory to the terminal
func Print(s string) {
	os.Stdout.Write([]byte(s + "\n"))
}
