package main

import (
	"exercise/dataStruct"
	"fmt"
)

func main() {
	h := dataStruct.Heap{}

	h.Push(9)
	h.Push(8)
	h.Push(7)
	h.Push(6)
	h.Push(5)

	h.Print()

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())

	//BFS부분
	// tree := dataStruct.NewBinaryTree(5)
	// tree.Root.AddNode(3)
	// tree.Root.AddNode(2)
	// tree.Root.AddNode(4)
	// tree.Root.AddNode(8)
	// tree.Root.AddNode(7)
	// tree.Root.AddNode(6)
	// tree.Root.AddNode(10)
	// tree.Root.AddNode(9)

	// tree.Print()
	// fmt.Println()

	// if found, cnt := tree.Search(6); found {
	// 	fmt.Println("found 6 cnt", cnt)
	// } else {
	// 	fmt.Println("NOt Found cnt", cnt)
	// }

	//DFS 부분
	// tree := dataStruct.Tree{}

	// val:=1

	// tree.AddNode(val)
	// val++

	// for i:=0; i<3; i++ {
	// 	tree.Root.AddNode(val)
	// 	val++
	// }

	// for i:=0; i < len(tree.Root.Childs); i++ {
	// 	for j :=0; j <2; j++ {
	// 		tree.Root.Childs[i].AddNode(val)
	// 		val++
	// 	}
	// }
	// tree.DFS_Recursion()
	// fmt.Println()
	// tree.DFS_Stack()

	// fmt.Println()
	// tree.BFS()
}
