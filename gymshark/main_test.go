package main

import (
	"reflect"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	// Test Case 1: case with a small number of items
	result := calculatePacks(200)
	expected := map[int]int{250: 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test_Case_1: calculatePacks(200) = %v; want %v", result, expected)
	}

	// Test Case 2: case with a larger number of items
	result = calculatePacks(5500)
	expected = map[int]int{5000: 1, 500: 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test_Case_2: calculatePacks(5500) = %v; want %v", result, expected)
	}

	// Test Case 3: Case where no additional pack is needed
	result = calculatePacks(5000)
	expected = map[int]int{5000: 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test_Case_3: calculatePacks(5000) = %v; want %v", result, expected)
	}

}
