package plalindrome

import "fmt"

// 获取字符串中包含回文字符串的数量
// s[i][j] : 从i到j是否是字符串, 如果s[i]==s[j], 那么判断s[i+1][j-1]是否是回文字符串
func Num(str string) int {
	length := len(str)
	var dp = make([][]bool, length, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length, length)
	}
	bools := 0
	for i := length - 1; i >= 0; i-- { //反着来,不然d[i][j]依赖dp[i+1][j-1]时,dp[i+1][j-1]还没计算出结果
		for j := i; j < length; j++ {
			if str[i] != str[j] {
				dp[i][j] = false
				continue
			}
			if j-i <= 2 {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i+1][j-1]
			}
			if dp[i][j] {
				fmt.Println(str[i : j+1])
				bools++
			}
		}
	}
	return bools
}
