package utils

import "strings"

func ConcatenateStrings(args ...string) string {
	var builder strings.Builder
	for _, str := range args {
		builder.WriteString(str)
	}
	return builder.String()
}
