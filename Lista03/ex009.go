package main
import "fmt"
import "math"

func main() {
    var tests int
    var x1, y1, z1, x2, y2, z2, distance float64
    
    fmt.Scan(&tests)
    fmt.Scan(&x1, &y1, &z1)
    
    for i := 1;i < tests;i++ {
        fmt.Scan(&x2, &y2, &z2)
        
        distance = math.Abs((x1-x2)*(x1-x2)) + math.Abs((y1-y2)*(y1-y2)) + math.Abs((z1-z2)*(z1-z2))
        distance = math.Sqrt(distance)
        fmt.Printf("%0.2f\n", distance)
        
        x1, y1, z1 = x2, y2, z2
    }

}