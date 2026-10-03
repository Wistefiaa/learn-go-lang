package function

import "fmt"

func Hello(name string) {
	fmt.Println(name)
}
func Add(a int, b int) int /*return type*/ {
	return a + b
}
func main() {
	fmt.Println("Hello CEDT")
}
