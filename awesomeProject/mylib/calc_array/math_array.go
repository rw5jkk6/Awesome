package calc_array

func Avg(a []float64) float64{
	var sum float64 = 0.0 
	for _, v := range a{
		sum += v
	}
	return  sum / float64(len(a))
}