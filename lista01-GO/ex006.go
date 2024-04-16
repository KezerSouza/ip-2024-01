package main
import "fmt"

func main() {
  var n int
  var fahrenheit, celsius float64
  
  fmt.Scan(&n)
  
  for i := n; i > 0; i-- {
    fmt.Scan(&fahrenheit)
    celsius = 5*(fahrenheit-32)/9
    if (i != 1) {
        fmt.Printf("%0.2f FAHRENHEIT EQUIVALE A %0.2f CELIUS \n Próximo: ", fahrenheit, celsius)
    }else {
        fmt.Printf("%0.2f FAHRENHEIT EQUIVALE A %0.2f CELIUS \n ", fahrenheit, celsius)
    }
    
    
  }
  
}