package median_finder
// 如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。
//如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

type MedianFinder struct {
	store []int
}

func Constructor() MedianFinder {
	return MedianFinder{
		store: []int{},
	}
}


func (this *MedianFinder) AddNum(num int)  {

}


func (this *MedianFinder) FindMedian() float64 {
	return 0
}

