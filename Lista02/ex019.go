package main
import "fmt"

func main() {
    var n, last int
    var sequencia, null []int
    sequencia = append(sequencia, 1)
    last = 1
    
    fmt.Scan(&n)
    
    for i := 1; i <= n; i++ {
        fmt.Scan(&n)
        
        fmt.Printf("%v*%v*%v = ", i, i, i)
        
        for j := 0; j < i; j++ {
            if(i != 1) {
                last += 2
            }
            sequencia = append(sequencia, last)
        }
        
        for j := 0; j < i; j++ {
            if(i != j+1) {
                fmt.Printf("%v+", sequencia[j])
            }else {
                fmt.Printf("%v", sequencia[j])
            }
        }
        fmt.Printf("\n")
        
        sequencia = null
    }
    
    
}