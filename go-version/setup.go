package main

import "sort"

/*
nextProduct takes a list of integers, and a length argument, and returns a generative function.
Each time the resulting function is called, it returns the next permutation.
Don't call directly, use the helper function Permutations.
Taken from here:
https://play.golang.org/p/ZfTy6C-lApN
https://www.reddit.com/r/golang/comments/ls3fz0/how_to_generate_all_patterns_for_number/gorvkf1/
*/
func nextProduct(values *[]int, length int) func() *[]int {
	p := make([]int, length)
	x := make([]int, length)
	return func() *[]int {
		p := p[:len(x)]
		for i, xi := range x {
			p[i] = (*values)[xi]
		}
		for i := len(x) - 1; i >= 0; i-- {
			x[i]++
			if x[i] < len(*values) {
				break
			}
			x[i] = 0
			if i <= 0 {
				x = x[0:0]
				break
			}
		}
		return &p
	}
}

/*
RepeatingPermutations creates a list of lists, each list being a permutation.
*/
func RepeatingPermutations(values *[]int, length int) [][]int {
	np := nextProduct(values, length)
	permuts := make([][]int, 0)

	for {
		product := np()

		if len(*product) == 0 {
			break
		}
		c := make([]int, BoardSize)
		copy(c, *product)
		permuts = append(permuts, c)

	}
	return permuts
}

/*
Permutations takes a length argument and outputs a list of all integer Permutations (non-repeating)
*/
func Permutations(input []int) [][]int {
	var helper func([]int, int)
	result := [][]int{}

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

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}

func sortedKeys(m *map[int][]int) *[]int {
	keys := make([]int, len(*m))
	i := 0
	for k := range *m {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	return &keys
}

func sliceToMap(input [][]int) *map[int][]int {
	deleterMap := make(map[int][]int)
	for idx := range input {
		deleterMap[sliceToInt(input[idx])] = input[idx]
	}
	return &deleterMap
}

func chunkSlice(slice []int, chunkSize int) [][]int {
	var chunks [][]int
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func unchunkSlice(chunks [][]int) []int {
	ret := make([]int, 0)

	for _, elements := range chunks {
		ret = append(ret, elements...)
	}

	return ret
}
