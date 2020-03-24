package parser

import (
	"gogostudy/crawler/fetcher"
	"testing"
)

func TestCityList(t *testing.T) {
	contents,err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if err != nil{
		panic(err)
	}
	result := CityList(contents)

	const resultSize  = 470

	if len(result.Requests) != resultSize{
		t.Errorf("result should have %d"+
			"request:but had %d",resultSize,len(result.Requests))
	}
	if len(result.Items) != resultSize{
		t.Errorf("result should have %d"+
			"request:but had %d",resultSize,len(result.Items))
	}
}
