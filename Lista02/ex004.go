package main
import "fmt"

func main() {
	var n, i, s, B float64
	var k int
	
	fmt.Scan(&n, &i, &k, &s)
    
    
    fmt.Println("Tabuada de soma:")
    for c := 0; c < k; c++ {
        B = i+float64(c)*s
        
        fmt.Printf("%v + %0.2f = %0.2f\n", n, B, n+B)
    }
    
    
    fmt.Println("Tabuada de subtração:")
    for c := 0; c < k; c++ {
        B = i+float64(c)*s

        fmt.Printf("%v - %0.2f = %0.2f\n", n, B, n-B)
    }
    
    
    fmt.Println("Tabuada de multiplicação:")
    for c := 0; c < k; c++ {
        B = i+float64(c)*s

        fmt.Printf("%v * %0.2f = %0.2f\n", n, B, n*B)
    }
    
    
    fmt.Println("Tabuada de divisão")
    for c := 0; c < k; c++ {
        B = i+float64(c)*s

        fmt.Printf("%v / %0.2f = %0.2f\n", n, B, n/B)
    }
    
    
}