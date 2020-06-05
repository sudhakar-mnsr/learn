/*
	// This is the API you need to build for these tests. You will need to
	// change the import path in this test to point to your code.
	package vlq
	// DecodeVarint takes a variable length VLQ based integer and
	// decodes it into a 32 bit integer.
	func DecodeVarint(input []byte) (uint32, error)
	// EncodeVarint takes a 32 bit integer and encodes it into
	// a variable length VLQ based integer.
	func EncodeVarint(n uint32) []byte
*/

package vlq

import (
	"bytes"
	"testing"
)

func TestEncodeDecodeVarint(t *testing.T) {
	testCases := []struct {
		input  []byte
		output uint32
	}{
		0:  {[]byte{0x7F}, 127},
		1:  {[]byte{0x81, 0x00}, 128},
		2:  {[]byte{0xC0, 0x00}, 8192},
		3:  {[]byte{0xFF, 0x7F}, 16383},
		4:  {[]byte{0x81, 0x80, 0x00}, 16384},
