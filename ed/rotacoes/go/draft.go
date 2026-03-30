package main
import "fmt"
func main() {
    
    var T, R int
    fmt.Scan(&T, &R)

    vetor := make([]int, T)
    for i := 0; i < T; i++ {
        fmt.Scan(&vetor[i])
    }

    resultado := make([]int, T)
    for i := 0; i < T ; i++ {
        posicaonova := (i + R) % T
        resultado[posicaonova] = vetor[i]
    }

    fmt.Print("[ ")
    for _, v := range resultado {
        fmt.Print(v, " ")
    }


    fmt.Println("]")

}
