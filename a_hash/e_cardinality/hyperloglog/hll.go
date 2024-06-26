package main

import (
	"fmt"
	"hash/fnv"
	"math"
)

//NOTE: This code was generated by GPT.

const numRegisters = 16 // Number of registers (adjust for precision)

type HyperLogLog struct {
	buckets []int32
}

func NewHyperLogLog() *HyperLogLog {
	return &HyperLogLog{
		buckets: make([]int32, numRegisters),
	}
}

func (h *HyperLogLog) Add(input string) {
	hashed := hashString(input)
	index := hashed % uint32(numRegisters) // Determine register index

	leadingZeros := countLeadingZeros(hashed)
	if leadingZeros > h.buckets[index] {
		h.buckets[index] = leadingZeros
	}
}

func (h *HyperLogLog) Estimate() int32 {
	// Link:https://github.com/Gookuruto/HyperLogLog_go/blob/7b6020b5979e23fd23bafa5fa1d7b5e3c3270906/hll.go#L64
	alpha := 0.7213 / (1.0 + 1.079/float64(numRegisters)) // Constant for bias correction

	harmonicMean := 0.0
	for _, r := range h.buckets {
		harmonicMean += math.Pow(2.0, float64(-r))
	}
	harmonicMean = 1.0 / harmonicMean

	estimate := alpha * math.Pow(float64(len(h.buckets)), 2) * harmonicMean

	//TODO: skipping some extra correction steps
	return int32(estimate)
}

func countLeadingZeros(hash uint32) int32 {
	if hash == 0 { // Handle the case of hash being 0
		return 32
	}

	count := int32(0)
	for hash > 0 {
		hash = hash / 2
		count++
	}
	return 32 - count
}

func hashString(input string) uint32 {
	hasher := fnv.New32a()
	_, _ = hasher.Write([]byte(input))
	return hasher.Sum32()
}

func main() {
	hll := NewHyperLogLog()
	hll.Add("hello")
	hll.Add("world")
	hll.Add("hello") // Duplicate entry
	fmt.Println("Estimated unique count:", hll.Estimate())
}
