package main
import ("fmt")

func main() {
  var number1, number2, number3 float64
  fmt.Scan(&number1, &number2, &number3)
  var media = (number1+number2+number3)/3
  
  fmt.Printf("MEDIA = %0.2f \n", media)
  if (media >= 6.0) {
      fmt.Printf("APROVADO \n")
  }else {
      fmt.Printf("REPROVADO \n")
  }
  
}