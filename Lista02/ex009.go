package main
import "fmt"

func main() {
    var n int
    var s string //Para verificar se a entrada está correta
    var primo = true
    
    fmt.Scanln(&n, &s)
    //Número inválido
    if(n <= 0 || s != "") {
        fmt.Println("Número inválido!")
    }else {
    //Todo resto do programa 
        //Se par
        if(n % 2 == 0 &&  n != 2) {
            primo = false
        }else if(n == 2) {
            primo = true
        }else if(n % 2 != 0) {
            for i := 3; i <= n/3; i++ {
                if(n % i == 0) {
                    primo = false
                    break
                }
            }
        }
        if(primo) {
            fmt.Println("PRIMO")
        }else {
            fmt.Println("NÃO PRIMO")
        }
    }
    
    
    
}
