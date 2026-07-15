package lib

import (
	"bufio"
	"os"
)
func Input() string{
		scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return scanner.Text() // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		return ""
	}
	return ""
}