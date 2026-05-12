package main
import (
    "bufio"
    "fmt"
    "os"
)
func main() {

    in := bufio.NewReader(os.Stdin)

    var entrada string
    fmt.Fscan(in, &entrada)

    esquerda := []rune{}
    direita := []rune{}

    for _, c := range entrada {

        switch c {

        case 'R':
            esquerda = append(esquerda, '\n')

        case 'B':
            if len(esquerda) > 0 {
                esquerda = esquerda[:len(esquerda)-1]
            }

        case 'D':
            if len(direita) > 0 {
                direita = direita[1:]
            }

        case '>':
            if len(direita) > 0 {
                esquerda = append(esquerda, direita[0])
                direita = direita[1:]
            }

        case '<':
            if len(esquerda) > 0 {
                ultimo := esquerda[len(esquerda)-1]
                esquerda = esquerda[:len(esquerda)-1]
                direita = append([]rune{ultimo}, direita...)
            }

        default:
            esquerda = append(esquerda, c)
        }
    }

    fmt.Print(string(esquerda) + "|" + string(direita) + "\n")

}
