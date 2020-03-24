package parser

import (
	"gogostudy/crawler/engine"
	"gogostudy/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([0-9]+)岁</div>`)
var heightRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([0-9]+)cm</div>`)

func Profile(contents []byte) engine.ParseResult{
	profile := model.Profile{}
	age,err := strconv.Atoi(extractString(contents,ageRe))
	if err == nil{
		profile.Age = age
	}
	height,err := strconv.Atoi(extractString(contents,heightRe))
	if err == nil{
		profile.Height = height
	}
	result :=engine.ParseResult{
		Items:    []interface{}{profile},
	}
	return result
}

func extractString(contents []byte,re * regexp.Regexp) string{
	match := re.FindSubmatch(contents)
	if len(match) >=2{
		return string(match[1])
	}else{
		return ""
	}

}
