package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path"
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

// EqualSlices takes two int slices and compares them for equality
func EqualSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// RemoveRotations converts input slice pointer to a map,
// removes redundant rotations and returns a slice pointer.
func RemoveRotations(input *[][]int) *[][]int {
	// Create a map and populate
	deleterMap := make(map[[BoardSize]int]bool)
	var tmp [BoardSize]int
	for idx := range *input {
		copy(tmp[:], (*input)[idx])
		deleterMap[tmp] = true
	}

	// Create and rotate boards, then delete them
	board := make([]int, 0)
	var deleteValue [BoardSize]int
	for key := range deleterMap {
		//fmt.Println(key)
		board = key[:]
		for i := 1; i < 4; i++ {
			deleteTmp := RotateBoard(&board, i)
			// fmt.Println(deleterMap)
			copy(deleteValue[:], *deleteTmp)
			delete(deleterMap, deleteValue)
			// fmt.Println("Deleting", deleteTmp)
			// fmt.Println(deleterMap)
			// fmt.Println()

		}

	}
	fmt.Println(len(deleterMap))

	//fmt.Println(deleterMap, len(deleterMap))
	// Finally, return to slice of slices
	keys := [][]int{}
	for key := range deleterMap {
		keys = append(keys, key[:])
	}

	return &keys
}

// Rethink:
// for any given board, generate the tree rotations.
// Then, go from either beginning (special case), or from last known position.
// Remove the three offending rotations (or reach the end).
// If you reach the end, you need to ignore what you get back and start from next iteration.
// How about a map?
// Trying a map

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
		for _, v := range elements {
			ret = append(ret, v)
		}
	}

	return ret
}

// RotateBoard is a utility function - it iteratively calls rotateBoard to get correct rotation
func RotateBoard(board *[]int, rotations int) *[]int {

	if rotations == 0 {
		return board
	}

	for i := 0; i < rotations; i++ {
		board = rotateBoard(board)
	}

	return board
}

func rotateBoard(board *[]int) *[]int {
	// Adapted from the python solution above
	// https://stackoverflow.com/questions/42519/how-do-you-rotate-a-two-dimensional-array/35438327#35438327

	rotated_board := make([]int, 0)

	size := int(math.Floor(math.Sqrt(float64(len(*board)))))
	//Rotating layers below
	layer_count := size / 2

	split_board := chunkSlice(*board, size)

	for i := 0; i < layer_count; i++ {
		first := i
		last := size - first - 1

		for j := first; j < last; j++ {
			offset := j - first

			top := split_board[first][j]
			right_side := split_board[j][last]
			bottom := split_board[last][last-offset]
			left_side := split_board[last-offset][first]

			split_board[first][j] = left_side
			split_board[j][last] = top
			split_board[last][last-offset] = right_side
			split_board[last-offset][first] = bottom

		}
	}

	rotated_board = unchunkSlice(split_board)

	return &rotated_board
}

func WriteBoards(filePath string, values *[][]int) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, value := range *values {
		fmt.Fprintln(f, value)
	}

	return nil
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
	//board_values := []int{1, 2, 3, 4}
	boards := Permutations(board_values)
	if len(boards) < 1000 {
		fmt.Println(boards, len(boards))
	} else {
		fmt.Println(len(boards))
		fmt.Println(boards[len(boards)-5:])
	}

	base_path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(base_path)
	err = WriteBoards(path.Join(base_path, "/cmd/drillespil/boards.txt"), &boards)
	if err != nil {
		fmt.Println(err)
	}

	boards = *RemoveRotations(&boards)
	// clean_boards := &boards
	// removeRotation(&boards, &[]int{2, 1, 3, 4, 5, 6, 7, 8, 9})
	if len(boards) < 1000 {
		fmt.Println(boards, len(boards))
	} else {
		fmt.Println(len(boards))
		fmt.Println((boards)[len(boards)-5:])
	}

	WriteBoards(path.Join(base_path, "/cmd/drillespil/clean_boards.txt"), &boards)

}
