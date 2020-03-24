package engine

import (
	"gogostudy/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _,r:= range seeds{
		requests = append(requests,r)
	}

	for len(requests) >0{
		r:=requests[0]
		requests = requests[1:]
		log.Printf("fetching %s",r.Url)
		body,err := fetcher.Fetch(r.Url)
		if err != nil{
			log.Panicf("Fetcher:err"+"fetching url %s:%v",r.Url,err)
			continue
		}
		ParseResult :=r.ParserFunc(body)
		requests = append(requests,ParseResult.Requests...)
		for _,item := range ParseResult.Items{
			log.Printf("Got item %v",item)
		}
	}
}
