package test

import "testing"

func TestDivision(t *testing.T) {
	var a, b, exp float64 = 10, 2, 5

	actual, err := division(a, b)

	if err != nil {
		t.Errorf("a = %f, b = %f, expect = %f, err %v", a, b, exp, err)
		return
	}

	if actual != exp {
		t.Errorf("a = %f, b = %f, expect = %f, actual = %f", a, b, exp, actual)
	}
}
