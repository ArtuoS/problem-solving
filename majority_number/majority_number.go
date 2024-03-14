package majoritynumber

func MajorityElement(nums []int) int {
	data := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if data[nums[i]] == 0 {
			data[nums[i]] = 1
			continue
		}

		currValue := data[nums[i]]
		data[nums[i]] = currValue + 1
	}

	higherValue, higherNumber := 0, 0
	for k, v := range data {
		if v > higherValue {
			higherValue = v
			higherNumber = k
		}
	}

	return higherNumber
}
