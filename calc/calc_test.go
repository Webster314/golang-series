package calc

import "testing"

func TestAdd(t *testing.T) {
	ans := Add(1, 2)
	if ans != 3 {
		t.Error("add(1, 2) should be equal to 3")
	}
	quo, rem := Div(7, 2)
	if quo != 3 || rem != 1 {
		t.Error("div(7, 2) shoul be equal to 3, 1")
	}
}
