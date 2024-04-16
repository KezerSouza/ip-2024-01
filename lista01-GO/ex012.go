package main
import "fmt"

func main() {
  var horas, valor float64
  
  fmt.Scan(&horas)
  
  if (horas < 3) {
      valor = horas*5
  }else if(horas >= 3) {
      valor = (horas-2)/3*10 + 10
  }
  
  fmt.Printf("O VALOR A PAGAR E = %0.2f", valor)
  
}