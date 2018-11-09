package main

import (
	"engine"
	"zhenai/parser"
)
const startUrl = `http://www.zhenai.com/zhenghun`
func main() {
	engine.Run(engine.Request{Url:startUrl, Parserfunc:parser.ParseCityList})
}