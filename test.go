package main

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	//var x, y int
	//p :=&x
	//p1 :=&x
	//fmt.Println(&x == &x, &x == &y, p  == p1)
	//fmt.Println(p)

	//freq := rand.Float64() * 3.0
	//fmt.Println(freq)

	//var x,y = 1,3
	//fmt.Println(&x == &x, &x == &y, &x == nil)
	//fmt.Println(&x)
	//fmt.Println(&y)

	//s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	//reverse(s[:2])
	//reverse(s[2:])
	//reverse(s)
	//fmt.Println(s) // "[2 3 4 5 0 1]"

	//x := []int{1,2,3}
	//xlen :=  len(x)
	//xcap := cap(x)
	//copy(y,x)
	//fmt.Printf("%T cap=%d\n ",x,cap(x))
	//fmt.Printf("%T cap=%d\n",y,cap(y))

	//for i:=0;i<10;i++{
	//	x = append(x,i)
	//	fmt.Println(x)
	//	fmt.Printf("%T cap=%d\n ",x,cap(x))
	//}
}
