package board

/*
bitboard.go
A file to keep all of the data associated with the bitboard type
*/

type Bitboard uint64

func (b *Bitboard) setBit(pos uint) {
	*b |= (1 << pos)
}

func (b *Bitboard) clearBit(pos uint) {
	*b &^= (1 << pos)
}

func (b Bitboard) hasBit(pos uint) bool {
	val := b & (1 << pos)
	return (val > 0)
}

func (b Bitboard) getOccupiedSqrs() []uint {
	var sqrKeys []uint
	for i := uint(0); i < 64; i++ {
		if b.hasBit(uint(i)) {
			sqrKeys = append(sqrKeys, i)
		}
	}
	return sqrKeys
}

func combine(boards []Bitboard) Bitboard {
	var result Bitboard = 0
	for _, board := range boards {
		result = result | board
	}
	return result
}
