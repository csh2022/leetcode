package problems

import (
	"algo/lib"
	"testing"
)

func twoSlicesEqual(a, b []int, cnt int) bool {
	if cnt != len(b) {
		return false
	}
	for i := 0; i < cnt; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestTwoRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		expected []int
	}{
		{
			name:     "test1",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: []int{2, 2},
		},
		{
			name:     "test2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: []int{0, 1, 4, 0, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			oldNums := make([]int, len(test.nums))
			copy(oldNums, test.nums)
			cnt := RemoveElement(test.nums, test.val)
			if !lib.IsTwoIntSlicesEqual(test.nums[:cnt], test.expected) {
				t.Errorf("RemoveElement(%v, %d) = %d, expected %v", oldNums, test.val, test.nums[:cnt], test.expected)
			}
		})
	}
}
