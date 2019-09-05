package test

import (
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	var a, exp float64 = -10, 10

	actual := math.Abs(a)

	if actual != exp {
		name := t.Name()
		t.Fatalf("func name: %s a = %f, actual = %f, expected = %f", name, a, actual, exp)
	}
}
