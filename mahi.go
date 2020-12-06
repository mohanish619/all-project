package main
import "fmt"
func main(){
	x,y := 10,5
	fmt.Println(calc(x,y))
}
func calc(num,num2 int)(int,int){
	var out = num + num2
	var outt = num - num2
	return out, outt
}




