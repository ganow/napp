package main

import (
	"fmt"
	"os"
)

const RED = "\x1b[31m"
const GREEN = "\x1b[32m"
const YELLOW = "\x1b[33m"
const NORMAL = "\x1b[39m"

func logError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s", RED)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintf(os.Stderr, "%s", NORMAL)
}

func logWarning(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s", YELLOW)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintf(os.Stderr, "%s", NORMAL)
}

func logStatus(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s", GREEN)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintf(os.Stderr, "%s", NORMAL)
}
