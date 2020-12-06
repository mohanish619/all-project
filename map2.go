package main
import "fmt"
func main() {
	StudentAge := make(map[string]int)

	StudentAge["mohanish"] = 23
	StudentAge["shrikar"] = 23
	StudentAge["chintu"] = 24
	StudentAge["shrinand"] = 22
	StudentAge["mahi"] = 21
	fmt.Println(StudentAge["mohanish"])

}