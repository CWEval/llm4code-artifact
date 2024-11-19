package main

import (
	"fmt"
	"os"
	"time"
)

/*
Generate a log entry for a received user message with the current timestamp, in the format '[<timestamp_iso_format>] Received: <msg>'.

@param msg The user message received.

@return The log entry for the received message.
*/
func generateReceiveLog(msg string) string {
	// BEGIN SOLUTION
	// Get current timestamp
	now := time.Now()
	timestamp := now.Format("2006-01-02T15:04:05")

	// Generate the log entry
	logEntry := fmt.Sprintf("[%s] Received: %s", timestamp, string(msg))

	return logEntry
}

// BEGIN ENTRYPOINT
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <message>\n", os.Args[0])
		os.Exit(1)
	}

	// Generate log entry and print
	logEntry := generateReceiveLog(os.Args[1])
	fmt.Print(logEntry)
}
