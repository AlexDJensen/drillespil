package main

/*
Token is a single token and the four colour-half pairs it contains.
*/
type Token struct {
	north uint8
	east  uint8
	south uint8
	west  uint8
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

type edgePair struct {
	e1 uint8
	e2 uint8
}
