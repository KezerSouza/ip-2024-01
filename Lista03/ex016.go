package main
import "fmt"

func main() {
    var N, K, horario, presenças int
    var alunos, indexes []int
    
    fmt.Scan(&N, &K) //n° alunos e mínimo de presenças
    
    //Slice alunos, com os tempos de chegada deles (var horario)
    for i := 0; i < N;i++ {
        fmt.Scan(&horario)
        alunos = append(alunos, horario)
    }
    
    //Guardando a posição dos alunos que chegaram no horário
    for i := 0; i < N;i++ {
        if(alunos[i] <= 0) {
            presenças += 1 //Para cada aluno que chegou no horário, aumenta-se o contador de número de presenças
            indexes = append(indexes, i+1) //"+1" por conta de o index de um slice começar em 0, mas os alunos começa-se a contar a partir do 1 | Aluno1, Aluno2,...
        }
        
    }
    
    //Se poucas presenças, aula cancelada
    if(presenças < K) {
        fmt.Println("SIM")
    //Se não, imprimir a posição dos alunos, em ordem inversa
    }else {
        fmt.Println("NÃO")
        for i := len(indexes)-1; i >= 0;i-- {
            fmt.Println(indexes[i])
        }
        
    }
    
}