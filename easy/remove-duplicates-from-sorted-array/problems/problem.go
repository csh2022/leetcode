package problems

func RemoveDuplicatesFromSortedArray1(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	i, j := 0, 0
	for i < size && j < size {
		if i == j {
			j++
			continue
		}

		if nums[i] == nums[j] {
			j++
			continue
		}

		if nums[i] != nums[j] {
			if i+1 == j {
				i++
				j++
				continue
			}
			nums[i+1] = nums[j]
			i++
			j++
			continue
		}
	}

	return i + 1
}

func RemoveDuplicatesFromSortedArray2(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	i, j := 0, 0
	for i < size && j < size {
		if i == j {
			j++
			continue
		}

		if i+1 == j {
			if nums[i] == nums[j] {
				j++
				continue
			} else {
				i++
				j++
				continue
			}
		}

		if nums[i] == nums[j] {
			j++
			continue
		}

		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
			j++
			continue
		}
	}

	return i + 1
}

var RemoveDuplicatesFromSortedArray = RemoveDuplicatesFromSortedArray2
