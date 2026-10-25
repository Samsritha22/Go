package main
import "fmt"
type contact struct {
Name string
Email string
Message string
}
func main() {
var contact contact
fmt.Print("Enter Your Name:")
fmt.Scan(&contact.Name)
fmt.Print("Enter Your Email:")
fmt.Scan(&contact.Email)
fmt.Print("Enter Your Message:")
fmt.Scan(&contact.Message)
fmt.Print("Contact Details:")
fmt.Println("Name:",contact.Name)
fmt.Println("Email:",contact.Email)
fmt.Println("Message:",contact.Message)
}