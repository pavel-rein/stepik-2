package main

import "fmt"
import "os"
import "bufio"
import "unicode"
import "strings"

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text2 := []rune(text)
	var right int
	if unicode.IsUpper(text2[0]) {
		right++
	}

	if  strings.Contains(string(text2[len(text2)-2]), ".") {
		right++
	}
	if right == 2 {
		fmt.Print("Right", right)
	} else {
		fmt.Print("Wrong", right)
	}
}
