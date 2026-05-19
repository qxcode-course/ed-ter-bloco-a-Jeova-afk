package main

import (
    "bufio"
    "fmt"
    "os"
)

func valido(seq []rune, pos int, valor rune, L int) bool {
    for i := pos - L; i <= pos+L; i++ {

        if i < 0 || i >= len(seq) || i == pos {
            continue
        }

        if seq[i] == valor {
            return false
        }
    }

    return true
}

func resolver(seq []rune, pos int, L int) bool {
    if pos == len(seq) {
        return true
    }

    if seq[pos] != '.' {
        return resolver(seq, pos+1, L)
    }

    for num := 0; num <= L; num++ {
        valor := rune(num + '0')

        if valido(seq, pos, valor, L) {
            seq[pos] = valor

            if resolver(seq, pos+1, L) {
                return true
            }

            seq[pos] = '.'
        }
    }

    return false
}

func main() {
    in := bufio.NewReader(os.Stdin)

    var entrada string
    var L int

    fmt.Fscan(in, &entrada)
    fmt.Fscan(in, &L)

    seq := []rune(entrada)

    resolver(seq, 0, L)

    fmt.Println(string(seq))
}
