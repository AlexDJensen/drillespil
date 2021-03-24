package main

import (
	"fmt"
)

var colours = map[int]string{
	1: "Blue-Top",
	2: "Blue-Bottom",
}

/*
nextProduct takes a list of integers, and a length argument, and returns a generative function.
Each time the resulting function is called, it returns the next permutation.
Don't call directly, use the helper function Permutations.
Taken from here:
https://play.golang.org/p/ZfTy6C-lApN
https://www.reddit.com/r/golang/comments/ls3fz0/how_to_generate_all_patterns_for_number/gorvkf1/
*/
func nextProduct(values []int, length int) func() []int {
	p := make([]int, length)
	x := make([]int, length)
	return func() []int {
		p := p[:len(x)]
		for i, xi := range x {
			p[i] = values[xi]
		}
		for i := len(x) - 1; i >= 0; i-- {
			x[i]++
			if x[i] < len(values) {
				break
			}
			x[i] = 0
			if i <= 0 {
				x = x[0:0]
				break
			}
		}
		return p
	}
}

/*
RepeatingPermutations creates a list of lists, each list being a permutation.
*/
func RepeatingPermutations(values []int, length int) (permuts [][]int) {
	np := nextProduct(values, length)
	permuts = make([][]int, 0)

	for {
		product := np()

		if len(product) == 0 {
			break
		}
		c := make([]int, BoardSize)
		copy(c, product)
		permuts = append(permuts, c)

	}
	return permuts
}

/*
Permutations takes a length argument and outputs a list of all integer Permutations (non-repeating)
*/
func Permutations(input []int) (result [][]int) {
	var helper func([]int, int)
	result = [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			result = append(result, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(input, len(input))
	return result
}

//Set to 3 for now during testing, should be 9
const BoardSize = 9

func main() {
	rotation_values := []int{0, 1, 2, 3}
	rotations := RepeatingPermutations(rotation_values, BoardSize)
	if len(rotations) < 1000 {
		fmt.Println(rotations, len(rotations))
	} else {
		fmt.Println(len(rotations))
		fmt.Println(rotations[len(rotations)-5:])
	}
	board_values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	boards := Permutations(board_values)
	if len(boards) < 1000 {
		fmt.Println(boards, len(boards))
	} else {
		fmt.Println(len(boards))
		fmt.Println(boards[len(boards)-5:])
	}

}
