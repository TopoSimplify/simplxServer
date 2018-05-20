package shapefile

func partRanges(parts []int32, N int32) [][2]int {
	var indices = make([]int, len(parts))
	for i, v := range parts {
		indices[i] = int(v)
	}
	var q = int(N)
	var n = len(indices) - 1
	if indices[n] < q {
		indices = append(indices, q)
	}
	var ranges [][2]int
	for k := 0; k < len(indices)-1; k++ {
		i, j := indices[k], indices[k+1]
		ranges = append(ranges, [2]int{i, j})
	}
	return ranges
}
