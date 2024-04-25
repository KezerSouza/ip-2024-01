package main
import "fmt"

func main() {
    var n, number, segmento, segmentomax, last_number int
    
    fmt.Scanln(&n)
    fmt.Scan(&last_number)
    segmento = 1
    
    for i := 1; i < n;i++ {
        fmt.Scan(&number)
        if(number == last_number) {
            segmento += 1
            if(segmentomax < segmento) {
                segmentomax = segmento
            }
        }else {
            segmento = 1
        }
        last_number = number
    }
    fmt.Printf("A maior subsequência de elementos iguais possui %v elementos.", segmentomax)
}