package commands

import (
	"fmt"
	"os"
)

func Execute() {
	println("Execute")
	if err := Router.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err) // TODO: use logger
		if err != nil {
			panic(err) // TODO: use logger
		}
		os.Exit(1)
	}
}
