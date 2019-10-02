package helferlein

import "strings"

//StringInSlice checks if a string is in a slice of strings
func StringInSlice(s string, list []string) bool {
	for _, b := range list {
		if b == s {
			return true
		}
	}
	return false
}

// Checks if a string contains a substring form a slice of strings
func Contains(input string, words []string) bool {
	for _, word := range words {
		if strings.Index(input, word) > -1 {
			return true
		}
	}
	return false
}
