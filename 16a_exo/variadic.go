package main

import "fmt"

func main(){
	lop := claudia(23,20,10,30,40)
	fmt.Println(lop)
	fmt.Println((true && false) || (false && true) || !(false && false))
}



func claudia(sf ...int) int{
	var largest int
	for _, v := range sf{
		if v > largest{
			largest = v
		}

	}
	return largest
}




