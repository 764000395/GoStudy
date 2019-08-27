package main

import "fmt"

func plus(a int, b int) int{
	return a + b
}

//多返回值函数
func vals() (int, int){
	return 3, 7
}

//变参函数
func sum(nums ...int){
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// 闭包
func intSeq() func() int{
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main(){
	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

	fmt.Println("--- 闭包 ---")
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())


	fmt.Println(" ----- ")

	newInts := intSeq()
	fmt.Println(newInts())
}
