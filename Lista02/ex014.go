package main
import "fmt"

func main() {
    var linhas, colunas int
    fmt.Scan(&linhas, &colunas)
    
    for i := 2; i <= linhas;i++ {
        fmt.Printf("(%v, 1)", i)
        //Looping enquanto menor que as colunas
        for j := 2; j <= colunas;j++ {
            if( i > j) {
                fmt.Printf(" - (%v, %v)", i, j)
            }
            
        }
        fmt.Printf("\n")
    }
    
}