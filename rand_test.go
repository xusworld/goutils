package goutils

import "testing"

func TestRandInt(t *testing.T) {

	for i := 0; i < 5; i++ {
		t.Log("RandInt() = ", RandInt())
	}
}

func TestRandIntn(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log("RandInt() = ", RandIntn(100))
	}
}
