package main

import (
	"fmt"
	"math"
)

type BangunDatar interface {
	Luas()
	Keliling()
}

// STRUCT PERSEGI, PERSEGI PANJANG, DAN LINGKARAN
type Persegi struct {
	Sisi float64
}

type Persegipjg struct {
	Panjang float64
	Lebar float64
}

type Lingkaran struct {
	Jari2 float64
}

type Modulo struct {
	A float64
	B float64
	C float64
	X float64
	Y float64
	Z float64
}

// RUMUS PERSEGI
func (p Persegi) Luas() float64 {
	return p.Sisi * p.Sisi
}

func (p Persegi) Keliling() float64 {
	return 4 * p.Sisi
}

// RUMUS PERSEGI PANJANG
func (pp Persegipjg) Keliling() float64 {
	return pp.Panjang * pp.Lebar
}

func (pp Persegipjg) Luas() float64 {
	return 2 * pp.Panjang * pp.Lebar
}


// RUMUS LINGKARAN
func (r Lingkaran) Luas() float64 {						// math.Pi --> pi = 22/7 atau 3.14
	return math.Round(math.Pi * r.Jari2 * r.Jari2)		// math.Round() --> untuk membulatkan bilangan
}

func (r Lingkaran) Keliling() float64 {
	return math.Round(2 * math.Pi * r.Jari2)
}

//RUMUS MATH
func (m Modulo) Luas() float64 {
	return math.Round(m.A + m.B / m.C - m.Y * m.X)
}

func (m Modulo) Keliling() float64 {
	return math.Round(m.X * m.B + m.A - m.C / m.Z)
}

// PRINT LUAS DAN KELILING
func hitungPersegi(p Persegi) {
	fmt.Println(p.Keliling())
	fmt.Println(p.Luas())
}

func hitungPersegipjg(pp Persegipjg) {
	fmt.Println(pp.Keliling())
	fmt.Println(pp.Luas())
}

func hitungLingkaran(r Lingkaran) {
	fmt.Println(r.Keliling())
	fmt.Println(r.Luas())
}

func hitungModulo(m Modulo) {
	fmt.Println(m.Keliling())
	fmt.Println(m.Luas())
}

// INSERT VALUE
func main() {
	hitungPersegi(Persegi{Sisi : 6})
	hitungPersegipjg(Persegipjg{Panjang: 8, Lebar:6})
	hitungLingkaran(Lingkaran{Jari2: 14})
	hitungModulo(Modulo{A: 1, B:3, C:7, X:2, Y:10, Z:6})
}