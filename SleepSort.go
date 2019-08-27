package main

import "fmt"
import "time"

func sleepSort(num int){
	time.Sleep(time.Second * time.Duration(num))
	fmt.Print(num, "\t")
}

func main(){
	nums := []int{2, 5, 3, 4, 1, 7, 6}

	for _, num := range nums {
		go sleepSort(num)
	}

	fmt.Scanln()
	fmt.Println("done")
}
