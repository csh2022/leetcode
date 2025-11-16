package problems

import (
	"algo/lib"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "test1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "test2",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "test3",
			nums:     []int{1, 10, 2, 5, 100},
			target:   101,
			expected: []int{0, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := TwoSum(test.nums, test.target)
			if !lib.IsTwoIntSlicesEqual(result, test.expected) {
				t.Errorf("TwoSum(%v, %d) = %d, expected %v", test.nums, test.target, result, test.expected)
			}
		})
	}
}
