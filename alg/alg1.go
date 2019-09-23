package alg

// Alg1 :
func Alg1(pools [][]string) [][]string {
	f := [][]string{}
	npools := len(pools)
	indices := make([]int, npools)

	result := make([]string, npools)
	for i := range result {
		result[i] = pools[i][0]
	}

	// fmt.Println(result)
	f = deepAppend(f, result)

	for {
		i := npools - 1
		for ; i >= 0; i-- {
			pool := pools[i]
			indices[i]++

			if indices[i] == len(pool) {
				indices[i] = 0
				result[i] = pool[0]
			} else {
				result[i] = pool[indices[i]]
				break
			}

		}

		if i < 0 {
			// fmt.Println(f)
			return f
		}

		// fmt.Println(result)
		f = deepAppend(f, result)
	}
}

func deepAppend(f [][]string, r []string) [][]string {
	cpy := make([]string, len(r))
	copy(cpy, r)
	return append(f, cpy)
}
