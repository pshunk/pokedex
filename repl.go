package main

import (
	"strings"
)

func cleanInput(text string) []string {
	cleanedText := strings.Split(strings.ToLower(text), " ")
	return cleanedText
}
