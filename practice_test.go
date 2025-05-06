package main

import "testing"

// TestBubbleSort tests the BubbleSort function
func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 3, 8, 1, 2}, []int{1, 2, 3, 5, 8}},
		{[]int{10, 7, 8, 9, 1}, []int{1, 7, 8, 9, 10}},
		{[]int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
	}

	for _, test := range tests {
		result := BubbleSort(test.input)
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("BubbleSort(%v) = %v; want %v", test.input, result, test.expected)
			}
		}
	}
}

// TestIsPalindrome tests the IsPalindrome function
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"racecar", true},
		{"hello", false},
	}
	for _, test := range tests {
		result := IsPalindrome(test.input)
		if result != test.expected {
			t.Errorf("IsPalindrome(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
