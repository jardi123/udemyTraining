package main

import "fmt"

func claudy (fag int) (int, bool) {
	return fag/2, fag%2 == 0
}

func main() {
	h,even := claudy(7.0)
	fmt.Print(h,even)
}