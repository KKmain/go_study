package slice

import "fmt"

func slice_1() {

	//演示切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	//声明一个切片
	//slice := intArr[1:3]
	//1.slice就是切片名
	//2.intArr[1:3] 表示slice引用到ineArr这个数组
	//3.引用intArr数组的起始下标为1，最后的下标为3(但是不包含3)

	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是=", slice)       // 22,33
	fmt.Println("slice 的元素个数=", len(slice)) //2
	fmt.Println("slice 的容量=", cap(slice))   //切片的容量是可以动态变化
}

/*
我们需要一个数组用于保存学生的成绩，但是学生的个数是不确定的，请问怎么办？使用切片

基本介绍：
1.英文是slice
2.切片是数组的一个引用，因此切片是引用类型
3.切片的使用和数组类似，遍历切片，访问切片的元素和求切片长度len(slice)都一样
4.切片的长度是可以变化的，因此切片是一个可以动态变化的数组
5.切片定义的基本语法：
var 切片名[不需要填长度]类型
 如：var a [] int
 6.cap相当于容器为扩容的最大值，一旦len超过cap，就会触发扩容机制
*/
