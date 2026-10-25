package main
import (
"fmt"
"sort"
)
func main() {
number:=[]int{2,4,7,1}
sort.Sort(sort.Reverse(sort.IntSlice(number)))
fmt.Println("Integers in descending order:",number)
names:=[]string{"Ravi","Anu","Sita","Kiran"}
sort.Sort(sort.Reverse(sort.Reverse(sort.StringSlice(names))))
fmt.Println("Strings in reverse alphabetical order:",names)
}