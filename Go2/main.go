package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Multi(a int, b int) int {
	return a * b
}

func Div(a int, b int) int {
	if b == 0 {
		panic("除数不能为0")
	} else {
		return a / b
	}
}

func Calculator(num1 int, num2 int, CMD func(int, int) int) int {
	return CMD(num1, num2)
}

func main() {
	var a, b int
	var operator string
	var result int

	fmt.Println("请输入算式，字符与字符间需要插入空格：")
	fmt.Scanln(&a, &operator, &b)

	switch operator {
	case "+":
		result = Calculator(a, b, Add)
	case "-":
		result = Calculator(a, b, Sub)
	case "*":
		result = Calculator(a, b, Multi)
	case "/":
		result = Calculator(a, b, Div)
	default:
		panic("不支持的符号")
	}

	fmt.Println(result)
}
