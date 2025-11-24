package problems

func MajorityElement(nums []int) int {
	hashTable := make(map[int]int)
	for _, num := range nums {
		if j, ok := hashTable[num]; ok {
			hashTable[num] = j + 1
		} else {
			hashTable[num] = 1
		}
	}

	var targetNum int
	maxCount := 0
	for k, v := range hashTable {
		if v > maxCount {
			targetNum = k
			maxCount = v
		}
	}

	return targetNum
}

// 哈希表
// 排序
