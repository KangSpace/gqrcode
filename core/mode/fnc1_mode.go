package mode

// Define FNC1 mode handle
// FNC1 mode is used for messages containing specific data formats.
// In the "1st position" it designates data formatted in accordance with the GS1 General Specifications.
// In the "2nd position" it designates data formatted in accordance with a specific industry application previously agreed with AIM Inc.
// FNC1 mode applies to the entire symbol and is not affected by subsequent mode indicators.
// NOTE "1st position" and "2nd position" do not refer to actual locations but are based on the positions of the character in Code 128 symbols,when used in an equivalent manner.
//
// Tips:
// 1. FNC1 mode is not available for Micro QR Code Symbol.
//


type FNC1 struct {
	*AbstractMode
}

// TODO NOT IMPLEMENT