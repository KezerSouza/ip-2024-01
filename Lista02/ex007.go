package main
import "fmt"

func main() {
	var numbers []float64
	var mp, mi float64
	var pares, impares float64
	var continuar = true
	
	//Colocando os números numa sequência 
	for i := 0; continuar == true; i++ {
	    var n float64
	    fmt.Scan(&n)
	    if(n != 0) {
	        numbers = append(numbers, n)
	        //Se par
	        if(int(numbers[i]) % 2 == 0) {
	            mp += numbers[i]
	            pares += 1
	        }else {
	            mi += numbers[i]
	            impares += 1
	        }
	        
	    }else {
	        continuar = false
	    }
	    
	}
	mp = mp/pares
	mi = mi/impares
	
	//média par
	
	fmt.Printf("MEDIA PAR: %0.2f\n", mp)
	fmt.Printf("MEDIA IMPAR: %0.2f", mi)
	
}