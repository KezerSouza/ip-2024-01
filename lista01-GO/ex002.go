
package main
import "fmt"

func main() {
    var num float64
    var pessoas, pessoas_P, pessoas_G, pessoas_A, pessoas_C float64
    
    
    fmt.Scan(&num)
    
    for n := num; n > 0; n--{
        fmt.Scan(&pessoas, &pessoas_P, &pessoas_G, &pessoas_A, &pessoas_C)
        renda_P := pessoas*pessoas_P/100
        renda_G := pessoas*pessoas_G/100*5
        renda_A := pessoas*pessoas_A/100*10
        renda_C := pessoas*pessoas_C/100*20
        fmt.Printf("A renda do jogo E = %0.2f \n", renda_P+renda_G+renda_A+renda_C)
    }
    
}