package main

import (
	"fmt"
)

func f(from string){
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main(){
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	_, _ = fmt.Scanln(&input)

	// 如果把上边按任意键结束注释，创建的协程不会进行输出
	// 是因为主进程结束了，协程被迫结束
	//time.Sleep(10)
	fmt.Println("done")
}
