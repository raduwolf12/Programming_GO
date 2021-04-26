package main

import "fmt"

func verifyConfig(table [9][9]int) bool {
	var a [3][9]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			x := table[i][j] - 1
			t := 1 << x
			if a[0][i]&t != 0 {
				return false
			}
			a[0][i] |= t
			if a[1][j]&t != 0 {
				return false
			}
			a[1][j] |= t

			k := i/3*3 + j/3
			if a[2][k]&t != 0 {
				return false
			}
			a[2][k] |= t

		}
	}
	return true
}

func main() {
	table := [9][9]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{9, 1, 2, 3, 4, 5, 6, 7, 8},
		{3, 4, 5, 6, 7, 8, 9, 1, 2},
		{6, 7, 8, 9, 1, 2, 3, 4, 5},
		{8, 9, 1, 2, 3, 4, 5, 6, 7},
		{2, 3, 4, 5, 6, 7, 8, 9, 1},
		{5, 6, 7, 8, 9, 1, 2, 3, 4}}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(table[i][j], " ")
		}
		fmt.Println()
	}

	if verifyConfig(table) == true {
		fmt.Print("Configuratie buna!")
	} else {
		fmt.Print("Configuratie gresita!")
	}

}
