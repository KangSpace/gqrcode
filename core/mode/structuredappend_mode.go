package mode

// Define Structured Append mode handle
// Structured Append mode is used to split the encoding of the data from a message over a number of QR Code symbols.
// All of the symbols require to be read and the data message an be reconstructed in the correct sequence.
// The Structured Append header is encoded in each symbol to identify the length of the sequence and the symbol's position in it,
// 	and verify that all the symbols read belong to the same message.
// Refer to 8for details of encoding in Structured Append mode.
//
// Tips:
// 1. Structured Append mode is not available for Micro QR Code Symbol.
//


type StructuredAppend struct {
	*AbstractMode
}

// TODO NOT IMPLEMENT