package main
import "fmt"

func main(){
	var x int
	var mem bool
	
	fmt.Scan(&x,&mem)
	fmt.Print(diskon(x,mem))
}

func diskon(x int, mem bool) int {
	if x > 100000 && mem {
		return x - ((x * 10)/100)
	} else if x > 100000 && !mem {
		return x - ((x*5)/100)
	} else {
		return x
	}
}