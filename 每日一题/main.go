package main

import "fmt"

// 盈利计划
// https://leetcode-cn.com/problems/profitable-schemes/
// func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
// 	const mod int = 1e9 + 7
// 	l, wait, res := len(group), make([]int, n+1), 0
// 	wait[0] = 1
// 	for i := 0; i < l; i++ { // 遍历所有的工作
// 		for k := 0; k <= n; k++ { // 遍历所有能出的人数

// 		}
// 	}
// 	for k, v := range profit {
// 		if group[k] <= n {
// 			tmp := [][]int{{group[k], v}}
// 			for _, w := range wait {
// 				if w[0]+group[k] <= n {
// 					tmp = append(tmp, []int{w[0] + group[k], w[1] + v})
// 				}
// 			}
// 			wait = append(wait, tmp...)
// 		}
// 	}
// 	if minProfit == 0 {
// 		wait = append(wait, []int{0, 0})
// 	}
// 	num := 0
// 	for _, v := range wait {
// 		if v[1] >= minProfit {
// 			num++
// 		}
// 	}
// 	return num % mod
// }
func profitableSchemes(n, minProfit int, group, profit []int) (sum int) {
	const mod int = 1e9 + 7
	ng := len(group)
	dp := make([][][]int, ng+1)
	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, minProfit+1)
		}
	}
	dp[0][0][0] = 1
	for i, members := range group {
		earn := profit[i]
		for j := 0; j <= n; j++ {
			for k := 0; k <= minProfit; k++ {
				if j < members {
					dp[i+1][j][k] = dp[i][j][k]
				} else {
					dp[i+1][j][k] = (dp[i][j][k] + dp[i][j-members][max(0, k-earn)]) % mod
				}
			}
		}
	}
	for _, d := range dp[ng] {
		sum = (sum + d[minProfit]) % mod
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(profitableSchemes(10, 5, []int{2, 3, 2, 3}, []int{3, 7, 8, 9}))
}
