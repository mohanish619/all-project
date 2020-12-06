package main
import "fmt"
func main() {
	defer firstrun()
	secondrun()
defer thirdrun()
}
func firstrun(){
	fmt.Println("i execute first")
}
func secondrun(){
	fmt.Println("i execute second")
}
func thirdrun(){
	fmt.Println("i execute third")
}
