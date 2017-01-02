// Dup2 prints the count and text of lines that appear more than once
// in the input. It rests from stdin or from a list of named files.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileCounts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fileCounts["stdin"] = make(map[string]int)
		countLines(os.Stdin, fileCounts["stdin"])
	} else {
		for _, arg := range files {
			fileCounts[arg] = make(map[string]int)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, fileCounts[arg])
			f.Close()
		}
	}
	for file, counts := range fileCounts {
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", file, n, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	//NOTE: ignoring potential errors from input.Err()
}
