package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t.Left == nil {
		ch <- t.Value
		if t.Right != nil {
			Walk(t.Right, ch)
		}
	} else {
		Walk(t.Left, ch)
		ch <- t.Value
		if t.Right != nil {
			Walk(t.Right, ch)
		}
	}
}

func throwAndClose(tr *tree.Tree, ch chan int) {
	Walk(tr, ch)
	close(ch)
}

func compareTrees(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go throwAndClose(t1, ch1)
	go throwAndClose(t2, ch2)
	for i := range ch1 {
		fmt.Println(i)
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	tr1 := tree.New(1)
	tr2 := tree.New(2)
	fmt.Println(compareTrees(tr1, tr2))
}
