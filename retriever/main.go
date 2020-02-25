package main

import (
	"fmt"
	"gogostudy/retriever/mock"
	reals "gogostudy/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string{
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is fake imooc.com"}
	inspect(r)

	//type assertion
	mockRetriever := r.(mock.Retriever)
	fmt.Println(mockRetriever.Contents)
	//fmt.Printf("%T %v\n",r,r)


	r = &reals.Retriever{UserAgent:"Mozilla/5.0",TimeOut:time.Minute}
	//fmt.Printf("%T %v\n",r,r)
	//fmt.Println(download(r))
	inspect(r)

	//type assertion
	if realRetriever,ok := r.(*reals.Retriever);ok{
		fmt.Println(realRetriever.TimeOut)
	}else{
		fmt.Println("not a mock retriever")
	}
}

func inspect(r Retriever)  {
	fmt.Printf("%T %v\n",r,r)
	switch v:=r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:",v.Contents)
	case *reals.Retriever:
		fmt.Println("UserAgent:",v.UserAgent)
	}
}
