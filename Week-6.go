package main
import "fmt"
func main() {
var str1, str2, result string
fmt.Print("Enter the str1:")
fmt.Scan(&str1)
fmt.Print("Enter the str2:")
fmt.Scan(&str2)
result=str1+str2
fmt.Print(result)
}