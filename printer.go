package printer

import (
	"fmt"

	nestedDep "github.com/stehlik-jan/go-nested-dep"
)

func Hello() int {
	printerMessage := "Simple message"
	bytes_printed, _ := fmt.Print(printerMessage)

	nestedDep.Diff(printerMessage)

    return bytes_printed
}