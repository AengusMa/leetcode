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
	"github.com/bndr/gotabulate"
)

func testSort(name string, f func([]int)) {
	nums := utils.GetArray(20, 50)
	f(nums)
	fmt.Println(name, nums)
}

func testPerformance(name string, f func([]int), len, max int) {
	nums := utils.GetArray(len, max)
	start := time.Now()
	f(nums)
	cost := time.Since(start)
	fmt.Print(name, "时间(无序数组)cost=", cost)
	nums = utils.GetSortArray(len, false)
	start = time.Now()
	f(nums)
	cost = time.Since(start)
	fmt.Print(name, "时间(有序数组)cost=", cost)
	nums = utils.GetSameArray(len)
	start = time.Now()
	f(nums)
	cost = time.Since(start)
	fmt.Println(name, "时间(重复数组)cost=", cost)
}
func testTable() {
	row1 := []interface{}{"john", 20, "ready"}
	row2 := []interface{}{"bndr", 23, "ready"}
	// Create an object from 2D interface array
	t := gotabulate.Create([][]interface{}{row1, row2})
	// Set the Headers (optional)
	t.SetHeaders([]string{"age", "status"})
	// Set the Empty String (optional)
	t.SetEmptyString("None")
	// Set Align (Optional)
	t.SetAlign("right")
	// Print the result: grid, or simple
	fmt.Println(t.Render("grid"))
}
func testBatchPerformance(data map[string]func([]int), len, max int) {
	var tableResult [][]string
	for key, fun := range data {
		var tmp []string
		tmp = append(tmp, key)
		nums := utils.GetArray(len, max)
		start := time.Now()
		fun(nums)
		cost := time.Since(start)
		tmp = append(tmp, cost.String())

		nums = utils.GetSortArray(len, false)
		start = time.Now()
		fun(nums)
		cost = time.Since(start)
		tmp = append(tmp, cost.String())

		nums = utils.GetSameArray(len)
		start = time.Now()
		fun(nums)
		cost = time.Since(start)
		tmp = append(tmp, cost.String())
		tableResult = append(tableResult, tmp)
	}
	t := gotabulate.Create(tableResult)
	t.SetHeaders([]string{"未排序数组", "排序数组", "重复数组"})
	fmt.Println(t.Render("simple"))
}

func testCorrect() {
	testSort("堆排序", heapSort.Sort)
	testSort("归并排序", mergeSort.Sort)
	testSort("希尔排序", shellSort.Sort)
	testSort("选择排序", simpleSort.SelectSort)
	testSort("插入排序", simpleSort.InsertSort)
	testSort("冒泡排序", simpleSort.BubbleSort)
	testSort("快速排序V1", quickV1.Sort)
	testSort("快速排序V2", quickV2.Sort)
}
func main() {
	// 测试正确性
	testCorrect()
	//2000000000
	// testPerformance("快速排序V1", quickV1.Sort, 20, 100)
	data := make(map[string]func([]int))
	data["快速排序V2"] = quickV2.Sort
	data["自带排序"] = sort.Ints
	data["希尔排序"] = shellSort.Sort
	testBatchPerformance(data, 200000, 50000000)
	// testTable()
	// testMy()
}
func testMy() {
	data := [][]string{{"1ew", "中文中文中卫", "2we"}, {"3242", "wefw22222223232", "2we"}}
	utils.Print(data)
}
