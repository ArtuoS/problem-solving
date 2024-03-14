package single_number

func SingleNumber(nums []int) int {
	data := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if data[nums[i]] == 0 {
			data[nums[i]] = 1
			continue
		}

		currValue := data[nums[i]]
		data[nums[i]] = currValue + 1
	}

	for k, v := range data {
		if v == 1 {
			return k
		}
	}

	return 0
}
