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
