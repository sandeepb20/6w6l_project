package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	} else if t.Left != nil {
		Walk(t.Left, ch)
	} else {
		ch <- t.Value
		if t.Right != nil {
			Walk(t.Right, ch)
		}
		return
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.

func Walker(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	//close the channel to avoid panic
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	var b bool
	done := make(chan bool)

	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walker(t1, ch1)
	go Walker(t2, ch2)
	go func() {
		for i := range ch1 {
			if i == <-ch2 {
				b = true
			} else {
				b = false
				break
			}
		}
		done <- true
	}()
	<-done
	return b
}
func main() {
	tree1 := tree.New(10)
	tree2 := tree.New(10)
	fmt.Println("Tree1 is as follows:", tree1)
	fmt.Println("Tree2 is as follows:", tree2)
	if Same(tree1, tree2) == true {

		fmt.Println("Yes they both trees (tree1 and tree2) are same ")
	} else {
		fmt.Println("No they both trees (tree1 and tree2) are not same ")
	}
	fmt.Println("\n------ \n")
	tree3 := tree.New(9)
	fmt.Println("Tree1 is as follows:", tree1)
	fmt.Println("Tree3 is as follows:", tree3)
	if Same(tree1, tree3) == true {

		fmt.Println("Yes they both trees (tree1 and tree3) are same ")
	} else {
		fmt.Println("No they both trees (tree1 and tree3) are not same ")
	}
	fmt.Println("")
}
