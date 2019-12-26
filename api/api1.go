package api

import (
	"fmt"
	"studyManage/gutil"
)

func Api() string {
	return "hello api-------!"
}

func ApiArr() {
	var list = []int{95, 45, 15, 78, 84, 51, 24, 12}
	var list32 = []int32{95, 45, 15, 78, 84, 51, 24, 12}
	var list64 = []int64{95, 45, 15, 78, 84, 51, 24, 12}
	////对数组进行排序
	fmt.Println("list-----", gutil.IntInsertionSortAsc(list))
	fmt.Println("list32-----", gutil.Int32InsertionSortAsc(list32))
	fmt.Println("list64-----", gutil.Int64InsertionSortAsc(list64))
}
