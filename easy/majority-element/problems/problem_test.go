package problems

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "test1",
			nums:     []int{1, 2, 2, 3},
			expected: 2,
		},
		{
			name:     "test2",
			nums:     []int{1, 1, 2, 3},
			expected: 1,
		},
		{
			name:     "test3",
			nums:     []int{8, 9, 9, 9},
			expected: 9,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MajorityElement(test.nums)
			if result != test.expected {
				t.Errorf("MajorityElement(%v) = %d, expected %d", test.nums, result, test.expected)
			}
		})
	}
}
