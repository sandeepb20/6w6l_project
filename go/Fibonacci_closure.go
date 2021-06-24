package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var num1, num2 int = 0, 1
	return func() int {
		var temp int = num1
		num1 = num2
		num2 += temp
		return temp
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
