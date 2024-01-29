package mystrings

func Reverse(s string) string {
	result := ""

	for _, c := range s {
		result = string(c) + result
	}

	return result
}
