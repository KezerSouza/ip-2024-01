package main
import "fmt"

func main() {
  var fahrenheit, pol, celsius, chuva float64
  
  fmt.Scan(&fahrenheit, &pol)
  
  celsius = (5*fahrenheit - 160)/9
  chuva = pol*25.4
  
  fmt.Printf("O VALOR EM CELSIUS = %0.2f \n", celsius)
  fmt.Printf("A QUANTIDADE DE CHUVA E = %0.2f \n", chuva)
}