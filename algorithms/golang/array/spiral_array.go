package array

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := 0
	if m > 0 {
		n = len(matrix[0])
	}
	count := m * n
	//directions := []map[int][int]{[]int{1,0}, []int{0,-1}, []int{-1,0}, []int{0, 1}}
	curDirection := 0
	rowMin := 0
	rowMax := m - 1
	coluwnMin := 0
	coluwnMax := n - 1
	res := make([]int, 0)
	row := 0
	coluwn := 0
	for i := 0; i < count; i++ {
		num := matrix[row][coluwn]
		res = append(res, num)
		curDirection = curDirection % 4
		if curDirection == 0 { // right
			if coluwn < coluwnMax {
				coluwn++
			} else if coluwn == coluwnMax {
				rowMin++
				row++
				curDirection++
			}
		} else if curDirection == 1 { // down
			if row < rowMax {
				row++
			} else if row == rowMax {
				coluwnMax--
				coluwn--
				curDirection++
			}
		} else if curDirection == 2 { // left
			if coluwn > coluwnMin {
				coluwn--
			} else if coluwn == coluwnMin {
				rowMax--
				row--
				curDirection++
			}
		} else if curDirection == 3 { //up
			if row > rowMin {
				row--
			} else if row == rowMin {
				coluwnMin++
				coluwn++
				curDirection++
			}
		}
	}
	return res
}
