package board

import (
	"strconv"
	"strings"
)

/*
We're using bitboards a way of representing a chess board using an unsigned 64 bit number
see : https://pages.cs.wisc.edu/~psilord/blog/data/chess-pages/index.html
each square on the chess board is mapped to a position on the bitboard with a1 being the LSB and h8 being the MSB
*/

type HalfBoard struct {
	// represents the pieces of one of the players
	pawns   Bitboard
	rooks   Bitboard
	knights Bitboard
	bishops Bitboard
	queens  Bitboard
	king    Bitboard

	kingsideCastle  bool
	queensideCastle bool
}

type Board struct {
	whiteMove     bool
	enPassantSqr  int // The square behind a pawn that has just moved two squares, used for the en passant rule set to -1 if empty
	halfMoveClock int // The number of half moves since the last capture or pawn advance, used for the fifty move rule
	fullMoveClock int // The total number of moves so far in the game idrk what it's used for but it's in the FEN standard so it must be important
	white         HalfBoard
	black         HalfBoard
}

func CreateBoard(fen string) Board {
	// Converts a FEN (https://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation) to a Board struct
	emptyHalfBoard := HalfBoard{
		pawns:           0,
		rooks:           0,
		knights:         0,
		bishops:         0,
		queens:          0,
		king:            0,
		kingsideCastle:  false,
		queensideCastle: false,
	}
	board := Board{
		whiteMove:     true,
		enPassantSqr:  -1,
		halfMoveClock: 0,
		fullMoveClock: 0,
		white:         emptyHalfBoard,
		black:         emptyHalfBoard,
	}
	parts := strings.Split(fen, " ")
	var i uint = 0
	for _, row := range strings.Split(parts[0], "/") {
		for _, c := range row {
			switch c {

			case 49, 50, 51, 52, 53, 54, 55, 56:
				// if c is [1 - 8] jump the pointer accordingly
				i += uint(c) - 49

			// White Pieces

			case 112:
				// white pawn
				board.white.pawns.setBit(i)
			case 114:
				// white rook
				board.white.rooks.setBit(i)
			case 110:
				// white knight
				board.white.knights.setBit(i)
			case 98:
				// white bishop
				board.white.bishops.setBit(i)
			case 113:
				// white queen
				board.white.queens.setBit(i)
			case 107:
				// white king
				board.white.king.setBit(i)

			// Black Pieces

			case 80:
				// black pawn
				board.black.pawns.setBit(i)
			case 82:
				// black rook
				board.black.rooks.setBit(i)
			case 78:
				// black knight
				board.black.knights.setBit(i)
			case 66:
				// black bishop
				board.black.bishops.setBit(i)
			case 81:
				// black queen
				board.black.queens.setBit(i)
			case 75:
				// black king
				board.black.king.setBit(i)

			}
			i++
		}
	}

	board.whiteMove = parts[1] == "w"
	board.white.kingsideCastle = strings.Contains(parts[2], "K")
	board.black.kingsideCastle = strings.Contains(parts[2], "k")
	board.white.queensideCastle = strings.Contains(parts[2], "Q")
	board.black.queensideCastle = strings.Contains(parts[2], "q")

	board.enPassantSqr = getIndex(bitMap, parts[3])

	board.halfMoveClock, _ = strconv.Atoi(parts[4])
	board.fullMoveClock, _ = strconv.Atoi(parts[5])

	return board
}

func getHalfOccupiedSquares(board HalfBoard) Bitboard {
	return board.pawns | board.rooks | board.knights | board.bishops | board.queens | board.king
}

func getAllOccupiedSquares(board Board) Bitboard {
	return getHalfOccupiedSquares(board.white) | getHalfOccupiedSquares(board.black)
}

func GenerateMoves(board Board) []Move {
	return []Move{}
}

/*
func generatePsuedoLegalPawnMoves(board Board) []Move {
	var myTeam HalfBoard
	var enemyTeam HalfBoard
	var startRank Bitboard
	var secondLRank Bitboard
	var vector int
	if board.whiteMove {
		myTeam = board.white
		enemyTeam = board.black
		startRank = row2
		secondLRank = row7

		vector = 8

	} else {
		myTeam = board.black
		enemyTeam = board.white
		secondLRank = row2
		startRank = row7
		vector = -8
	}
	pawns := myTeam.pawns.getOccupiedSqrs()

	for _, pawn := range pawns {
		pbb := Bitboard(0)
		pbb.setBit(pawn) // bitboard representing the square the pawn occupies

		// Generate Captures
		enemySquares := getHalfOccupiedSquares(enemyTeam)
		if board.enPassantSqr >= 0 {
			enPassantBB := Bitboard(0)
			enPassantBB.setBit(uint(board.enPassantSqr))
			enemySquares = enemySquares | enPassantBB
		}
		if aFile&pbb == 0 {
			attackSqr := Bitboard(0)
			attackSqr.setBit(uint(int(pawn) + vector + 1))

			if (enemySquares)&attackSqr > 0 {

			}

		}
	}

}
*/
