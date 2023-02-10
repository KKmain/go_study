package SortAndFind

import "fmt"

func find_1() {
	/*
		有一个数列：白眉，金毛，紫衫，青翼
		要求：从键盘中任意输入一个名称，判断数列中是否包含此名称{顺序查找}
		思路
		定义一个数组：白眉，金毛，紫衫，青翼 字符串数组
		从控制台接收一个名字，依次比较，如果发现有，提示
	*/
	names := [4]string{"白眉", "金毛", "紫衫", "青翼"}
	var heroName = ""
	fmt.Println("请输入要查找的人名")
	fmt.Scanln(&heroName)

	//顺序查找：第一种方式
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			fmt.Printf("找到%v,下标%v\n", heroName, i)
			break
		} else if i == (len(names) - 1) {
			fmt.Printf("没有找到 %v\n", heroName)
		}
	}
}
