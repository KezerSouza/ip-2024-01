package main
import "fmt"

func main() {
    var n, value, answer int
    var vector []int
    
    fmt.Scan(&n)
    
    for i := 0; i < n;i++ {
        fmt.Scan(&value)
        vector = append(vector, value)
        answer += value
    }
    
    fmt.Println(answer)
    
}