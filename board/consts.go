package board

import "sort"

/*
A file used to store constant values associated with boards
*/

// Arrays and slices can't be consts because they aren't immutable but the data will never change so it's going here along with the rest of our constants
var bitMap = []string{
	"a1",
	"b1",
	"c1",
	"d1",
	"e1",
	"f1",
	"g1",
	"h1",

	"a2",
	"b2",
	"c2",
	"d2",
	"e2",
	"f2",
	"g2",
	"h2",

	"a3",
	"b3",
	"c3",
	"d3",
	"e3",
	"f3",
	"g3",
	"h3",

	"a4",
	"b4",
	"c4",
	"d4",
	"e4",
	"f4",
	"g4",
	"h4",

	"a5",
	"b5",
	"c5",
	"d5",
	"e5",
	"f5",
	"g5",
	"h5",

	"a6",
	"b6",
	"c6",
	"d6",
	"e6",
	"f6",
	"g6",
	"h6",

	"a7",
	"b7",
	"c7",
	"d7",
	"e7",
	"f7",
	"g7",
	"h7",

	"a8",
	"b8",
	"c8",
	"d8",
	"e8",
	"f8",
	"g8",
	"h8",
}

const aFile Bitboard = 0b_10000000_10000000_10000000_10000000_10000000_10000000_10000000_10000000
const bFile Bitboard = 0b_01000000_01000000_01000000_01000000_01000000_01000000_01000000_01000000
const cFile Bitboard = 0b_00100000_00100000_00100000_00100000_00100000_00100000_00100000_00100000
const dFile Bitboard = 0b_00010000_00010000_00010000_00010000_00010000_00010000_00010000_00010000
const eFile Bitboard = 0b_00001000_00001000_00001000_00001000_00001000_00001000_00001000_00001000
const fFile Bitboard = 0b_00000100_00000100_00000100_00000100_00000100_00000100_00000100_00000100
const gFile Bitboard = 0b_00000010_00000010_00000010_00000010_00000010_00000010_00000010_00000010
const hFile Bitboard = 0b_00000001_00000001_00000001_00000001_00000001_00000001_00000001_00000001

const row2 Bitboard = 0b_00000000_00000000_00000000_00000000_00000000_00000000_11111111_00000000
const row7 Bitboard = 0b_00000000_11111111_00000000_00000000_00000000_00000000_00000000_00000000

func getIndex(slice []string, value string) int {
	return sort.StringSlice(slice).Search(value)
}
