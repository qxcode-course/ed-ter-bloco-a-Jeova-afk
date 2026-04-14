package main
import "fmt"
func main() {
    

    var qtd_possui, qtd_total int
    fmt.Scan(&qtd_total, &qtd_possui)

    aux := make(map[int] bool)
    var repetidas []int
    var faltando []int

     var temp int

    for i := 0; i < qtd_possui; i++ {
        fmt.Scan(&temp)
        if aux[temp] {
            repetidas = append(repetidas, temp)
        }

        aux[temp] = true 
    }

    for i := 1; i <= qtd_total; i++ {
        if !aux[i] {
            faltando = append(faltando, i)
        }
    }

    if len(repetidas) > 0 {
        for i, num := range repetidas {
            if i > 0 {
                fmt.Printf(" ")
            }

            fmt.Print(num)
        }
        
        fmt.Println()
    } else {
        fmt.Println("N")
    }

    if len(faltando) > 0 {
        for i, num := range faltando {
            if i > 0 {
                fmt.Printf(" ")
            }

            fmt.Print(num)
        }

        fmt.Println()
    } else {
        fmt.Println("N")
    }

}
