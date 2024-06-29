package main
import(
    "fmt"
    "bufio"
    "os"
    "strings"
)


func main() {
    in := bufio.NewReader(os.Stdin)
    
    for  {
        line, _ := in.ReadString('\n')
        line = strings.TrimSpace(line)

		if line == "#" {
			break
		}else{
		    for i := len(line) - 1; i >= 0; i-- {
			    fmt.Printf("%c", line[i])
		    }
		    fmt.Print("\n")
		}

    }
    
}
