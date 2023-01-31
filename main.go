package main

import "fmt"
import "bufio"
import "os"

func main() {
    text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    fmt.Print(text)
}