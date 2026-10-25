package main
import "fmt"
func main() {
var str string
var left, right, n int
fmt.Print("Enter a String:")
fmt.Scan(&str)
n=len(str)
left=0
right=n-1
for left<=right {
if str[left] != str[right] {
fmt.Println("Not a Palindrome")
return
}
left++
right--
}
fmt.Println("Palindrome")
}