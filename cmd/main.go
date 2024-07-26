package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/your-username/title/internal/titlecase"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Please provide a string to convert to title case.")
		os.Exit(1)
	}

	input := strings.Join(args, " ")
	titleCase := titlecase.ToTitleCase(input)
	fmt.Println(titleCase)
}