package engine

type Request struct {
	Url string
	Parserfunc func([]byte) ParserResult
}
type ParserResult struct {
	Requests []Request  // 循环定义了？？？ 这种循环如何理解
	Items    []interface{}  // 任何类型
}
// 空的解析函数
func NilParser([]byte) ParserResult {
	return ParserResult{}  // do nothing
}