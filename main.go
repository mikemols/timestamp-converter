package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/atotto/clipboard"
)

func convertTimestamp(input string) (string, error) {
	// Parses the input as integer
	timestamp, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return "", fmt.Errorf("Invalid number: %s", input)
	}

	// Auto-detect: seconds vs milliseconds based on digit count
	var t time.Time
	switch len(input) {
	case 10:
		// Seconds
		t = time.Unix(timestamp, 0)
	case 13:
		// Milliseconds (like your example: 1752574424823)
		t = time.Unix(0, timestamp*int64(time.Millisecond))
	default:
		return "", fmt.Errorf("timestamp must be 10 (seconds) or 13 (milliseconds) digits")
	}

	// Format the timestamp in local time
	return t.UTC().Format("Monday, January 2, 2006 15:04:05 MST"), nil
}

func main() {
	var (
		timestamp string
		copyToCb  bool
	)

	for _, arg := range os.Args[1:] {
		if arg == "-copy" || arg == "-c" || arg == "--copy" {
			copyToCb = true
		} else {
			timestamp = arg
		}
	}

	// Check for required args
	if timestamp == "" {
		fmt.Println(`Usage:
  timestamp-converter <timestamp> [options]

Arguments:
  <timestamp>         Unix timestamp (10 or 13 digits)

Options:
  -c, -copy, --copy   Copy the converted time to clipboard
`)
		os.Exit(1)
	}

	input := flag.Arg(0)

	result, err := convertTimestamp(input)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)

	if copyToCb {
		err := clipboard.WriteAll(result)
		if err != nil {
			fmt.Printf("⚠️  Failed to copy to clipboard: %v\n", err)
		} else {
			fmt.Println("📋 Copied to clipboard.")
		}
	}
}
