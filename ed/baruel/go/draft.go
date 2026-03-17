package main
import "fmt"
func main() {
    
    var total, possui int 
    fmt.Scan(&total)
    fmt.Scan(&possui)

   var contar [101]int 
   
   for i := 0; i < possui; i++ {
        var figurinha int 
        fmt.Scan(&figurinha);
        contar[figurinha]++
   }

   temreptd := false 
   primeiro := true 

   for i := 1; i <= 100; i++ {
        if contar[i] > 1 {
            for j := 0; j < contar[i]-1; j++ {

                if !primeiro {
                    fmt.Print(" ")
                }

                fmt.Print(i)
                primeiro = false 
                temreptd = true 
            }
        }
   }

   if !temreptd {
    fmt.Println("N")
   } else {
    fmt.Println()
   }

   temfaltando := false 
   primeiro = true 

   for i := 1; i <= total; i++ {
        if contar[i] == 0 {
            if !primeiro {
                fmt.Print(" ")
            }

            fmt.Print(i)
            primeiro = false 
            temfaltando = true 
        }
    }

    if !temfaltando {
        fmt.Println("N")
    } else {
        fmt.Println()
    }
    
}