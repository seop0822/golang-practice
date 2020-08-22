package main

import "fmt"

type Node struct {
	next *Node
	val int
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (l* LinkedList) RemoveNode(node *Node){
	if node == l.root{
		l.root = l.root.next
		if l.root == nil{
			l.tail =nil
		}
		node.next = nil
		return
	}
	//이전노드 구하기
	prev := l.root
	for prev.next != node {
		prev = prev.next
	}

	if node == l.tail {
		prev.next = nil
		l.tail = prev
	} else {	//root, tail이 아닐때
		prev.next = prev.next.next
	}
	node.next = nil
}

func (l *LinkedList) AddNode(val int) {
	if l.root == nil {
		l.root = &Node{val:val}
		l.tail = l.root
		return
	}
	l.tail.next = &Node{val:val}
	l.tail = l.tail.next
}

func main () {
	list := &LinkedList{}
	list.AddNode(0)

	for i:=1; i<10; i++ {
		list.AddNode(i)
	}

	list.PrintNodes()
	list.RemoveNode(list.root.next)
	list.PrintNodes()
	list.RemoveNode(list.root)
	list.PrintNodes()
	fmt.Printf("tail: %d",list.tail.val)

}

func (l *LinkedList) PrintNodes() {
	node := l.root
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



