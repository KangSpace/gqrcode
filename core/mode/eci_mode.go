package mode

// Define ECI mode handle
// The Extended Channel Interpretation(ECI) protocol defined in the AIM Inc.
// International Technical Specification Extended Channel Interpretations, allows the output data stream to have interpretations different from that of the default character set.
// The ECI protocol is defined consistently across a number of symbologies.
// The ECI protocol provides a consistent method to specify particular interpretations of byte values before printing and after decoding.
// The ECI protocol is not supported in Micro QR Code symbols.
// The default interpretation for QR Code is ECI 000003 representing the ISO/ECI 8859-1 character set.
//
// The ECI header(if present) shall comprise:
// --- ECI Mode Indicator(4 bits)
// --- ECI Designator(8,16 or 24 bits)
//


type ECI struct {
	*AbstractMode
}