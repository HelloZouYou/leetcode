// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2i30g/
package main

import (
	"fmt"
)

// 删除排序数组中的重复项
// 双指针
func removeDuplicates(nums []int) int {
	// 先处理边界
	l := len(nums)
	if l == 0 {
		return 0
	}
	left := 0

	// 双指针
	// 左指针永远指向已经确定的最后一个不重复项，并原地等待
	// 右指针不停去找下一个不重复项，找到后送到左指针面前，并让左指针右移继续等待
	for right := 1; right < l; right++ {
		if nums[right] != nums[left] {
			left++
			nums[left] = nums[right]
		}
	}
	return left + 1
}

// 买卖股票的最佳时机 II
// 贪心算法
func maxProfit(prices []int) int {
	// 处理边界，不足两天的没得赚
	l := len(prices)
	if l < 2 {
		return 0
	}
	want := 0
	// 最大利润即是每个上涨的日子都完成交易，每个下跌的日子都能避开
	// 因此只要某天价格比前天大，就加到利润want中
	for i := 1; i < l; i++ {
		if prices[i] > prices[i-1] {
			want += prices[i] - prices[i-1]
		}
	}
	return want
	// 处理边界
	// l := len(prices)
	// if l < 2 {
	// 	return 0
	// }
	// 与上不同，没有直接加上每天的收益，而是用max存储阶段的最大收益，最后统一结算
	// min, max, want := prices[0], prices[0], 0
	// for i := 1; i < l; i++ {
	// 	if prices[i] < max { // 找到一个比记录max小的，则进行一次结算
	// 		want += max - min
	// 		min, max = prices[i], prices[i]
	// 	} else {
	// 		max = prices[i]
	// 	}
	// }
	// return want + max - min
}

// 旋转数组
func rotate(nums []int, k int) {
	// 处理边界
	// l := len(nums)
	// k = k % l // 移动个数超过数组长度，取余
	// if k == 0 {
	// 	return
	// }
	// 直接截取前后部分调换
	// i := l - k
	// tmp := append([]int{}, nums[0:i]...)
	// nums = append(nums[:0], nums[i:]...)
	// nums = append(nums, tmp...)

	// 处理边界，余数判断实际是否需要移动
	// l := len(nums)
	// if k%l == 0 {
	// 	return
	// }
	// 依次跳跃移动，每移动一个，用tmp记录被挤出来的值，并为其找要移动到的位置
	// 遇到循环，即已经处理过的坑位，则用visited记录，然后位置v+1找下一个循环
	// tmp, i := nums[0], 0
	// visited, v := map[int]struct{}{}, 0
	// for v < l {
	// 	i = (i + k) % l
	// 	if _, ok := visited[i]; ok {
	// 		i++
	// 		tmp = nums[i]
	// 		continue
	// 	}
	// 	nums[i], tmp = tmp, nums[i]
	// 	visited[i] = struct{}{}
	// 	v++
	// }

	// 需要先找出规律
	// 先全部翻转，再翻转前k个，再翻转后l-k个
	// l := len(nums)
	// k = k % l
	// if k == 0 {
	// 	return
	// }
	// for i, e := 0, l-1; i < e; {
	// 	nums[i], nums[e] = nums[e], nums[i]
	// 	i++
	// 	e = l - i - 1
	// }
	// for i, e := 0, k-1; i < e; {
	// 	nums[i], nums[e] = nums[e], nums[i]
	// 	i++
	// 	e = k - i - 1
	// }
	// for i, e := k, l-1; i < e; {
	// 	nums[i], nums[e] = nums[e], nums[i]
	// 	i++
	// 	e = l - i - 1 + k
	// }

	l := len(nums)
	k = k % l
	if k == 0 {
		return
	}
	nums = append(nums[:0], append(nums[l-k:], nums[:l-k]...)...)
}

// 存在重复元素
func containsDuplicate(nums []int) bool {
	// 全部遍历，用临时map记录，发现已被记录过的即为重复元素
	v := map[int]struct{}{}
	for _, i := range nums {
		if _, ok := v[i]; ok {
			return true
		}
		v[i] = struct{}{}
	}
	return false
}

// 只出现一次的数字
func singleNumber(nums []int) int {
	// 处理边界
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l <= 2 {
		return nums[0]
	}
	// 根据异或的规律，一个数异或自己为0
	// 0异或一个数则为该数
	// 因此所有元素异或后的值，则一定为结果
	var res int
	for i := 0; i < l; i++ {
		res ^= nums[i]
	}
	return res
}

// 两个数组的交集 II
func intersect(nums1 []int, nums2 []int) []int {
	// 处理边界
	l1 := len(nums1)
	l2 := len(nums2)
	if l1 == 0 || l2 == 0 {
		return []int{}
	}

	// 先记录a数组所有元素的出现次数，用于给b的元素一一对应
	tmp := make(map[int]int)
	for _, item := range nums1 {
		tmp[item]++
	}

	// 然后遍历b数组，找不到的直接原地删掉
	// 每找到一个存在的元素，将其对应个数减一
	for i := 0; i < l2; i++ {
		if v, ok := tmp[nums2[i]]; ok && v > 0 {
			tmp[nums2[i]]--
		} else {
			nums2 = append(nums2[:i], nums2[i+1:]...)
			l2--
			i--
		}
	}
	return nums2
}

// 有效的数独
func isValidSudoku(board [][]byte) bool {
	// 处理边界
	l := len(board)
	if l != 9 {
		return false
	}
	// 用三个map分别存储横向、竖向、3*3
	h := map[int]map[byte]struct{}{}
	s := map[int]map[byte]struct{}{}
	q := map[int]map[byte]struct{}{}
	// 全部初始化好
	for i := 0; i < l; i++ {
		h[i] = make(map[byte]struct{})
		s[i] = make(map[byte]struct{})
		q[i] = make(map[byte]struct{})
	}
	for i, g := range board { // i为行
		for k, v := range g { // k为列， v为值
			if _, ok := h[i][v]; v != '.' && ok { // 判断数字在该行是否已经出现过，出现过则无效
				return false
			}
			h[i][v] = struct{}{}
			if _, ok := s[k][v]; v != '.' && ok { // 判断数字在该列是否已经出现过，出现过则无效
				return false
			}
			s[k][v] = struct{}{}
			if _, ok := q[(i/3)*3+k/3][v]; v != '.' && ok { // 判断数字是否在当前所属的3*3出现过
				return false
			}
			// i/3*3+k/3 用于确定处于哪个3*3
			q[(i/3)*3+k/3][v] = struct{}{}
		}
	}
	return true
}

// 加一
func plusOne(digits []int) []int {
	l := len(digits)
	if l == 0 {
		return []int{}
	}

	// 从最后面的数依次判断，tmp记录是否需要进1
	tmp := 1
	for i := l - 1; tmp != 0 && i >= 0; i-- {
		digits[i] += tmp
		tmp = digits[i] / 10
		digits[i] %= 10
	}
	// 最后补充判断是否还需要进1
	if tmp != 0 {
		res := make([]int, 0, l+1)
		digits = append(append(res, tmp), digits...)
	}
	return digits
}

// 移动零
func moveZeroes(nums []int) {
	l := len(nums)
	if l == 1 {
		return
	}
	// 遍历判断，存在0则原地截取后，末位补0
	for k, i := 0, 0; i < l; i++ {
		if nums[k] == 0 {
			nums = append(nums[:k], nums[k+1:]...)
			nums = append(nums, 0)
			k-- // 截断后需要保持指针不动
		}
		k++
	}
}

// 两数之和
func twoSum(nums []int, target int) []int {
	var (
		tmp = map[int]int{}
	)
	// 遍历判断，用tmp留下自己的目标值，以及自己的位置k
	for k, v := range nums {
		if key, ok := tmp[v]; ok { // 找到目标值v，则返回目标值的位置key和自己的位置k
			return []int{key, k}
		}
		tmp[target-v] = k // 用目标数记录自己的位置k
	}
	return []int{}
}

// 旋转图像
func rotate1(matrix [][]int) {
	l := len(matrix)
	if l <= 1 {
		return
	}
	// 先上下交换
	for i := 0; i < l/2; i++ {
		matrix[i], matrix[l-i-1] = matrix[l-i-1], matrix[i]
	}
	// 再对角线交换
	for i := 1; i < l; i++ {
		for k := 0; k < i; k++ {
			matrix[i][k], matrix[k][i] = matrix[k][i], matrix[i][k]
		}
	}
	fmt.Println(matrix)
}
