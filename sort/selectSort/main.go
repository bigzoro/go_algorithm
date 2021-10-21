package main

import "fmt"

func SelectSort1(arr *[5]int) {

	// 选出第一大的元素：假设第0个位置的元素是最大的，然后和后面的元素进行比较，
	//  如果后面的元素比第0个位置的元素大，就替换
	max := arr[0]
	maxIndex := 0

	// 比较
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			maxIndex = i
		}
	}

	// 替换
	if maxIndex != 0 {
		arr[0], arr[maxIndex] = arr[maxIndex], arr[0]
	}

	// 选出第二大元素
	max = arr[1]
	maxIndex = 1
	for i := 2; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			maxIndex = i
		}
	}

	if maxIndex != 1 {
		arr[1], arr[maxIndex] = arr[maxIndex], arr[1]
	}

	// 选出第三大元素
	max = arr[2]
	maxIndex = 2
	for i := 3; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			maxIndex = i
		}
	}

	if maxIndex != 2 {
		arr[2], arr[maxIndex] = arr[maxIndex], arr[2]
	}

	// 选出第四大元素
	max = arr[3]
	maxIndex = 3

	for i := 4; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			maxIndex = i
		}
	}

	if maxIndex != 3 {
		arr[3], arr[maxIndex] = arr[maxIndex], arr[3]
	}

	// 元素排序完成

	fmt.Println(arr)
}

func SelectSort2(arr *[5]int) {

	for i := 0; i < len(arr); i++ {
		// 假设第i个位置的元素是最大的
		max := arr[i]
		maxIndex := i
		// 与第i个位置的元素进行比较
		for j := i; j < len(arr); j++ {
			if max < arr[j] {
				max = arr[j]
				maxIndex = j
			}
		}

		// 替换
		if maxIndex != i {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		}
	}

	fmt.Println(arr)
}
func main() {
	arr := [5]int{10, 34, 19, 100, 80}
	// SelectSort1(&arr)
	SelectSort2(&arr)
}
