package main
import "fmt"

func main() {
    var horas, por_hora, salario float64
    var matricula int
    fmt.Scan(&matricula, &horas, &por_hora)
 
    if(matricula != 0) {
        salario = horas*por_hora
        fmt.Printf("%v %0.2f\n", matricula, salario)
        main()
    }    
    
}
