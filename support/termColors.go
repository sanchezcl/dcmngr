package support

import "fmt"

const ColorReset = "\033[0m"
const ColorRed = "\033[31m"
const ColorGreen = "\033[32m"
const ColorYellow = "\033[33m"
const ColorBlue = "\033[34m"
const ColorPurple = "\033[35m"
const ColorCyan = "\033[36m"
const ColorGray = "\033[37m"
const ColorWhite = "\033[97m"

func PrintError(msg string) {
	fmt.Printf("%sError%s: %s\n", ColorRed, ColorReset, msg)
}

func PrintSword() {
	fmt.Printf("     %s[%s\n","\033[38;5;245m", "\033[0m")
	fmt.Printf("%s@%sxxxx%s[%s{{%s:::::::::::::::::::::::>\n", "\033[97m", "\033[38;5;137m", "\033[38;5;247m", "\033[38;5;220m", "\033[0m")
	fmt.Printf("     %s[%s\n","\033[38;5;245m", "\033[0m")
}