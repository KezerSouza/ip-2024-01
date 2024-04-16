/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main
import (
    "fmt"
    "math"
)

func main() {
    var altura, aresta, Ab, volume float64
    
    fmt.Scan(&altura, &aresta)
    
    Ab = (3*aresta*aresta*math.Sqrt(3))/2
    volume = Ab*altura*1/3
    
    fmt.Println(volume)
    
    
}