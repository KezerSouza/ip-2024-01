package main
import "fmt"

func main() {
	var n, resultado int
	
	fmt.Scan(&n)
    resultado = n
    
    for i := n-1; i > 1; i-- {
        resultado = resultado*i
    }
	
	fmt.Printf("%v! = %v", n, resultado)
    
}