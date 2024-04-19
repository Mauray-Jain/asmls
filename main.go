package main

import (
	"log"
	"os"

	"github.com/Mauray-Jain/asmls/lsp"
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

	logger := log.New(logOut, "[asmls] ", log.Ldate | log.Ltime | log.Lshortfile)

	server := lsp.NewServer(os.Stdin, os.Stdout, logger)

	for {
		shouldShutdown := server.HandleMsg()
		if shouldShutdown {
			break
		}
	}
}
