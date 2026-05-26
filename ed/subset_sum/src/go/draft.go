package main
import "fmt"
func main() {
    var n, target int
    fmt.Scan(&n, &target)

    nums := make([]int, n)

    for i := 0; i < n; i++ {
        fmt.Scan(&nums[i])
    } 

    var backtrack func(index, sum int) bool

    backtrack = func(index, sum int) bool {
        if sum == target {
            return true
        }

        if index >= n {
            return false
        }

        if sum > target {
            return false
        }

        take := backtrack(index+1, sum+nums[index])

        skip := backtrack(index+1, sum)

        return take || skip
    }

    if backtrack(0, 0){
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }

}
