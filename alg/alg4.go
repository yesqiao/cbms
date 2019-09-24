package alg

// Alg4 : same to alg3, but save some memories, more slower
func Alg4(in [][]string) [][]string {
	result := [][]string{}
	for _, v := range in {
		lenBlock := len(result)
		for idx, i := range v {
			result = replicaS(result, lenBlock, idx, i)
		}
	}
	return result
}

func replicaS(res [][]string, lenBlock int, idx int, toAdd string) [][]string {
	lenRes := len(res)
	if lenRes == 0 {
		res = [][]string{
			[]string{toAdd},
		}
		return res
	}
	if lenBlock == 0 {
		res = append(res, []string{toAdd})
		return res
	}
	if idx == 0 {
		for i, v := range res {
			res[i] = append(v, toAdd)
		}
	}
	if idx != 0 {
		block := deepCopy(res, lenBlock)
		lenLine := len(block[0])
		for i := range block {
			block[i][lenLine-1] = toAdd
		}
		res = append(res, block...)
	}
	return res
}

func deepCopy(s [][]string, length int) [][]string {
	if length == 0 {
		return [][]string{}
	}
	block := make([][]string, length)
	for i := 0; i < length; i++ {
		vv := make([]string, len(s[i]))
		copy(vv, s[i])
		block[i] = vv
	}
	return block
}
