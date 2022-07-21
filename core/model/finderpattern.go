package model

import "github.com/KangSpace/gqrcode/util"

// Define Finder Pattern here

// FinderPattern :
// QRCode:
// There are 3 identical Finder pattern locate at the upper left,upper right and low left corner of the symbol.
//      Each finder pattern may be viewed as 3 superimposed concentric squares:
//        dark 7 x 7 modules
//        light 5 x 5 modules
//        dark 3 x 3 modules.
//      The ratio of module widths in each finder pattern is 1:1:3:1:1.
// Micro QRCode:
// A single finder pattern, as defined [QRCode]
//

// FinderPatternPosition :
type FinderPatternPosition struct {
	// value in (TOP_LEFT, TOP_RIGHT,BOTTOM_LEFT)
	Position int
	Axes     *PositionAxes
}
type PositionAxes struct {
	X int
	Y int
}

type PositionAxesRange struct {
	From *PositionAxes
	To   *PositionAxes
}

type Position = int

// FinderPattern Position constants
const (
	TOP_LEFT    Position = iota // 0
	TOP_RIGHT                   // 1
	BOTTOM_LEFT                 // 2
)
const (
	// QRCODE_BASE_MODULE_SIZE : Total forty sizes of QR Code symbol,
	// referred to as Version 1,Version 2...Version 40.
	// Version 1 measures 21 x 21 modules, Version 2 measures 25 x 25 modules and so on
	// in creasing in steps of 4 modules per side up to Version 40 which measures 177 x 177 modules.
	QRCODE_BASE_MODULE_SIZE = 21

	// MICRO_QRCODE_BASE_MODULE_SIZE : Total four sizes of Micro QR Code symbol,
	// referred to as Versions M1 to M4.
	// Version M1 measures 11 x 11 modules, Version M2 measures 13 x 13 modules, Version M3 measures 15 x 15 modules,
	// Version M4 measures 17 x 17 modules, increasing in steps of 2 modules per side.
	MICRO_QRCODE_BASE_MODULE_SIZE = 11

	FINDER_PATTERN_MODULE_SIZE = 7
)

type FinderPattern struct {
	Positions []FinderPatternPosition
}

func NewFinderPattern(version VersionId) *FinderPattern {
	finderPattern := new(FinderPattern)
	// QR Code
	if version > 0 {
		finderPattern.Positions = []FinderPatternPosition{
			{TOP_LEFT, GetPositionAxes(TOP_LEFT, version)},
			{TOP_RIGHT, GetPositionAxes(TOP_RIGHT, version)},
			{BOTTOM_LEFT, GetPositionAxes(BOTTOM_LEFT, version)},
		}
	} else {
		//Micro QR Code
		finderPattern.Positions = []FinderPatternPosition{
			{TOP_LEFT, GetPositionAxes(TOP_LEFT, version)},
		}
	}
	return finderPattern
}

// GetPositionAxes :
// The size of a QR code can be calculated with the formula (((V-1)*4)+21), where V is the QR code version. For example, version 32 is (((32-1)*4)+21) or 145 modules by 145 modules. Therefore, the positions of the finder patterns can be generalized as follows:
// The top-left finder pattern's top left corner is always placed at (0,0).
// The top-right finder pattern's top LEFT corner is always placed at ([(((V-1)*4)+21) - 7], 0)
// The bottom-left finder pattern's top LEFT corner is always placed at (0,[(((V-1)*4)+21) - 7])
func GetPositionAxes(p Position, version VersionId) *PositionAxes {
	axes := new(PositionAxes)
	switch p {
	case TOP_LEFT:
		axes.X = 0
		axes.Y = 0
	case TOP_RIGHT:
		axes.X = (version-1)*4 + QRCODE_BASE_MODULE_SIZE - FINDER_PATTERN_MODULE_SIZE
		axes.Y = 0
	case BOTTOM_LEFT:
		axes.X = 0
		axes.Y = (version-1)*4 + QRCODE_BASE_MODULE_SIZE - FINDER_PATTERN_MODULE_SIZE
	}
	return axes
}

// GetModules : 1：1:3:1:1 ,
func (fp *FinderPattern) GetModules() [][]util.Module {
	modules := [][]util.Module{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}
	return modules[:]
}
