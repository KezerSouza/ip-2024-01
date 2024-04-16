/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main
import "fmt"

func main() {
    var n int
    
    fmt.Scan(&n)
    
    for i := 2; i <= n; i += 2 {
        fmt.Printf("%v^2 = %v \n",i, i*i)
    }
    
}