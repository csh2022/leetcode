package problems

func RemoveElement(nums []int, val int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		if i == j {
			if nums[i] == val {
				j--
			} else {
				i++
			}
			continue
		}

		if nums[i] == val && nums[j] == val {
			j--
			continue
		}

		if nums[i] != val && nums[j] != val {
			i++
			continue
		}

		if nums[i] == val && nums[j] != val {
			nums[i] = nums[j]
			i++
			j--
			continue
		}

		if nums[i] != val && nums[j] == val {
			i++
			j--
			continue
		}
	}

	return i
}
