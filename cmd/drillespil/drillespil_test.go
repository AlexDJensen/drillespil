package main

import (
	"testing"
)

var length = 9
var boardValues []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var rotationValues []int = []int{0, 1, 2, 3}
var layouts [][]int = Permutations(boardValues)
var board_layouts [][]int = RemoveRotations((layouts))
var rotations [][]int = RepeatingPermutations(&rotationValues, length)
var testBoards []Board = MakeBoardsLimited(board_layouts, rotations, 100000)

var known_good_board = []int{1, 4, 9, 6, 2, 5, 7, 8, 3}
var known_good_rotat = []int{3, 0, 1, 1, 2, 3, 1, 3, 3}

func TestPermutations(t *testing.T) {
	permuts := Permutations(boardValues)
	if len(permuts) != 362880 {
		t.Errorf("Permutations(%v) = %v, want %v", boardValues, len(permuts), 7)
	}
}

func BenchmarkPermutations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Permutations(boardValues)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	permuts := Permutations(boardValues)
	pre_length := len(permuts)
	permuts = RemoveRotations(permuts)
	switch len(permuts) {
	case pre_length:
		t.Error("Nothing was removed :(")
	case 90720:
		t.Log("All is well")
	default:
		t.Error("Something has gone wrong in a new and exciting way")
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveRotations(rotations)
	}

}

func TestRotations(t *testing.T) {
	permuts := RepeatingPermutations(&rotationValues, 9)
	if len(permuts) != 262144 {
		t.Errorf("Permutations(%v) = %v, want %v", boardValues, len(permuts), 7)
	}
}

func BenchmarkRotations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatingPermutations(&rotationValues, 9)
	}
}

func TestMakeBoard(t *testing.T) {
	board := MakeBoard(board_layouts[0], rotations[0])
	t.Log(board)
}

func BenchmarkMakeBoard(b *testing.B) {

	for i := 0; i < b.N && i < len(board_layouts)-1; i++ {
		MakeBoard(board_layouts[i], rotations[i])
	}
}

func TestCheckBoard(t *testing.T) {
	board := MakeBoard(board_layouts[0], rotations[0])
	t.Log(board)
	result := CheckBoard(&board)
	t.Log(result)
	if result {
		t.Error("CheckBoard failed")
	} else {
		t.Log("CheckBoard succeeded")
	}

}

func BenchmarkCheckBoard(b *testing.B) {
	for i := 0; i < b.N && i < len(testBoards)-1; i++ {
		CheckBoard(&testBoards[i])
	}
}

func TestKnownSolutions(t *testing.T) {
	t.Log(known_good_board)
	t.Log(known_good_rotat)
	board := MakeBoard(known_good_board, known_good_rotat)
	t.Log(board)
	for idx, piece := range known_good_board {
		t.Logf("Position %v - Token %v before rotating %v times: %v \n", idx+1, piece, known_good_rotat[idx], tokens[piece])
		t.Logf("In actuality %v \n", tokenToText(tokens[piece]))
		rotated, _ := Rotate(tokens[piece], known_good_rotat[idx])
		t.Logf("Position %v - Token %v after rotating %v times: %v \n", idx+1, piece, known_good_rotat[idx], rotated)
		t.Logf("In actuality %v \n", tokenToText(rotated))
		t.Log("----------------")
	}
	res := CheckBoard(&board)
	edges := edgeFinder(&board)
	pairs := edgePairs(edges)
	t.Log(pairs)
	for idx, pair := range pairs {
		t.Log("----------------")
		t.Logf("Pair %v is numbers %v\n", idx, pair)
		t.Logf("Colours are: %v %v \n", designs[pair.e1], designs[pair.e2])
	}
	t.Logf("Result of good solution is %v\n", res)

	if !res {
		t.Error("Failed to deal with known good solution")
	} else {
		t.Log("Accepted known good solution")
	}
}

func TestCreateAndCheckBoards(t *testing.T) {
	var orders = make([][]int, 0)
	var rotations = make([][]int, 0)
	orders = append(orders, known_good_board)
	rotations = append(rotations, known_good_rotat)

	CreateAndCheckBoards(orders, rotations)

	b := MakeBoard(orders[0], rotations[0])
	t.Log(b)
	e := edgeFinder(&b)
	ep := edgePairs(e)
	t.Log(ep)
}
