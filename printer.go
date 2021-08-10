package public_printer

import (
	"fmt"
)

func Hello() int {
	printerMessage := "Simple message from public repo"
	bytes_printed, _ := fmt.Print(printerMessage)

    return bytes_printed
}