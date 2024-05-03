package main
import "fmt"

func main() {
    var vector_size, number_of_searches int
    var  vector_value, searches_value int
    var vector,  searches []int
    var achei = false
    
    fmt.Scan(&vector_size) //Primeira linha, quantos valores
    
    //Colocando os valores de entrada no vetor
    for i := 0; i < vector_size;i++ {
        fmt.Scan(&vector_value)
        vector = append(vector, vector_value)
    }
    
    //Quantas buscas serão feitas?
    fmt.Scan(&number_of_searches)
    
    //Colocando os valores de busca na slice searches
    for i := 0; i < number_of_searches;i++ {
        fmt.Scan(&searches_value)
        searches = append(searches, searches_value)
    }
    
    //Comparando as duas searches com o vetor
    for i := 0; i < number_of_searches;i++ {
        for j := 0; j < vector_size;j++ { //Looping duplo, para cada elemento de searches buscar por todo vetor
            if(searches[i] == vector[j]) {
                achei = true
            }
        }
        if(achei == true) {
            fmt.Printf("ACHEI\n")
        }else {
            fmt.Printf("NÃO ACHEI\n")
        }
        
        achei = false
        
    }
    
}