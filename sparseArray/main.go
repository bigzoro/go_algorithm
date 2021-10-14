package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	row   int
	col   int
	value int
}

// 普通二维数组转化为稀疏数组
func arrayToSparseArray(array [][]int, filePath string) (err error) {

	// 1. 获取稀疏数组的行和列，及非零值的个数
	rowL := len(array)
	colL := len(array[0])
	num := 0
	// 遍历二维数组, 找非零值的数量
	for _, v1 := range array {
		for _, v2 := range v1 {
			if v2 != 0 {
				num++
			}
		}
	}

	// 声明一个Node的切片，为Node切片类型
	//  Node切片有row、col、value三个属性，用来存储每个非零值元素的信息
	var sparseArr []Node

	// 初始化数组
	node := Node{
		rowL,
		colL,
		num,
	}
	sparseArr = append(sparseArr, node)

	// 遍历稀疏数组，得到每个非零值信息，并保存到array
	for i, v1 := range array {
		for j, v2 := range v1 {
			if v2 != 0 {
				node = Node{i, j, v2}
				sparseArr = append(sparseArr, node)
			}
		}
	}

	fmt.Println("转换后的数组为：")
	for _, v := range sparseArr {
		fmt.Printf("%d %d %d\n", v.row, v.col, v.value)
	}

	// 把转换后的保存到本地
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, value := range sparseArr {
		data := fmt.Sprintf("%d %d %d\n", value.row, value.col, value.value)
		_, err = writer.WriteString(data)
		if err != nil {
			return
		}
	}
	writer.Flush()

	fmt.Println("写入数据成功")
	return
}

// 将稀疏数组转化为二维数组
func sparseArrToArr(filePath string) (array [][]int, error error) {
	// 读取本地文件, 获取到稀疏数组
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		error = err
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	node := Node{}
	var row, col, value int
	var sparseArr []Node
	for {
		data, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			error = err
			return
		}
		if err == io.EOF {
			break
		}

		data2 := strings.Fields(data)

		row, err = strconv.Atoi(data2[0])
		if err != nil {
			error = err
			return
		}
		col, err = strconv.Atoi(data2[1])
		if err != nil {
			error = err
			return
		}
		value, err = strconv.Atoi(data2[2])
		if err != nil {
			error = err
			return
		}

		node = Node{row, col, value}
		sparseArr = append(sparseArr, node)
	}

	// 把稀疏数组转化为二维切片
	// 初始化二维数组
	for i, v := range sparseArr {
		if i == 0 {
			row = v.row
			col = v.col
			// fmt.Println("v.row", v.row)
			// fmt.Println("v.col", v.col)
			array = make([][]int, row)
			for i := range array {
				array[i] = make([]int, col)
			}

			// 跳过本次循环
			continue
		}
		array[v.row][v.col] = v.value
	}

	return array, err
}
func main() {

	filePath := "./sparseArray.txt"

	// 测试数组
	array := [][]int{
		{0, 1, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
	}

	err := arrayToSparseArray(array, filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := sparseArrToArr(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range data {
		fmt.Println(v)
	}
}
