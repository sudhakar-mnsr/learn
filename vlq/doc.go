/*
Package vlq implements VLQ encoding/decoding.
In short, the goal of this encoding is to save encode integer values in
a way that would save bytes. Only the first 7 bits of each byte is significant
(right-justified; sort of like an ASCII byte). So, if you have a 32-bit value,
you have to unpack it into a series of 7-bit bytes. Of course, you will have
a variable number of bytes depending upon your integer. To indicate which
is the last byte of the series, you leave bit #7 clear. In all of the
preceding bytes, you set bit #7.
