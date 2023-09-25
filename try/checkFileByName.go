package main
import (
	"fmt"
    "log"
    "os"
)
func main() {
    fileInfo, error := os.Stat("create.go")

    if error != nil {
        log.Fatal(error)
    } else{
		fmt.Println(fileInfo)
	}
}