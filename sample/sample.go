package main

import (
	"fmt"
	"os"
	"strings"

	rearrange "github.com/tanaikech/go-rearrange"
)

// main : main function
func main() {
	data := []string{
		"sample1",
		"sample2",
		"sample3",
		"sample4",
		"sample5",
		"sample6",
		"sample7",
		"sample8",
		"sample9",
		"sample10",
	}
	result, history, err := rearrange.Do(data, 3, false, false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(strings.Join(result, "\n"))
	fmt.Printf("\nHistory of selected values: %+v\n", history)
}
