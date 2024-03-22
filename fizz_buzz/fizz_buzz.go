package fizzbuzz

import "strconv"

func FizzBuzz(n int) []string {
	if n == 1 {
		return []string{"1"}
	}
	answer := []string{}
	for i := 1; i < n+1; i++ {
		if i%3 == 0 && i%5 == 0 {
			answer = append(answer, "FizzBuzz")
			continue
		}

		if i%3 == 0 {
			answer = append(answer, "Fizz")
			continue
		}

		if i%5 == 0 {
			answer = append(answer, "Buzz")
			continue
		}

		answer = append(answer, strconv.Itoa(i))
	}

	return answer
}
