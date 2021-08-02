package teste

import (
	"fmt"
	"sync"
)

func main(){

	numeros := []int{1,2,3,4}

	var wg sync.WaitGroup
	wg.Add(1)
	
	go teste(&wg, numeros...)
	
	wg.Wait()
}

func teste(wg *sync.WaitGroup, x ...int){
	defer wg.Done()
	for i := range x{
		fmt.Println("numero:", x[i])
		if  y := x[i]%2; y == 0{
			fmt.Println("Par: ", x[i])
			
		} else{
			fmt.Println("Impar: ", x[i])
		}
	}
}
