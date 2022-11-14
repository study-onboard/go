package ch05

import "testing"

func TestCalculateArea(t *testing.T) {
	area := CalculateArea(10, 20)
	if area != 200 {
		t.Error("Calculate area wrong")
	}
}
