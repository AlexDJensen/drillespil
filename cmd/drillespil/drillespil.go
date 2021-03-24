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
Permutations creates a list of lists, each list being a permutation.
*/
func Permutations(values []int) (permuts [][]int) {
	np := nextProduct(values, BoardSize)
	permuts = make([][]int, 0)

	for {
		product := np()
		fmt.Println(product)
		if len(product) == 0 {
			break
		}
		permuts = append(permuts, product)

	}
	return permuts
}

//Set to 3 for now during testing, should be 9
const BoardSize = 3

func main() {
	val := []int{0, 1}
	perms := Permutations(val)
	fmt.Println(perms, len(perms))
}
