
package main
import "fmt"

func main() {
    var num, x float64
    
    fmt.Scan(&num)
    
    for num > 0 {
        x += 1/num
        num--
    }
    fmt.Printf("%0.6f", x)
    
}