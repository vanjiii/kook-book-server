package main

import (
	"fmt"
	"os"

	"vanjiii/kook-book-server/internal/cmd"
)

func main() {
	err := cmd.NewRootCmd().Execute()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unsuccessful command, err: %+v", err)

		os.Exit(1)
	}
}
