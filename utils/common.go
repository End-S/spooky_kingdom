package utils

// Contains function checks if an array of strings contains the given string
func Contains(s []string, str string) bool {
	for _, ele := range s {
		if ele == str {
			return true
		}
	}
	return false
}

// Includes function checks if an array of strings contains one or more of the given strings
func Includes(s []string, inc []string) bool {
	for _, ele := range inc {
		if Contains(s, ele) {
			return true
		}
	}
	return false
}
