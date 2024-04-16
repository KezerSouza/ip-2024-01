package main
import "fmt"

func main() {
    var n1, n2, n3, resultado int
    
    fmt.Scan(&n1, &n2, &n3)
    
    if(n1 > 9 || n2 > 9 || n2 > 9) {
        fmt.Println("DIGITO INVALIDO")
    }
    
    resultado = n1*100 + n2*10 + n3
    fmt.Printf("%v%v%v, %v", n1, n2, n3, resultado*resultado)
    
}