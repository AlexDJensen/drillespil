package main

import (
	"fmt"
	"math"
	"os"
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

func sliceToInt(s *[]int) int {
	res := 0
	op := 1
	for i := len(*s) - 1; i >= 0; i-- {
		res += (*s)[i] * op
		op *= 10
	}
	return res
}

func sortedKeys(m map[int][]int) []int {
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	return keys
}

func sliceToMap(input *[][]int) map[int][]int {
	deleterMap := make(map[int][]int, 0)
	for idx := range *input {
		deleterMap[sliceToInt(&(*input)[idx])] = (*input)[idx]
	}
	return deleterMap
}

// RemoveRotations converts input slice pointer to a map,
// removes redundant rotations and returns a slice pointer.
func RemoveRotations(input *[][]int) *[][]int {
	// Create a map and populate
	deleterMap := sliceToMap(input)
	//fmt.Println(len(deleterMap), deleterMap)

	//Idea - generate rotations and add to slice, then go to a separate loop to delete them?
	// Create and rotate boards, then delete them

	//for _, val := range deleterMap {
	map_keys := sortedKeys(deleterMap)
	for _, k := range map_keys {
		var deleteValues []int
		//fmt.Println(len(map_keys))

		for i := 1; i < 4; i++ {
			val := (deleterMap)[k]
			deleteVal := RotateBoard(val, i)
			deleteTmp := make([]int, BoardSize)
			copy(deleteTmp, deleteVal)
			deleteValues = append(deleteValues, sliceToInt(&deleteTmp))

		}
		fmt.Println((deleteValues))
		//fmt.Println(deleteValues)
		for _, val := range deleteValues {
			// _, ok := deleterMap[val]
			// fmt.Println(ok)
			delete(deleterMap, val)
			//fmt.Println(deleterMap, len(deleterMap))
		}
		//fmt.Println(deleteValues)
	}
	//fmt.Println("After deleting from the map, down to", len(deleterMap))

	//fmt.Println(deleterMap, len(deleterMap))
	// Finally, return to slice of slices
	map_keys = sortedKeys(deleterMap)
	keys := [][]int{}
	for _, k := range map_keys {
		//for _, value := range deleterMap {
		//keys = append(keys, value)
		keys = append(keys, deleterMap[k])
	}
	//fmt.Println(len(keys))
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
func RotateBoard(board []int, rotations int) []int {

	if rotations == 0 {
		return board
	}

	for i := 0; i < rotations; i++ {
		board = rotateBoard(board)
	}

	return board
}

func rotateBoard(board []int) []int {
	// Adapted from the python solution above
	// https://stackoverflow.com/questions/42519/how-do-you-rotate-a-two-dimensional-array/35438327#35438327

	rotatedBoard := make([]int, 0)

	size := int(math.Floor(math.Sqrt(float64(len(board)))))
	//Rotating layers below
	layerCount := size / 2

	splitBoard := chunkSlice(board, size)

	for i := 0; i < layerCount; i++ {
		first := i
		last := size - first - 1

		for j := first; j < last; j++ {
			offset := j - first

			top := splitBoard[first][j]
			rightSide := splitBoard[j][last]
			bottom := splitBoard[last][last-offset]
			leftSide := splitBoard[last-offset][first]

			splitBoard[first][j] = leftSide
			splitBoard[j][last] = top
			splitBoard[last][last-offset] = rightSide
			splitBoard[last-offset][first] = bottom

		}
	}

	rotatedBoard = unchunkSlice(splitBoard)

	return rotatedBoard
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
const Print = true

func main() {
	// rotationValues := []int{0, 1, 2, 3}
	// rotations := RepeatingPermutations(rotationValues, BoardSize)
	// if Print {
	// 	if len(rotations) < 1000 {
	// 		fmt.Println(rotations, len(rotations))
	// 	} else {
	// 		fmt.Println(len(rotations))
	// 		fmt.Println(rotations[len(rotations)-5:])
	// 	}
	// }
	boardValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//boardValues := []int{1, 2, 3, 4}
	boards := Permutations(boardValues)
	if Print {
		if len(boards) < 1000 {
			fmt.Println(boards, len(boards))
		} else {
			fmt.Println(len(boards))
			fmt.Println(boards[len(boards)-5:])
		}
	}

	// basePath, err := os.Getwd()
	// if err != nil {
	// 	log.Println(err)
	// }
	// //fmt.Println(basePath)
	// err = WriteBoards(path.Join(basePath, "/cmd/drillespil/boards.txt"), &boards)
	// if err != nil {
	// 	log.Println(err)
	// }
	original := sliceToInt(&(boards[0]))
	fmt.Println(original)
	rotated := rotateBoard(boards[0])
	fmt.Println(sliceToInt(&rotated))
	rotated1 := rotateBoard(rotated)
	fmt.Println(sliceToInt(&rotated1))
	rotated2 := rotateBoard(rotated1)
	fmt.Println(sliceToInt(&rotated2))

	//fmt.Println(original, sliceToInt(&rotated), sliceToInt(&rotated1), sliceToInt(&rotated2))
	fmt.Println("All done now")
	// boards = *RemoveRotations(&boards)
	// if Print {
	// 	if len(boards) < 1000 {
	// 		fmt.Println(boards, len(boards))
	// 	} else {
	// 		fmt.Println(len(boards))
	// 		fmt.Println((boards)[len(boards)-5:])
	// 	}
	// }

	// err = WriteBoards(path.Join(basePath, "/cmd/drillespil/clean_boards.txt"), &boards)
	// if err != nil {
	// 	log.Println(err)
	// }
}
