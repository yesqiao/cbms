package alg

// Alg0 : have bug
func Alg0(in [][]string) [][]string {
	ll := 1
	for _, v := range in {
		ll = ll * len(v)
	}
	res := make([][]string, ll)
	for _, v := range in {
		insert(res, v)
	}
	return res
}

func insert(res [][]string, b []string) {
	ll := len(res)
	lb := len(b)
	for li := 0; li < ll; li++ {
		indexInB := li % lb
		res[li] = append(res[li], b[indexInB])
	}
}
