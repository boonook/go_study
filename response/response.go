package response

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
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
	test, err := ReadAll("data/test.txt")
	result := strings.Replace(string(test), "\n", "", 1)
	fmt.Println("test_string", result)
	WriteFile("data/test.txt")
	readExcel("data/boonook.xlsx")
	_dir := "logs/boonook.xlsx"
	exist, err := isFileExist(_dir)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return
	}
	if exist {
		fmt.Printf("has dir![%v]\n", _dir)
	} else {
		fmt.Printf("--------------------vv--------")
		fmt.Printf("no dir![%v]\n", _dir)
		// 创建文件夹
		err := os.Mkdir(_dir, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
	file_name := "logs/a.txt"
	CreateFile(file_name)
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

/////读取文件内容
func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

/////将内容写入到文件中去
func WriteFile(fileName string) {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		// 打开文件失败处理

	} else {
		content := "写入的文件内容\n"

		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, 2)

		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
}

/////读取excel
func readExcel(fileName string) {
	xlsx, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println("open excel error,", err.Error())
		os.Exit(1)
	}
	// rows, err := xlsx.GetRows(xlsx.GetSheetName(xlsx.GetActiveSheetIndex()))
	cell := xlsx.GetCellValue("Sheet1", "B2")
	fmt.Println(cell)
	rows := xlsx.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println("row", row)
	}
}

////判断文件是否存在
func isFileExist(path string) (bool, error) {
	fileInfo, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	}
	//我这里判断了如果是0也算不存在
	if fileInfo.Size() == 0 {
		return false, nil
	}
	if err == nil {
		return true, nil
	}
	return false, err
}

/////创建文件
func CreateFile(file_name string) {
	//创建文件
	f, err := os.Create(file_name)
	//判断是否出错
	if err != nil {
		fmt.Println(err)
	}
	//打印文件名称
	fmt.Println(f.Name())
	// 　　 defer f.close()
}
