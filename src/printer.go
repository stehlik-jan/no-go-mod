package printer

import "fmt"

func printer() int {
	bytes_printed, _ := fmt.Print("Simple message")
    return bytes_printed
}