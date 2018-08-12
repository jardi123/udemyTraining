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
