package main
import "fmt"

func main() {
  var raio, altura, At float64
  pi := 3.14159
  
  fmt.Scan(&raio, &altura)
  
  At = 2*pi*raio * (raio+altura)
  
  fmt.Printf("O VALOR DO CUSTO E = %0.2f \n", At*100)
}