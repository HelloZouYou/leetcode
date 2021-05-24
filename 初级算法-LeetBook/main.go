// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2gy9m/
package main

import (
	"fmt"
	"math"
)

// 删除排序数组中的重复项
// 双指针
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	k := 0
	for i := 1; i < l; i++ {
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}

// 买卖股票的最佳时机 II
// 贪心算法
func maxProfit(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}
	want := 0
	for i := 1; i < l; i++ {
		if prices[i] > prices[i-1] {
			want += prices[i] - prices[i-1]
		}
	}
	return want
	// l := len(prices)
	// if l < 2 {
	// 	return 0
	// }
	// min, max, want := prices[0], prices[0], 0
	// for i := 1; i < l; i++ {
	// 	if prices[i] < max {
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
	// l := len(nums)
	// k = k % l
	// if k == 0 {
	// 	return
	// }
	// i := l - k
	// tmp := append([]int{}, nums[0:i]...)
	// nums = append(nums[:0], nums[i:]...)
	// nums = append(nums, tmp...)

	// l := len(nums)
	// if k%l == 0 {
	// 	return
	// }
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
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l <= 2 {
		return nums[0]
	}
	var res int
	for i := 0; i < l; i++ {
		res ^= nums[i]
	}
	return res
}

// 两个数组的交集 II
func intersect(nums1 []int, nums2 []int) []int {

	return nil
}

// 加一
func plusOne(digits []int) []int {
	l := len(digits)
	if l == 0 {
		return []int{}
	}
	tmp := 1
	for i := l - 1; tmp != 0 && i >= 0; i-- {
		digits[i] += tmp
		tmp = digits[i] / 10
		digits[i] %= 10
	}
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
	for k, i := 0, 0; i < l; i++ {
		if nums[k] == 0 {
			nums = append(nums[:k], nums[k+1:]...)
			nums = append(nums, 0)
			k--
		}
		k++
	}
}

// 两数之和
func twoSum(nums []int, target int) []int {
	var (
		tmp = map[int]int{}
	)

	for k, v := range nums {
		if key, ok := tmp[v]; ok {
			return []int{key, k}
		} else {
			tmp[target-v] = k
		}
	}
	return []int{}
}

// 旋转图像
func rotate2(matrix [][]int) {
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

// 反转字符串
func reverseString(s []byte) {
	l := len(s)
	if l <= 1 {
		return
	}
	k := l / 2
	for i := 0; i < k; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}

// 整数反转
func reverse(x int) int {
	if x < 10 && x > -10 {
		return x
	}
	res := 0
	for x != 0 {
		s := x % 10
		m := res*10 + s
		if m >= math.MaxInt32 || m <= math.MinInt32 {
			return 0
		}
		x, res = x/10, m
	}
	return res
}

// 字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	l := len(s)
	if l == 1 {
		return 0
	}
	// m, t := map[byte]int{}, 0
	// for i := l - 1; i > 0; i-- {
	// 	if v, ok := m[s[i]]; ok {
	// 		delete(m, s[i])
	// 		continue
	// 	}
	// 	m[s[i]] = 1
	// }
	// return m[t]
}

func main() {
	fmt.Println(firstUniqChar("abcdefg"))
}
