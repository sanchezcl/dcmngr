package support

import "fmt"

const ColorReset = "\033[0m"
const ColorRed = "\033[31m"
const ColorGreen = "\033[32m"
const ColorYellow = "\033[33m"
const ColorBlue = "\033[34m"
const ColorPurple = "\033[35m"
const ColorCyan = "\033[36m"
const ColorWhite = "\033[37m"

func PrintError(msg string) {
	fmt.Printf("%sError%s: %s\n", ColorRed, ColorReset, msg)
}
