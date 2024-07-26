package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/thefury/title/internal/titlecase"

	"golang.org/x/term"
)

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1:]
		fmt.Println(titlecase.ToTitleCase(strings.Join(args, " ")))
	} else if !term.IsTerminal(int(os.Stdin.Fd())) {
		// Handle input from standard input
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(titlecase.ToTitleCase(scanner.Text()))
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			os.Exit(1)
		}
	} else {
		fmt.Fprintln(os.Stderr, "No input provided. Please provide input as arguments or through stdin.")
		os.Exit(1)
	}
}
