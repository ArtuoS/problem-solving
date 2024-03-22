package reversestring

func ReverseString(s []byte) {
	for i := len(s) - 1; i >= 0; i-- {
		s = append(s, s[i])
		s = append(s[:i], s[i+1:]...)
	}
}
