package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func qSort(arr []int) {
	if arr == nil || len(arr) == 0 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}
func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	// 数据分为三个区
	posL, posR := partition(arr, low, high)
	quickSort(arr, low, posL)
	quickSort(arr, posR, high)

}

func partition(arr []int, left, right int) (int, int) {
	p_index := choosePivotFirst(arr, left, right)
	if p_index != left {
		arr[p_index], arr[left] = arr[left], arr[p_index]
	}
	p := arr[left]
	i := left + 1
	sameLen := 1
	for j := left + 1; j <= right; j++ {
		if arr[j] <= p {
			arr[j], arr[i] = arr[i], arr[j]
			// 相等的元素放在头部
			if arr[i] == p {
				arr[i], arr[left+sameLen] = arr[left+sameLen], arr[i]
				sameLen++
			}
			i++
		}
	}
	posRight := i
	// 相等的元素移动中间
	pRight := 0
	fmt.Println(arr[left : right+1])
	for ; i > left+1 && arr[i-1] != p && arr[left+pRight] == p; pRight++ {
		arr[left+pRight], arr[i-1] = arr[i-1], arr[left+pRight]
		i--
		if p == 4 {
			fmt.Println("===", i)
		}
	}
	fmt.Println(arr[left : right+1])
	return posRight - sameLen - 1, posRight
}
func choosePivotFirst(arr []int, left, right int) int {
	return left
}
func choosePivotMedianOfThree(arr []int, left, right int) int {
	mid := left + ((right - left) >> 1)
	if (arr[mid] > arr[right] && arr[mid] < arr[left]) || (arr[mid] > arr[left] && arr[mid] < arr[right]) {
		return mid
	} else if (arr[left] > arr[right] && arr[left] < arr[mid]) || (arr[left] > arr[mid] && arr[left] < arr[right]) {
		return left
	} else {
		return right
	}
}
func main() {
	// nums := []int{1, 1, 2, 2, 2, 2, 2, 3, 2, 2, 1, 1, 1, 1, 1, 1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 1, 1, 14, 4, 4, 1, 1, 1, 4, 4, 4, 4, 4, 2, 5, 4, 4, 7, 4, 4, 4, 4, 4, 3, 1, 1, 1, 1, 1}
	// nums := getArray(20, 20)
	nums := []int{8, 15, 4, 15, 11, 2, 13, 12, 4, 14, 0, 10, 6, 18, 9, 15, 6, 13, 12, 14}
	fmt.Println("原始：", nums)
	qSort(nums)
	fmt.Println(nums)
	// testQuick(100)2000000000
	// testSortInts(10000000)
}
func testQuick(len int) {
	nums := getArray(len, 2000000000)
	start := time.Now()
	qSort(nums)
	cost := time.Since(start)
	fmt.Printf("快速排序时间(无序)cost=[%s]", cost)
	print("=========")
	nums = getSortArray(len, false)
	start = time.Now()
	qSort(nums)
	cost = time.Since(start)
	fmt.Printf("快速排序时间(有序数组)cost=[%s]", cost)
	print("=========")
	nums = getSameArray(len)
	start = time.Now()
	qSort(nums)
	cost = time.Since(start)
	fmt.Printf("快速排序时间(重复数组)cost=[%s]", cost)
	println()
}
func testSortInts(len int) {
	nums := getArray(len, 2000000000)
	start := time.Now()
	sort.Ints(nums)
	cost := time.Since(start)
	fmt.Printf("自带排序时间(无序)cost=[%s]", cost)
	print("=========")
	nums = getSortArray(len, false)
	start = time.Now()
	sort.Ints(nums)
	cost = time.Since(start)
	fmt.Printf("自带排序时间(有序数组)cost=[%s]", cost)
	println()
}
func getSortArray(len int, desc bool) []int {
	result := []int{}
	if desc {
		for i := len - 1; i > 0; i-- {
			result = append(result, i)
		}
	} else {
		for i := 0; i < len; i++ {
			result = append(result, i)
		}
	}
	return result
}

func getArray(len, max int) []int {
	rand.Seed(time.Now().Unix())
	result := []int{}
	for i := 0; i < len; i++ {
		result = append(result, rand.Intn(max))
	}
	return result
}

func getSameArray(len int) []int {
	result := []int{}
	for i := 0; i < len; i++ {
		result = append(result, 10)
	}
	return result
}
