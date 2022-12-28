package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"
)

const BoardSize = 9

//var known_good_board = []int{1, 4, 9, 6, 2, 5, 7, 8, 3}
//var known_good_rotat = []int{3, 0, 1, 1, 2, 3, 1, 3, 3}

/*
Token is a single token and the four colour-half pairs it contains.
*/
type Token struct {
	north uint8
	east  uint8
	south uint8
	west  uint8
}

var designs = map[uint8]string{
	1: "blue_bottom",
	2: "pink_bottom",
	3: "orange_boots",
	4: "green_boots",
	5: "yellow_top",
	6: "white_top",
	7: "pink_top",
	8: "blue_top",
}

var tokens = map[int]Token{
	//  It goes
	//  n, e, s, w
	1: {5, 7, 3, 1},
	2: {5, 3, 2, 8},
	3: {7, 2, 3, 6},
	4: {3, 8, 7, 4},
	5: {7, 3, 1, 5},
	6: {4, 2, 6, 8},
	7: {4, 1, 7, 6},
	8: {6, 8, 2, 4},
	9: {1, 6, 7, 4},
}

/*
Board is the controlling structure in this game - it places tokens and compares edges
*/
type Board struct {
	p1          Token
	p2          Token
	p3          Token
	p4          Token
	p5          Token
	p6          Token
	p7          Token
	p8          Token
	p9          Token
	og_order    []int
	og_rotation []int
}

type edges struct {
	p1e uint8
	p1s uint8
	p2e uint8
	p2s uint8
	p2w uint8
	p3s uint8
	p3w uint8
	p4n uint8
	p4e uint8
	p4s uint8
	p5n uint8
	p5e uint8
	p5s uint8
	p5w uint8
	p6n uint8
	p6s uint8
	p6w uint8
	p7n uint8
	p7e uint8
	p8n uint8
	p8e uint8
	p8w uint8
	p9n uint8
	p9w uint8
}

/*
edgeFinder takes in a board, and returns a structure of the edges that exist on a given board.
The board is expected to consist of rotated tokens
*/
func edgeFinder(b *Board) (e *edges) {
	e = &edges{}
	e.p1e = b.p1.east
	e.p1s = b.p1.south
	e.p2e = b.p2.east
	e.p2s = b.p2.south
	e.p2w = b.p2.west
	e.p3s = b.p3.south
	e.p3w = b.p3.west
	e.p4n = b.p4.north
	e.p4e = b.p4.east
	e.p4s = b.p4.south
	e.p5n = b.p5.north
	e.p5e = b.p5.east
	e.p5s = b.p5.south
	e.p5w = b.p5.west
	e.p6n = b.p6.north
	e.p6s = b.p6.south
	e.p6w = b.p6.west
	e.p7n = b.p7.north
	e.p7e = b.p7.east
	e.p8n = b.p8.north
	e.p8e = b.p8.east
	e.p8w = b.p8.west
	e.p9n = b.p9.north
	e.p9w = b.p9.west
	return e
}

type edgePair struct {
	e1 uint8
	e2 uint8
}

func edgePairs(e *edges) []edgePair {
	p1 := edgePair{e.p1e, e.p2w}
	p2 := edgePair{e.p1s, e.p4n}
	p3 := edgePair{e.p2e, e.p3w}
	p4 := edgePair{e.p2s, e.p5n}
	p5 := edgePair{e.p3s, e.p6n}
	p6 := edgePair{e.p4e, e.p5w}
	p7 := edgePair{e.p4s, e.p7n}
	p8 := edgePair{e.p6w, e.p5e}
	p9 := edgePair{e.p6s, e.p9n}
	p10 := edgePair{e.p7e, e.p8w}
	p11 := edgePair{e.p8n, e.p5s}
	p12 := edgePair{e.p8e, e.p9w}
	return []edgePair{p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12}
}

// Match checks if edgepair is valid, returns bool
func Match(e *edgePair) bool {
	return e.e1+e.e2 == 9
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

func rotateBoard(board []int) []int {
	// Adapted from the python solution above
	// https://stackoverflow.com/questions/42519/how-do-you-rotate-a-two-dimensional-array/35438327#35438327

	size := int(math.Floor(math.Sqrt(float64(len(board)))))

	splitBoard := chunkSlice(board, size)
	// fmt.Printf("splitBoard pre rotate: %v \n", splitBoard)
	//fmt.Println(splitBoard)

	// Transpose it, i.e. turn columns to rows
	for i := 0; i < size; i++ {
		for j := 0; j < i; j++ {
			splitBoard[i][j], splitBoard[j][i] = splitBoard[j][i], splitBoard[i][j]
		}
	}

	// Reverse the matrix
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		splitBoard[i], splitBoard[j] = splitBoard[j], splitBoard[i]
	}

	return unchunkSlice(splitBoard)

}

// RotateBoard is a utility function - it iteratively calls rotateBoard to get correct rotation
func RotateBoard(board []int, rotations int) []int {
	newBoard := make([]int, len(board))

	if rotations == 0 {
		return board
	}

	derefBoard := board

	for i := 0; i < rotations; i++ {
		derefBoard = rotateBoard(derefBoard)
	}

	copy(newBoard, derefBoard)
	return newBoard
}

// RemoveRotations converts input slice pointer to a map,
// removes redundant rotations and returns a slice pointer.
func RemoveRotations(input [][]int) [][]int {
	// Create a map and populate
	deleterMap := sliceToMap(input)

	// Create rotations and add to a slice
	map_keys := sortedKeys(deleterMap)
	for _, k := range *map_keys {
		var deleteValues []int

		val := (*deleterMap)[k]
		for i := 1; i < 4; i++ {
			deleteVal := RotateBoard(val, i)

			deleteTmp := make([]int, BoardSize)
			copy(deleteTmp, deleteVal)
			deleteValues = append(deleteValues, sliceToInt(deleteTmp))

		}

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
	return keys
}

/*
Rotate takes a token and an optional rotation,
and returns the corresponding colour-half pair in the board-edges for that rotation.
Rotation is number of clockwise turns off 0 (meaning 0 is token north = board north)
*/
func Rotate(t Token, rotation int) (r Token, err error) {
	switch rotation {
	case 0:
		r = t
		err = nil
	case 1:
		r.north = t.east
		r.east = t.south
		r.south = t.west
		r.west = t.north
		err = nil
	case 2:
		r.north = t.south
		r.east = t.west
		r.south = t.north
		r.west = t.east
		err = nil
	case 3:
		r.north = t.west
		r.east = t.north
		r.south = t.east
		r.west = t.south
		err = nil
	default:
		newRotation := rotation % 4
		r, err = Rotate(t, newRotation)
		fmt.Println(err)
		errorMessage := ("Rotation argument is too high (" +
			strconv.Itoa(rotation) +
			"), running with " +
			strconv.Itoa(newRotation) +
			" instead")
		err = errors.New(errorMessage)
	}

	return r, err
}

func MakeBoard(order []int, rotation []int) Board {
	hold := make([]Token, 0)
	for idx, val := range order {
		t := tokens[val]
		t, _ = Rotate(t, rotation[idx])
		hold = append(hold, t)
	}

	return Board{hold[0], hold[1], hold[2], hold[3], hold[4], hold[5], hold[6], hold[7], hold[8], order, rotation}
}

func MakeBoards(orders [][]int, rotations [][]int) []Board {
	boards := make([]Board, len(rotations)-1*len(orders)-1)

	for i := 0; i < len(orders); i++ {
		for j := 0; j < len(rotations); j++ {
			board := MakeBoard(orders[i], rotations[j])
			boards = append(boards, board)
		}
	}
	return boards
}

func MakeBoardsLimited(orders [][]int, rotations [][]int, n int) []Board {
	boards := make([]Board, 0)
	x := 0
	for i := 0; i < len(orders) && x < n; i++ {
		for j := 0; j < len(rotations) && x < n; j++ {
			x++
			board := MakeBoard(orders[i], rotations[j])
			boards = append(boards, board)
		}
	}
	return boards
}

func CheckBoard(board *Board) bool {
	edges := edgeFinder(board)
	pairs := edgePairs(edges)
	for _, pair := range pairs {
		valid := Match(&pair)
		if !valid {
			return false
		}
	}
	return true
}

func printBoard(b Board) {
	for _, keyword := range []string{"Order", "Rotation"} {
		fmt.Println(keyword)
		switch keyword {
		case "Order":
			rows := chunkSlice(b.og_order, 3)
			for _, row := range rows {
				fmt.Printf("%v\n", row)
			}
		case "Rotation":
			rows := chunkSlice(b.og_rotation, 3)
			for _, row := range rows {
				fmt.Printf("%v\n", row)
			}
		}

	}
}

func makeAndCheckBoard(order []int, rotation []int, ch chan Board) {
	board := MakeBoard(order, rotation)
	//fmt.Println(board)
	res := CheckBoard(&board)
	// Create a method to remove redundant elements from orders
	// If for instance the check fails at the first edge,
	//then remove all orders with the same start.
	if res {
		ch <- board
	}
}

func CreateAndCheckBoards(orders [][]int, rotations [][]int) {
	valids := make(chan Board, 100)
	start_time := time.Now()
	for i := 0; i < len(orders); i++ {
		if i%1000 == 0 {
			fmt.Printf("Currently at order iteration %d, after %v time \n", i, time.Since(start_time))
		}
		for j := 0; j < len(rotations); j++ {
			order := orders[i]
			rotation := rotations[j]
			go makeAndCheckBoard(order, rotation, valids)
		}

	}
	fmt.Printf("Length of valids is %v\n", len(valids))
	time.Sleep(2 * time.Second)
	close(valids)
	for elem := range valids {
		printBoard(elem)
	}
	//return valids
}

// Uanset rotation, så kan brik x og brik y aldrig matche.
// Derfor fjern fra boards alle entries, hvor de to er i positioner imod hinandend
//

func main() {
	rotationValues := []int{0, 1, 2, 3}
	rotations := RepeatingPermutations(&rotationValues, BoardSize)
	boardValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// rotations[0] = known_good_rotat
	// rotations = rotations[0:1]
	boards := Permutations(boardValues)

	boards = RemoveRotations(boards)
	//boards[0] = known_good_board
	//boards = boards[0:1]
	fmt.Println("Prep work done")
	fmt.Printf("Size of boards: %d\n", len(boards))
	fmt.Printf("Size of rotations: %d\n", len(rotations))
	fmt.Printf("Total objects to process: %d\n", len(rotations)*len(boards))

	CreateAndCheckBoards(boards, rotations)
	//fmt.Printf("Valid boards are: %v\n", valids)
	// time.Sleep(time.Second * 5)

}
