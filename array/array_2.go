package array

import "fmt"

func array_2() {
	//数组可以存放多个同一类型数据，数组也是一种数据类型，在go中，数组是值类型
	//用数组解决养鸡场问题
	//定义一个数组
	var hens [6]float64
	//给数组的每个元素赋值，元素的下标是从0开始的
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0
	//遍历数组求出总体重
	totalWeight2 := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight2 += hens[i]
	}

	//求出平均体重
	avgWeight2 := fmt.Sprintf("%.2f", totalWeight2/float64(len(hens)))
	fmt.Printf("totalWeight2=%v avgWeight2=%v", totalWeight2, avgWeight2)
}
