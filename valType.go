package main

import "fmt"

/*
value.(type)的应用
 */
 func main(){
 	var val interface{} = int8(-110)
	 switch v:=val.(type){//需要重新声明一个变量等于value.(type)，且必须放到switch case里
	 case int8:
	 	//此时编译器看来v就是int8类型
	 	fmt.Println(byte(v))
	 case int16:
	 	//此时编译其看来v就是int16类型
	 }
 }