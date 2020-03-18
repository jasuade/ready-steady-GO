package main

import "golang.org/x/tour/tree"

// Binary tree Struct
// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int)
ch <- t.Value
if t.Left != nil{
	Walk(t.Left)
}else if t.Right != nil{
	Walk(t.Right)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool

func main() {
	ch := make(chan int)
	
}
