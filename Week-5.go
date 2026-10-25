package main
import "fmt"
func main() {
var rows int
num:= 1
var i, j int
fmt.Print("Enter number of rows:")
fmt.Scan(&rows)
for i=0;i<rows;i++ {
for j=0;j<=i;j++ {
fmt.Print(num)
num++
}
fmt.Println()
}
}