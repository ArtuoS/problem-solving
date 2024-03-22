package validanagram

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	storage := []map[string]int{}
	searchLetters := func(str string) map[string]int {
		letterOccurrance := map[string]int{}
		for _, v := range str {
			occurrences := letterOccurrance[string(v)] + 1
			letterOccurrance[string(v)] = occurrences
		}
		return letterOccurrance
	}

	storage = append(storage, searchLetters(s))
	storage = append(storage, searchLetters(t))

	for k, v := range storage[0] {
		if storage[1][k] != v {
			return false
		}
	}

	return true
}
