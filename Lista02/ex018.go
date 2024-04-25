package main
import "fmt"
import "math"

func main() {
    var n1, n2, n3, pn1, pn2, pn3, multiplo int
    mmc := 1
    
    fmt.Scan(&n1, &n2, &n3)
    
    
    for n1 != 1 && n2 != 1 && n3 != 1 {
        max_num := math.Max(float64(n1), math.Max(float64(n2), float64(n3)))
        for i := 2; i <= int(max_num);i++ {
            if(n1 % i == 0 || n2 % i == 0 || n3 % i == 0) {
                pn1, pn2, pn3 = n1, n2, n3
               if(n1 % i == 0) {
                   pn1 = n1/i
                   multiplo = i
               }
               if(n2 % i == 0) {
                   pn2 = n2/i
                   multiplo = i
               }
               if(n3 % i == 0) {
                   pn3 = n3/i
                   multiplo = i
               }
               mmc = mmc*multiplo
               fmt.Printf("%v %v %v :%v\n", n1, n2, n3, multiplo)
               n1, n2, n3 = pn1, pn2, pn3
            }

        }
    }
    fmt.Printf("MMC: %v", mmc)
}