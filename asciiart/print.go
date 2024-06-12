package asciiart

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func AsciiPrint(text string, style string) {
	copy := strings.ReplaceAll(text, string(byte(10)), "\\n")
	res := strings.Split(copy, "\\n")
	flag := true
	for _, el := range res {
		if el != "" {
			flag = false
		}
	}
	if flag {
		res = res[1:]
	}
	for _, word := range res {
		if word != "" {
			PrintASCII(word, style)
		} else {
			fmt.Print("\n")
		}
	}
}

func PrintASCII(text string, style string) {
	var filePath string = style + ".txt"
	input, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	defer input.Close()
	lines := []string{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	res := [8]string{}
	string_num := 0
	for i := 0; i < len(text); i++ {
		string_num = (((int(text[i]) - 32) * 8) + int(text[i]) - 31)
		for j := 0; j < 8; j++ {
			if i != 0 {
				res[j] = res[j] + lines[string_num+j]
			} else {
				res[j] += lines[string_num+j]
			}
		}

	}
	byteSize := len([]byte(res[0]))
	width, _, err := getTerminalSize()
	// diff := (width - byteSize) / 2
	if err == nil {
		if byteSize > width {
			log.Printf("\nРазмер текса: %d\nРазмер терминала: %d\nНайдите текст поменьше, или экран пошире\n", byteSize, width)
			return
		}
	} else {
		log.Print(err)
	}
	for _, ascii := range res {
		fmt.Println(ascii)
	}

	// if diff >= 0 {
	// 	for _, ascii := range res {
	// 		for diff != 0 {
	// 			fmt.Print(" ")
	// 			diff--
	// 		}
	// 		fmt.Println(ascii)
	// 		diff = (width - byteSize) / 2
	// 	}
	// }
}

func getTerminalSize() (width, height int, err error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	size := strings.Split(strings.TrimSpace(string(out)), " ")
	if len(size) != 2 {
		return 0, 0, fmt.Errorf("неверный формат вывода stty size")
	}
	fmt.Sscanf(size[0], "%d", &height)
	fmt.Sscanf(size[1], "%d", &width)
	return width, height, nil
}
