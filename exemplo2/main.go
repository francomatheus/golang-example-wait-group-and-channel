package main

import (
	"fmt"
)

func main(){

	numeros := [10]int{1,2,3,4,5,6,7,8,9,10}
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	
	go teste(par, impar, quit, numeros)
	go print(par, impar, quit)

	<-quit
	
}

func teste(par chan int, impar chan int, quit chan bool, x [10]int){
	for i := range x{

		fmt.Println("numero:", x[i])
		if  y := x[i]%2; y == 0{
			par <- x[i]
			
		} else{
			impar <- x[i]
		}
	}
	quit <- true
}

func print(par chan int, impar chan int, quit chan bool){
	for {
		select{
		case x := <-par:
			fmt.Println("Par: ", x)
		case y := <-impar:
			fmt.Println("Impar: ", y)
		case <-quit:
			return
		}
	}
}