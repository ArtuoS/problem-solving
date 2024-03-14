package mergestringalternatively

func MergeAlternately(word1 string, word2 string) string {
	length := getBiggerLen(word1, word2)
	var concatedString string
	for i := 0; i < length; i++ {
		if i < len(word1) {
			concatedString += string(word1[i])
		}
		if i < len(word2) {
			concatedString += string(word2[i])
		}
	}
	return concatedString
}

func getBiggerLen(word1 string, word2 string) int {
	if len(word1) > len(word2) {
		return len(word1)
	}

	return len(word2)
}
