package heap

const (
	MaxHeap = "MAX_HEAP"
	MinHeap = "MIN_HEAP"
)

func BuildHeap(raw []int, heapType string) []int {
	switch heapType {
	case MaxHeap:
		buildMaxHeap(raw)
	case MinHeap:
		buildMinHeap(raw)
	default:
	}
	return raw
}

func buildMaxHeap(raw []int) {
	for i := len(raw) / 2; i >= 0; i-- {
		Heapify(raw, i, len(raw), MaxHeap)
	}
}

func buildMinHeap(raw []int) {
	for i := len(raw) / 2; i >= 0; i-- {
		Heapify(raw, i, len(raw), MinHeap)
	}
}

func Heapify(raw []int, index, heapSize int, heapType string) {
	left := 2*index + 1
	right := 2*index + 2

	switch heapType {
	case MaxHeap:
		larger := index
		if left < heapSize && raw[left] > raw[larger] {
			larger = left
		}
		if right < heapSize && raw[right] > raw[larger] {
			larger = right
		}
		if larger != index {
			raw[larger], raw[index] = raw[index], raw[larger]
			Heapify(raw, larger, heapSize, heapType)
		}

	case MinHeap:
		smaller := index
		if left < heapSize && raw[left] < raw[smaller] {
			smaller = left
		}
		if right < heapSize && raw[right] < raw[smaller] {
			smaller = right
		}
		if smaller != index {
			raw[smaller], raw[index] = raw[index], raw[smaller]
			Heapify(raw, smaller, heapSize, heapType)
		}
	}
}
