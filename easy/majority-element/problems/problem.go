package problems

// 解法一：哈希表
// func MajorityElement(nums []int) int {
// 	hashTable := make(map[int]int)
// 	for _, num := range nums {
// 		if j, ok := hashTable[num]; ok {
// 			hashTable[num] = j + 1
// 		} else {
// 			hashTable[num] = 1
// 		}
// 	}
//
// 	var targetNum int
// 	maxCount := 0
// 	for k, v := range hashTable {
// 		if v > maxCount {
// 			targetNum = k
// 			maxCount = v
// 		}
// 	}
//
// 	return targetNum
// }

// 更简洁的写法
func MajorityElement1(nums []int) int {
	hashTable := make(map[int]int)
	half := len(nums) / 2
	for _, num := range nums {
		hashTable[num] += 1
		if hashTable[num] > half {
			return num
		}
	}

	return -1
}

// 解法二：排序

// 解法三：摩尔投票
func MajorityElement3(nums []int) int {
	var candidate int
	score := 0
	for _, num := range nums {
		if score == 0 {
			candidate = num
			score = 1
		} else if num == candidate {
			score += 1
		} else {
			score -= 1
		}
	}

	return candidate
}

var MajorityElement = MajorityElement3
