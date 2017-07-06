package main

import "fmt"

const boilingF  = 212.0

func main()  {
	var f  =  boilingF
	var c  =  (f - 32)*5/9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)

	x := 1
	var p *int = &x//指针
	y := *p //取指针所指向的值赋值给变量y
	*p = 2 //修改指针所指向的内存中的值
	fmt.Println(y)
	fmt.Println(x)
	fmt.Println(p)
	fmt.Println(*p)

}
