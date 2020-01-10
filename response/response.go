package response

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/huichen/pinyin"
	"io/ioutil"
	"os"
	"strings"
	"time"
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
	readExcel("data/boonook.xlsx")
	t := time.Now()
	////时间转换
	nowTime := t.Format("2006-01-02")
	file_name := "logs/" + nowTime + ".log"
	CreateFile(file_name)
	files()
	timeDaXiao()
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
func WriteFile(fileName string, content string) {
	fmt.Println("content----------", content)
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		// 打开文件失败处理

	} else {
		content := content + "\n"

		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, 2)

		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
}

/**
	读取excel
	param fileName{string} ///文件路径
**/
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

/***
	////判断文件是否存在
	param path{string}  ///文件路径
***/
func isFileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/***
	创建文件
	param path{string}  ///文件路径
***/
func CreateFile(file_name string) {
	//创建文件
	_dir := file_name
	exist, err := isFileExist(_dir)
	t := time.Now()
	////时间转换
	nowTime := t.Format("2006-01-02 15:04:05")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		////获取文件信息失败
		fmt.Printf("get dir error![%v]\n", err)
		return
	}
	if exist {
		////文件已经存在，无法创建文件
		fmt.Printf("has dir![%v]\n", _dir)
		WriteFile(file_name, nowTime+" [succ] 111111111111111")
	} else {
		////文件不存在我们可以创建文件了
		f, err := os.Create(file_name)
		//判断是否出错
		if err != nil {
			fmt.Println(err)
		}
		//打印文件名称
		fmt.Println(f.Name())
		WriteFile(file_name, nowTime+" [succ] 111111111111111")
		// 　　 defer f.close()
	}
}

////获取目录下的所有文件并进行截取，同时获取十五天以内的文件
func files() {
	files, err := ioutil.ReadDir(`logs/`)
	if err != nil {
		panic(err)
	}
	// 获取文件，并输出它们的名字
	// 保留15天以内的文件
	for _, file := range files {
		println(file.Name())
		////对文件名进行截取
		s := strings.Split(file.Name(), ".")[0]
		// fmt.Println("s-----------", s)
		///当前时间的前15天
		d, _ := time.ParseDuration("-360h")
		format := "2006-01-02 15:04:05"
		now := time.Now()
		d1 := now.Add(d).Format("2006-01-02 15:04:05")
		_today := now.Format(format)
		////获取15天之前的日期
		fmt.Println("d1----------", _today)
		// a, _ := time.Parse(s, "2019-03-10 11:00:00")
		fmt.Println("a----------", d1)
		stringTime := s
		loc, _ := time.LoadLocation("Local")
		the_time, err := time.ParseInLocation("2006-01-02", stringTime, loc)
		if err == nil {
			unix_time := the_time.Unix()                        //1504082441
			timeNow := time.Unix(unix_time, 0)                  //2017-08-30 16:19:19 +0800 CST
			unix_time2 := timeNow.Format("2006-01-02 15:04:05") //2015-06-15 08:52:32
			fmt.Println("----------------------------------", unix_time2)
			///文件名对应的时刻
			fileTime, _ := time.Parse(format, unix_time2)
			////当前时间的前15天的时刻
			a, _ := time.Parse(format, d1)
			status := fileTime.Before(a)
			if status {
				////说明文件名对应的时间不在15天以内
				fmt.Println("now  aaaaa---------   After: ", status)
				err := os.Remove("logs/" + file.Name())
				if err != nil {
					t := time.Now()
					////时间转换
					nowTime := t.Format("2006-01-02")
					file_name := "logs/" + nowTime + ".log"
					WriteFile(file_name, nowTime+" [erro]"+err.Error())
				} else {
					fmt.Println("now  success---------: ", "文件删除成功")
					t := time.Now()
					////时间转换
					nowTime := t.Format("2006-01-02")
					file_name := "logs/" + nowTime + ".log"
					newDateTime := t.Format(format)
					WriteFile(file_name, newDateTime+" [succ] 文件删除成功"+file_name)
				}
			} else {
				////说明文件名对应的时间在15天以内
				fmt.Println("now  a   Before: ", status)

			}
		}

		// d2, _ := time.Parse(d1, "2019-03-10 11:00:00")
	}
}

func timeDaXiao() {
	format := "2006-01-02 15:04:05"
	now := time.Now()
	//now, _ := time.Parse(format, time.Now().Format(format))
	a, _ := time.Parse(format, "2019-03-10 11:00:00")
	b, _ := time.Parse(format, "2015-03-10 16:00:00")

	fmt.Println("now:", now.Format(format), "\na:", a.Format(format), "\nb:", b.Format(format))
	fmt.Println("---    a > now  >  b   -----------")
	fmt.Println("now  a   After: ", now.After(a))
	fmt.Println("now  a   Before:", now.Before(a))
	fmt.Println("now  b   After:", now.After(b))
	fmt.Println("now  b   Before:", now.Before(b))
	fmt.Println("a  b   After:", a.After(b))
	fmt.Println("a  b   Before:", a.Before(b))
	fmt.Println(now.Format(format), a.Format(format), b.Format(format))
	fmt.Println(now.Unix(), a.Unix(), b.Unix())
}

////将年月日2019-09-08转成2019-09-08 00:00:00
