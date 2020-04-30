package main

import "strings"

// basename removes directory components and file extensions from a filepath
// Ex. a -> a; a.go -> a; a/b/c.go -> c; a/b/c.go.gz -> c
func basename1(s string) string {
	// discard anything up to, and including, the final '/'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// preserve everything befor the last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// Simpler version using strings library
func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if no "/"
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
