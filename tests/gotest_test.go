package tests

import "testing"

func TestDivision1(t *testing.T) {
	if res, err := Division(6, 2); res != 3 || err != nil {
		t.Error("division function tests do not pass")
	} else {
		t.Log("First test passed")
	}
}

func TestDivision2(t *testing.T) {
	if _, err := Division(6, 0); err == nil {
		t.Error("division function tests do not pass")
	} else {
		t.Log("First test passed")
	}
}

func TestHello(t *testing.T) {
	if res := Hello("I am strong!"); res != "I am strong! kappa" {
		t.Error(res)
	} else {
		t.Log("string test passed")
	}
}
