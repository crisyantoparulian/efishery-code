package util

import "sort"

func FindMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func CalcMedian(n []int) float64 {
	sort.Ints(n)
	mNumber := len(n) / 2

	if IsOdd(len(n)) {
		return float64(n[mNumber])
	}

	return float64(n[mNumber-1]+n[mNumber]) / 2.0
}

func IsOdd(n int) bool {
	if n%2 == 0 {
		return false
	}

	return true
}

func CalcAvg(n []int) float64 {

	var t int
	for _, v := range n {
		t += v
	}
	return float64(t) / float64(len(n))
}

func CalcMinMax(n []int) (int, int) {
	var (
		min int
		max int
	)
	sort.Ints(n)

	lenWeeks := len(n)
	if lenWeeks > 0 {
		min = n[0]
		max = n[lenWeeks-1]
	}

	return min, max

}

func CalcMinMaxMedAvg(n []int) (int, int, float64, float64) {
	var (
		min int
		max int
		med float64
		avg float64
	)
	min, max = CalcMinMax(n)

	med = CalcMedian(n)
	avg = CalcAvg(n)

	return min, max, med, avg
}
