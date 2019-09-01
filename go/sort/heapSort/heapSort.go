package heapsort

import "fmt"

// Sort 排序
// 大顶堆：arr[i] >= arr[2i+1] && arr[i] >= arr[2i+2]
// i结点的父结点下标就为(i – 1) / 2。它的左右子结点下标分别为2 * i + 1和2 * i + 2。
func Sort(nums []int) {
	heapSort(nums)
}
func heapSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		siftUp(nums, i)
	}
	// fmt.Println("构建堆后：", nums)
	for n := len(nums) - 1; n > 0; n-- {
		// 因为每次完成堆的构造后, 根节点为最大值节点
		// 将根节点与最后一个节点交换
		nums[0], nums[n] = nums[n], nums[0]
		for i := 0; i < n; i++ {
			siftUp(nums, i)
		}
	}
}
func heapSort1(nums []int) {
	// 构建堆，因为堆是完全二叉树的特性, 所以下标小于等于 array.length / 2 的节点为非叶子节点
	// 采用下沉的方式 从下往上构建子堆
	for i := len(nums) / 2; i >= 0; i-- {
		siftDown(nums, i, len(nums))
	}
	// 因为每次完成堆的构造后, 根节点为最大值节点
	// 将根节点与最后一个节点交换
	fmt.Println("构建堆后：", nums)
	for n := len(nums) - 1; n > 0; n-- {
		nums[0], nums[n] = nums[n], nums[0]
		for i := n / 2; i >= 0; i-- {
			// 排除有序的节点
			// 重新构造堆
			siftDown(nums, i, n)
		}
	}
}

// 下沉
func siftDown(nums []int, k, len int) {
	// 获取左子节点索引
	child := 2*k + 1
	// 判断是否存在子节点
	for child < len {
		// 判断左右子节点 查找最大的子节点
		if child+1 < len && nums[child+1] > nums[child] {
			child++
		}
		if nums[k] >= nums[child] {
			break
		} else {
			nums[k], nums[child] = nums[child], nums[k]
			// 修改k的指向进行下一次下沉
			k = child
			child = 2*k + 1
		}
	}
}

// 上浮
func siftUp(nums []int, k int) {
	// k == 0 时表明上浮到根节点，结束上浮操作
	for k > 0 {
		// 获取父节点索引
		parent := (k - 1) >> 1
		// 小于等于父节点时退出，结束上浮操作
		if nums[k] <= nums[parent] {
			break
		} else {
			nums[k], nums[parent] = nums[parent], nums[k]
			k = parent
		}

	}
}
