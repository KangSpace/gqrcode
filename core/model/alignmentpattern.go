package model

import "github.com/gqrcode/util"

// Define Alignment Pattern here

// AlignmentPattern :
// QR codes that are version 2 and larger are required to have alignment patterns.
// An alignment pattern:
// 	consists of a 5 x 5 modules black square,
// 	an inner 3 module by 3 module white square,
// 	and a single black module in the center.
type AlignmentPattern struct {
	// Center modules axes.
	Positions []*PositionAxes
}

// VersionAlignmentPatternLocationsMap :
var VersionAlignmentPatternLocationsMap = map[VersionId][]int{
	VERSION2: {6, 18},	VERSION3:  {6, 22},	VERSION4:  {6, 26},	VERSION5:  {6, 30},	VERSION6:  {6, 34},	VERSION7:  {6, 22, 38},	VERSION8:  {6, 24, 42},	VERSION9:  {6, 26, 46},	VERSION10: {6, 28, 50},	VERSION11: {6, 30, 54},	VERSION12: {6, 32, 58},	VERSION13: {6, 34, 62},	VERSION14: {6, 26, 46, 66},	VERSION15: {6, 26, 48, 70},	VERSION16: {6, 26, 50, 74},	VERSION17: {6, 30, 54, 78},	VERSION18: {6, 30, 56, 82},	VERSION19: {6, 30, 58, 86},	VERSION20: {6, 34, 62, 90},	VERSION21: {6, 28, 50, 72, 94},	VERSION22: {6, 26, 50, 74, 98},	VERSION23: {6, 30, 54, 78, 102},	VERSION24: {6, 28, 54, 80, 106},	VERSION25: {6, 32, 58, 84, 110},	VERSION26: {6, 30, 58, 86, 114},	VERSION27: {6, 34, 62, 90, 118},	VERSION28: {6, 26, 50, 74, 98, 122},	VERSION29: {6, 30, 54, 78, 102, 126},	VERSION30: {6, 26, 52, 78, 104, 130},	VERSION31: {6, 30, 56, 82, 108, 134},	VERSION32: {6, 34, 60, 86, 112, 138},	VERSION33: {6, 30, 58, 86, 114, 142},	VERSION34: {6, 34, 62, 90, 118, 146},	VERSION35: {6, 30, 54, 78, 102, 126, 150},	VERSION36: {6, 24, 50, 76, 102, 128, 154},	VERSION37: {6, 28, 54, 80, 106, 132, 158},	VERSION38: {6, 32, 58, 84, 110, 136, 162},	VERSION39: {6, 26, 54, 82, 110, 138, 166},	VERSION40: {6, 30, 58, 86, 114, 142, 170},
}

func NewAlignmentPattern(version *Version) *AlignmentPattern{
	ap := new(AlignmentPattern)
	ap.Positions = GetAlignmentPatternPositions(version)
	return ap
}

// GetAlignmentPatternPositions :get alignment pattern positions for specified version
func GetAlignmentPatternPositions(version *Version) []*PositionAxes{
	versionId := version.Id
	aplm := VersionAlignmentPatternLocationsMap[versionId]
	return getAvailablePositionAxes(version.GetFinderPattern(),aplm)
}

func getAvailablePositionAxes(finderPattern *FinderPattern,locArr []int) []*PositionAxes{
	locArrLen := len(locArr)
	positions := make([]*PositionAxes,0)
	for i:=0; i<locArrLen; i++ {
		for j:=0; j<locArrLen; j++ {
			pos := &PositionAxes{locArr[i],locArr[j]}
			if !IsInFinderPatternSeparatorZone(finderPattern,pos){
				positions = append(positions,pos)
			}
		}
	}
	return positions
}

// IsInFinderPatternSeparatorZone : AlignmentPattern
func IsInFinderPatternSeparatorZone(finderPattern *FinderPattern,axes *PositionAxes) bool{
	fpPositionsLen := len(finderPattern.Positions)
	for i:=0; i< fpPositionsLen; i++{
		fpp:= finderPattern.Positions[i]
		if (fpp.Axes.X == 0 && fpp.Axes.Y == 0 && // top_left
			axes.X <= fpp.Axes.X + FINDER_PATTERN_MODULE_SIZE + 1 &&
			axes.Y <= fpp.Axes.Y + FINDER_PATTERN_MODULE_SIZE + 1)||
			(fpp.Axes.X >0 && // top_right
				axes.X >= fpp.Axes.X - 1 &&
				axes.Y <= fpp.Axes.Y + FINDER_PATTERN_MODULE_SIZE + 1)||
			(fpp.Axes.X == 0 && fpp.Axes.Y >0 && //bottom_left
			axes.X <= fpp.Axes.X + FINDER_PATTERN_MODULE_SIZE + 1 &&
			axes.Y >= fpp.Axes.Y - 1){
			return true
		}
	}
	return false
}


// GetModules : 1:1:1:1:1 ,or 5:3
func (ap *AlignmentPattern) GetModules() [][]util.Module{
	modules := [][]util.Module{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}
	return modules[:]
}
