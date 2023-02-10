package array

import "fmt"

// func test01(arr [3]int) {
// 	arr[0] = 40
// }

func test02(arr *[3]int) {
	(*arr)[0] = 99
}

func array_7() {

	/*
		//数组是多个相同类型数据的组合，一个数组一旦声明/定义了，其长度是固定的，不能动态变化
		var arr01 [3]int
		arr01[0] = 1
		arr01[1] = 2
		arr01[2] = 3
		arr03[2] = 3.3 //这里会报错,因为数组的类型是int，就不能给float类型
		arr04[3] = 80  //长度是固定的，不能动态变化，否则报越界
		fmt.Println(arr01)
	*/

	//数组创建后，如果没有赋值，有默认值
	var arr01 [3]int
	var arr02 [3]float32
	var arr03 [3]string
	var arr04 [3]bool
	fmt.Printf("arr01=%v arr02=%v arr03=%v arr04=%v\n", arr01, arr02, arr03, arr04)

	//数组属值类型，默认情况下是值传递，因此会进行值拷贝
	arr := [3]int{10, 20, 30}
	//test01(arr)
	test02(&arr)
	//fmt.Println("main arr=", arr)
	fmt.Println("main arr=", arr)

}
