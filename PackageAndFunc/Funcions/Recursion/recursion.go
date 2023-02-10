package recursion

import (
	"fmt"
)

func recursion() {
	test(4)
}

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n", n)
}

/*
一个函数在函数体内又调用了本身，称递归调用
栈：先入后出
*/
