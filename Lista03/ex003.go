package main
import "fmt"

func main() {
    var n, value int
    var vector []int
    
    fmt.Scan(&n)
    
    //Atribuindo os valores do vetor
    for i := 0; i < n;i++ {
        fmt.Scan(&value)
        vector = append(vector, value)
    }
    
    //Escrevendo os valores na ordem inversa
    for i := n-1; i >= 0;i-- {
        fmt.Printf("%v ", vector[i])
    }
    
}