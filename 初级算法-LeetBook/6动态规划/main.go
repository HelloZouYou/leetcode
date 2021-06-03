// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2qpms/
package main

import (
	"fmt"
)

// 爬楼梯
func climbStairs(n int) int {
	return ClimbStairs(n, 1, 1)
}
func ClimbStairs(n, a, b int) int {
	if n == 1 {
		return b
	}
	return ClimbStairs(n-1, b, a+b)
}

// 买卖股票的最佳时机
func maxProfit(prices []int) int {
	min := 99999
	res := 0
	length := len(prices)
	for i := 0; i < length; i++ {
		if prices[i] < min {
			min = prices[i]
		} else if profit := prices[i] - min; profit > res {
			res = profit
		}
	}
	return res
}

// 最大子序和
func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	//取出第一位为基准值
	sum, res := nums[0], nums[0]
	//下标从1开始
	for i := 1; i < l; i++ {
		// 对当前i而言，前面如果已经加为负值，则直接丢弃，以当前作为新的开始
		// 而若i之前相加就是正值，则继续加
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		// 总是存储历史最大值
		if sum > res {
			res = sum
		}
	}
	return res
}

// 打家劫舍
func rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	// 光顾完的每一家的最大金额，可能劫了也可能没劫
	got1, got2 := 0, nums[0]
	// 从第二家开始，第一家定义上述两种可能
	for i := 1; i < l; i++ {
		// 进入迭代，got1和got2表示上一家的两种可能下的积累最大金额
		tmp := got2 // tmp暂存最大值
		if got1 > got2 {
			tmp = got1
		}
		// got2表示如果干掉这第i家，则上一家得没有干掉过
		// got1表示如果不干掉第i家，则上一家劫没劫都行，取上一家两种可能的最大值
		got2, got1 = nums[i]+got1, tmp
	}
	// 最后那种情况的积累金额大则采用该方案
	if got1 > got2 {
		return got1
	}
	return got2
}

func main() {
	fmt.Println(maxProfit([]int{4, 11, 2, 1, 7}))
}
