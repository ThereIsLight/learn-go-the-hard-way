package parser

import (
	"engine"
	"regexp"
)
const cityRe= `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
func ParseCity(contents []byte) engine.ParserResult{
	compile := regexp.MustCompile(cityRe)
	match := compile.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range match {
		// fmt.Println(m)  //为什么这样就是输出一个字符切片，%s就不会输出切片。%s将其转化为string了吗？？
		// fmt.Printf("%s\n", m)
		result.Items = append(result.Items, "City: " + string(m[2]))  // 这里充分的体现出了Interface{}类型的优势
		// 在处理不同的网页的时候，得到的Item是不同，但是讲功能相似的代码抽象为函数之后，需要对这些不同类型的Item进行处理。
		// 这时使用接口类型就可以包含所有的数据类型。
		result.Requests = append(result.Requests, engine.Request{Url:string(m[1]), Parserfunc:engine.NilParser})  // 函数作为参数
	}
	return result
}
