package main


import (
	"regexp"
	"fmt"
)

/*
正则表达式根据语法规则匹配文本中的字符串
*/
const test = `
My eamil is ccmouse@gmail.com@gmail.com.cn
My eamil is abc@def.org
My eamil is kkk@qq.com
`

func main() {
	//compile, e := regexp.Compile("ccmouse@gmail.com")  // ctrl+shift+alt+t 快捷键自动补全函数的返回值
	//mustCompile := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	mustCompile := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9]+)`)
	s := mustCompile.FindString(test)  // 寻找单个匹配的字符串
	ss := mustCompile.FindAllString(test, -1)  //寻找多个匹配的字符串,-1代表匹配所有。
	sss := mustCompile.FindAllStringSubmatch(test, -1)  //括号里面的内容匹配出来
	fmt.Println(s)
	fmt.Println(ss)
	fmt.Println(sss)
	// ctrl+鼠标左键 查看函数的声明，点击查看函数的具体代码。
}
/*
简单的正则表达式：
.  匹配任意的字符
+  1~N个字符
*  0~N个字符
[a-zA-z0-9]  匹配大括号里面的字符
` `不会转义
 */