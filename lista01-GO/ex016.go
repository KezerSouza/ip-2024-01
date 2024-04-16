/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main
import "fmt"

func main() {
    var salario, resultado float64
    
    fmt.Scan(&salario)
    
    if (salario <= 300) {
        resultado = salario*1.5
    }else if(salario > 300) {
        resultado = salario*1.3
    }
    fmt.Printf("SALARIO COM REAJUSTE = %0.2f", resultado)
    
}