package numberof1bites

import (
	"math/bits"
)

func HammingWeight(num uint32) int {
	return bits.OnesCount32(num)
}
