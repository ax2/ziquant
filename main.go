package main

import (
	"fmt"
	"os"

	"github.com/ax2/ziquant/cmd"
)

func main() {
	rootCmd := cmd.RootCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
