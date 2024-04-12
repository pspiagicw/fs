/*
Provides file-system helper functions

- All exposed functions, don't return anything.
- Either function completes execution without error
- Or function terminates execution with error-code 1.
*/
package fs

import (
	"errors"
	"io/fs"
	"os"

	"github.com/pspiagicw/goreland"
)

// Ensure given path exists on the file-system.
//   - It creates directory if it doesn't exist.
//   - It expects absolute path.
//   - It can't resolve home directories.
//   - It creates all parent directories if needed.
func EnsurePathExists(location string) {
	if !dirExists(location) {
		err := os.MkdirAll(location, 0755)
		if err != nil {
			goreland.LogFatal("Error creating directory: %s, %v", location, err)
		}
	}
}

// Read file contents from the given ReadFile
//   - Expects absolute path.
//   - Can't resolve home directories.
//   - Returns file contents as byte-array
//   - Terminates execution if file can't be read
func ReadFile(file string) []byte {
	contents, err := os.ReadFile(file)
	if err != nil {
		goreland.LogFatal("Error reading file: %s, %v", file, err)
	}

	return contents
}

func dirExists(dir string) bool {
	_, err := os.Stat(dir)
	if errors.Is(err, fs.ErrNotExist) {
		return false
	} else if err != nil {
		goreland.LogFatal("Error stating directory: %v", err)
	}
	return true
}
