package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		matriks           [100][100]int64
		i, j, m           int64
		sum1, sum2, diff1 int64
		diff2             float64
		operator          string
	)

	fmt.Print("input matriks NxN :")
	fmt.Scanln(&m)

	fmt.Println("input element matriks [(-100 ) - (100)]")
	for i = 0; i < m; i++ {
		for j = 0; j < m; j++ {
			fmt.Printf("line %v column %v :", i, j)
			fmt.Scan(&matriks[i][j])
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
	fmt.Print("\ninput operator (+) or (-) or (*) or (/) :")
	fmt.Scanln(&operator)
	if operator == "+" {
		// plus left to right
		for i = 0; i < m; i++ {
			j = i
			sum1 += matriks[i][j]
		}
		fmt.Printf("\nsum diagonal left to right =%v\n", sum1)

		//plus right to left
		j = m - 1
		for i = 0; i < m; i++ {
			sum2 += matriks[i][j]
			j--
		}
		fmt.Printf("sum diagonal right to left=%v\n", sum2)
	} else if operator == "-" {
		// min left to right
		for i = 0; i < m; i++ {
			j = i
			sum1 -= matriks[i][j]
		}
		fmt.Printf("\nsum diagonal left to right =%v\n", sum1)

		//min right to left
		j = m - 1
		for i = 0; i < m; i++ {
			sum2 -= matriks[i][j]
			j--
		}
		fmt.Printf("sum diagonal right to left=%v\n", sum2)
	} else if operator == "*" {
		// * left to right
		for i = 0; i < m; i++ {
			j = i
			if i == 0 {
				sum1 = matriks[i][j]
			} else {
				sum1 *= matriks[i][j]
			}
		}
		fmt.Printf("\nsum diagonal left to right =%v\n", sum1)

		//* right to left
		j = m - 1
		for i = 0; i < m; i++ {
			if i == 0 {
				sum2 = matriks[i][j]
			} else {
				sum2 *= matriks[i][j]
			}
			j--
		}
		fmt.Printf("sum diagonal right to left=%v\n", sum2)
	} else if operator == "/" {
		// / left to right
		for i = 0; i < m; i++ {
			j = i
			if i == 0 {
				sum1 = matriks[i][j]
			} else {
				sum1 /= matriks[i][j]
			}
		}
		fmt.Printf("\nsum diagonal left to right =%v\n", sum1)

		// / right to left
		j = m - 1
		for i = 0; i < m; i++ {
			if i == 0 {
				sum2 = matriks[i][j]
			} else {
				sum2 /= matriks[i][j]
			}
			j--
		}
		fmt.Printf("sum diagonal right to left=%v\n", sum2)
	}

	//different
	diff1 = sum1 - sum2
	diff2 = float64(diff1)
	diff2 = math.Abs(diff2)
	fmt.Printf("\ndifferent=%v\n", diff2)
}
