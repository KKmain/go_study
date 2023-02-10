package array

import "fmt"

func array_3() {
	heroes := [...]string{"宋江", "吴用", "卢俊义"}

	for i, v := range heroes {
		fmt.Printf("i=%v v=%v\n", i, v)
		fmt.Printf("heroes[%d]=%v\n", i, heroes)
	}

	for _, v := range heroes {
		fmt.Printf("元素的值=%v\n", v)
	}
}

/*基本语法
for index,value := range array01{
    ...
}

说明
1、第一个返回值index是数组的下标
2、第二个value是在该下标位置的值
3、他们都是仅在for循环内部可见的局部变量
4、遍历数组元素的时候，如果不想使用下标index，可以直接把下标index标为下划线_
5、index和value的名称不是固定的
*/
