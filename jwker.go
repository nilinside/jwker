package jwker

import (
	"fmt"
	"os"
)

func throwParseError(message string) {
	fmt.Fprintf(os.Stderr, "Could not parse key: %v\n", message)
	os.Exit(1)
}

func stopOnParseError(err error) {
	if err != nil {
		throwParseError(err.Error())
	}
}
