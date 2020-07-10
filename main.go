package main

import "fmt"

/**
Version of Delve is too old for this version of Go
Delve 是调试工具，由于版本过低，不能使用debug模式

*/

func main() {
	const helloworld string = "hello world"
	fmt.Println("这是第一个go程序，向世界打个招呼吧，" + helloworld)
	fmt.Println(f)
	f := "变更"
	fmt.Println(f)
	name = "hellochan"
	age = 12
	fmt.Println(name)
	fmt.Println(age)
	if success {
		fmt.Println("成功进入")
	}
}

const a string = "a"        // 常量 const 常量名称 数据类型 = 变量值
const b = "test"            // 自动判断类型
const c, d, e = "c", "d", 1 // 多常量声明
// 变量
var f = "f"
var (
	name = "chen"
	age  int
)
var success bool = false

/*
	算数运算符: + - * / % ++ --
	关系运算符: == != > < >= <=
	逻辑运算符: && || !
	赋值运算符: = += -= *= /= %= <<= >>= &= |= 指定运算后赋值
*/

/*
	if 布尔表达式 {} else {}
	switch var {
		case "":
		case "":
		case "":
		default:
	}
	select {


	}
*/

/*
	for init; condition; post { }
	for condition { }
	for { }
	init： 一般为赋值表达式，给控制变量赋初值；
	condition： 关系表达式或逻辑表达式，循环控制条件；
	post： 一般为赋值表达式，给控制变量增量或减量。
*/
