package main

import "fmt"

func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t%b \t %#x \t %q \n", i, i, i, i)
	}
}
// %d decimal %b binary %#x hexadecimal %q utf-8