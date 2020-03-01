package statistic

func Average(xs []float64) (float64, float64) {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs)), total
}

func Sum(s []float64) float64 {
	var sumtotal float64
	for _, x := range s {
		sumtotal += x
	}
	return sumtotal
}
