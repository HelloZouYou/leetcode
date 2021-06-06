// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2kje3/
package main

import "fmt"

// 位1的个数
func hammingWeight(num uint32) int {
	nums := 0
	for ; num != 0; num >>= 1 {
		if num&1 == 1 {
			nums++
		}
	}
	return nums
}

// 汉明距离
func hammingDistance(x int, y int) int {
	num, nums := x^y, 0
	for ; num != 0; num >>= 1 {
		if num&1 == 1 {
			nums++
		}
	}
	return nums
}

// 颠倒二进制位
func reverseBits(num uint32) uint32 {
	var res uint32
	for k := 0; k < 32; k++ {
		res <<= 1
		res += num & 1
		num >>= 1
	}
	return res
}

// 杨辉三角
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		for k := 1; k < i; k++ {
			res[i][k] = res[i-1][k-1] + res[i-1][k]
		}
		res[i][0], res[i][i] = 1, 1
	}
	return res
}

// 有效的括号
func isValid(s string) bool {
	l := len(s)
	if l < 2 {
		return false
	}
	tmp := make([]rune, 0, l)
	for _, v := range s {
		k, t := len(tmp), '0'
		switch v {
		case '}':
			t = '{'
		case ']':
			t = '['
		case ')':
			t = '('
		}
		if k > 0 && tmp[k-1] == t {
			tmp = tmp[:k-1]
			continue
		}
		tmp = append(tmp, v)
	}
	return len(tmp) == 0
}

// 缺失数字
func missingNumber(nums []int) int {
	l := len(nums)
	res := 0
	for i := 0; i < l; i++ {
		res ^= i ^ nums[i]
	}
	res ^= l
	return res
}

func main() {
	fmt.Println(missingNumber([]int{1}))
}
