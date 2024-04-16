package main
import "fmt"

func main() {
  var hora, minuto, segundo int
  
  fmt.Scan(&hora, &minuto, &segundo)
  
  resultado := (hora*3600 + minuto*60 + segundo)
  
  fmt.Printf("O TEMPO EM SEGUNDOS E = %v", resultado)
  
}