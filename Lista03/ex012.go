package main
import "fmt"
import "sort"

func main() {
    var numbers []int
    var n, tests int //Para os valores de entrada
    
    fmt.Scan(&tests)
    
    for i := 0; i < tests;i++ {
        fmt.Scan(&n)
        numbers = append(numbers, n)
    }
    
    sort.Ints(numbers)
    
    for i := 0; i < tests;i++ {
        fmt.Printf("%v\n", numbers[i])
    }
    
}