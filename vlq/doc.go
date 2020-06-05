/*
Package vlq implements VLQ encoding/decoding.
In short, the goal of this encoding is to save encode integer values in
a way that would save bytes. Only the first 7 bits of each byte is significant
(right-justified; sort of like an ASCII byte). So, if you have a 32-bit value,
you have to unpack it into a series of 7-bit bytes. Of course, you will have
a variable number of bytes depending upon your integer. To indicate which
is the last byte of the series, you leave bit #7 clear. In all of the
preceding bytes, you set bit #7.

So, if an integer is between 0-127, it can be represented as one byte. The
largest integer allowed is 0FFFFFFF, which translates to 4 bytes variable
length. Here are examples of delta-times as 32-bit values, and the variable
length quantities that they translate to:

 NUMBER        VARIABLE QUANTITY
00000000              00
00000040              40
0000007F              7F
00000080             81 00
00002000             C0 00
00003FFF             FF 7F
00004000           81 80 00
00100000           C0 80 00
001FFFFF           FF FF 7F
00200000          81 80 80 00
08000000          C0 80 80 00
0FFFFFFF          FF FF FF 7F


Resources:
https://en.wikipedia.org/wiki/Variable-length_quantity
A variable-length quantity (VLQ) is a universal code that uses an arbitrary
number of binary octets (eight-bit bytes) to represent an arbitrarily large
integer. It was defined for use in the standard MIDI file format[1] to save
additional space for a resource constrained system, and is also used in the
later Extensible Music Format (XMF). A VLQ is essentially a base-128
representation of an unsigned integer with the addition of the eighth bit
to mark continuation of bytes. See the example below.

	Int:    16384
	IntHex: 0x00004000
	IntBin: 00000000 00000000 01000000 00000000
	VLQHex: 0x81 0x80 0x00
	VLQBin: 00000000 10000001 10000000 00000000
https://blogs.infosupport.com/a-primer-on-vlq/
