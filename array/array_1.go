package array

import "fmt"

func array_1() {
	/*
		一个养鸡场有6只鸡，他们的体重分别是3、5、1、3.4、2、50kg。
		请问这六只鸡的总体重是多少？平均体重是多少？

		思路分析：定义6个变量，分别表示六只鸡，然后求出和，再求平均值
	*/
	hen1 := 3.0
	hen2 := 5.0
	hen3 := 1.0
	hen4 := 3.4
	hen5 := 2.0
	hen6 := 50.0
	totalWeight := hen1 + hen2 + hen3 + hen4 + hen5 + hen6
	avgWeight := fmt.Sprintf("%.2f", totalWeight/6)
	//将totalWeight/6的结果四舍五入保留到小数点后两位
	fmt.Printf("totalWeight=%v avgWeight=%v", totalWeight, avgWeight)

	//如果养鸡场有600只鸡呢？？ 使用传统方法不利于数据的管理和维护，不够灵活，因此引出数组
}
