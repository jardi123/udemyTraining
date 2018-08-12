package main

import "fmt"

func devoir(fag int)(float64,bool){
	return float64(fag)/2, fag%2==0
}

func main(){
	numero, valeur := devoir(7)
	fmt.Println(numero, valeur)
}