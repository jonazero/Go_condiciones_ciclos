package main
import "fmt"
func main(){
	var e float64 = 1
	var aux float64 = 1
	for i:=1; i<10; i++{
		aux = aux/float64(i)
		e = e + aux
		fmt.Println(e)
	}
}