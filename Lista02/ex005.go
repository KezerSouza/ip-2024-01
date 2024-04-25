package main
import "fmt"

func main() {
    var n, number, segmento, segmentomax, last_number int
    
    fmt.Scanln(&n)
    fmt.Scan(&last_number)
    
    for i:= 0; i < n;i++ {
        fmt.Scan(&number)
        if(number > last_number) {
            segmento += 1
            if(segmentomax < segmento) {
                segmentomax = segmento
            }
        }else {
            segmento = 0
        }
        last_number = number
        
    }
    fmt.Printf("O comprimento do segmento crescente máximo é: %v", segmentomax)
}