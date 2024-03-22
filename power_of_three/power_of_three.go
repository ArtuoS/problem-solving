package powerofthree

func IsPowerOfThree(n int) bool {
	if n < 0 || n == 0 {
		return false
	}

	for {
		if n%3 == 0 {
			n = n / 3
		} else {
			break
		}
	}

	return n == 1
}
