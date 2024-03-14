package longestcommonprefix

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	var commonPrefix string
	var prefix byte
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			if len(strs[j]) > i {
				if j == 0 {
					prefix = strs[j][i]
				}

				if prefix != strs[j][i] {
					return commonPrefix
				}

				if j == len(strs)-1 {
					commonPrefix += string(strs[j][i])
				}
			} else {
				return commonPrefix
			}
		}
	}

	return commonPrefix
}
