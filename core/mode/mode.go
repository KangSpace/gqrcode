package mode

import (
	"errors"
	"github.com/KangSpace/gqrcode/core/cons"
	"github.com/KangSpace/gqrcode/core/logger"
	"github.com/KangSpace/gqrcode/core/model"
	"github.com/KangSpace/gqrcode/core/output"
	"github.com/KangSpace/gqrcode/util"
	"strconv"
)

// Mode : Contain all of the steps to handle input data to QRCode image in this file .
// Data encoding:
// 	Each mode segment shall begin with the first (most significant) bit of the mode indicator and end with the final(least significant) bit of the data bit stream.
// 		Mode Indicator
// 		Character count Indicator
// 		Data bit stream
type Mode interface {
	DataEncode(qr *QRCodeStruct) (dataStream *util.DataStream)
	GetMode() *AbstractMode
	IsSupport(data string) bool
}

type AbstractMode struct {
	Name cons.ModeType `json:"name"`
}

func (m *AbstractMode) isSupport(data string) bool {
	return false
}

// SupportModes :Array of All models
var SupportModes = []Mode{
	NewNumericMode(),
	NewAlphanumericMode(),
	NewKanjiModeMode(),
	NewByteModeMode(),
	//ECI_MODE,
	//STRUCTUREDAPPEND_MODE,
	//FNC1_MODE,
	//FNC1_MODE_P1,
	//FNC1_MODE_P2,
}

// unionIndicator : Struct for Mode Indicator and Number Of Count Indicator
type unionIndicator struct {
	Mode                         cons.ModeType
	modeIndicatorDetails         []*modeIndicatorDetail
	CountIndicatorBitsNumDetails []*CountIndicatorBitsNumberDetail
}

// modeIndicatorDetail : The ModeIndicator detail of Mode
type modeIndicatorDetail struct {
	Format cons.Format
	// model.VersionAll("0") for all
	Version             model.VersionId
	ModeIndicatorLength uint8
	// The Mode and Version not supported,if ModeIndicatorCode is null,
	ModeIndicatorCode []util.Bit
}

type CountIndicatorBitsNumberDetail struct {
	Versions []model.VersionId
	// Defined the bits length , not represent the full bits for max count , max count in Table7
	NumberOfBits int
	// Max count by NumberOfBits, value is math.Pow(2,NumberOfBits)-1
}

// UnionIndicators :
// 1. Mode Indicator: Table 2 - Mode indicators for QR Code
// 2. CountIndicatorBitsNumbers : Table3 - Number of bits in character count indicator for QR Code
// struct: []*UnionIndicator
var unionIndicators = []*unionIndicator{
	{cons.EciMode, []*modeIndicatorDetail{
		{cons.QRCODE, model.VersionAll, 4, []byte{0, 1, 1, 1}},
		{cons.MicroQrcode, model.VersionM1, 0, nil},
		{cons.MicroQrcode, model.VersionM2, 1, nil},
		{cons.MicroQrcode, model.VersionM3, 2, nil},
		{cons.MicroQrcode, model.VersionM4, 3, nil}}, nil},
	{cons.NumericMode, []*modeIndicatorDetail{
		{cons.QRCODE, model.VersionAll, 4, []byte{0, 0, 0, 1}},
		{cons.MicroQrcode, model.VersionM1, 0, nil},
		{cons.MicroQrcode, model.VersionM2, 1, []byte{0}},
		{cons.MicroQrcode, model.VersionM3, 2, []byte{0, 0}},
		{cons.MicroQrcode, model.VersionM4, 3, []byte{0, 0, 0}}},
		[]*CountIndicatorBitsNumberDetail{
			{[]model.VersionId{model.VersionM1}, 3},
			{[]model.VersionId{model.VersionM2}, 4},
			{[]model.VersionId{model.VersionM3}, 5},
			{[]model.VersionId{model.VersionM4}, 6},
			//VERSION_1 to VERSION_9
			{[]model.VersionId{model.VERSION1, model.VERSION2, model.VERSION3, model.VERSION4, model.VERSION5,
				model.VERSION6, model.VERSION7, model.VERSION8, model.VERSION9}, 10},
			//VERSION_10 to VERSION_26
			{[]model.VersionId{model.VERSION10, model.VERSION11, model.VERSION12, model.VERSION13, model.VERSION14,
				model.VERSION15, model.VERSION16, model.VERSION17, model.VERSION18, model.VERSION19, model.VERSION20,
				model.VERSION21, model.VERSION22, model.VERSION23, model.VERSION24, model.VERSION25, model.VERSION26}, 12},
			//VERSION_27 to VERSION_40
			{[]model.VersionId{model.VERSION27, model.VERSION28, model.VERSION29, model.VERSION30, model.VERSION31,
				model.VERSION32, model.VERSION33, model.VERSION34, model.VERSION35, model.VERSION36, model.VERSION37,
				model.VERSION38, model.VERSION39, model.VERSION40}, 14}}},
	{cons.AlphanumericMode, []*modeIndicatorDetail{
		{cons.QRCODE, model.VersionAll, 4, []byte{0, 0, 1, 0}},
		{cons.MicroQrcode, model.VersionM1, 0, nil},
		{cons.MicroQrcode, model.VersionM2, 1, []byte{1}},
		{cons.MicroQrcode, model.VersionM3, 2, []byte{0, 1}},
		{cons.MicroQrcode, model.VersionM4, 3, []byte{0, 0, 1}}},
		[]*CountIndicatorBitsNumberDetail{
			{[]model.VersionId{model.VersionM1}, 0},
			{[]model.VersionId{model.VersionM2}, 3},
			{[]model.VersionId{model.VersionM3}, 4},
			{[]model.VersionId{model.VersionM4}, 5},
			//VERSION_1 to VERSION_9
			{[]model.VersionId{model.VERSION1, model.VERSION2, model.VERSION3, model.VERSION4, model.VERSION5,
				model.VERSION6, model.VERSION7, model.VERSION8, model.VERSION9}, 9},
			//VERSION_10 to VERSION_26
			{[]model.VersionId{model.VERSION10, model.VERSION11, model.VERSION12, model.VERSION13, model.VERSION14,
				model.VERSION15, model.VERSION16, model.VERSION17, model.VERSION18, model.VERSION19, model.VERSION20,
				model.VERSION21, model.VERSION22, model.VERSION23, model.VERSION24, model.VERSION25, model.VERSION26}, 11},
			//VERSION_27 to VERSION_40
			{[]model.VersionId{model.VERSION27, model.VERSION28, model.VERSION29, model.VERSION30, model.VERSION31,
				model.VERSION32, model.VERSION33, model.VERSION34, model.VERSION35, model.VERSION36, model.VERSION37,
				model.VERSION38, model.VERSION39, model.VERSION40}, 13}}},
	{cons.ByteMode, []*modeIndicatorDetail{
		{cons.QRCODE, model.VersionAll, 4, []byte{0, 1, 0, 0}},
		{cons.MicroQrcode, model.VersionM1, 0, nil},
		{cons.MicroQrcode, model.VersionM2, 1, nil},
		{cons.MicroQrcode, model.VersionM3, 2, []byte{1, 0}},
		{cons.MicroQrcode, model.VersionM4, 3, []byte{0, 1, 0}}},
		[]*CountIndicatorBitsNumberDetail{
			{[]model.VersionId{model.VersionM1}, 0},
			{[]model.VersionId{model.VersionM2}, 0},
			{[]model.VersionId{model.VersionM3}, 4},
			{[]model.VersionId{model.VersionM4}, 5},
			//VERSION_1 to VERSION_9
			{[]model.VersionId{model.VERSION1, model.VERSION2, model.VERSION3, model.VERSION4, model.VERSION5,
				model.VERSION6, model.VERSION7, model.VERSION8, model.VERSION9}, 8},
			//VERSION_10 to VERSION_26
			{[]model.VersionId{model.VERSION10, model.VERSION11, model.VERSION12, model.VERSION13, model.VERSION14,
				model.VERSION15, model.VERSION16, model.VERSION17, model.VERSION18, model.VERSION19, model.VERSION20,
				model.VERSION21, model.VERSION22, model.VERSION23, model.VERSION24, model.VERSION25, model.VERSION26}, 16},
			//VERSION_27 to VERSION_40
			{[]model.VersionId{model.VERSION27, model.VERSION28, model.VERSION29, model.VERSION30, model.VERSION31,
				model.VERSION32, model.VERSION33, model.VERSION34, model.VERSION35, model.VERSION36, model.VERSION37,
				model.VERSION38, model.VERSION39, model.VERSION40}, 16}}},
	{cons.KanjiMode, []*modeIndicatorDetail{
		{cons.QRCODE, model.VersionAll, 4, []byte{1, 0, 0, 0}},
		{cons.MicroQrcode, model.VersionM1, 0, nil},
		{cons.MicroQrcode, model.VersionM2, 1, nil},
		{cons.MicroQrcode, model.VersionM3, 2, []byte{1, 1}},
		{cons.MicroQrcode, model.VersionM4, 3, []byte{0, 1, 1}}},
		[]*CountIndicatorBitsNumberDetail{
			{[]model.VersionId{model.VersionM1}, 0},
			{[]model.VersionId{model.VersionM2}, 0},
			{[]model.VersionId{model.VersionM3}, 3},
			{[]model.VersionId{model.VersionM4}, 4},
			{[]model.VersionId{model.VERSION1, model.VERSION2, model.VERSION3, model.VERSION4, model.VERSION5,
				model.VERSION6, model.VERSION7, model.VERSION8, model.VERSION9}, 8},
			//VERSION_10 to VERSION_26
			{[]model.VersionId{model.VERSION10, model.VERSION11, model.VERSION12, model.VERSION13, model.VERSION14,
				model.VERSION15, model.VERSION16, model.VERSION17, model.VERSION18, model.VERSION19, model.VERSION20,
				model.VERSION21, model.VERSION22, model.VERSION23, model.VERSION24, model.VERSION25, model.VERSION26}, 10},
			//VERSION_27 to VERSION_40
			{[]model.VersionId{model.VERSION27, model.VERSION28, model.VERSION29, model.VERSION30, model.VERSION31,
				model.VERSION32, model.VERSION33, model.VERSION34, model.VERSION35, model.VERSION36, model.VERSION37,
				model.VERSION38, model.VERSION39, model.VERSION40}, 12}}},
	{cons.StructuredAppendMode, []*modeIndicatorDetail{
		{cons.QRCODE, model.VersionAll, 4, []byte{0, 0, 1, 1}},
		{cons.MicroQrcode, model.VersionM1, 0, nil},
		{cons.MicroQrcode, model.VersionM2, 1, nil},
		{cons.MicroQrcode, model.VersionM3, 2, nil},
		{cons.MicroQrcode, model.VersionM4, 3, nil}}, nil},
	{cons.Fnc1ModeP1, []*modeIndicatorDetail{
		{cons.QRCODE, model.VersionAll, 4, []byte{0, 1, 0, 1}},
		{cons.MicroQrcode, model.VersionM1, 0, nil},
		{cons.MicroQrcode, model.VersionM2, 1, nil},
		{cons.MicroQrcode, model.VersionM3, 2, nil},
		{cons.MicroQrcode, model.VersionM4, 3, nil}}, nil},
	{cons.Fnc1ModeP2, []*modeIndicatorDetail{
		{cons.QRCODE, model.VersionAll, 4, []byte{1, 0, 0, 1}},
		{cons.MicroQrcode, model.VersionM1, 0, nil},
		{cons.MicroQrcode, model.VersionM2, 1, nil},
		{cons.MicroQrcode, model.VersionM3, 2, nil},
		{cons.MicroQrcode, model.VersionM4, 3, nil}}, nil},
}

func (m *AbstractMode) BuildEncodeData(qr *QRCodeStruct, buildDataBits func(qr *QRCodeStruct, dataStream *util.DataStream) (dataBitLen int)) (dataStream *util.DataStream) {
	dataStream = util.NewDataStream(16)
	// Mode indicator bits
	m.buildModeIndicator(qr, dataStream)
	// Data character count indicator bits
	_, _, numberOfDataBits := m.buildCharacterCountIndicator(qr, dataStream)
	// Data bits: implement by various modes
	buildDataBits(qr, dataStream)

	terminatorBitsLen := numberOfDataBits - dataStream.GetCount()
	// add Terminator
	m.buildTerminators(qr, dataStream, terminatorBitsLen)
	return dataStream
}

// GetModeIndicator :Get Mode Indicator by Iterate ModeIndicators with Mode and Version
func (m *AbstractMode) GetModeIndicator(qr *QRCodeStruct) *modeIndicatorDetail {
	mode := m.Name
	version := qr.Version.Id
	if !qr.Version.IsMicroQRCode() {
		// for VERSION_ALL
		version = model.VersionAll
	}
	for _, unionIndicator := range unionIndicators {
		if mode == unionIndicator.Mode {
			for _, detail := range unionIndicator.modeIndicatorDetails {
				if version == detail.Version {
					return detail
				}
			}
		}
	}
	return nil
}

// GetCharacterCountIndicatorBitsNumber :Get Character count indicator bits number.
// return: numberOfCountIndicatorBits.
// return: numberOfDataBits,  the max count of the data bits(mode indicator bits + character count indicator bits + data bits).
// return: modeDataCapacity,  the max count of the version data capacity.
func (m *AbstractMode) GetCharacterCountIndicatorBitsNumber(qr *QRCodeStruct) (numberOfCountIndicatorBits int, modeDataCapacity int, numberOfDataBits int) {
	mode := m.Name
	version := qr.Version
	versionId := version.Id
	ec := qr.ErrorCorrection
	for _, unionIndicator := range unionIndicators {
		if mode == unionIndicator.Mode {
			for _, detail := range unionIndicator.CountIndicatorBitsNumDetails {
				if util.IntIn(versionId, detail.Versions) {
					versionDataCapacity := version.GetVersionSymbolCharsAndInputDataCapacity(ec.Level)
					modeDataCapacity := versionDataCapacity.DataCapacity[mode]
					return detail.NumberOfBits, modeDataCapacity, versionDataCapacity.NumberOfDataBits
				}
			}
		}
	}
	return 0, 0, 0
}

// Get Mode Indicator bits
func (m *AbstractMode) buildModeIndicator(qr *QRCodeStruct, data *util.DataStream) (modeIndicatorBitLen int) {
	indicator := m.GetModeIndicator(qr)
	modeIndicatorBitLen = int(indicator.ModeIndicatorLength)
	//data := util.NewDataStream(int(length))
	data.AddBit(indicator.ModeIndicatorCode, modeIndicatorBitLen)
	//return data
	return modeIndicatorBitLen
}

// Get Data character count indicator bits
func (m *AbstractMode) buildCharacterCountIndicator(qr *QRCodeStruct, dataStream *util.DataStream) (countIndicatorBitLen, maxDataCapacity int, numberOfDataBits int) {
	numberOfBits, modeDataCapacity, numberOfDataBits := m.GetCharacterCountIndicatorBitsNumber(qr)
	dataLength := len(qr.Data)
	maxDataCapacity = modeDataCapacity
	if dataLength > modeDataCapacity {
		dataLength = modeDataCapacity
		qr.Data = qr.Data[:dataLength]
		logger.Warn("buildCharacterCountIndicator: Data length is too long, it will be trim to max count by rule \"Table3 - Number of bits in character count indicator for QR Code\", max count is :" + strconv.Itoa(modeDataCapacity))
	}
	//data = util.NewDataStream(numberOfBits)
	dataStream.AddIntBit(dataLength, numberOfBits)
	return numberOfBits, maxDataCapacity, numberOfDataBits
}

// buildTerminators :build Terminator bits by rules
func (m *AbstractMode) buildTerminators(qr *QRCodeStruct, dataStream *util.DataStream, fillBitsLen int) (terminatorLen int) {
	if fillBitsLen > 0 {
		terminatorBits := m.GetTerminalBits(qr.Version)
		terminatorLen = len(terminatorBits)
		dataStream.AddBit(terminatorBits, terminatorLen)
	}
	return terminatorLen
}

// BuildFinalErrorCorrectionCodewords : build the error correction codewords
func (m *AbstractMode) BuildFinalErrorCorrectionCodewords(qr *QRCodeStruct, dataStream *util.DataStream) []util.Bit {
	ecBlockList := m.ConvertDataBitsToECBlocks(qr, dataStream)
	return m.InterleaveECBlocks(qr, ecBlockList)
}

// calculatePixelSizePerModule :calculate pixel size for per module, default is cons.DefaultPixelSizePerModule.
// Generate as clear an image as possible, avoid resize.
// return: finalImageSize, the final output image size
// return: pixelSize, pixel size for per module.
// return: resize, need or not resize the image.
// return: quietZonePixels, quiet zone pixels for width/height
func calculatePixelSizePerModule(imageSize int, moduleSize int, quietZoneSize int) (pixelSize int) {
	// Calculate the max pixelSize for per module of QRCode.
	pixelSize = cons.DefaultPixelSizePerModule
	// image auto
	if imageSize == output.AUTO_SIZE {
		// finalImageSize = (moduleSize + quietZoneSize * 2) * pixelSize
		// image size is already init in Output.Init()
		return pixelSize
	}
	totalModuleSize := moduleSize + quietZoneSize
	pixelSize = imageSize / totalModuleSize
	// Image size can not full the qrcode by single-pixel size
	if pixelSize < 1 {
		panic(errors.New("image size:" + strconv.Itoa(imageSize) + " can not accommodate an effective qrcode"))
	}
	return pixelSize
}

// newQRCodeMaskOutputGroup : make a new output group for mask (QRCode by 0-7,Micro QRCode by 0-3).
// param: maskCount, the mask count number(QRCode: 8,Micro QRCode: 4).
// param: from, 0 by param "from",and 1-7 by from.Clone() take.
func newQRCodeMaskOutputGroup(maskCount int, from output.Output) []output.Output {
	outputs := make([]output.Output, maskCount)
	outputs[0] = from
	for i := 1; i < maskCount; i++ {
		outputs[i] = from.Clone()
	}
	return outputs
}

// selectRightPenaltyMaskOut :
// 1. QRCode: Evaluate mask penalty for mask (0-7 for QRCode) and select the pattern with the lowest penalty points score.
// 2. Micro QRCode: Evaluate mask penalty for mask (0-3 for Micro QRCode) and select the pattern with the highest score.
func selectRightPenaltyMaskOut(qr *QRCodeStruct, outs []output.Output, moduleSize int) (out output.Output, mask int) {
	if qr.IsMicroQRCode() {
		return selectHighestScoreMaskOut(outs, moduleSize)
	}
	return selectLowestPenaltyMaskOut(outs, moduleSize)
}

// selectLowestPenaltyMaskOut :For QRCode
func selectLowestPenaltyMaskOut(outs []output.Output, moduleSize int) (out output.Output, mask int) {
	penalty := ^uint(0)
	var penaltyOut output.Output = nil
	for i, out_ := range outs {
		var tempPenalty uint
		tempPenalty = out_.GetBaseOutput().EvalPenalty(moduleSize)
		if tempPenalty < penalty {
			penalty = tempPenalty
			penaltyOut = out_
			mask = i
		}
	}
	return penaltyOut, mask
}

// selectHighestScoreMaskOut :For Micro QRCode
func selectHighestScoreMaskOut(outs []output.Output, moduleSize int) (out output.Output, mask int) {
	penalty := uint(0)
	var penaltyOut output.Output = nil
	for i, out_ := range outs {
		var tempPenalty uint
		tempPenalty = out_.GetBaseOutput().EvalMicroQRCodePenalty(moduleSize)
		if tempPenalty > penalty {
			penalty = tempPenalty
			penaltyOut = out_
			mask = i
		}
	}
	return penaltyOut, mask
}

// BuildModuleInMatrix : Page 54,7.7 Codeword placement in matrix
// Place the codeword modules in the matrix together with the finder pattern,separators,timing pattern,and (if required) alignment patterns.
func (m *AbstractMode) BuildModuleInMatrix(qr *QRCodeStruct, codewordsBits []util.Bit, out output.Output) output.Output {
	version := qr.Version
	imageSize := out.GetBaseOutput().Size
	moduleSize := version.GetModuleSize()
	quietZoneSize := qr.QuietZone.GetQuietZoneSize()
	pixelSize := calculatePixelSizePerModule(imageSize, moduleSize, quietZoneSize)
	maskOutputGroup := newQRCodeMaskOutputGroup(qr.GetMaskCount(), out)

	// output group with mask for common
	outputGroupOutMask := func(x int, y int, val bool, hasMask bool, part cons.QRCodeStructPart) {
		srcVal := val
		for i, out_ := range maskOutputGroup {
			newVal := srcVal
			if hasMask {
				if version.Id > 0 {
					newVal = getQRCodeMaskVal(x, y, srcVal, i)
				} else {
					newVal = getMircoQRCodeMaskVal(x, y, srcVal, i)
				}
			}
			out_.WriteModule(x, y, newVal, pixelSize, part)
		}
	}
	// output group no mask for common
	outputGroupOut := func(x int, y int, val bool, part cons.QRCodeStructPart) {
		outputGroupOutMask(x, y, val, false, part)
	}
	// output group with mask for format info
	outputGroupFormatMaskOut := func(x int, y int, maskBits map[int][]util.Bit, maskBitIdx int, part cons.QRCodeStructPart) {
		for i, out_ := range maskOutputGroup {
			out_.WriteModule(x, y, maskBits[i][maskBitIdx] == 1, pixelSize, part)
		}
	}

	// draw finder pattern and separators
	drawFinderPatternAndSeparator(version, outputGroupOut)
	// draw alignment pattern,version 2 or larger must contain alignment pattern.
	drawAlignmentPattern(qr.Version, qr.AlignmentPattern, outputGroupOut)
	// draw timing pattern
	drawTimingPattern(qr.TimingPattern, moduleSize, outputGroupOut)
	// draw Dark Block for QRCode
	drawDarkBlock(version.Id, outputGroupOut)
	// draw version information for version 7+
	drawVersionInformation(qr, outputGroupOut)
	// format zone axes for QRCode module2
	drawFormatInformation(qr, outputGroupFormatMaskOut)
	// draw data
	drawData(version, moduleSize, codewordsBits, out, outputGroupOutMask)

	// evaluate mask penalty for mask 0-7 or 0-3
	lowestPenaltyOut, mask := selectRightPenaltyMaskOut(qr, maskOutputGroup, moduleSize)
	qr.SetMask(mask)
	if out != lowestPenaltyOut {
		out = lowestPenaltyOut
	}
	// draw with quiet zone
	out.ResizeToFit(moduleSize, quietZoneSize, pixelSize)
	return out
}

// BuildModuleInMatrixMicroQRCode :
// Deprecated
func (m *AbstractMode) BuildModuleInMatrixMicroQRCode(qr *QRCodeStruct, codewordsBits []util.Bit, out output.Output) output.Output {
	version := qr.Version
	imageSize := out.GetBaseOutput().Size
	moduleSize := version.GetModuleSize()
	quietZoneSize := qr.QuietZone.GetQuietZoneSize()
	pixelSize := calculatePixelSizePerModule(imageSize, moduleSize, quietZoneSize)
	maskOutputGroup := newQRCodeMaskOutputGroup(qr.GetMaskCount(), out)

	// output group with mask for common
	outputGroupOutMask := func(x int, y int, val bool, hasMask bool, part cons.QRCodeStructPart) {
		srcVal := val
		for i, out_ := range maskOutputGroup {
			newVal := srcVal
			if hasMask {
				if version.Id > 0 {
					newVal = getQRCodeMaskVal(x, y, srcVal, i)
				} else {
					newVal = getMircoQRCodeMaskVal(x, y, srcVal, i)
				}
			}
			out_.WriteModule(x, y, newVal, pixelSize, part)
		}
	}
	// output group no mask for common
	outputGroupOut := func(x int, y int, val bool, part cons.QRCodeStructPart) {
		outputGroupOutMask(x, y, val, false, part)
	}
	// output group with mask for format info
	outputGroupFormatMaskOut := func(x int, y int, maskBits map[int][]util.Bit, maskBitIdx int, part cons.QRCodeStructPart) {
		for i, out_ := range maskOutputGroup {
			out_.WriteModule(x, y, maskBits[i][maskBitIdx] == 1, pixelSize, part)
		}
	}

	// draw finder pattern and separators
	drawFinderPatternAndSeparator(version, outputGroupOut)
	// draw timing pattern
	drawTimingPattern(qr.TimingPattern, moduleSize, outputGroupOut)
	// format zone axes for QRCode module2
	drawFormatInformation(qr, outputGroupFormatMaskOut)
	// draw data
	drawData(version, moduleSize, codewordsBits, out, outputGroupOutMask)

	lowestPenaltyOut, mask := selectRightPenaltyMaskOut(qr, maskOutputGroup, moduleSize)
	qr.SetMask(mask)
	if out != lowestPenaltyOut {
		out = lowestPenaltyOut
	}
	// draw with quiet zone
	out.ResizeToFit(moduleSize, quietZoneSize, pixelSize)
	return out
}

func drawDarkBlock(version model.VersionId, outputGroupOut func(x int, y int, val bool, part cons.QRCodeStructPart)) {
	// Dark Block: All QR codes have a dark module beside the bottom left finder pattern. More specifically,
	// the dark module is always located at the coordinate ([(4 * V) + 9], 8) where V is the version of the QR code.
	if version > 0 {
		//out.WriteModule(model.FINDER_PATTERN_MODULE_SIZE + 1, 4 * version + 9,true,pixelSize)
		outputGroupOut(model.FINDER_PATTERN_MODULE_SIZE+1, 4*version+9, true, cons.FinderPatternPart)
	}
}

// drawFinderPatternAndSeparator :draw finder pattern and separators
// Reserved Areas:
// A strip of modules beside the separators must be reserved for the format information area as follows:
//  Near the top-left finder pattern, a one-module strip must be reserved below and to the right of the separator.
//	Near the top-right finder pattern, a one-module strip must be reserved below the separator.
//	Near the bottom-left finder pattern, a one-module strip must be reserved to the right of the separator.
func drawFinderPatternAndSeparator(version *model.Version, outputGroupOut func(x int, y int, val bool, part cons.QRCodeStructPart)) {
	finderPatterns := version.GetFinderPattern()
	fpModules := finderPatterns.GetModules()
	separatorLen := model.FINDER_PATTERN_MODULE_SIZE + 1

	for _, pos := range finderPatterns.Positions {
		for row, bits := range fpModules {
			for col, bit := range bits {
				//out.WriteModule(pos.Axes.X+row, pos.Axes.Y + col,bit == 1,pixelSize)
				outputGroupOut(pos.Axes.X+row, pos.Axes.Y+col, bit == 1, cons.FinderPatternPart)
			}
		}
		// draw separators
		for lop := 0; lop < separatorLen; lop++ {
			separatorHX := pos.Axes.X
			separatorHY := pos.Axes.Y + model.FINDER_PATTERN_MODULE_SIZE
			separatorVX := pos.Axes.X + model.FINDER_PATTERN_MODULE_SIZE
			separatorVY := pos.Axes.Y
			if pos.Position == model.TOP_RIGHT {
				separatorHX = pos.Axes.X - 1
				separatorVX = pos.Axes.X - 1
			}
			if pos.Position == model.BOTTOM_LEFT {
				separatorHY = pos.Axes.Y - 1
				// TOP_LEFT by (0,0) is ok, need handle TOP_RIGHT,BOTTOM_LEFT draw
				separatorVY = pos.Axes.Y - 1
			}
			// draw horizontal separators
			//out.WriteModule(separatorHX + lop,separatorHY ,false,pixelSize)
			outputGroupOut(separatorHX+lop, separatorHY, false, cons.FinderPatternPart)
			// draw vertical separators
			//out.WriteModule(separatorVX ,separatorVY + lop,false,pixelSize)
			outputGroupOut(separatorVX, separatorVY+lop, false, cons.FinderPatternPart)
		}
	}
}

// drawAlignmentPattern :draw alignment pattern,version 2 or larger must contain alignment pattern.
func drawAlignmentPattern(version *model.Version, pattern *model.AlignmentPattern, outputGroupOut func(x int, y int, val bool, part cons.QRCodeStructPart)) {
	if version.Id < 2 {
		return
	}
	apModules := pattern.GetModules()
	for _, pos := range pattern.Positions {
		for row, bits := range apModules {
			for col, bit := range bits {
				// pos is alignment center point axes,need transform to left_top point axes.
				//out.WriteModule(quietZoneSize + pos.X - 2 + row,quietZoneSize + pos.Y - 2+col,bit == 1,pixelSize)
				//out.WriteModule(pos.X - 2 + row, pos.Y - 2 + col,bit == 1,pixelSize)
				outputGroupOut(pos.X-2+row, pos.Y-2+col, bit == 1, cons.AlignmentPart)
			}
		}
	}
}

func drawTimingPattern(timingPattern *model.TimingPattern, moduleSize int, outputGroupOut func(x int, y int, val bool, part cons.QRCodeStructPart)) {
	//startPos := model.FINDER_PATTERN_MODULE_SIZE
	// max loop count is moduleSize - 16
	//for i:=0; i< moduleSize - model.FINDER_PATTERN_MODULE_SIZE * 2 - 2; i++ {
	//	// draw horizontal timing pattern
	//	//out.WriteModule(startPos + 1 + i , startPos - 1, i%2 == 0,pixelSize)
	//	outputGroupOut(startPos + 1 + i , startPos - 1, i%2 == 0)
	//	// draw vertical timing pattern
	//	//out.WriteModule(startPos - 1,startPos + 1 + i, i%2 == 0,pixelSize)
	//	outputGroupOut(startPos - 1,startPos + 1 + i, i%2 == 0)
	//}
	h := timingPattern.H
	w := timingPattern.W
	loopLen := h.To.X - h.From.X
	for i := 0; i < loopLen; i++ {
		// draw horizontal timing pattern
		outputGroupOut(h.From.X+i, h.From.Y, i%2 == 0, cons.TimingPatternPart)
		// draw vertical timing pattern
		outputGroupOut(w.From.X, w.From.Y+i, i%2 == 0, cons.TimingPatternPart)
	}
}

// drawFormatInformation :
// Page 63,7.9 Format information.
// The format information is a 15-bit sequence containing 5 data bits,with 10 error correction bits calculated using the (15,5) BHC code.
func drawFormatInformation(qr *QRCodeStruct, outputGroupMaskOut func(x int, y int, maskBits map[int][]util.Bit, maskBitIdx int, part cons.QRCodeStructPart)) {
	moduleSize := qr.Version.GetModuleSize()
	version := qr.Version.Id
	level := qr.ErrorCorrection.Level
	if qr.IsMicroQRCode() {
		drawMicroQRCodeFormatInformation(version, level, moduleSize, outputGroupMaskOut)
	} else {
		drawQRCodeFormatInformation(level, moduleSize, outputGroupMaskOut)
	}
}

// Draw QRCode format information
func drawQRCodeFormatInformation(level int, moduleSize int, outputGroupMaskOut func(x int, y int, maskBits map[int][]util.Bit, maskBitIdx int, part cons.QRCodeStructPart)) {
	formatInfoBits := cons.FormatInformationBitsMap[level]
	for i := 0; i < 15; i++ {
		// 0 - 6
		if i < 7 {
			x := i
			if i >= 6 {
				x++
			}
			//out.WriteModule(x, 8 , formatInfoBits[i] == 1 ,pixelSize)
			outputGroupMaskOut(x, 8, formatInfoBits, i, cons.FormatPart)
			if i <= 6 {
				//out.WriteModule(8, moduleSize - i - 1, formatInfoBits[i] == 1 ,pixelSize)
				outputGroupMaskOut(8, moduleSize-i-1, formatInfoBits, i, cons.FormatPart)
			}
		} else {
			// 7-14
			//out.WriteModule(moduleSize - (15 - i), 8 , formatInfoBits[i] == 1 ,pixelSize)
			outputGroupMaskOut(moduleSize-(15-i), 8, formatInfoBits, i, cons.FormatPart)
			y := 15 - i
			if y <= 6 {
				y--
			}
			//out.WriteModule(8, y , formatInfoBits[i] == 1 ,pixelSize)
			outputGroupMaskOut(8, y, formatInfoBits, i, cons.FormatPart)
		}
	}
}

// Draw Micro QRCode format information
// Page 65, 7.9.2 Micro QRCode Symbol.
// The format information is a 15-bit sequence contaning 5 data bits,with 10 error correction bits calculated using the
//  (15,5) BCH code. The first three data bits contain the symbol number(in binary). which identifies the version
//   and error correction level, as shown in Table 13.
func drawMicroQRCodeFormatInformation(version model.VersionId, level int, moduleSize int, outputGroupMaskOut func(x int, y int, maskBits map[int][]util.Bit, maskBitIdx int, part cons.QRCodeStructPart)) {
	var formatInfoBits = cons.MicroQRCodeFormatInformationBitsMap[version][level]
	for i := 0; i < 15; i++ {
		if i < 8 {
			// h: 14 - 7
			outputGroupMaskOut(1+i, model.FINDER_PATTERN_MODULE_SIZE+1, formatInfoBits, i, cons.FormatPart)
		} else {
			// w:  0 - 6
			outputGroupMaskOut(model.FINDER_PATTERN_MODULE_SIZE+1, 15-i, formatInfoBits, i, cons.FormatPart)
		}
	}
}

// version zone axes:
// Page 66,7.10 Version information: It consists of an 18-bit sequence containing 6 data bits,with 12 error correction bits calculated using the (18,6) Golay code.
//
// QR codes versions 7 and larger must contain two areas where version information bits are placed. The areas are a 6x3 block above the bottom-left finder pattern and a 3x6 block to the left of the top-right finder pattern. The following images show the locations of the reserved areas in blue.
// Two position version info block:
// 1. Bottom Left Version Information Block:
// arrangement:
// 0 3 6 9  12 15   row = 0 , col = 0 ,1 ,2 ,3
// 1 4 7 10 13 16
// 2 5 8 11 14 17
// 2. Top Right Version Information Block
// arrangement:
// 0  1  2			col = 0 row = 0
// 3  4  5			col = 0 row = 1
// 6  7  8			col = 0 row = 2
// 9  10 11			...
// 12 13 14
// 15 16 17
func drawVersionInformation(qr *QRCodeStruct, outputGroupOut func(x int, y int, val bool, part cons.QRCodeStructPart)) {
	version := qr.Version.Id
	if version < 7 {
		return
	}
	moduleSize := qr.Version.GetModuleSize()
	versionInfoBits := cons.VersionInformationBitsMap[version]
	for row := 0; row < 3; row++ {
		for col := 0; col < 6; col++ {
			// bottom_left
			//out.WriteModule(col, moduleSize - 8 - 3 + row , versionInfoBits[ col * 3 + row] == 1 ,pixelSize)
			outputGroupOut(col, moduleSize-8-3+row, versionInfoBits[cons.VersionInformationBitsLen-1-(col*3+row)] == 1, cons.VersionPart)
			// top_right
			//out.WriteModule(moduleSize - 8 - 3 + row, col , versionInfoBits[ col * 3 + row] == 1 ,pixelSize)
			outputGroupOut(moduleSize-8-3+row, col, versionInfoBits[cons.VersionInformationBitsLen-1-(col*3+row)] == 1, cons.VersionPart)
		}
	}
}

// drawData : draw data and set mask.
// Mask rule in Page 58,7.8 Data Masking.
func drawData(version *model.Version, moduleSize int, codewordsBits []util.Bit, out output.Output, outputGroupOutMask func(x int, y int, val bool, hasMask bool, part cons.QRCodeStructPart)) {
	var moduleBitIdx int
	for pos := range iterateModulesPlacement(version, moduleSize, out.IsModuleSet) {
		var bit bool
		if moduleBitIdx < len(codewordsBits) {
			bit = codewordsBits[moduleBitIdx] == 1
		}
		//out.WriteModule(pos.X ,pos.Y ,bit,pixelSize)
		outputGroupOutMask(pos.X, pos.Y, bit, true, cons.DataPart)
		moduleBitIdx++
	}
}

// iterateModules : get empty position by upward/downward sort.
// return: <-*model.PositionAxes, empty position
// Module set direction from right-bottom corner for two way:
// upward:
// D6 D5
// D4 D3
// D2 D1
//
// bit set in Module by upward(the most significant bit,shown as 7),e.g.: D1
// 0 1
// 2 3
// 4 5
// 6 7
//
// downward:
// D8 D7
// D10 D9
// D12 D11
// D14 D13
// bit set in Module by downward(the most significant bit,shown as 7), e.g.: D7
// 6 7
// 4 5
// 2 3
// 0 1
func iterateModulesPlacement(version *model.Version, moduleSize int, isModuleSet func(x int, y int) bool) <-chan *model.PositionAxes {
	allModuleBitPos := make(chan *model.PositionAxes)
	//remainBitLen := 6
	//if version.IsMicroQRM1M3Code(){
	//	remainBitLen = 4
	//}
	go func() {
		isUpward := true
		x := moduleSize - 1
		y := moduleSize - 1
		for x >= 0 {
			allModuleBitPos <- &model.PositionAxes{X: x, Y: y}
			if x > 0 {
				allModuleBitPos <- &model.PositionAxes{X: x - 1, Y: y}
			}
			if isUpward {
				y--
				// turn to next columns group (2 columns)
				if y < 0 {
					y = 0
					x -= 2
					//if x == remainBitLen{
					//	x--
					//}
					isUpward = false
				}
			} else {
				y++
				if y >= moduleSize {
					y = moduleSize - 1
					x -= 2
					//if x == remainBitLen{
					//	x --
					//}
					isUpward = true
				}
			}

		}
		close(allModuleBitPos)
	}()
	moduleBitPos := make(chan *model.PositionAxes)
	go func() {
		for mb := range allModuleBitPos {
			// check the module whether or not be set
			if !isModuleSet(mb.X, mb.Y) {
				moduleBitPos <- mb
			}
		}
		close(moduleBitPos)
	}()
	return moduleBitPos
}

// getQRCodeMaskVal :A mask pattern changes which modules are dark and which are light according to a particular rule. The purpose of this step is to modify the QR code to make it as easy for a QR code reader to scan as possible.
// Page 58,7.8.2 Data mask patterns.
// Have 8 types in mask,value in( 0 ... 7).
func getQRCodeMaskVal(x, y int, val bool, mask int) bool {
	switch mask {
	case 0:
		val = val != (((y + x) % 2) == 0)
		break
	case 1:
		val = val != ((y % 2) == 0)
		break
	case 2:
		val = val != ((x % 3) == 0)
		break
	case 3:
		val = val != (((y + x) % 3) == 0)
		break
	case 4:
		val = val != (((y/2 + x/3) % 2) == 0)
		break
	case 5:
		val = val != (((y*x)%2)+((y*x)%3) == 0)
		break
	case 6:
		val = val != ((((y*x)%2)+((y*x)%3))%2 == 0)
		break
	case 7:
		val = val != ((((y+x)%2)+((y*x)%3))%2 == 0)
	}
	return val
}

func getMircoQRCodeMaskVal(x, y int, val bool, mask int) bool {
	switch mask {
	case 0:
		val = val != ((y % 2) == 0)
		break
	case 1:
		val = val != (((y/2 + x/3) % 2) == 0)
		break
	case 2:
		val = val != ((((y*x)%2)+((y*x)%3))%2 == 0)
		break
	case 3:
		val = val != ((((y+x)%2)+((y*x)%3))%2 == 0)
		break
	}
	return val
}

// GetMode : Get first supported mode in SupportModes
// param: data
func GetMode(data string) (Mode, error) {
	for _, mode := range SupportModes {
		if mode.IsSupport(data) {
			return mode, nil
		}
	}
	return nil, errors.New("please check the input data,can not find a valid Mode for data:" + data)
}

// GetSupportedModes : Get all supported modes for input data
// param: data
func GetSupportedModes(data string) (modes []Mode, err error) {
	modes = make([]Mode, 0)
	for _, mode := range SupportModes {
		if mode.IsSupport(data) {
			modes = append(modes, mode)
		}
	}
	if len(modes) < 1 {
		return nil, errors.New("please check the input data,can not find a valid Mode for data:" + data)
	}
	return modes, err
}
