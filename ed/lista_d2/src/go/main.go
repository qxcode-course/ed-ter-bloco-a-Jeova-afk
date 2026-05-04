package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	Value int
	next *Node
	prev *Node
	root *Node
}

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}

	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}

	return n.prev

}

type LList struct {
	root *Node 
	size int 

}

func (l *LList) Front() *Node {
	if l.size == 0 {
		return nil
	}

	return l.root.next 
}

func (l *LList) Back() *Node {
	if l.size == 0 {
		return nil
	}

	return l.root.prev
}

func NewLList() *LList {
	root := &Node{}

	root.next = root
	root.prev = root
	root.root = root

	return &LList {
		root: root,
		size: 0,
	}
}

func (l *LList) createNode(value int) *Node {
	return &Node {
		Value: value,
		root: l.root,
	}
}

func (l *LList) PushFront(value int) {
	node := l.createNode(value)

	first := l.root.next
	
	node.next = first
	node.prev = l.root

	first.prev = node
	l.root.next = node

	l.size++

}

func (l *LList) PushBack(value int) {
	node := l.createNode(value)

	last := l.root.prev

	node.prev = last
	node.next = l.root

	last.next = node
	l.root.prev = node 

	l.size++

}

func (l *LList) Search(value int) *Node {
	for node := l.Front(); node != nil; node = node.Next() {
		if node.Value == value {
			return node
		}
	}

	return nil

}

func (l *LList) Insert(node *Node, value int){
	newNode := l.createNode(value)

	prev := node.prev
	
	newNode.next = node
	newNode.prev = prev

	prev.next = newNode
	node.prev = newNode

	l.size++

}

func (l *LList) Remove(node *Node) *Node {
	if node == nil || node == l.root {
		return nil
	}

	next := node.next
	prev := node.prev

	prev.next = next
	next.prev = prev

	l.size--

	if next == l.root {
		return nil
	}

	return next
}

func (l *LList) Size() int {
	return l.size
}

func (l *LList) String() string {
	result := "[" 
	for node :=l.Front(); node != nil; node = node.Next() {
		result += fmt.Sprintf("%d", node.Value)
		if node.Next() != nil {
			result += ", "
		}
	}

	result += "]"
	return result
}

func (l *LList) PopFront() {
	if l.size == 0 {
		return
	}

	first := l.root.next
	second := first.next

	l.root.next = second
	second.prev = l.root

	l.size--
}

func (l *LList) PopBack() {
	if l.size == 0 {
		return 
	}

	last := l.root.prev
	before := last.prev

	before.next = l.root
	l.root.prev = before

	l.size--

}

func (l *LList) Clear() {
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			 fmt.Println(ll.String())
		case "size":
			 fmt.Println(ll.Size())
		case "push_back":
			 for _, v := range args[1:] {
			 	num, _ := strconv.Atoi(v)
			 	ll.PushBack(num)
			 }
		case "push_front":
			 for _, v := range args[1:] {
			 	num, _ := strconv.Atoi(v)
			 	ll.PushFront(num)
			 }
		case "pop_back":
			 ll.PopBack()
		case "pop_front":
			 ll.PopFront()
		case "clear":
			 ll.Clear()
		case "walk":
			 fmt.Print("[ ")
			 for node := ll.Front(); node != nil; node = node.Next() {
			 	fmt.Printf("%v ", node.Value)
			 }
			 fmt.Print("]\n[ ")
			 for node := ll.Back(); node != nil; node = node.Prev() {
			 	fmt.Printf("%v ", node.Value)
			 }
			 fmt.Println("]")
		case "replace":
			 oldvalue, _ := strconv.Atoi(args[1])
			 newvalue, _ := strconv.Atoi(args[2])
			 node := ll.Search(oldvalue)
			 if node != nil {
			 	node.Value = newvalue
			 } else {
			 	fmt.Println("fail: not found")
			 }
		case "insert":
			 oldvalue, _ := strconv.Atoi(args[1])
			 newvalue, _ := strconv.Atoi(args[2])
			 node := ll.Search(oldvalue)
			 if node != nil {
			 	ll.Insert(node, newvalue)
			 } else {
			 	fmt.Println("fail: not found")
			 }
		case "remove":
			 oldvalue, _ := strconv.Atoi(args[1])
			 node := ll.Search(oldvalue)
			 if node != nil {
			 	ll.Remove(node)
			 } else {
			 	fmt.Println("fail: not found")
			 }
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
