package test

import "errors"

func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("b is fall")
	}

	return a/b, nil
}
