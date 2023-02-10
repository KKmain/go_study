package slice

import "fmt"

func slice_3() {
	//用append内置函数，可以对切片进行动态追加
	var slice3 []int = []int{100, 200, 300}
	//slice4 := append(slice3, 400, 500, 600)
	slice3 = append(slice3, 400, 500, 600) //通过append直接给slice3追加具体的元素

	//通过append将切片slice3追加给slice3，这个slice3也可以是别的切片,后面的只能是切片，不能是数组， ...必须带上
	slice3 = append(slice3, slice3...)

	//fmt.Println("slice4", slice4)
	fmt.Println("slice3", slice3)

	fmt.Println()
	/////切片使用copy内置函数完成拷贝////
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)

	copy(slice5, slice4) //说明copy(para1,para2)：para1和para2都是切片类型
	fmt.Println("slice=4", slice4)
	fmt.Println("slice=5", slice5)

}

/*
1.切片出hi话时 var slice=arr[startIndex:endIndex]
说明：从arr数组下标为startIndex开始，取到下标为endIndex的元素(不包含arr[endIndex])
切片初始化时，仍然不能越界。范围在[0-len(arr)]之间，但是可以动态增长.
var slice = arr[0:end] 可以简写 var slice = arr[:end]
var slice = arr[start:len(arr)] 可以简写 var slice = arr[start:]
var slice = arr[0:len(arr)] 可以简写 var slice = arr[:]

cap是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素
切片定义完后，还不能使用，因为本身是一个空的，需要让其引用到一个数组，或者make一个空间供切片来使用

切片可以继续切片
------------------------------------

用append内置函数，可以对切片进行动态追加
append将元素追加到切片的末尾，若它有足够的容量，其目标就hi重新切片以容纳新的元素，否则，就会
分配一个新的基本数组。append返回更新后的切片，因此必须存储追加后的结果

切片append操作的底层原理分析：

1.切片append操作的本质就是对数据扩容
2.go底层会创建一下新的数组newArr(安装扩容后大小)
3.将slice原来包含的元素拷贝到新的数组newArr
4.slice重新引用到newArr
5.注意newArr是在底层来维护的，程序员不可见

切片的拷贝操作
切片使用copy内置函数完成拷贝

 string底层是一个byte数组，因此string也可以进行切片处理
 string是不可变的，也就说不能通过str[0]='z'方式来修改字符串
 如果需要修改字符串，可以先将string->[]byte/或者[]rune->修改->重写转成string
*/
