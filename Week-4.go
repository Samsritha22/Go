

package main
import (
"fmt"
"math"
)
func main() {
var nums = []float64{10, 20, 30, 40, 50}
var sum, mean, variance, stdDev float64
var i int
for i=0;i<len(nums);i++ {
sum += nums[i]
}
mean=sum/float64(len(nums))
for i=0;i<len(nums);i++{
variance += math.Pow(nums[i]-mean,2)
}
variance = variance/float64(len(nums))
stdDev = math.Sqrt(variance)
fmt.Println("Mean=",mean)
fmt.Println("Variance=",variance)
fmt.Println("Standard Deviation=",stdDev)
}