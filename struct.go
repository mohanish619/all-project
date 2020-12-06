package main

import "fmt"
type Student struct{
	Rollno int
	name string
	marks int
}
func main() {
	var s1 = Student{64,"mohanish",64}
	fmt.Println(s1)
	fmt.Println(s1.name)
}
