package config

import "testing"

func TestConnectDb(t *testing.T) {
	if res := ConnectDb(); res == nil {
		t.Error("Failed to connect database")
	} else {
		t.Log("config test passed")
	}
}
