package main
import "fmt"
func main() {
    
    var h, p, f, d int 
    fmt.Scan(&h, &p, &f, &d)

    pos := f
    
    for {
        pos = (pos + d + 16) % 16

        if pos == h {
            fmt.Println("S")
            break
        }

        if pos == h {
            fmt.Println("S")
            break
        }

        if pos == p {
            fmt.Println("N")
            break
        }
    }
}
