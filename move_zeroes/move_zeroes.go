package movezeroes

import "fmt"

func MoveZeroes(nums []int) {
	zeroes := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroes++
			nums = append(nums[:i], nums[i+1:]...)

			fmt.Println(nums)
		}
	}

	fmt.Println(nums)
}
