package main
import "fmt"

func main() {
    var n, resultado int
    
    fmt.Scan(&n)
    
    resultado =  n*32 + n*62 + n
    
    fmt.Printf("%v", resultado)
    
}
