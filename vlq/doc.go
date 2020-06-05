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
