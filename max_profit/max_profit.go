package max_profit

import "math"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxprofit := 0
	histroyMin := 0
	for _, i := range prices {
		histroyMin = int(math.Min(float64(histroyMin), float64(i)))
		maxprofit = int(math.Max(float64(maxprofit), float64(i - histroyMin)))
	}

	return maxprofit
}
