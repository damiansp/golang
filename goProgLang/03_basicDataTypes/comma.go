package main

import "fmt"

// commaRecursive inserts commas in long numbers
func commaRecursive(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaRecursive(s[:n-3]) + "," + s[n-3:]
}

// comma inserts commas in long numbers
func comma(s string) string {
	out := ""
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		out = string(s[i]) + out
		if (n-i)%3 == 0 && i != 0 {
			out = "," + out
		}
	}
	return out
}

func main() {
	fmt.Printf("%s\n", commaRecursive("123"))
	fmt.Printf("%s\n", comma("123"))
	fmt.Printf("%s\n", commaRecursive("12345"))
	fmt.Printf("%s\n", comma("12345"))
	fmt.Printf("%s\n", commaRecursive("123456789"))
	fmt.Printf("%s\n", comma("123456789"))
}
