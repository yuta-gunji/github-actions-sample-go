package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result1 := EvenOrOdd(10)
	if result1 != "even" {
		t.Errorf("expected: even, actual: %s", result1)
	}

	result2 := EvenOrOdd(11)
	if result2 != "odd" {
		t.Errorf("expected: odd, actual: %s", result2)
	}
}
