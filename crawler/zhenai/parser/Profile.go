package parser

import (
	"gogostudy/crawler/engine"
	"gogostudy/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)岁</div>`)

var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)cm</div>`)
//var otherRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)

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
	re := regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
	matchs := re.FindAllSubmatch(contents,-1)
	if matchs != nil{
		for _,v:=range matchs{
			profile.Label = append(profile.Label,string(v[1]))
		}
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
