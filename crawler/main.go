package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func printCityList(contents []byte){
	//contents = "<a data-v-5e16505f="" href="http://www.zhenai.com/zhenghun/aba">阿坝</a>"
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	//matchs := re.FindAll(contents,-1)
	matchs := re.FindAllSubmatch(contents,-1)
	//fmt.Println(matchs)
	for _,m := range matchs{
			fmt.Printf("City: %s, URL: %s \n", m[2],m[1])
		}
	//fmt.Printf("%s\n",m)
	fmt.Printf("Matches found :%d\n",len(matchs))
}

func main() {
	resp,err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Println("Error:status code",resp.StatusCode)
		return
	}
	//utf8Reader := transform.NewReader(resp.Body,simplifiedchinese.GBK)
	all,err :=ioutil.ReadAll(resp.Body)
	if err != nil{
		panic(err)
	}
	//fmt.Printf("%s\n",all)
	printCityList(all)
}
