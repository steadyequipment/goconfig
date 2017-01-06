package goutility

import (
	"io"
	"os"

	"path/filepath"
)

// CopyFile copy a file from src to dst
func CopyFile(src string, dst string) (err error) {
	// http://stackoverflow.com/a/21067803
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

// FileExists check if a file exists
func FileExists(path string) (bool, error) {
	// http://stackoverflow.com/a/10510783

	absPath, absPathError := filepath.Abs(path)
	if absPathError != nil {
		return false, absPathError
	}

	_, err := os.Stat(absPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
