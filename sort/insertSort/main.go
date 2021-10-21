package main

import "fmt"

func InsertSort1(arr [5]int) {
	// 从后面选择一个元素，然后在排序过后的数组中，从后向前比较，找到合适的位置就插入

	// 先确定插入的位置和要插入的值
	insertVal := arr[1]
	// 待插入元素位置的前一个位置
	insertIndex := 1 - 1

	// 开始循环比较
	// 满足这些条件才循环
	for insertIndex >= 0 && arr[insertIndex] < insertVal {
		// 数据后移一位
		arr[insertIndex+1] = arr[insertIndex]
		insertIndex--
	}

	// 找到了合适的位置，进行插入
	// 判断是否在最后
	if insertIndex+1 != 1 {
		arr[insertIndex+1] = insertVal
	}

	// 进行第二次插入排序，排第三个元素
	insertVal = arr[2]
	insertIndex = 2 - 1

	for insertIndex >= 0 && arr[insertIndex] < insertVal {
		arr[insertIndex+1] = arr[insertIndex]
		insertIndex--
	}

	if insertIndex+1 != 2 {
		arr[insertIndex+1] = insertVal
	}

	// 进行第三次插入排序，排第四个元素
	insertVal = arr[3]
	insertIndex = 3 - 1

	for insertIndex >= 0 && arr[insertIndex] < insertVal {
		arr[insertIndex+1] = arr[insertIndex]
		insertIndex--
	}

	if insertIndex+1 != 3 {
		arr[insertIndex+1] = insertVal
	}

	// 进行第四次插入排序，排第五个元素
	insertVal = arr[4]
	insertIndex = 4 - 1

	for insertIndex >= 0 && arr[insertIndex] < insertVal {
		arr[insertIndex+1] = arr[insertIndex]
		insertIndex--
	}

	if insertIndex+1 != 4 {
		arr[insertIndex+1] = insertVal
	}
	fmt.Println(arr)
}

func InsertSort2(arr [5]int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex != i {
			arr[insertIndex+1] = insertVal
		}
	}

	fmt.Println(arr)
}

func main() {

	arr := [5]int{23, 0, 12, 56, 34}

	// InsertSort1(arr)
	InsertSort2(arr)
}
