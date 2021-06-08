package output

import (
	"github.com/gqrcode/core/model"
	"image"
	"image/color"
	"math"
)

// Define Output here

type Type = int
type OptionName = string

// Option : Output Option
type Option struct {
	Name OptionName
	Value string
}

// defined OutputType constants
const (
	JPG Type= iota // 0
	PNG // 1
	GIF // 2
	SVG // 3

	// AUTO_SIZE : set auto size for qrcode , per module fill by 4 pixels
	AUTO_SIZE = 0

	LogoOptionName OptionName = "logo"
)

type BaseOutput struct {
	// image type
	Type Type
	// image width/height
	Size int
	Options []*Option
	modules [][]*bool
}


// LogoOption :  Option for add logo image at center of QRCode
func LogoOption(logoImage string) *Option{
	return &Option{Name: LogoOptionName,Value: logoImage}
}

// containLogoOption : Check whether contain LogoOption option or not
func (out *BaseOutput) containLogoOption() *Option{
	for _,opt:=range out.Options{
		if opt.Name == LogoOptionName{
			return opt
		}
	}
	return nil
}

func (out *BaseOutput) AddOption(options... *Option){
	out.Options = append(out.Options,options...)
}

func (out *BaseOutput) IsModuleSet(x int,y int) bool{
	return out.modules[x][y] != nil
}

func (out *BaseOutput) GetModule(x int,y int) bool{
	return out.modules[x][y] != nil && *out.modules[x][y]
}

// Output : the output interface for qrcode print
type Output interface {
	GetBaseOutput() *BaseOutput
	// Init :init for output when size is AUTO_SIZE
	Init(version *model.Version,qz *model.QuietZone)
	Write(x int,y int, black bool)
	// WriteModule :write per module by pixelSize
	WriteModule(x int,y int, black bool,pixelSize int)
	WriteModuleColor(x int,y int, dark bool, setColor color.Color,pixelSize int)
	// IsModuleSet : check the module whether or not be set
	IsModuleSet(x int,y int) bool
	// GetModule : x,y is module axes , not pixel axes.
	GetModule(x int,y int) bool
	GetImage() *image.NRGBA
	Clone() Output
	ResizeToFit(moduleSize int,quietZoneSize int,pixelSize int)
	Save(fileName string) error
	SaveToBase64() (string,error)
}


// EvalPenalty :Evaluate Penalty for QRCode.
// param: moduleSize, not contains quiet zone size.
func (out *BaseOutput) EvalPenalty(moduleSize int) uint {
	return out.evalPenaltyRule1(moduleSize) +
		out.evalPenaltyRule2(moduleSize) +
		out.evalPenaltyRule3(moduleSize) +
		out.evalPenaltyRule4(moduleSize)
}

// EvalMicroQRCodePenalty :Evaluate Penalty for Micro QRCode.
// Page 62, 7.8.3.2 Evaluation of Micro QRCode Symbols.
// if SUM1 <= SUM2
// 	Evaluation Score = SUM1 x 16 + SUM2
// if SUM1 > SUM2
// 	Evaluation Score = SUM2 x 16 + SUM2
// where:
//  SUM1 number of dark modules in right side edge
//  SUM2 number of dark modules in lower side edge
func (out *BaseOutput) EvalMicroQRCodePenalty(moduleSize int) uint {
	var sum1,sum2,resultPoint uint
	for i:=0;i<moduleSize;i++ {
		if out.GetModule(moduleSize - 1, i){
			sum1++
		}
		if out.GetModule(i, moduleSize - 1){
			sum2++
		}
	}
	if sum1 <= sum2 {
		resultPoint = sum1 * 16 + sum2
	}else{
		resultPoint = sum2 * 16 + sum1
	}
	return resultPoint
}

// evalPenaltyRule1 :
// (N1 = 3,N2 = 3,N3 = 40,N4 = 10)
// Feature											Evaluation condition	 Points
// Adjacent modules in row/column in same color ,	No.of modules = (5 + i), N1 + i
//
// NOTE1:
// Check the blocks consisting of light(white) or dark(blank) modules of more then five in a row both laterally and
//   vertically for the evaluation of data masking results. The rule of this calculation is that 3 penalty points shall
//	 be added to each block of five consecutive modules, 4 penalty points for each block of six consecutive modules and so on,
//   with scoring by 1 point each time the number of modules increases. For example, impose 5 penalty points on the block of
//	 "dark:dark:dark:dark:dark:dark:dark" module pattern, where a series of seven consecutive modules is counted as on block.
//   However, do not double-count the point. The penalty point for a seven-module block, for example, shall be 5 not the sum of
//   3(for a five-module block) + 4(for a six-module block) + 5(for a seven-module block) = 12
func (out *BaseOutput) evalPenaltyRule1(moduleSize int) uint {
	var resultPoint uint
	for row := 0; row < moduleSize; row++ {
		rowCheck := false
		colCheck := false
		var rowCheckCnt uint = 0
		var colCheckCnt uint = 0
		for col := 0; col < moduleSize; col++ {
			currModule := out.GetModule(row,col)
			if currModule == rowCheck{
				rowCheckCnt++
			}else{
				rowCheck = !rowCheck
				rowCheckCnt = 1
				if rowCheckCnt >= 5{
					resultPoint += 3 + (rowCheckCnt - 5)
				}
			}
			if currModule == colCheck{
				colCheckCnt++
			}else{
				colCheck = !colCheck
				colCheckCnt = 1
				if colCheckCnt >= 5{
					resultPoint += 3 + (colCheckCnt - 5)
				}
			}
		}
		if rowCheckCnt >= 5{
			resultPoint += 3 + (rowCheckCnt - 5)
		}
		if colCheckCnt >= 5{
			resultPoint += 3 + (colCheckCnt - 5)
		}
	}
	//fmt.Printf("%p evalPenaltyRule1 %d\n:",out,resultPoint)
	return resultPoint
}


// evalPenaltyRule2 :
// (N1 = 3,N2 = 3,N3 = 40,N4 = 10)
// Feature											Evaluation condition	 Points
// Module blocks in the same colour ,			Block size = 2 x 2			  N2
//
// The penalty point shall be equal to the number of blocks with 2 x 2 light or dark modules. Take a block consisting of
//	 3x3 dark modules for an example. Considering that up to four 2 x 2 dark modules can be included in this block, the penalty
// 	 applied to this block shall be calculated as 4(blocks) x 3(points) = 12 points.
func (out *BaseOutput) evalPenaltyRule2(moduleSize int) uint {
	var result uint
	for x := 0; x < moduleSize-1; x++ {
		for y := 0; y < moduleSize-1; y++ {
			val := out.GetModule(x, y)
			if out.GetModule(x, y+1) == val && out.GetModule(x+1, y) == val && out.GetModule(x+1, y+1) == val {
				result += 3
			}
		}
	}
	//fmt.Printf("%p evalPenaltyRule2 %d\n:",out,result)
	return result
}

// evalPenaltyRule3 :
// (N1 = 3,N2 = 3,N3 = 40,N4 = 10)
// Feature											Evaluation condition	 Points
// 1:1:3:1:1 ratio(dark:light:dark:light:dark) 		Existence of the pattern	N3
// pattern in row/column,preceded or followed by
// light are 4 modules wide
//
// If the light area of more than 4 module wide exists after or before a 1:1:3:1:1 ratio(dark:light:dark:light:dark) pattern,
// the imposed penalty shall be 40 points.
func (out *BaseOutput) evalPenaltyRule3(moduleSize int) uint {
	leftDarkPattern := []bool{true, false, true, true, true, false, true, false, false, false, false}
	rightDarkPattern := []bool{false, false, false, false, true, false, true, true, true, false, true}
	patternLen := len(leftDarkPattern)
	var result uint
	for x := 0; x <= moduleSize - patternLen; x++ {
		for y := 0; y < moduleSize; y++ {
			leftDarkPatternXFound := true
			rightDarkPatternXFound := true
			leftDarkPatternYFound := true
			rightDarkPatternYFound := true

			for i := 0; i < patternLen; i++ {
				iv := out.GetModule(x+i, y)
				if iv != leftDarkPattern[i] {
					leftDarkPatternXFound = false
				}
				if iv != rightDarkPattern[i] {
					rightDarkPatternXFound = false
				}
				iv = out.GetModule(y, x+i)
				if iv != leftDarkPattern[i] {
					leftDarkPatternYFound = false
				}
				if iv != rightDarkPattern[i] {
					rightDarkPatternYFound = false
				}
			}
			if leftDarkPatternXFound || rightDarkPatternXFound ||
				leftDarkPatternYFound || rightDarkPatternYFound{
				result += 40
			}
		}
	}
	//fmt.Printf("%p evalPenaltyRule2 %d\n:",out,result)
	return result
}

// evalPenaltyRule4 :
// (N1 = 3,N2 = 3,N3 = 40,N4 = 10)
// Feature											Evaluation condition	 Points
// Proportion of dark modules in entire symbol, 	50x(5+k)% to 50x(tx(k+1))%	N4 x k
//
// Add 10 points to deviation of 5% increment or decrement in the proportion ratio of dark module from the referential
// 50%or 0 point) level. For example, assign 0 points as a penalty if the ratio of dark module is between 45%
// and 55%, or 1- points if the ratio of dark module is between 40% and 60%.
func (out *BaseOutput) evalPenaltyRule4(moduleSize int) uint {
	darkCount := 0
	for  row:= 0; row < moduleSize; row++ {
		for col := 0; col < moduleSize; col++ {
			if out.GetModule(row,col) {
				darkCount++
			}
		}
	}
	darkRate := float64(darkCount) / float64(moduleSize * moduleSize)  * 100
	floor := math.Abs(math.Floor(darkRate/5) - 10)
	ceil := math.Abs(math.Ceil(darkRate/5) - 10)
	//fmt.Printf("%p evalPenaltyRule4 floor:%f ceil:%f \n:",out,floor,ceil)
	return uint(math.Min(floor, ceil) * 10)
}

// GetRecommendSize :Get recommend size for QRCode
// return: the array of two recommend sizes
func (out *BaseOutput) GetRecommendSize(moduleSize int) []int{
	pixelSizePerModule := out.Size / moduleSize
	recSize1 := moduleSize * pixelSizePerModule
	recSize2 := moduleSize * (pixelSizePerModule + 1)
	return []int{recSize1,recSize2}
}