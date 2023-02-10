package array

import (
	"fmt"
	"math/rand"
	"time"
)

func array_6() {
	//1).创建一个byte类型的26个元素的数组，分别放置'A'-'Z'
	//使用for循环访问所有元素并打印出来。提示：字符数据运算'A'+ 1 -> 'B'

	//思路
	//1.声明一个数组 var mychars [26]byte
	//2.使用for循环，利用字符可以进行运算的特点来赋值'A'+ 1 -> 'B'
	//3.使用for打印即可

	var mychars [26]byte
	for i := 0; i < 26; i++ {
		mychars[i] = 'A' + byte(i) //注意需要将i =>byte
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c \n", mychars[i])
	}

	//2).请求出一个数组的最大值，并得到对应的下标

	//思路
	//1.声明一个数组var intArr[5]=[...]int{1,-1,9,90,11}
	//2.假定第一个元素就是最大值，下标就是0
	//3.然后从第二个元素开始循环比较，如果发现有更大的，则交换
	var intArr = [...]int{1, -1, 59, 79, 90, 4000, -905, 10}
	maxVal := intArr[0]
	maxValindex := 0

	for i := 1; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValindex = i
		}

	}
	fmt.Printf("maxVal=%v maxValindex=%v\n\n", maxVal, maxValindex)

	//3).请求出一个数组的和和平均值 for-range
	//1.声明一个数组 var intArr[5]=[...]int{1,-1,9,90,11}
	//2.求出和sum
	//3.求出平均值
	var intArr2 [5]int = [...]int{1, -1, 9, 90, 12}
	sum := 0
	for _, val := range intArr2 {
		//累计求和
		sum += val
	}

	//如何让平均值保留到小数
	fmt.Printf("sum=%v 平均值=%v\n\n", sum, float64(sum)/float64(len(intArr2)))

	//要求：随机生成五个数，并将其反转打印
	//1.随机生成五个数，rand.Intn（）函数
	//2.当我们得到随机数后，就放到一个数组int数组
	//3.反转打印，交换的次数是len/2 ，倒数第一个和第一个元素交换，倒数第二个和第二个元素交换

	var intArr3 [5]int
	//为了每次生成的随机数不一样，我们需要给一个seed值
	len := len(intArr3)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		intArr3[i] = rand.Intn(100) //0<=n<100
	}
	fmt.Println("交换前=", intArr3)
	//反转打印，交换的次数是len/2
	//倒数第一个和第一个元素交换，倒数第二个和第二个元素交换
	temp := 0 //做一个临时变量
	for i := 0; i < len/2; i++ {
		temp = intArr3[len-1-i]
		intArr3[len-1-i] = intArr3[i]
		intArr3[i] = temp
	}
	fmt.Println("交换后=", intArr3)
}
