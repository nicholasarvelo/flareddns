package util

import "strings"

const minStringLength = 3
const maskingLengthOffset = 2
const maskingCharacter = "*"

// ObfuscateVariable masks the middle characters of a string, preserving
// the first and last characters. Short strings are returned as-is.
func ObfuscateVariable(variable string) string {
	if len(variable) < minStringLength {
		return variable
	}
	maskingLength := len(variable) - maskingLengthOffset
	return string(variable[0]) + strings.Repeat(
		maskingCharacter,
		maskingLength,
	) + string(variable[len(variable)-1])
}
