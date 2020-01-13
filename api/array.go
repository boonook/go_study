package api

import (
	"encoding/json"
	"fmt"
	"strings"
)

// 操作数组
func HandleArray() {
	fmt.Println("------------------------------------操作数组-----------------------------------")
	var _arr1 = [3]int{1, 2, 3}
	fmt.Println(_arr1)
	var _arr2 = [3]string{"a", "b", "c"}
	fmt.Println("string数组", _arr2)
	/////float32与float64的区别在于小数的精度
	var _arr3 = [3]float32{1.11, 2.22, 3.33}
	fmt.Println("float数组", _arr3)
	///////遍历数组
	fmt.Println(_arr3[0])
	fmt.Println(_arr3[2])
	for k, v := range _arr2 {
		fmt.Println(k, v)
	}
	for k, v := range _arr1 {
		fmt.Println(k, v)
	}
	slice := []int{1, 2, 3}
	/////函数接受数组改为接受切片，返回数组改为返回切片
	a := getAverage(_arr1, len(_arr1), slice)
	fmt.Println("myNum切片-------------", a)
}

///向函数中传递数组
func getAverage(arr [3]int, size int, num []int) []int {
	fmt.Printf("nums: %v,nums addr: %p\n", num, &num)
	fmt.Println("size-------------", size)
	b, err := json.Marshal(arr)
	if err != nil {
		panic(err)
	}
	var result = string(b)
	fmt.Println("result-------------", result)
	////将数组转成用逗号凭借的字符串
	var temp = make([]string, len(arr))
	for k, v := range arr {
		temp[k] = fmt.Sprintf("%d", v)
	}
	var result2 = strings.Join(temp, ",")
	fmt.Println("result2-------------", result2)
	myNum := []int{10, 20, 30, 40, 50}
	// 改变索引为 1 的元素的值
	myNum[1] = 25
	fmt.Println("myNum切片-------------", myNum)
	return myNum
}
