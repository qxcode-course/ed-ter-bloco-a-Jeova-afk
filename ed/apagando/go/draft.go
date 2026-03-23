package main
import "fmt"
func main() {
	var n int 
	fmt.Scan(&n)

	fila := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&fila[i])
	}

	var m int 
	fmt.Scan(&m)

	saiu := make([]int, m)
	
	for i := 0; i < m; i++ {
		fmt.Scan(&saiu[i])
	}

	for i := 0; i < n; i++ {
		removido := false 

		for j := 0; j < m; j++ {
			if fila[i] == saiu[j] {
				removido = true 
			}
		}

		if removido == false {
			fmt.Print(fila[i], " ")
		}
	}

	fmt.Println()
}
