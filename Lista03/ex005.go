package main
import "fmt"

func main() {
    var n, value, max, max_index int
    
    fmt.Scan(&n)
    
    //Executar enquanto tiver casos de teste
    for n != 0 {
        
        //Verificar número máximo da sequência e seu index
        for i := 0; i < n;i++ {
            fmt.Scan(&value)
            //Se valor atual for o maior, seu valor é atribuído a max
            if(value > max) {
                max = value
                max_index = i
            }
            
        }
        
        fmt.Printf("%v %v\n", max_index, max)
        max_index, max, value = 0, 0, 0 //Zerando as variáveis para o próximo de teste
        
        //Próximo caso de teste
        fmt.Scan(&n)
        
    }
    
}