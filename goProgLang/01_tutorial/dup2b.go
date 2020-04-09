package main

// prints the count and text of lines that appear more than 1x in input.
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	locations := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, locations, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, locations, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Printf("%s\n", locations[line])
		}
	}
}

func countLines(
	f *os.File, counts map[string]int, locations map[string]string,
	filename string) {

	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		locations[input.Text()] += (" " + filename)
	}
	// NOTE: ignoring potential errors from input.Err()
}
