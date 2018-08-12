//Improved code proposed by Jaden Weis

package main

import "fmt"

func main() {
	fmt.Println(FibEuleur())
}

func FibEuleur() int {
	sum := 0
	for cur, prev := 1, 1; cur <= 4000000; cur, prev = cur+prev, cur {
		// it sets cur to cur + prev and prev to cur at the same time
		if cur%2 == 0 {
			sum += cur
		}
	}
	return sum
}

/*
i%2== 0 is the same as i&1==0 (bitwise and)

Because

1 in binary as a uint8 is `00000001`
2 in binary as a uint8 is `00000010`
3 in binary as a uint8 is `00000011`

so binary wise any odd number has a 1 in first place

i&1 is faster to execute however it can hurt readability



*/
