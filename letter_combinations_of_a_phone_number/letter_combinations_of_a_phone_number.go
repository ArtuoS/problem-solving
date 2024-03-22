package lettercombinationsofaphonenumber

import (
	"fmt"
)

func LetterCombinations(digits string) []string {
	keyboard := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	arr := []string{}
	if len(digits) == 1 {
		for _, v := range keyboard[rune(digits[0])] {
			arr = append(arr, string(v))
		}

		fmt.Println(arr)
		return arr
	}

	return arr
}
