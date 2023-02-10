package slice

import "fmt"

func test(slice3 []int) {
	slice3[0] = 100
}

func slice_4() {

	var a []int = []int{1, 2, 3, 4, 5}
	var slice = make([]int, 1)
	fmt.Println(slice) //[0]
	copy(slice, a)
	fmt.Println(slice) //[1]

	fmt.Println()

	//切片是引用类型，在传递时，遵守引用传递机制
	var slice1 []int
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	slice1 = arr[:]
	var slice2 = slice1
	slice2[0] = 10

	fmt.Println("slice2", slice2) //[10,2,3,4,5]
	fmt.Println("slice1", slice1) //[10,2,3,4,5]
	fmt.Println("arr", arr)       //[10,2,3,4,5]

	fmt.Println()

	var slice3 = []int{1, 2, 3, 4}
	fmt.Println("slice3=", slice3) //[1,2,3,4]
	test(slice3)
	fmt.Println("slice3=", slice3) //[100,2,3,4]
}
