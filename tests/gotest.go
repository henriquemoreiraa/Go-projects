package tests

import "errors"

func main() {

}

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisor can not be 0")
	}

	return a / b, nil
}

func Hello(s string) string {
	s += " kappa"

	return s
}
