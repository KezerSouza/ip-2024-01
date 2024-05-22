package main
import "fmt"
import "sort"

func main() {
    var T, n int
    var odd, even []int
    
    fmt.Scan(&T)
    
    for i := 0; i < T;i++ {
        fmt.Scan(&n)
        
        if(n % 2 == 0) {
            even = append(even, n)
        }else{
            odd = append(odd, n)
        }
        
    }
    
    sort.Ints(even)
    sort.Sort(sort.Reverse(sort.IntSlice(odd)))
    
    for i := 0; i < len(even)-1;i++ {
        fmt.Println(even[i])
    }
    
    for i := 0;i < len(odd)-1;i++ {
        fmt.Println(odd[i])
    }
    
}