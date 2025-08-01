package main

import "testing"

func TestGenerateSecret(t *testing.T) {
	for i := 0; i < 1000; i++ {
		s := generateSecret()
		if s < 1 || s > 100 {
			t.Errorf("generateSecret() = %d, out of range", s)
		}
	}
}
