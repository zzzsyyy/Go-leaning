// test1
package main

import (
	"fmt" //标准输入输出库
)

func main() {
	fmt.Println("Hello World!") //Println-换行打印
	test1()
}

func test1() {
	s := "cello"    //字符串-不能直接改变
	c := []byte(s)  //将其转换成[]byte变量
	c[0] = 'h'      //将c的第一位改变
	ls := string(c) //将[]byte转变回字符串
	fmt.Printf(ls)  //打印改变后的字符串
}

func test2() {

}
