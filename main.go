package main

import (
	"os"
)

// Treesitter parser maybe?

func main() {
	os.Stderr.Write([]byte("Starting AsmLs\n"))
	logOut, err := os.Create("/tmp/asmls_log")
	if err != nil {
		os.Stderr.WriteString("Unable to open log file")
		os.Exit(1)
	}
	defer logOut.Close()

}
