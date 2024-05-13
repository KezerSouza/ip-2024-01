package main
import "fmt"

func main() {
    var vector, inverted_vector []int
    var vector_size, n, max_num int
    
    fmt.Scan(&vector_size)
    
    //Preenchendo o slice com os números
    for i := 0;i < vector_size;i++ {
        fmt.Scan(&n)
        vector = append(vector, n)
    }
    
    //Invertendo vetor
    for i := vector_size-1;i >= 0;i-- {
        inverted_vector = append(inverted_vector, vector[i])
    }
    
    min_num := vector[0]
    //Passada por todos os números
    for i := 0;i<len(vector);i++ {
        
        //Qual será o maior número
        if(vector[i] > max_num) {
            max_num = vector[i]
        }
        
        //Qual será o menor número
        if(vector[i] < min_num) {
            min_num = vector[i]
        }
        
        //Imprimindo os vetores
        fmt.Printf("%v ", vector[i])
    }
    fmt.Printf("\n")
    for i:= 0;i<len(vector);i++ {
        fmt.Printf("%v ", inverted_vector[i])
    }
    
    //Imprimindo valor mínimo e máximo
    fmt.Printf("\n%v\n%v\n", max_num, min_num)
    
}