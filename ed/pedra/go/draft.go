package main
import "fmt"
func main() {
    
    var a, b, n int 
    fmt.Scan(&n)

    ganhador := -1
    melhordiferenca := 101

    for i := 0; i < n; i++ {
        fmt.Scan(&a, &b)

        if a >= 10 && b >= 10 {
            dif := a - b
            if dif < 0 {
                dif = -dif
            }

            if dif < melhordiferenca {
                melhordiferenca = dif
                ganhador = i
            }
        }
    }

    if ganhador == -1 {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(ganhador)
    }
      
}
