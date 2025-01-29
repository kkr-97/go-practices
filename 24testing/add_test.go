package main

import "testing"

type Test[T Number] struct {
	name   string
	a      T
	b      T
	output T
}

func TestAdderTableDriven(t *testing.T) {
	tests := []Test[int]{
		{"Test 1", 1, 2, 3},
		{"Test 2", 2, 3, 5},
		{"Test 3", 10, 20, 30},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if result := add(tc.a, tc.b); result != tc.output {
				t.Errorf("expected %d, got %d", tc.output, result)
			} else {
				t.Logf("expected %d, got %d", tc.output, result)
			}
		})
	}

}

func TestAdder(t *testing.T) {
	result := add(4, 5)
	expected := 9

	if result != expected {
		t.Errorf("Result was incorrect, got: %v, want: %v.", result, expected)
	}
}
