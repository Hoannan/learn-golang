package math

import (
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	var a, exp float64 = -10, 10

	actual := math.Abs(a)

	if actual != exp {
		t.Fatalf("a = %f, actual = %f, expected = %f", a, actual, exp)
	}
}
