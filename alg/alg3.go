package alg

// Alg3 :
func Alg3(in [][]string) [][]string {
	result := [][]string{}
	for _, v := range in {
		block := [][]string{}
		for _, i := range v {
			tmpBlock := replica(result, i)
			block = append(block, tmpBlock...)
		}
		result = make([][]string, len(block))
		copy(result, block)
	}
	return result
}

func replica(res [][]string, toAdd string) [][]string {
	if len(res) == 0 {
		return [][]string{
			[]string{toAdd},
		}
	}
	block := make([][]string, len(res))
	copy(block, res)
	for i, v := range block {
		block[i] = append(v, toAdd)
	}
	return block
}
