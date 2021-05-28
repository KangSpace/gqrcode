package mode

// Define Kanji mode handle
// The Kanji mode efficiently encodes Kanji characters in accordance with the Shift JIS system based on JIS X 0209.
// The Shift JIS values are shifted from the JIS X 0208 values.
// JIS X 0208 gives details of the shift coded representation.
// Each two-byte character value is compacted to a 13-bit binary codeword.
//
// Tips:
// 1. Kanji mode is not available in Version M1 Or M2 Micro QR Code Symbol.
//


type Kanji struct {
	*AbstractMode
}