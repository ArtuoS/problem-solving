package plusone

import (
	"math/big"
	"strconv"
)

func PlusOne(digits []int) []int {
	var stringNumber string
	for _, v := range digits {
		stringNumber += strconv.Itoa(v)
	}

	number := new(big.Int)
	number.SetString(stringNumber, 10)
	number.Add(number, big.NewInt(1))
	arr := []int{}
	for _, v := range number.String() {
		arr = append(arr, int(v-'0'))
	}

	return arr
}
