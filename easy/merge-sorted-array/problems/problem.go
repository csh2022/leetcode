package problems

func MergeTwoSortedArray(nums1 []int, m int, nums2 []int, n int) {
	for i, j, tail := m-1, n-1, m+n-1; i >= 0 || j >= 0; tail-- {
		if i < 0 {
			nums1[tail] = nums2[j]
			j--
			continue
		}

		if j < 0 {
			nums1[tail] = nums1[i]
			i--
			continue
		}

		if nums1[i] < nums2[j] {
			nums1[tail] = nums2[j]
			j--
		} else {
			nums1[tail] = nums1[i]
			i--
		}
	}
}

// type TwoSumFunc func([]int, int) []int

// var TwoSum TwoSumFunc = TwoSum1

// var TwoSum func([]int, int) []int = TwoSum1

// var TwoSum = TwoSum1

// var TwoSum = TwoSum2

// // brute force, time complexity: O(n^2), space complexity: O(1)
// func TwoSum1(nums []int, target int) []int {
// 	for i, _ := range nums {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				if i > j {
// 					return []int{j, i}
// 				}
// 				return []int{i, j}
// 			}
// 		}
// 	}
//
// 	return nil
// }
//
// // hash table, time complexity: O(n), space complexity: O(n)
// func TwoSum2(nums []int, target int) []int {
// 	hashTable := make(map[int]int)
// 	for i, x := range nums {
// 		if j, ok := hashTable[target-x]; ok {
// 			if i > j {
// 				return []int{j, i}
// 			}
// 			return []int{i, j}
// 		}
// 		hashTable[x] = i
// 	}
// 	return nil
// }
//
// // type TwoSumFunc func([]int, int) []int
//
// // var TwoSum TwoSumFunc = TwoSum1
//
// // var TwoSum func([]int, int) []int = TwoSum1
//
// // var TwoSum = TwoSum1
//
// var TwoSum = TwoSum2
