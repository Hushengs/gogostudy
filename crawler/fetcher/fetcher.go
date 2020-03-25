package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string)([]byte,error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err!=nil {
		return nil,err
	}
	request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.137 Safari/537.36 LBBROWSER")
	resp, err :=http.DefaultClient.Do(request)
	if  err!=nil{
		return nil,err
	}
	defer resp.Body.Close()
	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return nil,fmt.Errorf("wrong status code: %d",resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
