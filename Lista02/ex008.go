package main
import "fmt"

func main() {
    var n int
    final := 1
    var times []string
    fmt.Scan(&n)
    
    for i := 1; i <= n;i++ {
        times = append(times, fmt.Sprintf("Time%v", i))
        
    }
    
    for i := 0; i < n;i++ {
        for j := 0; j < n;j++ {
            if(i != j && j > i) {
                fmt.Printf("Final %v: %v X %v\n", final, times[i], times[j])
                final += 1
            }
        }
    }
    
}