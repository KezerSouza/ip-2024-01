package main
import "fmt"

func main() {
	var a, b float64
    var contador int
	
	fmt.Scan(&a, &b)
	
	for i := 1; a < b; i++{
	    a += a*0.03
	    b += b*0.015
	    contador = i
	}
	fmt.Printf("ANOS = %v", contador)
    
}