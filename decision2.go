package main

import "fmt"

func main() {
	age := 18
	switch age{
	case 16: fmt.Println("prepare for clg")
	case 18: fmt.Println("don't run after the girls")
	case 20: fmt.Println("get yourself a job")
	default: fmt.Println("are you even alive")
	}
}
