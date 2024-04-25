package main
import "fmt"

func main() {
    var valorIngresso, valorInicial, valorFinal, valorAtual float64
    var vezes int
    var vendas, lucro float64
    var melhorValor, melhorLucro, melhorVendas float64
    
    fmt.Scan(&valorIngresso, &valorInicial, &valorFinal)
    
    if(valorFinal <= valorInicial) {
        fmt.Printf("Dígito Inválido")
        
    }else {
        vezes_f := valorFinal - valorInicial + 1
        vezes = int(vezes_f)
        valorAtual = valorInicial
        
        for i := 0; i < vezes;i++ {
            if(valorAtual < valorIngresso) {
                vendas = 120 + (valorIngresso-valorAtual)*50 //Quanto diminuiu vezes 50
            }
            if(valorAtual > valorIngresso) {
                vendas = 120 - (valorAtual-valorIngresso)*60
            }
            if(valorAtual == valorIngresso) {
                vendas = 120
            }
            lucro = vendas*valorAtual - (200+0.05*vendas)
            if(melhorLucro < lucro) {
                melhorValor = valorAtual
                melhorLucro = lucro
                melhorVendas = vendas
            }
            fmt.Printf("V: %v, N: %v, L: %v\n", valorAtual, vendas, lucro)
            valorAtual += 1
            
        }
        fmt.Printf("Melhor valor final: %0.2f\nLucro: %0.2f\nNúmero de ingressos: %v", melhorValor, melhorLucro, melhorVendas)
        
    }
    
}