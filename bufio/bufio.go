package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {

	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	bw.WriteByte('H')
	bw.WriteByte('e')
	bw.WriteByte('l')
	bw.WriteByte('l')
	bw.WriteByte('o')
	bw.WriteByte(' ')
	bw.WriteRune('世')
	bw.WriteRune('界')
	bw.WriteRune('！')
	bw.Flush()
	fmt.Println(b) // Hello 世界！


	//times := "1234567898765432123456789"
	//r := strings.NewReader(times)
	//br := bufio.NewReaderSize(r,5)
	//b :=  make([]byte,5)
	//for i:=0;i<10;i++{
	//	n,err := br.Read(b)
	//	fmt.Printf("%v\n", n)
	//	fmt.Printf("%s\n", string(b))
	//	fmt.Printf("%s\n", err)
	//}
	//br.ReadBytes()
	//br.ReadString()
	//b, err := br.Peek(50)


	//n1,err := br.Read(b)
	//fmt.Printf("%v\n", n1)
	//fmt.Printf("%s\n", string(b))
	//fmt.Printf("%s\n", err)
	//b1, err := br.Peek(50)
	//fmt.Printf("%s\n", err)
	//fmt.Printf("%s\n", string(b1))

	// b, _ := br.ReadByte()
	//for k,v := range b {
	//	fmt.Println(k,string(v))
	//}
	//if err != nil{
	//	fmt.Printf("%s\n", err)
	//}
}
