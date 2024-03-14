package main

import (
	"fmt"

	searchinsertposition "github.com/ArtuoS/problem-solving/search_insert_position"
)

func main() {
	fmt.Println(searchinsertposition.SearchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchinsertposition.SearchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchinsertposition.SearchInsert([]int{1, 3, 5, 6}, 7))
}
