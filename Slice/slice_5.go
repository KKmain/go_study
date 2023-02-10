package slice

import "fmt"

func slice_5() {
	//使用常规的for循环遍历切片
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4] //20,30,40
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v\n", i, slice[i])
	}

	fmt.Println()

	//使用for-range 遍历切片
	for i, v := range slice {
		fmt.Printf("i=%v v=%v \n", i, v)
	}

	slice2 := slice[1:2] //   slice[20,30,40]  [30]
	slice2[0] = 100      // 因为arr slice和slice2 指向的数据空间是同一个，因此slice2[0]=100
	fmt.Println("slice2=", slice2)
	fmt.Println("slice=", slice)
	fmt.Println("arr=", arr)

	// fmt.Println(slice[3]) 只能取到slice的 下标为 0 1 2  也就是 20 30 40
}

/*
slice是一个引用类型
slice从底层来说，其实就是一个数据结构 struct结构体
type slice struct{
    ptr *[2]int //指针
    len int  //切片的长度
    cap int  //切片的容量
}
*/
