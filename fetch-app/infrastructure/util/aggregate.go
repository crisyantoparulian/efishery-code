package util

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
