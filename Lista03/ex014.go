package main
import "fmt"
import "sort"

func main() {
    var n, x int
    var mediana float64
    var numbers []int
    
    fmt.Scan(&n)
    
    for i := 0; i < n;i++ {
        fmt.Scan(&x)
        numbers = append(numbers, x)
    }
    
    sort.Ints(numbers) //Ordenação dos números
    
    //Mediana e conversão para float64
    if(len(numbers) % 2 == 0) {
        mediana = float64(numbers[len(numbers)/2] + numbers[len(numbers)/2-1])/2 //Se não tiver valor central, mediana é a média dos dois centrais
    }else { 
        mediana = float64(numbers[len(numbers)/2]) //Se tiver valor central, mediana é ele
    }
    
    fmt.Printf("%0.2f", mediana)
    
}