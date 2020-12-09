package support

import (
	"fmt"
	"os"
)

func FileExist(filePathName string) bool {
	fmt.Println(filePathName)
	if _, err := os.Stat(filePathName); err == nil {
		return true
	}
	return false
}