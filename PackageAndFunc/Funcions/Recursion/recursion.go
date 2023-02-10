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
