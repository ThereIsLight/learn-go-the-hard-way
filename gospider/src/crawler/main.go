package main

import (
	"engine"
	"zhenai/parser"
)
const startUrl = `http://www.zhenai.com/zhenghun`
func main() {
	engine.Run(engine.Request{Url:startUrl, Parserfunc:parser.ParseCityList})
}

// 20181112 目前来说最大的问题： 怎么从request中找到函数，并且执行？？？