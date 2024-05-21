package main
import "fmt"

func main() {
    var n, number, added int
    var vector, vector2 []int
    var add bool = true
    
    fmt.Scan(&n)
    
    for i := 0; i < n;i++ {
        fmt.Scan(&number)
        vector = append(vector, number)
    }
    fmt.Println(vector)
    
    for i := 0; i < len(vector);i++ {
        for j := 0; j < len(vector);j++ {
            if(vector[i] != added) {
                add = true
                added = vector[i]
            }
            
        }
        
        if(add == true) {
            vector2 = append(vector2, vector[i])
        }
        add = false
        
    }
    
    
    fmt.Println(vector2)
}