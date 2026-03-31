package main

import (
	"bufio"
	"sort"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	resultado := []int{}
	for _, v := range vet {
		if v > 0 {
			resultado = append(resultado, v)
		}
	}

	return resultado

}

func getCalmWomen(vet []int) []int {
	resultadoM := []int{}
	for _, v := range vet {
		if v < 0 && v * -1 < 10 {
			resultadoM = append(resultadoM, v)
		}
	}

	return resultadoM
}

func sortVet(vet []int) []int {
	result := make([]int, len(vet))
	copy(result, vet)
	sort.Ints(result)
	return result
}

func sortStress(vet []int) []int {
	resultado := make([]int, len(vet))
	copy(resultado, vet)
	sort.Slice(resultado, func(i, j int) bool {
		si := resultado[i]
		sj := resultado[j]
		if si < 0 {
			si = si * -1
		}
		if sj < 0 {
			sj = sj * -1
		}

		return si < sj
	})

	return resultado
}

func reverse(vet []int) []int {
	resultado := []int{}
	for i := len(vet) - 1; i >= 0; i-- {
		resultado = append(resultado, vet[i])
	}

	return resultado
}

func unique(vet []int) []int {
	resultado := []int{}
	vistos := map[int]bool{}
	for _, v := range vet {
		if !vistos[v] {
			resultado = append(resultado, v)
			vistos[v] = true
		}
	}

	return resultado
	
}

func repeated(vet []int) []int {
	resultado := []int{}
	vistos := map[int]bool{}
	for _, v := range vet {
		if vistos[v] {
			resultado = append(resultado, v)
		}

		vistos[v] = true 
	}

	return resultado
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

