package main

import "fmt"

func Sum(s []int, c chan int) {
	var sum int = 0
	for _, c := range s {
		sum += c
	}
	c <- sum
}

func main() {
	s := []int{7, 5, 8, -9, 2, 4}

	c := make(chan int)
	go Sum(s[0:3], c)
	go Sum(s[3:6], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
