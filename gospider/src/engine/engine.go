package engine

import (
	"fetcher"
	"log"
)

func Run(seeds ...Request) {  //可变参数，可以传入任意数量的Request类型的参数
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests)>0 {
		r := requests[0]
		requests = requests[1:]  // 会不会出现溢出问题？？不会溢出
		log.Printf("Fetching %s\n", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher:error " + "fetching url %s : &v", r.Url,err)
			continue  // 报错并且处理下一个request
		}
		parserResult := r.Parserfunc(body)  // 这个函数是个抽象的函数，但是在哪里根据URL调用具体的函数呢？？
		// 这个函数实在main函数中指定的ParseCityList
		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items {
			log.Printf("Got item %v", item)
		}

	}
}
// 对于interface{}, %v会打印实际类型的值。