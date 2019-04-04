package sort

import (
	"fmt"
	"testing"

	heapSort "./heapSort"
	quickV3 "./quickV3"
	utils "./utils"
)

func TestHeapSort(t *testing.T) {
	nums := utils.GetArray(20, 50)
	heapSort.Sort(nums)
	if !utils.IsSorted(nums, false) {
		t.Error("出错了")
		return
	}
	fmt.Println(nums)
}

func BenchmarkQuickSortV3(b *testing.B) {
	nums := utils.GetArray(b.N, 20000000)
	quickV3.Sort(nums)
	if !utils.IsSorted(nums, false) {
		b.Error("出错了")
		return
	}
}

func TestMain(t *testing.T) {
	main()
}
