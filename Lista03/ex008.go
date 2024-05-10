package main

import "fmt"
import "os"
import "bufio"
import "strconv"

func main() {
    //Declaração de um slice numbers onde ficará os algarismos do número binário, e um slice vazio para resetar o numbers 
    var numbers, null []int64
    
    scanner := bufio.NewScanner(os.Stdin)
    
    //Looping para executar enquanto tiver linhas para ler na entrada
    for scanner.Scan() {
        //Conversão do valor de entrada para inteiro
        value, _ := strconv.ParseInt(scanner.Text(), 10, 64)
        
        //Se o valor de entrada for 0 já encerra aqui, respota = 0
        if(value == 0) {
            fmt.Println("0")
        //Se não for 0:
        }else {
            //Ficar dividindo o valor por 2 enquanto puder ser dividido
            for (value >= 1) {
                //Se for par
                if(value % 2 == 0) {
                    numbers = append(numbers, 0)
                    value = value/2
                }
                //Se for ímpar
                if(value % 2 != 0) {
                    numbers = append(numbers, 1)
                    value = value/2
                }
            
            }
            
            //Imprimir os algarismos, mas em ordem reversa
            for i := len(numbers)-1;i >= 0;i-- {
                fmt.Printf("%v", numbers[i])
            }
            
            //Resetando o slice numbers
            numbers = null
            
            fmt.Printf("\n")
            
        }
        
    }

}