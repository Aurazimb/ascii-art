package main

import (
	"fmt"
	"os"

	"main.go/asciiart"
)

func main() {
	style := ""
	if len(os.Args) == 2 {
		style = "standard"
	} else if len(os.Args) == 3 {
		style = os.Args[2]
	} else {
		return
	}
	text := os.Args[1]

	for _, ch := range text {
		if ch > 126 && ch < 32 {
			fmt.Println("Введите в аргумент только ASCII символы")
			return
		}
	}

	if CheckASCII(style) {
		asciiart.AsciiPrint(text, style)
	}
}

func CheckASCII(style string) bool {
	flag := false
	if style == "standard" || style == "shadow" || style == "thinkertoy" {
		flag = true
	}
	return flag
}
