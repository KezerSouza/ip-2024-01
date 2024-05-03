package main
import "fmt"

func main() {
    var n, k, value, answer int
    var vector []int
    
    fmt.Scan(&n)
    
    //Atribuir n até que o número seja válido
    for n < 1 || n >= 1000 {
        fmt.Scan(&n)
    }
    
    //Atribuindo os valores do vetor
    for i := 0; i < n;i++ {
        fmt.Scan(&value)
        vector = append(vector, value)
    }
    
    fmt.Scan(&k) //Número a ser comparado
    
    for i := 0; i < n;i++ {
        if(vector[i] >= k) {
            answer += 1
        }
    }
    
    fmt.Println(answer)
    
}