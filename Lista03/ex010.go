package main
import "fmt"

func main() {
    var tests, x, last_num, max_num, times, indice int
    var numbers []int
    
    fmt.Scan(&tests)
    
    //Preenchendo o slice com os números
    for i := 0;i < tests;i++ {
        fmt.Scan(&x)
        numbers = append(numbers, x)
    }
    
    //Último número
    last_num = numbers[tests-1]
    
    //Passada por todos os números
    for i := 0;i<len(numbers);i++ {
        
        //Qual será o maior número
        if(numbers[i] > max_num) {
            max_num = numbers[i]
            indice = i
        }
        
        //Quantas vezes o último número aparece?
        if(numbers[i] == last_num) {
            times += 1
        }
        
    }
    
    fmt.Printf("Nota %v, %v vezes\n", last_num, times)
    fmt.Printf("Nota %v, indice %v", max_num, indice)

}