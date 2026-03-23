package main
import "fmt"
func main() {
	var n int 
	fmt.Scan(&n)

	fila := make([]int , n)
	for i := 0; i < n; i++ {
		fmt.Scan(&fila[i])
	}

	var m int
	fmt.Scan(&m)

	saiu := make(map[int]bool)
	for i := 0; i < m; i++ {
		var x int 
		fmt.Scan(&x)
		saiu[x] = true 
	}

	for i := 0; i < n; i++ {
		if !saiu[fila[i]] {
			fmt.Print(fila[i], " ")
		}
	}

	fmt.Println()
}