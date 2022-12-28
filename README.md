# Drillespil

This is a repo of code to solve a puzzle. The repo will eventually contain several solutions spanning several languages, as well as images and rules to allow others to participate.

## Rules of the game:

### The manual way:

1. The 9 square tokens pictured below has to be arranged in a 3x3 grid, such that all pictures align (_TODO_: see below). Unfortunately, some of the sides can be hard to distinguish.
1. Two tokens align if the clothes on each side of an edge are of the same color, and both a top and bottom half is represented on an edge.
1. If all tokens align on all sides touching other tokens (outer edges don't matter), the solution is valid.

For ease of reproducability, the token definitions are found in `./data.json`.
Note that the selection of a north-face/orientation for the tokens is entirely arbitrary - it is merely there to have consistency.

## Current solutions:

- `./go-version` contains a solution in Golang. Code is not yet organized too prettily, and might never be.
