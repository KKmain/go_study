package slice

import "fmt"

func slice_2() {
	//方式2
	//演示切片的使用make
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 20
	fmt.Println(slice)
	fmt.Println("slice的size=", len(slice))
	fmt.Println("slice的cap=", cap(slice))

	fmt.Println()
	//方式3
	//定义一个切片，直接指定具体数组，使用原理类似make方式
	var strSlice []string = []string{"tom", "jack", "mary"}
	fmt.Println("strSlice=", strSlice)
	fmt.Println("strSlice siza=", len(strSlice))
	fmt.Println("strSlice cap=", cap(strSlice))

}

/*
方式1
定义一个切片，然后让切片去引用一个创建好的数组，切片基本介绍已经演示


方式2
通过make来创建切片
基本语法： var 切片名[]type=make([]type,len,[cap])
参数说明：type就是数据类型 len：大小  cap：指定切片容量(可选，如果分配了cap，则要求cap>=len)

总结：
1.通过make方式创建切片可以指定切片的大小和容量
2.如果没有给切片的各个元素赋值，那么就会使用默认值[int,float=>0 string=>"" bool=>false]
3.通过make方式创建的切片对应的数组是由make底层维护，对外不可见，只能通过slice去访问各个元素


方式3
定义一个切片，直接就指定具体数组，使用原理类似make的方式
*/
