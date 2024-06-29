package main
import(
    "fmt"
    "bufio"
    "os"
)


func main() {
    in := bufio.NewReader(os.Stdin)
    line, _ := in.ReadString('\n')
    
    for line != "#" {
        for i := len(line)-1; i >= 0; i-- {
            fmt.Printf("%c", line[i])
        }
        fmt.Print("\n")
        line, _ = in.ReadString('\n')
    }
    
}