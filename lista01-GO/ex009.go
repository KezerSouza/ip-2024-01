package main
import "fmt"

func main() {
  var a, b, c, delta float64
  
  fmt.Scan(&a, &b, &c)
  
  delta = b*b - 4*a*c
  fmt.Printf("O VALOR DE DELTA E = %0.2f", delta)
  
}