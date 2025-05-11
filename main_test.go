package main

import (
	"testing"
)

// TestGenerateRandomElements checks what happens if you pass 0 elements and what will happen if I transfer 10 elements normally
func TestGenerateRandomElements(t *testing.T) {
	_, err := generateRandomElements(0)
	if err == nil {
		t.Error("Expected error for size = 0")
	}

	data, err := generateRandomElements(10)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(data) != 10 {
		t.Errorf("Expected length 10, got %d", len(data))
	}
}

// TestMaximum checks for empty array protection, then checks for a single number array (minimum possible not empty), checks an array of multiple elements
func TestMaximum(t *testing.T) {
	_, err := maximum([]int{})
	if err == nil {
		t.Error("Expected error for empty slice")
	}

	value, err := maximum([]int{42})
	if err != nil || value != 42 {
		t.Errorf("Expected 42, got %d (err: %v)", value, err)
	}

	value, _ = maximum([]int{1, 3, 2, 10, 5})
	if value != 10 {
		t.Errorf("Expected 10, got %d", value)
	}
}

// TestMaxChunks checks for an error when passing an empty slice, that the array has fewer chunks than the number of threads, normal operation with an array of 8 elements
func TestMaxChunks(t *testing.T) {
	_, err := maxChunks([]int{})
	if err == nil {
		t.Error("Expected error for empty slice")
	}

	small := make([]int, CHUNKS-1)
	_, err = maxChunks(small)
	if err == nil {
		t.Error("Expected error for too small slice")
	}

	data := []int{1, 2, 3, 4, 5, 100, 10, 8}
	result, err := maxChunks(data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 100 {
		t.Errorf("Expected 100, got %d", result)
	}
}
