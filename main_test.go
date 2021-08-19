package main

import "testing"

var (
	persegi Persegi = Persegi{Sisi:4}
	luasExpected float64 = 16
	kelilingExpected float64 = 16
)

// var (
// 	modulo Modulo = Modulo{A: 5, B:2, C:1, X:3, Y:8, Z:1}
// 	luasExpected float64 = 
// 	kelilingExpected float64 = 
// )

func TestLuasPersegi(t *testing.T) {
	t.Logf("Luas persegi : %v", persegi.Luas())

	if persegi.Luas() != luasExpected {
		t.Errorf("Salah! Seharusnya : %v", luasExpected)
	}
}

func TestKelilingPersegi(t *testing.T) {
	t.Logf("Keliling persegi : %v", persegi.Keliling())

	if persegi.Keliling() != kelilingExpected {
		t.Errorf("Salah! Seharusnya : %v", kelilingExpected)
	}
}