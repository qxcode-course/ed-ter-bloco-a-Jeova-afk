package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data []int
	size int 
	capacity int

}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet {
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func(m *MultiSet) expand() {
	if m.capacity == 0 {
		m.capacity = 1
	} else {
		m.capacity *= 2
	}

	newData := make([]int, m.capacity)
	
	for i := 0; i < m.size; i++ {
		newData[i] = m.data[i]
	}

	m.data = newData

}

func (m *MultiSet) Insert(value int) {
	if m.size == m.capacity {
		m.expand()	
	}

	i := m.size - 1
	for i >= 0 && m.data[i] > value {
		m.data[i+1] = m.data[i]
		i--
	}

	m.data[i+1] = value 
	m.size++

}

func (m *MultiSet) String() string {
	return "[" + Join(m.data[:m.size], ", ") + "]"
}

func (m *MultiSet) Contains(value int) bool{
	for i := 0; i < m.size; i++ {
		if m.data[i] == value {
			return true
		}
	}

	return false

}

func (m *MultiSet) Erase(value int) error{
	pos := -1
	for i := 0; i < m.size; i++ {
		if m.data[i] == value {
			pos = i
			break
		}
	}

	if pos == -1 {
		return fmt.Errorf("value not found")
	}

	for i := pos; i < m.size-1; i++ {
		m.data[i] = m.data[i+1]
	}

	m.size--
	return nil
}

func (m *MultiSet) Count(value int) int {
	count := 0

	for i := 0; i < m.size; i++ {
		if m.data[i] == value {
			count++
		}
	}

	return count

}

func (m *MultiSet) Clear() {
	m.size = 0
}

func (m *MultiSet) Unique() int {
	if m.size == 0 {
		return 0
	}

	count := 1

	for i := 1; i < m.size; i++ {
		if m.data[i] != m.data[i-1] {
			count++
		}
	}

	return count

}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	 ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			 value, _ := strconv.Atoi(args[1])
			 ms = NewMultiSet(value)
		case "insert":
			 for _, part := range args[1:] {
			 	value, _ := strconv.Atoi(part)
				ms.Insert(value)
			 }
		case "show":
			fmt.Println(ms)
		case "erase":
			 value, _ := strconv.Atoi(args[1])
			 err := ms.Erase(value)
			 if err != nil {
				fmt.Println(err)
			 }
		case "contains":
			 value, _ := strconv.Atoi(args[1])
			 if ms.Contains(value) {
				fmt.Println("true")
			 } else {
				fmt.Println("false")
			 }
		case "count":
			 value, _ := strconv.Atoi(args[1])
			 fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
