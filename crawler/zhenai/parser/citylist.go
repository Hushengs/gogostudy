package parser

import (
	"gogostudy/crawler/engine"
	"regexp"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func CityList(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents,-1)
	result :=engine.ParseResult{}
	limit :=1
	for _,m := range matchs{
		result.Items = append(result.Items,"City "+string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url:        string(m[1]),
			ParserFunc: City,
		})
		limit --
		if limit == 0{
			break
		}
	}
	return result
}
