package main

import "fmt"

type Node struct {
	next *Node
	val int
}

func main () {
	var root *Node
	var tail * Node

	root = &Node{val:0}
	tail = root //

	for i:=1; i<2; i++ {
		tail = AddNode(tail, i)
	}

	PrintNodes(root)
	root , tail = RemoveNode(root, root, tail)
	PrintNodes(root)
	root , tail = RemoveNode(root, root, tail)
	PrintNodes(root)
	fmt.Printf("tail %d",tail.val)

}

func PrintNodes(root *Node) {
	node := root
	if node == nil{
		fmt.Println("Nothing")
	}else {
		for node.next != nil{
			fmt.Printf("%d ->", node.val)
			node = node.next
		}
		fmt.Printf("%d\n", node.val)
	}
}

func RemoveNode(node *Node, root *Node, tail *Node) (*Node, *Node){	//삭제할떄 root와 tail이 바뀌기 때문에 반환
	if node == root{
		root = root.next
		//삭제하려는 노드가 root 하나일때 tail도 처리해줘야함 root.next가 널이라는 것은 노드가 하나라는 것
		if root == nil {
			tail = nil
		}
		return root, tail
	}
	//이전노드 구하기
	prev := root
	for prev.next != node {
		prev = prev.next
	}

	if node == tail {
		prev.next = nil
		tail = prev
	} else {	//root, tail이 아닐때
		prev.next = prev.next.next
	}

	return root, tail
}

func AddNode(tail *Node, val int)  *Node{

	node := &Node{val:val}
	tail.next = node

	return node
}