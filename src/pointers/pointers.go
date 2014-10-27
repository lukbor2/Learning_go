package main 

import (
	"fmt"
)

func swapAndProduct1 (i, j, prod *int) {
	temp := *i
	*i = *j
	*j = temp
	*prod = *i * *j 
}

//Usually we would not really use pointers to get the same result as swapAndProduct1
func swapAndProduct2 (x, y int) (int, int, int){
	x , y = y , x
	return x, y, x*y
	
}

func main() {
	i:= 9
	j:= 5
	product := 0
	swapAndProduct1(&i, &j, &product) //NOTE we pass &i, &j, &product and the function's parameters are *int
	fmt.Println("First Function: ", i, j, product)
	
	i, j, product = swapAndProduct2(i,j)
	fmt.Println("Second Function: ", i, j, product)
}

