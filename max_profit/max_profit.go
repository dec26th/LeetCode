package max_profit

import (
	"math"
)

func maxProfitSingleProc(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxprofit := 0
	histroyMin := math.MaxInt32
	for _, i := range prices {
		histroyMin = int(math.Min(float64(histroyMin), float64(i)))
		maxprofit = int(math.Max(float64(maxprofit), float64(i-histroyMin)))
	}

	return maxprofit
}

func maxProfitMultiProcGreedy(prices []int) int {
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			maxProfit += diff
		}
	}

	return maxProfit
}

func maxProfitK(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	k = min(k, n/2)
	buy := make([][]int, n)
	sell := make([][]int, n)
	for i := range buy {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
	}
	buy[0][0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[0][i] = math.MinInt64 / 2
		sell[0][i] = math.MinInt64 / 2
	}

	for i := 1; i < n; i++ {
		buy[i][0] = max(buy[i-1][0], sell[i-1][0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[i][j] = max(buy[i-1][j], sell[i-1][j]-prices[i])
			sell[i][j] = max(sell[i-1][j], buy[i-1][j-1]+prices[i])
		}
	}
	return max(sell[n-1]...)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

func maxProfitMultiProcDP(prices []int) int {
	lenOfPrices := len(prices)
	type lenOne [2]int
	dp := make([]lenOne, lenOfPrices)

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < lenOfPrices; i++ {
		dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1]+prices[i])))
		dp[i][1] = int(math.Max(float64(dp[i-1][1]), float64(dp[i-1][0]-prices[i])))
	}

	return dp[lenOfPrices-1][0]
}

func maxProfitTwiceDP(prices []int) int {
	lenOfPrices := len(prices)
	type fiveType [5]int
	dp := make([]fiveType, lenOfPrices)

	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0]

	for i := 1; i < lenOfPrices; i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = int(math.Max(float64(dp[i-1][0]-prices[i]), float64(dp[i-1][1])))
		dp[i][2] = int(math.Max(float64(dp[i-1][1]+prices[i]), float64(dp[i-1][2])))
		dp[i][3] = int(math.Max(float64(dp[i-1][2]-prices[i]), float64(dp[i-1][3])))
		dp[i][4] = int(math.Max(float64(dp[i-1][3]+prices[i]), float64(dp[i-1][4])))
	}

	return int(math.Max(math.Max(float64(dp[lenOfPrices-1][0]), float64(dp[lenOfPrices-1][1])),
		math.Max(float64(dp[lenOfPrices-1][2]), math.Max(float64(dp[lenOfPrices-1][3]), float64(dp[lenOfPrices-1][4])))))
}
