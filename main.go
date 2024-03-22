package main

import (
	"fmt"

	rangesumqueryimmutable "github.com/ArtuoS/problem-solving/range_sum_query_immutable"
)

func main() {
	numArray := rangesumqueryimmutable.Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(numArray.SumRange(0, 2))
	fmt.Println(numArray.SumRange(2, 5))
	fmt.Println(numArray.SumRange(0, 5))
}
