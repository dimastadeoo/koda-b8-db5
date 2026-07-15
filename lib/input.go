package lib

import (
	"bufio"
	"fmt"
	"os"
)

func Input() string {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return scanner.Text() // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		return ""
	}
	return ""
}

func PressEnter(mss string) {
	fmt.Print(mss)
	fmt.Scanln()
}
