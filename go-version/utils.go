package main

import (
	"fmt"
	"math"
)

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

func tokenToText(t Token) string {
	return (" " +
		fmt.Sprint(designs[t.north]) + " " +
		fmt.Sprint(designs[t.east]) + " " +
		fmt.Sprint(designs[t.south]) + " " +
		fmt.Sprint(designs[t.west]) + " ")
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

func makeAndCheckBoard(order []int, rotation []int, ch chan Board) {
	board := MakeBoard(order, rotation)

	res := CheckBoard(&board)
	// Create a method to remove redundant elements from orders
	// If for instance the check fails at the first edge,
	//then remove all orders with the same start.
	if res {
		ch <- board
	}
}
