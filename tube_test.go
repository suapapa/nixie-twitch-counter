package main

import "testing"

func TestGetStDigit(t *testing.T) {
	a, b, c := getStDigits(1)
	if a != 0 || b != 0 || c != 1 {
		t.Errorf("expected 0 0 1, got %d %d %d", a, b, c)
	}
	a, b, c = getStDigits(50)
	if a != 0 || b != 5 || c != 0 {
		t.Errorf("expected 0 5 0, got %d %d %d", a, b, c)
	}
	a, b, c = getStDigits(3486)
	if a != 4 || b != 8 || c != 6 {
		t.Errorf("expected 4 8 6, got %d %d %d", a, b, c)
	}
}
