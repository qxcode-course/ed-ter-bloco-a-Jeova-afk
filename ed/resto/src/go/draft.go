package main
import "fmt"

func divisao(n int){
    if n == 0 {
        return 
    }

    
    div := n / 2
    resto := n % 2
    divisao(div)
    
    fmt.Println(div, resto)
}

func main() {
    var n int 
    fmt.Scan(&n)

    divisao(n)
}
