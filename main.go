package main

import (
	"encoding/json"
	"fmt"
	"studyManage/api"
	"studyManage/gutil"
	"studyManage/response"
)

type Users struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

////首字母大写才能转成json,如果不json:"id"，那么默认去的key值就是ID了
type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func main() {
	//结构体解析成json
	user1 := Users{"1", "user1", 22}
	s, err := json.Marshal(user1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(s))
	}

	var user2 Users
	err = json.Unmarshal(s, &user2)
	if err == nil {
		fmt.Println(user2)
	}

	//map解析成json
	m := make(map[string]interface{}, 2)
	m["userinfo"] = user2
	m["remark"] = "bb"
	// fmt.Println(m)
	var data []byte
	if data, err = json.Marshal(m); err == nil {
		fmt.Println("map解析成json", string(data))
	}

	//json解析成map
	// if err = json.Unmarshal(data, &m); err == nil {
	// 	fmt.Println(m)
	// }
	var a Users /* 声明 Book1 为 Books 类型 */
	a.ID = "123123"
	a.Name = "boonook"
	a.Age = 26
	aa, aErr := json.Marshal(a)
	if aErr != nil {
		fmt.Println(aErr)
	} else {
		fmt.Println(string(aa))
	}
	sss := []Response{
		Response{
			"1",
			"yy",
		},
		Response{
			"2",
			"yang",
		},
		Response{
			"3",
			"go",
		},
	}
	fmt.Println(sss)

	type Transport struct {
		Time  string `json:"time"`
		MAC   string `json:"mac"`
		Id    string `json:"id"`
		Rssid string `json:"rssid"`
	}

	type JsonReturn struct {
		Msg  string      `json:"message"`
		Code int         `json:"code"`
		Data interface{} `json:"data"` //Data字段需要设置为interface类型以便接收任意数据
		//json标签意义是定义此结构体解析为json或序列化输出json时value字段对应的key值,如不想此字段被解析可将标签设为`json:"-"`
	}

	var st []Transport
	for i := 1; i < 10; i++ {
		t1 := Transport{Time: "22", MAC: "33", Id: "44", Rssid: "55"}
		st = append(st, t1)
	}
	// t1 := Transport{Time: "22", MAC: "33", Id: "44", Rssid: "55"}

	// t2 := Transport{Time: "66", MAC: "77", Id: "88", Rssid: "99"}
	// st = append(st, t2)
	fmt.Println("st", st)
	buf, _ := json.Marshal(st)            ////数组转成json数组
	fmt.Println("go的json数组", string(buf)) ////string()它将该整型数字转换成ASCII码值等于该整形数字的字符。strconv.Itoa()函数的参数是一个整型数字，它可以将数字转换成对应的字符串类型的数字。
	var str = string(buf)
	var info JsonReturn
	info.Msg = "操作成功"
	info.Code = 200
	info.Data = string(buf)
	fmt.Println("info--------", info)
	var st1 []Transport
	errr := json.Unmarshal([]byte(str), &st1) ///将json字符串解码到相应的数据结构
	if errr != nil {
		fmt.Println("some error")
	}
	fmt.Println(st1)
	fmt.Println(st1[0].Time)

	var Msg []map[string]string
	json.Unmarshal([]byte(str), &Msg)
	fmt.Println(Msg)
	gutil.Ceshi()
	api.ApiArr()
	fmt.Println(api.Api2())
	response.Response()
	var data_get = api.Get("https://api.imjad.cn/cloudmusic/?type=song&id=28012031&br=128000")
	fmt.Println("data--------", data_get)
	api.GetToken()
	api.Mysql()
	api.HandleArray()
}
