package findtheindexofthefirstoccurrenceinastring

func StrStr(haystack string, needle string) int {
	var haystackIndex int
	for i := 0; i < len(haystack); i++ {
		if len(needle) > len(haystack) {
			return -1
		}

		if haystack[i] != needle[0] {
			continue
		}

		haystackIndex = i
		haystackIndex++
		found := true
		for j := 1; j < len(needle); j++ {
			if !found {
				continue
			}

			if haystackIndex >= len(haystack) {
				return -1
			}

			if needle[j] != haystack[haystackIndex] {
				found = false
			}

			haystackIndex++
		}

		if found {
			return i
		}
	}
	return -1
}
