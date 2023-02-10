package array

import "fmt"

func array_5() {
	//从终端循环输入5个成绩，保存到float64数组，并输出
	// var score [5]float64

	// for i := 0; i < len(score); i++ {
	// 	fmt.Printf("请输入第%d个元素的值\n", i+1)
	// 	fmt.Scanln(&score[i])
	// }

	// for i := 0; i < len(score); i++ {
	// 	fmt.Printf("score[%d]=%v\n", i, score[i])
	// }

	//数组初始化方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01=", numArr01)

	var numArr02 = [3]int{5, 6, 7}
	fmt.Println("numArr02", numArr02)
	//这里的{...}是规定的写法
	var numArr03 = [...]int{8, 9, 10}
	fmt.Println("numArr03=", numArr03)

	var numArr04 = [...]int{1: 800, 0: 900, 2: 888}
	fmt.Println("numArr04=", numArr04)

	strArr04 := [...]string{1: "tom", 0: "jim", 2: "json"}
	fmt.Println("strArr04=", strArr04)

}
