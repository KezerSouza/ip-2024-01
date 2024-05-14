package main
import "fmt"

func main() {
    var n, x, max_num int //n => quantidade de números a serem atribuídos || x => número a ser atribuído no momento
    var numbers, count []int
    
    fmt.Scan(&n)
    
    //Preenchimento do vetor numbers
    for i := 0; i < n;i++ {
        fmt.Scan(&x)
        numbers = append(numbers, x)
    }
    
    //Buscando pelo maior valor do vetor
    for i := 0; i < len(numbers);i++ {
        if(numbers[i] > max_num) {
            max_num = numbers[i]
        }
        
    }
    
    //Agora tem-se um novo vetor "count", com índices de 0 até o maior valor do vetor "numbers"
    for i := 0; i <= max_num;i++ {
        count = append(count, 0) //Tamanho vetor
        //Colocando os valores nos índices || A frequência de um número será o valor em "count" que tem por índice esse número
        //Exemplo: Se quero a frequência do número 3, a reposta está em count[3]
        for j := 0; j < len(numbers);j++ {
            if(numbers[j] == i) {
                count[i] += 1
            }
            
        }
        
    }
    
    value := 0
    frequency := 0
    //Uma busca no vetor count para saber qual número possui a maior frequência
    for i := 0; i < len(count);i++ {
        if(count[i] > frequency) {
            value = i
            frequency = count[i]
        }
        
    }
    
    fmt.Printf("%v\n%v", value, frequency)
}