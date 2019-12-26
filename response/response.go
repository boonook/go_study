package response

import (
	"encoding/json"
	"fmt"
	"github.com/huichen/pinyin"
)

func Response() {
	type JsonReturn struct {
		Msg  string      `json:"message"`
		Code int         `json:"code"`
		Data interface{} `json:"data"` //Data字段需要设置为interface类型以便接收任意数据
		//json标签意义是定义此结构体解析为json或序列化输出json时value字段对应的key值,如不想此字段被解析可将标签设为`json:"-"`
	}
	var info JsonReturn
	info.Msg = "操作成功"
	info.Code = 200
	type Transport struct {
		Time  string `json:"time"`
		MAC   string `json:"mac"`
		Id    string `json:"id"`
		Rssid string `json:"rssid"`
	}
	var st []Transport
	for i := 1; i < 10; i++ {
		t1 := Transport{Time: "22", MAC: "33", Id: "44", Rssid: "55"}
		st = append(st, t1)
	}
	////将数组转成数组对象
	buf, _ := json.Marshal(st)
	fmt.Println("---------------------------", string(buf))
	info.Data = st
	///将结构体转成数组
	b, err := json.Marshal(info)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("---------------------------", string(b))
	/////将数组转成结构体
	var info2 JsonReturn
	json.Unmarshal([]byte(string(b)), &info2)
	fmt.Println("将数组转成结构体----------------", info2)
	////获取结构体中的某个值
	fmt.Println("获取结构体中的某个值----------------", info2.Code)
	d2, err := json.Marshal(info2.Data)
	fmt.Println("将结构体中的数组转化成正常的数组", string(d2))
	/////集合
	var arr = []string{"hello", "hi", "world", "hi", "china", "hello", "hi"}
	fmt.Println(RemoveRepeatedElement(arr))
	// fmt.Println(JsonRemoveRepeatedElement(string(d2),"id"))
	var py pinyin.Pinyin

	// 初始化，载入汉字拼音映射文件
	py.Init("data/pinyin_table.txt")

	// 返回汉字的拼音
	// GetPinyin的第一个参数为单个汉字，第二个参数为是否返回带声调的拼音。
	// 比如下面两行的输出分别为 "zhōng" 和 "zhong"
	// 当该字无法识别或者不是汉字时返回空字符串。
	fmt.Println(py.GetPinyin('中', true))
	fmt.Println(py.GetPinyin('中', false))

	// 返回汉字的声调（整数），0为轻声，1为平声，依此类推。
	fmt.Println(py.GetNumericTone('中'))

	// 下面的输出分别为 "lǜ" 和 "lv"
	fmt.Println(py.GetPinyin('绿', true))
	fmt.Println(py.GetPinyin('绿', false))
	fmt.Println(py.GetPinyin('朱', false))
}

////数组去重
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
