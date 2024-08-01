package shannon_fano

import (
	"math"
)

type Generator struct{}

type CharStat map[rune]int

type encodingTable map[rune]code

type code struct {
	Char     rune
	Quantity int
	Bit      uint32
	Size     int
}

func NewGenerator() Generator {
	return Generator{}
}

//func (g Generator) NewTable(txt string) table.EncodingTable {
//	//stat := NewCharStat(txt)
//
//	//encoding table
//
//	//return table.encodingTable
//
//}

//func build(stat CharStat) encodingTable {
//	codes := make([]code, len(stat))
//
//	for ch, qty := range stat {
//		codes = append(codes, code{
//			Char:     ch,
//			Quantity: qty,
//		})
//	}
//
//	sort.Slice(codes, func(i, j int) bool {
//		if codes[i].Quantity > codes[j].Quantity {
//			return codes[i].Quantity > codes[j].Quantity
//		}
//
//		return codes[i].Char < codes[j].Char
//	})
//
//	assignCodes(codes)
//
//}

func assignCodes(codes []code) {
	if len(codes) < 2 {
		return
	}

}

func bestDividerPosition(codes []code) int {
	total := 0

	for _, code := range codes {
		total += code.Quantity
	}

	left := 0
	prefDiff := math.MaxInt
	bestPosition := 0
	for i := 0; i < len(codes)-1; i++ {
		right := total - left

		diff := abs(right - left)
		if diff >= prefDiff {
			break
		}
		prefDiff = diff

		bestPosition = i + 1
	}
	return bestPosition
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func NewCharStat(txt string) CharStat {
	res := make(CharStat)

	for _, ch := range txt {
		res[ch]++
	}

	return res
}
