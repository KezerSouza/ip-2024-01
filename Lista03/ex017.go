package main
import "fmt"

func main() {
    var n, number int
    var numbers, unique_numbers []int
    var add bool = true
    
    fmt.Scan(&n)
    
    //Número de elementos
    for i := 0; i < n; i++{
        fmt.Scan(&number)
        numbers = append(numbers, number)
    }
    
    //Looping duplo, comparação de cada elemento com todos os outros
    for i := 0; i < len(numbers); i++{
        for j := 0; j < len(numbers); j++{
            //Verificar se tem algum número igual || Não comparar com o próprio número (i != j)
            if(i != j) {
                if(numbers[i] == numbers[j]) {
                    //Se um número é igual a outro, variável add = false
                    add = false
                }
                
            }
            
        }
        //Só adicionará o número se add for true, ou seja, se ele for um elemento único
        if(add == true) {
            unique_numbers = append(unique_numbers, numbers[i])
        }
        //Reinicializando a variável add para true, para que se possa comparar a partir de um novo elemento (i)
        add = true
        
    }
    
    //Saída, quantidade de números no vetor de elementos únicos
    fmt.Println(len(unique_numbers))
    
}