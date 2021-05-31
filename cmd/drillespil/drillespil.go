package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path"
	"sort"
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
func RepeatingPermutations(values *[]int, length int) *[][]int {
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
	return &permuts
}

/*
Permutations takes a length argument and outputs a list of all integer Permutations (non-repeating)
*/
func Permutations(input []int) *[][]int {
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
	return &result
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

func sliceToInt(s *[]int) int {
	res := 0
	op := 1
	for i := len(*s) - 1; i >= 0; i-- {
		res += (*s)[i] * op
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

func sliceToMap(input *[][]int) *map[int][]int {
	deleterMap := make(map[int][]int, 0)
	for idx := range *input {
		deleterMap[sliceToInt(&(*input)[idx])] = (*input)[idx]
	}
	return &deleterMap
}

func chunkSlice(slice *[]int, chunkSize int) *[][]int {
	var chunks [][]int
	for i := 0; i < len(*slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(*slice) {
			end = len(*slice)
		}

		chunks = append(chunks, (*slice)[i:end])
	}

	return &chunks
}

func unchunkSlice(chunks *[][]int) *[]int {
	ret := make([]int, 0)

	for _, elements := range *chunks {
		for _, v := range elements {
			ret = append(ret, v)
		}
	}

	return &ret
}

func rotateBoard(board []int) []int {
	// Adapted from the python solution above
	// https://stackoverflow.com/questions/42519/how-do-you-rotate-a-two-dimensional-array/35438327#35438327

	rotatedBoard := make([]int, 0)
	// fmt.Printf("rotateBoard input: %v \n", &board)

	size := int(math.Floor(math.Sqrt(float64(len(board)))))

	splitBoard := chunkSlice(&board, size)
	// fmt.Printf("splitBoard pre rotate: %v \n", splitBoard)
	//fmt.Println(splitBoard)

	// Reverse the matrix
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		(*splitBoard)[i], (*splitBoard)[j] = (*splitBoard)[j], (*splitBoard)[i]
	}

	// fmt.Printf("after reverse %v \n", splitBoard)

	// Transpose it, i.e. turn columns to rows
	for i := 0; i < size; i++ {
		for j := 0; j < i; j++ {
			(*splitBoard)[i][j], (*splitBoard)[j][i] = (*splitBoard)[j][i], (*splitBoard)[i][j]
		}
	}

	//fmt.Println(splitBoard)
	// fmt.Printf("splitBoard after rotate: %v \n", splitBoard)
	rotatedBoard = *unchunkSlice(splitBoard)
	// fmt.Printf("Output of rotateBoard: %v \n", rotatedBoard)
	//fmt.Println(rotatedBoard)

	return rotatedBoard
}

// RotateBoard is a utility function - it iteratively calls rotateBoard to get correct rotation
func RotateBoard(board *[]int, rotations int) *[]int {
	newBoard := make([]int, len(*board))

	if rotations == 0 {
		return board
	}

	derefBoard := *board

	for i := 0; i < rotations; i++ {
		derefBoard = rotateBoard(derefBoard)
	}

	copy(newBoard, derefBoard)
	return &newBoard
}

// RemoveRotations converts input slice pointer to a map,
// removes redundant rotations and returns a slice pointer.
func RemoveRotations(input *[][]int) *[][]int {
	// Create a map and populate
	deleterMap := sliceToMap(input)

	// Create rotations and add to a slice
	map_keys := sortedKeys(deleterMap)
	for _, k := range *map_keys {
		var deleteValues []int
		//fmt.Println((*deleterMap)[k])
		val := (*deleterMap)[k]
		fmt.Printf("Starting loop with value %v \n", val)
		for i := 1; i < 4; i++ {
			fmt.Println(i)
			fmt.Println(val)
			deleteVal := RotateBoard(&val, i)

			// fmt.Println(deleteVal)
			deleteTmp := make([]int, BoardSize)
			copy(deleteTmp, *deleteVal)
			deleteValues = append(deleteValues, sliceToInt(&deleteTmp))

		}

		fmt.Println(deleteValues)

		// Remove the rotations
		for _, val := range deleteValues {
			delete(*deleterMap, val)
		}
	}
	// Finally, return to slice of slices
	map_keys = sortedKeys(deleterMap)
	keys := [][]int{}
	for _, k := range *map_keys {
		keys = append(keys, (*deleterMap)[k])
	}
	return &keys
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

func PrintList(obj *[][]int) {
	if Print {
		if len(*obj) < 1000 {
			fmt.Println(obj, len(*obj))
		} else {
			fmt.Println(len(*obj))
			fmt.Println((*obj)[len(*obj)-5:])
		}
	}
}

//Set to 3 for now during testing, should be 9
const BoardSize = 9
const Print = false

func main() {
	rotationValues := []int{0, 1, 2, 3}
	rotations := RepeatingPermutations(&rotationValues, BoardSize)
	PrintList(rotations)
	boardValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//boardValues := []int{1, 2, 3, 4}
	boards := Permutations(boardValues)
	PrintList(boards)

	basePath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(len(*boards))
	//fmt.Println(basePath)
	err = WriteBoards(path.Join(basePath, "/cmd/drillespil/boards.txt"), boards)
	if err != nil {
		log.Println(err)
	}

	// raw := (*boards)[0]
	// raw_int := sliceToInt(&raw)
	// boards = RemoveRotations(boards)
	// fmt.Println("Raw was:", raw_int)
	// fmt.Println("How rotating stuff")
	// rot0 := RotateBoard(&raw, 0)

	// rot1 := RotateBoard(&raw, 1)

	// rot2 := RotateBoard(&raw, 2)

	// rot3 := RotateBoard(&raw, 3)

	// fmt.Println("Now print stuff")
	// fmt.Println("rot0", sliceToInt(rot0))
	// fmt.Println("rot1", sliceToInt(rot1))
	// fmt.Println("rot2", sliceToInt(rot2))
	// fmt.Println("rot3", sliceToInt(rot3))
	// fmt.Println(*rot0)
	// fmt.Println(*rot1)
	// fmt.Println(*rot2)
	// fmt.Println(*rot3)
	// PrintList(boards)
	// fmt.Println(len(*boards))

	boards = RemoveRotations(boards)
	err = WriteBoards(path.Join(basePath, "/cmd/drillespil/clean_boards.txt"), boards)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("All done now")
}
