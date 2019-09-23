package alg

// Alg2 :Recursion solution
func Alg2(in [][]string) [][]string {
	ll := 1
	for _, v := range in {
		ll = ll * len(v)
	}
	res := make([][]string, ll)
	cur(res, in, 0, ll)
	return res
}

func cur(res [][]string, in [][]string, l int, r int) {
	if len(in) == l {
		return
	}
	// repeat := r
	repeat := r / len(in[l])
	currrepeat := repeat
	id := 0
	for i := 0; i < len(res); i++ {
		if currrepeat == 0 {
			currrepeat = repeat
			id++
		}
		res[i] = append(res[i], in[l][id%len(in[l])])
		currrepeat--
	}
	cur(res, in, l+1, repeat)
}
