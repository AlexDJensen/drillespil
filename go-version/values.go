package main

var designs = map[uint8]string{
	// This is correct
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
	// Rechecked - this is not the issue
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
