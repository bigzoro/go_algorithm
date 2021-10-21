package main

import (
	"fmt"
	"strconv"
)

/*
	思路：
		1. 创建两个栈，numStack、operStack
		2. numStack存放数、operStack存放操作符
		3. index := 0
		4. exp计算表达式是一个数字，则直接入numStack
		5. 如果发现是一个运算符
			a. 如果operStack是一个空栈，直接入栈
			b. 如果operStack不是一个空栈
				i. 如果发现operStack栈顶的运算符的优先级大于等于当前准备入栈的运算符的优先级，就从符号栈pop出，
					并从数栈也pop两个数，进行运算，运算后的结果再重新入栈到数栈，符号再入符号栈
				ii. 否则，运算符直接入栈
		6. 如果扫描表达式完毕，依次从符号栈取出符号，然后从数栈取出两个数，运算后的结果，入数栈，直到符号栈为空
*/
// 使用数组模拟栈
type Stack struct {
	// 栈的容量
	MaxTop int
	// 栈堆
	Top int
	// 数组模拟栈
	Array [20]int
}

func (s *Stack) Push(val int) {
	// fmt.Println(len(s.Array))
	// 判断栈是否满了
	if s.Top >= s.MaxTop-1 {
		fmt.Println("the stack is full")
		return
	}

	s.Top++
	// 入栈
	s.Array[s.Top] = val
	return
}

// 出栈
func (s *Stack) Pop() (data int) {
	// 判断栈是否为空
	if s.Top == -1 {
		fmt.Println("the stack is empty")
		return
	}

	data = s.Array[s.Top]
	s.Top--
	return
}

// 遍历栈
func (s *Stack) List() {
	// 判断栈是否满了
	if s.Top == -1 {
		fmt.Println("the stack is empty")
		return
	}
	for i := s.Top; i >= 0; i-- {
		fmt.Println(s.Array[i])
	}
}

// 判断一个字符是不是运算符[+ - * /]
func (s *Stack) IsOper(val int) bool {
	// 这些数字是加减乘除对于的ASSIC码
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

// 运算的方法
func (s *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误")
	}
	return res
}

// 返回某个运算符的优先级
func (s *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}

	return res
}
func main() {

	// 数栈
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	// 符号栈
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	// 表达式
	exp := "3+2*6-2"

	// 定义一个index, 帮助扫描exp
	index := 0

	// 辅助变量
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := ""

	for {
		// 循环遍历每一个字符
		ch := exp[index : index+1]
		// 字符转化为ASSIC码
		temp := int([]byte(ch)[0])
		if operStack.IsOper(temp) {
			// 说明是符号
			// 如果是一个空栈，直接入栈
			if operStack.Top == -1 {
				operStack.Push(temp)
			} else {
				if operStack.Priority(operStack.Array[operStack.Top]) >= operStack.Priority(temp) {
					num1 = numStack.Pop()
					num2 = numStack.Pop()
					oper = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)

					// 将计算结果重新入栈
					numStack.Push(result)

					// 将当前的符号压入符号栈
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}
		} else {
			// 说明是数字
			// 1.定义一个变量 keepNum string, 做多位数拼接
			keepNum += ch
			// 2. 每次要向index的后面字符测试一下，看看是不是运算符，然后处理
			//如果已经到表达最后，直接处理keepNum
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				//向index 后面测试看看是不是运算符 [index]
				if operStack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
		}

		//继续扫描
		//先判断index是否已经扫描到计算表达式的最后
		if index+1 == len(exp) {
			break
		}
		index++
	}
	//如果扫描表达式 完毕，依次从符号栈取出符号，然后从数栈取出两个数，
	//运算后的结果，入数栈，直到符号栈为空
	for {
		if operStack.Top == -1 {
			break //退出条件
		}
		num1 = numStack.Pop()
		num2 = numStack.Pop()
		oper = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)
		//将计算结果重新入数栈
		numStack.Push(result)

	}

	//如果我们的算法没有问题，表达式也是正确的，则结果就是numStack最后数
	res := numStack.Pop()
	fmt.Printf("表达式%s = %v", exp, res)

}
