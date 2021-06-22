package main

import "golang.org/x/tour/pic"

func Pic1(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		s0 := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			s0[j] = uint8((i ^ j) / 2)
		}
		s[i] = s0
	}
	println(s)
	return s
}
func Pic2(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		s0 := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			s0[j] = uint8((i + j) / 2)
		}
		s[i] = s0
	}
	println(s)
	return s
}
func Pic3(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		s0 := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			s0[j] = uint8((i * j) / 2)
		}
		s[i] = s0
	}
	println(s)
	return s
}

func main() {
	pic.Show(Pic1)
	pic.Show(Pic2)
	pic.Show(Pic3)
}
