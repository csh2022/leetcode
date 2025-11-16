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
