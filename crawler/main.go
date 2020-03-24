package main

import (
	"gogostudy/crawler/engine"
	"gogostudy/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.CityList,
	})
}
