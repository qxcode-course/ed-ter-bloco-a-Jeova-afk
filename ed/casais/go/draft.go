package main
import "fmt"
func main() {

    var n int 
    fmt.Scan(&n)

    sempar := make(map[int]int)
    pares := 0

    for i := 0; i < n; i++ {
        var animal int 
        fmt.Scan(&animal)

        par := -animal

        if sempar[par] > 0 {
            sempar[par]--
            pares++

        } else {
            sempar[animal]++
        }
    }

    fmt.Println(pares)
    
}