package parser

import (
	"engine"
	"regexp"
	"strconv"
	"model"
)
// 20181112 同视频中的网页相比，前端页面大幅度的修改。最主要的区别就是不再是表格形式，值前面再也没有关键字。
// 找不到性别找个值的位置，貌似是没有的。
// 写起正则表达式非常的困难
// 烦烦烦
var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08="">([\d]+)岁</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08="">([\d]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08="">([\d]+)kg</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08="">月收入([^<]+)</div>`)
// var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)  // 找不到
var xingzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-ff544c08="">([^座]+)座[^<]+</div>`)
// var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)  // 怎么找？？？
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
func ParseProfile(contents []byte) engine.ParserResult {
	profile := model.Profile{}
	// profile.Name = name
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	profile.Marriage = extractString(contents, marriageRe)
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}
	profile.Income = extractString(contents, incomeRe)
	// profile.Gender = extractString(contents, genderRe)
	profile.Xingzuo = extractString(contents, xingzuoRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Hukou = extractString(contents, hokouRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)
	result := engine.ParserResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)  // 切片中第一个值为匹配到的字符串，其余的内容是需要提取的部分。
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
