package main
import ("fmt")

func main() {
  var salario, kw float64
  fmt.Scan(&salario, &kw)
  var custo_por_kw = 0.7*salario/100
  
  fmt.Printf("Custo por kw: %0.2f \n", custo_por_kw)
  fmt.Printf("Custo pordo consumo: %0.2f \n", kw*custo_por_kw)
  fmt.Printf("Custo com desconto: %0.2f \n", kw*custo_por_kw*0.9)
}