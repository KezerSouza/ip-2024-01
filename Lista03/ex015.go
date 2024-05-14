package main
import "fmt"
import "math"

func main() {
    var T, n, x, comparisons int
    var diference float64
    
    fmt.Scan(&T)
    for i := 0;i < T;i++{
        
        fmt.Scan(&n)
        
        //Declaração do vetor numbers
        numbers := make([]int, n)
        for i := 0; i < n;i++ {
            fmt.Scan(&x)
            numbers[i] = x
        }
        
        final_diference := math.Abs(float64(numbers[0]-numbers[1])) //Declaração da resposta com a menor diferença, iniciando com a diferença entre 1° e 2° valor

        //Compara atual ao próximo, depois o próximo com o próximo do próximo
        for i := 0; i < len(numbers);i++ {
            for j := i+1; j < len(numbers);j++ {
                diference = math.Abs(float64(numbers[i]-numbers[j]))
                if(diference < final_diference) {
                    final_diference = diference
                }
                comparisons++
                
            }
        }
        fmt.Println(final_diference, comparisons)
        comparisons = 0
    }
    
}