package main

import (
	"fmt"
	"math"
)

type magicSquare struct {
	row [][]int32
}

func createMagicSquares() []magicSquare {
	var magicSquares []magicSquare
	magicSquares = append(magicSquares,
		magicSquare{
			[][]int32{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
		},
		magicSquare{
			[][]int32{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
		},
		magicSquare{
			[][]int32{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
		},
		magicSquare{
			[][]int32{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
		},
		magicSquare{
			[][]int32{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
		},
		magicSquare{
			[][]int32{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
		},
		magicSquare{
			[][]int32{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
		},
		magicSquare{
			[][]int32{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
		},
	)

	return magicSquares
}

func formingMagicSquare(s [][]int32) int32 {
	var min float64
	t := magicSquare{
		s,
	}

	for _, ms := range createMagicSquares() {
		var total float64
		for i, row := range ms.row {
			for j, e := range row {
				if e != t.row[i][j] {
					total += math.Abs(float64(e - t.row[i][j]))
				}
			}
		}
		if total == 0 {
			return int32(total)
		}

		if min == 0 || total < min {
			min = total
		}
	}
	return int32(min)
}

func main() {
	minCost := formingMagicSquare([][]int32{{3, 9, 2}, {3, 5, 7}, {8, 1, 6}})
	fmt.Println(minCost)
}
