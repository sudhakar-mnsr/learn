package vlq

import (
	"math"
	"math/bits"
)

// DecodeVarint takes a variable length VLQ based integer and
// decodes it into a 32 bit integer.
func DecodeVarint(input []byte) (uint32, error) {
	const lastBitSet = 0x80 // 1000 0000

	var d uint32
	var bitPos int

	for i := len(input) - 1; i >= 0; i-- {
		n := uint8(input[i])

		// Process the first 7 bits and ignore the 8th.
		for checkBit := 0; checkBit < 7; checkBit++ {

			// Rotate the last bit off and move it to the back.
			// Before: 0000 0001
			// After:  1000 0000
			n = bits.RotateLeft8(n, -1)

			// Calculate based on only those 1 bits that were rotated.
			// Convert the bitPos to base 10.

			if n >= lastBitSet {
				switch {
				case bitPos == 0:
					d++
				default:
					base10 := math.Pow(2, float64(bitPos))
					d += uint32(base10)
				}
			}
