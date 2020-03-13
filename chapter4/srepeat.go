package main

import "fmt"

func srepeat(s []string) []string {
	var ss []string
	for i:=0;i<len(s);i++  {
		if i==len(s)-1 {
			if s[i]!=s[i-1]{
				ss = append(ss,s[i])
			}
			break
		}
		if s[i]!=s[i+1]{
			ss = append(ss,s[i])
		}else if s[i]==s[i+1] && i==len(s)-2{
			ss = append(ss,s[i])
		}
	}
	return ss
}

func main() {
	s:=[]string{"hello","world","world","hello","hello","super","man","man","jap"}
	fmt.Println(srepeat(s))
}
