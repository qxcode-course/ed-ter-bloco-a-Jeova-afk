package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Set struct {
	data  []int
	size int 
	capacity int 
}

func NewSet(capacity int) *Set {
	return &Set{
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}


func(s *Set) reserve(newCapacity int) {
	newData := make([]int, newCapacity)

	for i := 0; i < s.size; i++ {
		newData[i] = s.data[i]
	}

	s.data = newData
	s.capacity = newCapacity
}

func (s *Set) String() string {
	return "[" + Join(s.data[:s.size], ", ") + "]"
}

func (s *Set) Contains(value int) bool {
	for i := 0; i < s.size; i++ {
		if s.data[i] == value {
			return true
		}
	}

	return false 
}

func (s *Set) Insert(value int) {
	if s.Contains(value) {
		return
	}

	if s.size == s.capacity {
		s.reserve(s.capacity * 2)
	}

	i := s.size - 1
	for i >= 0 && s.data[i] > value {
		s.data[i+1] = s.data[i]
		i--
	}

	s.data[i+1] = value 
	s.size++

}

func (s *Set) Erase(value int) error {
	pos := -1
	for i := 0; i < s.size; i++ {
		if s.data[i] == value {
			pos = i
			break
		}
	}

	if pos == -1 {
		return fmt.Errorf("value not found")
	}

	for i := pos; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}

	s.size--
	return nil
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

	 v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			 value, _ := strconv.Atoi(parts[1])
			 v = NewSet(value)
		case "insert":
			 for _, part := range parts[1:] {
			 	value, _ := strconv.Atoi(part)
				v.Insert(value)
			 }
		case "show":
			fmt.Println(v)
		case "erase":
			 value, _ := strconv.Atoi(parts[1])
			 err := v.Erase(value)
			 if err != nil {
				fmt.Println(err)
			 }
		case "contains":
			 value, _ := strconv.Atoi(parts[1])
			 if v.Contains(value) {
				fmt.Println("true")
			 } else {
				fmt.Println("false")
			 }
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
