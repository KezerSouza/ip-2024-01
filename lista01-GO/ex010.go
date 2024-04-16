package main
import "fmt"

func main() {
  var a, b, c, d, det float64
  
  fmt.Scan(&a, &b, &c, &d)
  
  det = a*d - b*c
  fmt.Printf("O VALOR DO DETERMINANTE E = %0.2f", det)
  
}