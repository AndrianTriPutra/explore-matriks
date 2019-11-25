package main

import (
	"fmt"
)

var (
	matriks [10][10]int
	i, j, m uint8

	sumL    [20]int
	largest int
)

func main() {
	m = 6

	fmt.Println("matriks 6x6")
	fmt.Println("input element matriks [(-9 ) - (9)]")

	for i = 0; i < m; i++ {
		for j = 0; j < m; j++ {
			fmt.Printf("row %v column %v :", i, j)
			fmt.Scan(&matriks[i][j])
			if matriks[i][j] < -9 || matriks[i][j] > 9 {
				if matriks[i][j] < -9 {
					fmt.Println("error cause value < -9")
				} else if matriks[i][j] > 9 {
					fmt.Println("error cause  value > 9")
				}
				return
			}
		}
	}
	fmt.Println()

	//print matriks
	for i = 0; i < m; i++ {
		for j = 0; j < m; j++ {
			fmt.Printf("%v ", matriks[i][j])
			if m-1 == j {
				fmt.Println()
			}
		}
	}
	fmt.Println()

	sumL[0] = hourglass(0, 3, 1, 1, 0, 3)
	sumL[1] = hourglass(0, 3, 1, 2, 1, 4)
	sumL[2] = hourglass(0, 3, 1, 3, 2, 5)
	sumL[3] = hourglass(0, 3, 1, 4, 3, 6)

	sumL[4] = hourglass(1, 4, 2, 1, 0, 3)
	sumL[5] = hourglass(1, 4, 2, 2, 1, 4)
	sumL[6] = hourglass(1, 4, 2, 3, 2, 5)
	sumL[7] = hourglass(1, 4, 2, 4, 3, 6)

	sumL[8] = hourglass(2, 5, 3, 1, 0, 3)
	sumL[9] = hourglass(2, 5, 3, 2, 1, 4)
	sumL[10] = hourglass(2, 5, 3, 3, 2, 5)
	sumL[11] = hourglass(2, 5, 3, 4, 3, 6)

	sumL[12] = hourglass(3, 6, 4, 1, 0, 3)
	sumL[13] = hourglass(3, 6, 4, 2, 1, 4)
	sumL[14] = hourglass(3, 6, 4, 3, 2, 5)
	sumL[15] = hourglass(3, 6, 4, 4, 3, 6)

	for i = 0; i < 20; i++ {
		if sumL[i] >= largest {
			largest = sumL[i]
		}
	}
	fmt.Printf("largest:%v\n\n", largest)
}

func hourglass(line, a, b, c, d, e uint8) (result int) {
	var sub [5][5]int

	switch line {
	case 0:
		for i = 0; i < a; i++ {
			if i == 1 {
				sub[1][1] = matriks[b][c]
			} else {
				for j = d; j < e; j++ {
					sub[i][j-d] = matriks[i][j]
				}
			}
		}

	case 1:
		for i = 1; i < a; i++ {
			if i == 2 {
				sub[1][1] = matriks[b][c]
			} else {
				for j = d; j < e; j++ {
					sub[i-1][j-d] = matriks[i][j]
				}
			}
		}

	case 2:
		for i = 2; i < a; i++ {
			if i == 3 {
				sub[1][1] = matriks[b][c]
			} else {
				for j = d; j < e; j++ {
					sub[i-2][j-d] = matriks[i][j]
				}
			}
		}

	case 3:
		for i = 3; i < a; i++ {
			if i == 4 {
				sub[1][1] = matriks[b][c]
			} else {
				for j = d; j < e; j++ {
					sub[i-3][j-d] = matriks[i][j]
				}
			}
		}
	}
	/*
		//print sub (uncomment this if you see sub matriks)
		for i = 0; i < 3; i++ {
			for j = 0; j < 3; j++ {
				if i == 1 {
					if j == 0 {
						fmt.Print("  ")
					} else if j == 1 {
						fmt.Printf("%v ", sub[i][j])
					} else {
						fmt.Println()
					}
				} else {
					fmt.Printf("%v ", sub[i][j])
					if j == 2 {
						fmt.Println()
					}
				}
			}
		}*/

	//sum
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if i == 1 && j == 1 {
				result += sub[i][j]
			} else {
				result += sub[i][j]
			}
		}
	}

	return result
}
