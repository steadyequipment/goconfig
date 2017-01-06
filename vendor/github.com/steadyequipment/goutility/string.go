package goutility

import (
	"strings"

	"math/rand"
)

// StringAppendWithJoin append two strings with a join string accounting for empty strings
func StringAppendWithJoin(left string, join string, right string) (result string) {
	if len(left) > 0 {
		if len(right) > 0 {
			result = left + join + right
		} else {
			result = left
		}
	} else {
		result = right
	}

	return
}

// PadStringToLength pad string with string until it is equal to or greater than length param and return the result
func PadStringToLength(original string, padWith string, length int) string {
	result := original
	for {
		if len(result) >= length {
			break
		}
		result = result + padWith
	}

	return result
}

// PathStrippedOfQuery return the path minus any query
func PathStrippedOfQuery(path string) string {
	splits := strings.Split(path, "{")
	if len(splits) > 0 {
		return splits[0]
	}

	return path
}

// FirstCharacter return the first character in a string
func FirstCharacter(in string) string {
	return string([]rune(in)[0])
}

// StringOfStringRepeated return a string composed of component repeated count times
func StringOfStringRepeated(component string, count int) string {
	components := []string{}
	for times := 1; times <= count; times++ {
		components = append(components, component)
	}

	return strings.Join(components[:], "")
}

// http://stackoverflow.com/a/31832326
var randomStringLetterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandomString create a random string of specified length
func RandomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = randomStringLetterRunes[rand.Intn(len(randomStringLetterRunes))]
	}
	return string(b)
}
