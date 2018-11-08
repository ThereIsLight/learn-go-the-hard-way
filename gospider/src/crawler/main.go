package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"regexp"
)

func main() {
	resp,err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()  // ??
	if resp.StatusCode == http.StatusOK {
		all, err := ioutil.ReadAll(resp.Body)  //change io.reader to []byte ???
		if err != nil {
			panic(err)
		}
		//fmt.Printf("%s\n", all)
		// 字符编码的处理，这里不需要，网页已经改成了utf-8的编码形式。
		// 正则表达式提取城市列表
	}
}
func printCityList(contents []byte) {  // byte格式
	regexp.Compile(`<a target="_blank" href="http://www.zhenai.com/zhenghun/shanghai" data-v-5450af50>上海</a>`)
}
