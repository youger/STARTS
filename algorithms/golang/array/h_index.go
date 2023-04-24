package array

func hIndex(citations []int) int {
	count := len(citations)
	if count <= 0 {
		return 0
	}
	if count == 1 && citations[0] >= 1 {
		return 1
	}
	quickSort(citations)
	h := 0
	for i := count - 1; i >= 0; i-- {
		if citations[i] > h {
			h++
			break
		}
	}
	return h
}

func swap(a *int, b *int) {
	tmp := a
	a = b
	b = tmp
}
