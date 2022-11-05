package main

import (
	"reflect"
	"testing"
)

var length = 9
var boardValues []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var rotationValues []int = []int{0, 1, 2, 3}
var layouts [][]int = Permutations(boardValues)
var board_layouts [][]int = RemoveRotations((layouts))
var rotations [][]int = RepeatingPermutations(&rotationValues, length)
var testBoards []Board = MakeBoardsLimited(board_layouts, rotations, 100000)

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

func TestObjectSizes(t *testing.T) {
	board := MakeBoard(board_layouts[0], rotations[0])
	t.Log(board)
	t.Logf("A single board has size %d", reflect.TypeOf(board).Size())

}

func TestLimitedBoards(t *testing.T) {
	param := 10000
	boards := MakeBoardsLimited(board_layouts, rotations, param)
	if len(boards) != param {
		t.Errorf("List of boards is size %d, want %d", len(boards), param)
	} else {
		t.Logf("List of boards is size %d, want %d", len(boards), param)
	}
}
