/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main
import "fmt"

func main() {
    var x, y int
    
    fmt.Scan(&x, &y)
    
    if (x%2 == 0) {
        fmt.Printf("%v ", x)
        for y > 1 {
            x += 2
            fmt.Printf("%v ", x)
            y--
        }
    }else {
        fmt.Println("O PRIMEIRO NUMERO NAO E PAR")
    }
    
}