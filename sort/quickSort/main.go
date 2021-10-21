package main

import "fmt"

/*
	快速排序（quickSort）是对冒泡排序的一种改进。基本思想是：通过一趟排序将要排序的数据分割成独立的两部分，其中
		一部分的所有数据都比另外一部分的所有数据都要小，然后再按此方法对这两部分别进行快速排序，整个排序过程可以递归进行
		，依此达到整个数据变成有序序列
*/

/*
	left: 数组左边的下标
	right：所在右边的下标
	array：要排序的数组
*/
func QuickSort(left int, right int, array *[6]int) {
	l := left
	r := right
	// 找到中间的数
	pivot := array[(left+right)/2]
	temp := 0
	// for循环的目的是将比pivot小的数放到左边，比pivot大的说放到右边
	for l < r {
		// 从pivot的左边找到大于等于pivot的值
		for array[l] < pivot {
			l++
		}
		// 从pivot的右边找到小于等于pivot的值
		for array[r] > pivot {
			r--
		}
		// 到中间位置了，本次循环完成
		if l >= r {
			break
		}
		// 交换
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		// 优化，相等的情况不需要交换
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}

	// 如果 l == r，再移动下
	if l == r {
		l++
		r--
	}

	// 向后递归，继续排序
	// 向左递归
	if left < r {
		QuickSort(left, r, array)
	}
	// 向右递归
	if right > l {
		QuickSort(l, right, array)
	}
}
func main() {
	arr := [6]int{-9, 78, 0, 23, -53, 70}
	QuickSort(0, len(arr)-1, &arr)
	fmt.Println(arr)
}
