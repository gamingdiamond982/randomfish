package board

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type MoveType uint

const (
	QUIET MoveType = iota
	CAPTURE
	EVASION
	ENPASSANT
	CASTLING
)

type PieceType uint

const (
	NullPiece PieceType = iota // A NULL Piece Type cus GoLang doesn't have a null type and I wan't to be able to use the nil type as null or smthing idk
	Pawm
	Knight
	Bishop
	Rook
	Queen
	King
)

type Piece struct {
	pieceType PieceType
	isWhite   bool
}

var PieceSymbols = []string{
	"",
	"p",
	"n",
	"b",
	"r",
	"q",
	"k",
}

func (p Piece) String() string {
	r := PieceSymbols[uint(p.pieceType)]
	if p.isWhite {
		return strings.ToUpper(r)
	}
	return r
}

func PieceFromSymbol(symbol string) Piece {
	var pieceType PieceType = PieceType(sort.StringSlice(PieceSymbols).Search(strings.ToLower(symbol)))
	var isWhite bool = strings.ToUpper(symbol) == symbol
	return Piece{
		pieceType: pieceType,
		isWhite:   isWhite,
	}
}

type Move struct {
	startSqr  int
	endSqr    int
	promotion PieceType // will be the NullPiece if not promoting
	drop      PieceType // will be the NullPiece if not used
}

func (move Move) uci() string {
	/*
		Returns a UCI string representing the move
		as specified at https://www.shredderchess.com/download/div/uci.zip

	*/
	if move.drop != NullPiece {
		return strings.ToUpper(PieceSymbols[move.drop]) + "@" + bitMap[move.endSqr]
	} else if move.promotion != NullPiece {
		return bitMap[move.startSqr] + bitMap[move.endSqr] + PieceSymbols[move.promotion]
	} else if move.startSqr != -1 && move.endSqr != -1 {
		return bitMap[move.startSqr] + bitMap[move.endSqr]
	} else {
		return "0000" // For Null moves UCI says to send 0000
	}
}

func MoveFromUci(uci string) (Move, error) {
	if uci == "0000" {
		return Move{-1, -1, NullPiece, NullPiece}, nil
	} else if len(uci) == 4 && uci[1:2] == "@" {
		drop := PieceType(getIndex(PieceSymbols, uci[0:1]))
		square := getIndex(bitMap, uci[2:])
		return Move{square, square, NullPiece, drop}, nil
	} else if 4 <= len(uci) && len(uci) <= 5 {
		from_square := getIndex(bitMap, uci[0:2])
		to_square := getIndex(bitMap, uci[2:4])
		var promotion PieceType
		if len(uci) == 5 {
			promotion = PieceType(getIndex(PieceSymbols, uci[4:5]))
		} else {
			promotion = NullPiece
		}
		if from_square == to_square {
			return Move{}, errors.New("Invalid UCI Move")
		}
		return Move{
			from_square,
			to_square,
			promotion,
			NullPiece,
		}, nil
	}
	return Move{}, errors.New(fmt.Sprintf("Expected uci string to between 4 and 5 instead got: %v\n of length: %v", uci, len(uci)))
}
