//go:build !windows
// +build !windows

package acme

import (
	"os"
)

// CheckFile checks file permissions and content size.
func CheckFile(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil && os.IsNotExist(err) {
		nf, err := os.Create(name)
		if err != nil {
			return false, err
		}
		defer nf.Close()
		return false, nil
	}
	if err != nil {
		return false, err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return false, err
	}

	return fi.Size() > 0, nil
}
