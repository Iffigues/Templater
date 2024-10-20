package main

import "strings"

type arrayFlag []string

// Implement the Set method for arrayFlag
func (a *arrayFlag) Set(value string) error {
	*a = append(*a, value)
	return nil
}

// Implement the String method for arrayFlag
func (a *arrayFlag) String() string {
	return strings.Join(*a, ",")
}
