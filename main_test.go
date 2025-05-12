package main

import (
	"math"
	"testing"
)

// TestGenerateRandomElements according to comments, corrected: made a table test and edge cases
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{"Zero size", 0, 0},
		{"Small size", 10, 10},
		{"Large size", 1000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)
			if len(data) != tt.expected {
				t.Errorf("expected length %d, got %d", tt.expected, len(data))
			}
		})
	}
}

// TestMaximum according to comments, corrected: made a table test and edge cases
func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{"One element", []int{42}, 42},
		{"Max at end", []int{1, 2, 3, 10}, 10},
		{"Max at start", []int{100, 1, 2, 3}, 100},
		{"Max in middle", []int{1, 50, 2, 3}, 50},
		{"With MaxInt64", []int{1, 2, math.MaxInt64}, math.MaxInt64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.data)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}

// TestMaxChunks according to comments, corrected: made a table test and edge cases
func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{"Exact CHUNKS", []int{1, 2, 3, 4, 5, 6, 100, 8}, 100},
		{"Large max at start", []int{math.MaxInt64, 1, 2, 3, 4, 5, 6, 7}, math.MaxInt64},
		{"Large max at end", []int{1, 2, 3, 4, 5, 6, 7, math.MaxInt64}, math.MaxInt64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxChunks(tt.data)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
