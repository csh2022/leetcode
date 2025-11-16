package problems

import (
	"algo/lib"
	"testing"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "test1",
			nums:     []int{1, 1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "test2",
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: []int{0, 1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			oldNums := make([]int, len(test.nums))
			copy(oldNums, test.nums)
			cnt := RemoveDuplicatesFromSortedArray(test.nums)
			if !lib.IsTwoIntSlicesEqual(test.nums[:cnt], test.expected) {
				t.Errorf("RemoveDuplicatesFromSortedArray(%v) = %v, expected %v", oldNums, test.nums[:cnt], test.expected)
			}
		})
	}
}
