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
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	root := &Node{}

	root.next = root
	root.prev = root

	return &LList{
		root: root, 
		size: 0,
	}
}

func (l *LList) PushFront(value int) {
	node := &Node{Value: value}

	first := l.root.next

	node.next = first
	node.prev = l.root

	first.prev = node
	l.root.next = node

	l.size++
}

func (l *LList) PushBack(value int) {
	node := &Node{Value: value}

	last := l.root.prev
	node.prev = last
	node.next = l.root

	last.next = node
	l.root.prev = node

	l.size++
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

func (l *LList) Size() int {
	return l.size
}

func (l *LList) Clear() {
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}

func (l *LList) String() string {
	result := "["
	node := l.root.next

	for node != l.root {
	result += fmt.Sprintf("%d", node.Value)

		if node.next != l.root {
			result += ", "
		}

		node = node.next
	}

	result += "]"
	return result
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
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
