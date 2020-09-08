package max_profit

import "math"

func maxProfitSingleProc(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxprofit := 0
	histroyMin := math.MaxInt32
	for _, i := range prices {
		histroyMin = int(math.Min(float64(histroyMin), float64(i)))
		maxprofit = int(math.Max(float64(maxprofit), float64(i - histroyMin)))
	}

	return maxprofit
}

func maxProfitMultiProcGreedy(prices []int) int {

	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i - 1]
		if  diff > 0 {
			maxProfit += diff
		}
	}

	return maxProfit
}