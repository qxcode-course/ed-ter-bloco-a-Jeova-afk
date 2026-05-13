package main

import (
	"fmt"
	"strings"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	values := []string{}

	for n := l.Front(); n != l.End(); n = n.next {
		if n == sword {
			values = append(values, fmt.Sprintf(">%v<", n.Value))
		} else {
			values = append(values, fmt.Sprintf("%v", n.Value))
		}
	}

	return "[" + strings.Join(values, ", ") + "]"
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	next := it.next

	if next == l.root {
		next = next.next
	}

	return next

}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for i := 0; i < chosen - 1; i++ {
		sword = Next(l, sword)
	}
	for i := 0; i < qtd - 1; i++ {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
