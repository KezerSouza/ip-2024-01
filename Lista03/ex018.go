package main
import "fmt"

func main() {
    var n, digit, soma, soma_2, soma_3 int
    var digits, digits_2, digits_3, null = make([]int, 11), make([]int, 11), make([]int, 11), make([]int, 11)
    var valid = true
    
    fmt.Scan(&n)
    
    //Para cada número
    for i := 0; i < n;i++ {
        //dígitos do cpf
        for i := 0; i < 11; i++{
            fmt.Scan(&digit)
            digits[i] = digit
            soma += digit
        }
        
        for i := 0; i < 9; i++{
            digits_2[i] = digits[i]*(i+1)
            soma_2 += digits_2[i]
            digits_3[i] = digits[i]*(9-i)
            soma_3 += digits_3[i]
        }
        
        if(soma % 11 != 0) {
            valid = false
        }
        //Décimo elemento
        if(digits[9] != soma_2 % 11) {
            valid = false
        }
        //Último elemento (décimo primeiro)
        if(digits[10] != soma_3 % 11) {
            valid = false
        }
        
        
        if(valid) {
            fmt.Println("CPF valido")
        }else {
            fmt.Println("CPF invalido")
        }
        //Reinicializando tudo
        valid = true
        digits = null
        for i := 0; i < 9;i++{
            digits_2[i] = 0
            digits_3[i] = 0
        }
        soma = 0
        soma_2 = 0
        soma_3 = 0
        
    }
    
    
}