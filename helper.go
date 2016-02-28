package gamadues

import "strings"

func getArray(input string) []string {
	if input == "" {
		return nil
	}
	return strings.Split(input, ",")
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
