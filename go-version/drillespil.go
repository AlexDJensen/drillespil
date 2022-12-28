package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const BoardSize = 9

//var known_good_board = []int{1, 4, 9, 6, 2, 5, 7, 8, 3}
//var known_good_rotat = []int{3, 0, 1, 1, 2, 3, 1, 3, 3}

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
		r.north = t.west
		r.east = t.north
		r.south = t.east
		r.west = t.south
		err = nil

	case 2:
		r.north = t.south
		r.east = t.west
		r.south = t.north
		r.west = t.east
		err = nil
	case 3:
		r.north = t.east
		r.east = t.south
		r.south = t.west
		r.west = t.north
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
