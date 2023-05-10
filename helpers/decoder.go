package helpers

import "strings"

func Decode(message string) []string {
	return strings.Split(message, " ")
}