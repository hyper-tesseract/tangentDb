package common

import (
	// "os"
)

func IsValidPath(path string) bool {
	if len(path) < 1 {
		return false
	}

	// _, err := os.Stat(path)
	// if err != nil {
	// 	return false
	// }

	// Add any more validations needed

	return true
}