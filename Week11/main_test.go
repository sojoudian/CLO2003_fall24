package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 6
	if result != expected {
		t.Fatalf("Add(2, 3) will result %d, but we got %d", result, expected)
	}
}
