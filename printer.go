package printer

import "fmt"

func Hello() int {
	bytes_printed, _ := fmt.Print("Simple message")
    return bytes_printed
}