package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/*
Token is a single token and the four colour-half pairs it contains.
*/
type Token struct {
	north string
	east  string
	south string
	west  string
}

/*
Board is the controlling structure in this game - it places tokens and compares edges
*/
type Board struct {
	p1 Token
	p2 Token
	p3 Token
	p4 Token
	p5 Token
	p6 Token
	p7 Token
	p8 Token
	p9 Token
}

type edges struct {
	p1e string
	p1s string
	p2e string
	p2s string
	p2w string
	p3s string
	p3w string
	p4n string
	p4e string
	p4s string
	p5n string
	p5e string
	p5s string
	p5w string
	p6n string
	p6s string
	p6w string
	p7n string
	p7e string
	p8n string
	p8e string
	p8w string
	p9n string
	p9w string
}

/*
edgeFinder takes in a board, and returns a structure of the edges that exist on a given board.
The board is expected to consist of rotated tokens
*/
func edgeFinder(b Board) (e edges) {
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
	e1 string
	e2 string
}

func edgePairs(e edges) (p []edgePair) {
	p1 := edgePair{e.p1e, e.p2w}
	p2 := edgePair{e.p1s, e.p4n}
	p3 := edgePair{e.p2e, e.p3w}
	p4 := edgePair{e.p2s, e.p5n}
	p5 := edgePair{e.p3s, e.p6n}
	p6 := edgePair{e.p4e, e.p5w}
	p7 := edgePair{e.p4s, e.p7n}
	p8 := edgePair{e.p6w, e.p5e}
	p9 := edgePair{e.p6s, e.p9n}
	p10 := edgePair{e.p7e, e.p8e}
	p11 := edgePair{e.p8n, e.p5s}
	p12 := edgePair{e.p8e, e.p9w}
	p = append(p, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12)
	return p
}

// Match checks if edgepair is valid, returns bool
func Match(e edgePair) bool {
	if e.e1 == e.e2 {
		return true
	}
	return false

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
		errorMessage := ("Rotation argument is too high (" +
			strconv.Itoa(rotation) +
			"), running with " +
			strconv.Itoa(newRotation) +
			" instead")
		err = errors.New(errorMessage)
	}

	return r, err
}

/*
Two is a test function.
*/
func Two() {

	rand.Seed(time.Now().UnixNano())

	brick := Token{"AB", "CD", "EF", "FG"}
	start := time.Now()

	runs := 1
	maxRuns := 1000000
	for runs < maxRuns {
		new, _ := Rotate(brick, rand.Intn(4))

		if runs == maxRuns/2 {
			fmt.Printf("%#v\n", new)
			fmt.Printf("Now on run %v\n", runs)
		}
		runs = runs + 1
	}

	elapsed := time.Since(start)
	fmt.Printf("Took %s\n", elapsed)

}

/*
HeapPermutation does something
*/
func HeapPermutation(a []int, size int) {

	for i := 0; i < size; i++ {
		HeapPermutation(a, size-1)

		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}

/*
PermutationsWithoutRepetition create permutations in a list.
As arguments, it takes an array of integers and a length argument, and returns an array of arrays of integers
*/
func PermutationsWithoutRepetition(L []int, r int) [][]int {
	if r == 1 {
		//Convert every item in L to List and
		//Append it to List of List
		temp := make([][]int, 0)
		for _, rr := range L {
			t := make([]int, 0)
			t = append(t, rr)
			temp = append(temp, [][]int{t}...)
		}
		return temp
	}
	res := make([][]int, 0)
	for i := 0; i < len(L); i++ {
		//Create List Without L[i] element
		perms := make([]int, 0)
		perms = append(perms, L[:i]...)
		perms = append(perms, L[i+1:]...)
		//Call recursively to Permutations
		for _, x := range PermutationsWithoutRepetition(perms, r-1) {
			t := append(x, L[i])
			res = append(res, [][]int{t}...)
		}
	}
	return res

}

func main() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//b := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	res := PermutationsWithoutRepetition(a, len(a))
	fmt.Println(len(res))

	fmt.Printf("%+v\n", res[len(res)-2])

	elapsed := time.Since(start)
	fmt.Printf("Took %s\n", elapsed)
}
