package main

import (
	"fmt"
	"sort"
	"time"

	heapSort "./heapSort"
	mergeSort "./mergeSort"
	quickV1 "./quickV1"
	quickV2 "./quickV2"
	shellSort "./shellSort"
	simpleSort "./simpleSort"
	utils "./utils"
)

func testQuickV1(len, max int) {
	nums := utils.GetArray(len, max)
	start := time.Now()
	quickV1.Sort(nums)
	cost := time.Since(start)
	fmt.Print("快速排序时间(无序)cost=", cost)
	nums = utils.GetSortArray(len, false)
	start = time.Now()
	quickV1.Sort(nums)
	cost = time.Since(start)
	fmt.Print("快速排序时间(有序数组)cost=", cost)
	nums = utils.GetSameArray(len)
	start = time.Now()
	quickV1.Sort(nums)
	cost = time.Since(start)
	fmt.Println("快速排序时间(重复数组)cost=", cost)
}

func testQuickV2(len, max int) {
	nums := utils.GetArray(len, max)
	start := time.Now()
	quickV2.Sort(nums)
	cost := time.Since(start)
	fmt.Print("快速排序时间(无序)cost=", cost)
	nums = utils.GetSortArray(len, false)
	start = time.Now()
	quickV2.Sort(nums)
	cost = time.Since(start)
	fmt.Print("快速排序时间(有序数组)cost=", cost)
	nums = utils.GetSameArray(len)
	start = time.Now()
	quickV2.Sort(nums)
	cost = time.Since(start)
	fmt.Println("快速排序时间(重复数组)cost=", cost)
}
func testInnerSort(len, max int) {
	nums := utils.GetArray(len, max)
	start := time.Now()
	sort.Ints(nums)
	cost := time.Since(start)
	fmt.Print("自带排序时间(无序)cost=", cost)
	nums = utils.GetSortArray(len, false)
	start = time.Now()
	sort.Ints(nums)
	cost = time.Since(start)
	fmt.Print("自带排序时间(有序数组)cost=", cost)
	nums = utils.GetSameArray(len)
	start = time.Now()
	sort.Ints(nums)
	cost = time.Since(start)
	fmt.Println("快速排序时间(重复数组)cost=", cost)
}
func testHeapSort() {
	nums := utils.GetArray(20, 50)
	heapSort.Sort(nums)
	fmt.Println("堆排序结果：", nums)
}
func testMergeSort() {
	nums := utils.GetArray(20, 50)
	mergeSort.Sort(nums)
	fmt.Println("归并排序结果：", nums)
}
func testShellSort() {
	nums := utils.GetArray(20, 50)
	shellSort.Sort(nums)
	fmt.Println("归并排序结果：", nums)
}
func testSimpleSort() {
	nums := utils.GetArray(20, 50)
	simpleSort.SelectSort(nums)
	fmt.Println("选择排序结果：", nums)

	nums = utils.GetArray(20, 50)
	simpleSort.InsertSort(nums)
	fmt.Println("插入排序结果：", nums)

	nums = utils.GetArray(20, 50)
	simpleSort.BubbleSort(nums)
	fmt.Println("冒泡排序结果：", nums)
}
func testQuickSort() {
	nums := utils.GetArray(20, 50)
	quickV1.Sort(nums)
	fmt.Println("快速排序V1结果：", nums)

	nums = utils.GetArray(20, 50)
	quickV2.Sort(nums)
	fmt.Println("快速排序V2结果：", nums)
}
func main() {
	testHeapSort()
	testMergeSort()
	testShellSort()
	testSimpleSort()
	testQuickSort()
	//2000000000
	// testQuickV1(20, 50000000)
	testQuickV2(200000, 50000000)
	testInnerSort(200000, 50000000)
}
