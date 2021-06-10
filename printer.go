package printer // import "github.com/stehlik-jan/go-printer"

import "fmt"

func Hello() int {
	bytes_printed, _ := fmt.Print("Simple message")
    return bytes_printed
}