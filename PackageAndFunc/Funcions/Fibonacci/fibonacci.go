package fibonacci

import "fmt"

func fibonacci() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci_2(i)
		fmt.Printf("fibonacci(%d)=%d\n", i, result)
	}

}

func fibonacci_2(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci_2(n-1) + fibonacci_2(n-2)
	}
	return
}
