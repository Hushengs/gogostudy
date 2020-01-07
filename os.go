package main

func main() {
	//var s,sep string
	//原始写法
	//for i :=1;i < len(os.Args);i++{
	//	s +=sep + os.Args[i]
	//	sep = ""
	//}
	//Go语言中这种情况的解决方法是用空标识符（blank identifier），即_（也就是下划线）。空标识符可用于任何语法需要变量名但程序逻辑不需要的时候
	//for _, arg := range os.Args[1:] {
	//	s += sep + arg
	//	sep = " "
	//}
	//3
	//fmt.Println(strings.Join(os.Args[1:], " "))

	//practice 1 修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字
	//for i :=0;i<len(os.Args);i++{
	//	s += sep + os.Args[i]
	//	sep = ""
	//}

	//for _,arg := range os.Args[0:] {
	//	s +=sep +arg
	//	sep = ""
	//}

	//fmt.Println(strings.Join(os.Args[0:],""))

	//practice 2 修改echo程序，使其打印每个参数的索引和值，每个一行。
	//for i:=1;i<len(os.Args);i++{
	//	fmt.Println(i,os.Args[i])
	//}

	//for inx,arg := range os.Args[1:]{
	//	fmt.Println(inx,arg)
	//}

	//practice 3做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异


	//fmt.Println(s)
}
