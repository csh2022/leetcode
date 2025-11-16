package problems

import (
	"algo/lib"
	"testing"
)

func TestMergeTwoSortedArray(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			name:     "test1",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:     "test2",
			nums1:    []int{1, 1, 1, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 2, 3},
			n:        3,
			expected: []int{1, 1, 1, 2, 2, 3},
		},
		{
			name:     "test3",
			nums1:    []int{1, 2, 2, 4, 0, 0, 0, 0},
			m:        4,
			nums2:    []int{1, 2, 5, 6},
			n:        4,
			expected: []int{1, 1, 2, 2, 2, 4, 5, 6},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			MergeTwoSortedArray(test.nums1, test.m, test.nums2, test.n)
			if !lib.IsTwoIntSlicesEqual(test.nums1, test.expected) {
				t.Errorf("TwoSum(%v, %d, %v, %d) = %v, expected %v", test.nums1, test.m, test.nums2, test.n, test.nums1, test.expected)
			}
		})
	}
}
