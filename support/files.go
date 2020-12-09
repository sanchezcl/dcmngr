package support

import (
	"os"
)

func FileExist(filePathName string) bool {
	if _, err := os.Stat(filePathName); err == nil {
		return true
	}
	return false
}