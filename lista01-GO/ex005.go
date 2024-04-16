package main
import "fmt"

func main() {
  var conta, tipo string
  var y, valor float64
  
  fmt.Scan(&conta, &y, &tipo)
  if(tipo == "R") {
      valor = 5 + 0.05*y
      
  }else if(tipo == "C") {
      if (y > 80) {
        valor = 500 + 0.25*(y-80)
      }
      if (y <= 80) {
        valor = 500
      }
      
  }else if (tipo == "I") {
      if (y > 100) {
        valor = 800 + 0.04*(y-100)
      }
      if (y <= 100) {
        valor = 800
      }
  }
  fmt.Printf("CONTA = %v \n", conta)
  fmt.Printf("VALOR DA CONTA = %0.2f", valor)
  
}