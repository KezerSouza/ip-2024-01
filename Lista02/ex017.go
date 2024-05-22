package main
import "fmt"

func main() {
    var produtos []int
    var produto int = 1
    var unidades, lucros []float64
    var lucro, custo_total, vendas_total, lucro_total float64
    
    var preço, custo, unidade float64
    
    for produto != 0 {
        ant := produto
        fmt.Scan(&produto, &custo, &preço, &unidade)
        if(ant == produto) {
            break
        }
        custo_total += custo*unidade
        vendas_total += preço*unidade
        lucro = preço*unidade - custo*unidade
        lucro = lucro/(custo*unidade)
        produtos = append(produtos, produto)
        unidades = append(unidades, unidade)
        lucros = append(lucros, lucro)
        
    }
    matricula_index := 0 
    unidade_index := 0
    menor_10 := 0
    maior_10 := 0
    maior_20 := 0
    for i := 0; i < len(produtos);i++ {
        //Maior lucro?
        if(i != 0) {
            if(lucros[i] > lucros[matricula_index]) {
                matricula_index = i 
            }  
        }else {
            matricula_index = 0
        }
        //Mais vendas?
        if(i != 0) {
            if(unidades[i] > unidades[i-1]) {
                unidade_index = i
            }
        }else {
            unidade_index = 0
        }
        //Quantidade de mercadorias que geraram lucro x
        if(lucros[i] < 0.1) {
            menor_10 += 1
        }else if(lucros[i] >= 0.1 && lucros[i] <= 0.2) {
            maior_10 += 1
        }else if(lucros[i] > 0.2) {
            maior_20 += 1
        }
        
    }
    
    lucro_total = vendas_total-custo_total
    lucro_total = lucro_total/custo_total
    
    fmt.Println(lucros)
    
    fmt.Printf("Quantidade de mercadorias que geraram lucro menor que 10%%: %v\n", menor_10)
    fmt.Printf("Quantidade de mercadorias que geraram lucro maior ou igual a 10%% e menor ou igual a 20%%: %v\n", maior_10)
    fmt.Printf("Quantidade de mercadorias que geraram lucro maior do que 20%%: %v\n", maior_20)
    fmt.Printf("Código da mercadoria que gerou maior lucro: %v\n", produtos[matricula_index])
    fmt.Printf("Codigo da mercadoria mais vendida: %v\n", produtos[unidade_index])
    fmt.Printf("Valor total de compras: %0.2f, valor total de vendas: %0.2f e percentual de lucro total: %0.2f%s", custo_total, vendas_total, lucro_total*100, "%")
    
    
}
