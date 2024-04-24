package main
import "fmt"

func main() {
    var T, n1, n2 int
    
    fmt.Scanln(&T)
    
    for i := 0; i < T; i++ {
        fmt.Scan(&n1, &n2)
        inverter(&n1)
        inverter(&n2)
        if(n1 >= n2) {
            fmt.Println(n1)
        }else {
            fmt.Println(n2)
        }
    }
    
}

func inverter(x *int) {
    var unidade, dezena, centena int
    unidade = *x/100
    dezena = (*x-unidade*100) / 10
    centena = *x-unidade*100-dezena*10
    *x = unidade + dezena*10 + centena*100
}
