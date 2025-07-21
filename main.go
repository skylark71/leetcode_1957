package main

import "fmt"

func main() {
	fmt.Println(makeFancyString("a"))
}

func makeFancyString(s string) string {
	if len(s) == 0 {
		return ""
	}

	var result []byte
	count := 1
	result = append(result, s[0])

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}

		if count < 3 {
			result = append(result, s[i])
		}
	}

	return string(result)
}
