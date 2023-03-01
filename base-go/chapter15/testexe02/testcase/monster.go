package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Store() bool {

	// 先序列号
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Printf("monster json err=%v\n", err)
		return false
	}

	// 写入文件中
	file1 := "e:/test2"
	err = ioutil.WriteFile(file1, data, 111)
	if err != nil {
		fmt.Printf("write file err=%v\n", err)
		return false
	}
	return true

}

func (this *Monster) ReStore() bool {

	// 读取文件
	file1 := "e:/test2"
	data, err := ioutil.ReadFile(file1)
	if err != nil {
		fmt.Printf("read file err=%v\n", err)
		return false
	}

	// 反序列化
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
		return false
	}
	return true
}
