package main

import (
	"fmt"
	"time"
)

func worker(done chan bool){
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main(){
	messages := make(chan string)

	go func() {
		messages <- "ping"
		messages <- "192.168.1.1"
	}()

	go func() {
		fmt.Println(<-messages)
	}()

	msg := <-messages
	fmt.Println(msg, "111")

	fmt.Println("\n ---- 通道缓冲 ---- \n")

	// 可缓存通道允许在没有对应接收方的情况下，缓存限定数量的值
	messages1 := make(chan string, 2)

	messages1 <- "buffered"
	messages1 <- "channel"

	fmt.Println(<-messages1)
	fmt.Println(<-messages1)


	fmt.Println("\n 通道同步 \n")

	fmt.Println(time.Second)

	done := make(chan bool, 1)
	go worker(done)

	<-done
	fmt.Println("阻塞结束了")

	fmt.Println("\n -----------通道方向-------- \n")
}
