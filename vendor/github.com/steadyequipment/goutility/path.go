package goutility

import (
	"os"
	"path/filepath"
)

// CreatePathIfDoesNotExist ensures the path exists, even if it has to create it. Returns an error if there was a problem checking that the path exists or creating the path. Returns true if the path had to be created, false if the path already existed.
func CreatePathIfDoesNotExist(path string) (bool, error) {

	absPath, absPathError := filepath.Abs(path)
	if absPathError != nil {
		return false, absPathError
	}

	exists, findError := FileExists(absPath)
	if findError != nil {
		return false, findError

	} else if exists == false {

		makeDirError := os.MkdirAll(absPath, os.ModePerm)
		if makeDirError != nil {
			return false, makeDirError
		}
	}

	return !exists, nil
}
