package main
import "fmt"
func main() {

    var N, E int 
    fmt.Scan(&N, &E)

    vivo := make([]bool, N)
    for i := 0; i < N; i++ {
        vivo[i] = true 
    }

    restantes := N 
    ind := E - 1
    for restantes > 1 {
        fmt.Print("[ ")
        primeiro := true 
        for i:= 0; i < N; i++ {
            if vivo[i] {
                if !primeiro {
                    fmt.Print(" ")
                }

                if i == ind {
                    fmt.Printf("%d>", i+1)
                } else {
                    fmt.Printf("%d", i+1)
                }

                primeiro = false 
            }
        }

        fmt.Println(" ]")

        prox := ind
        for {
            prox = (prox + 1) % N 
            if vivo[prox] {
                break
            }
        }

        vivo[prox] = false 
        restantes--

        ind = prox
        for {
            ind = (ind + 1) % N 
            if vivo[ind] {
                break
            }
        }
    }

    fmt.Print("[ ")
    for i := 0; i < N; i++ {
        if vivo[i] {
            fmt.Printf("%d>", i+1)
            break
        }
    }

    fmt.Println(" ]")

}