package model

// Define Timing Pattern here

// TimingPattern :
type TimingPattern struct {
	H *PositionAxesRange
	W *PositionAxesRange
}

func NewTimingPattern(version *Version) *TimingPattern {
	timingPattern := new(TimingPattern)
	positions := timingPattern.GetTimingPatternPositions(version)
	timingPattern.H = positions[0]
	timingPattern.W = positions[1]
	return timingPattern
}

// GetTimingPatternPositions :Get timing pattern positions for specified version
func (tp *TimingPattern) GetTimingPatternPositions(version *Version) []*PositionAxesRange {
	versionId := version.Id
	positionAxesRange := make([]*PositionAxesRange, 0, 2)
	moduleSize := version.GetModuleSize()
	p1 := FINDER_PATTERN_MODULE_SIZE + 1
	p2 := FINDER_PATTERN_MODULE_SIZE - 1
	p3 := moduleSize - FINDER_PATTERN_MODULE_SIZE - 1
	if versionId < 0 {
		// Micro QRCode Timing Pattern
		// Horizontal timing pattern.
		positionAxesRange = append(positionAxesRange, &PositionAxesRange{From: &PositionAxes{X: p1, Y: 0}, To: &PositionAxes{X: moduleSize, Y: 0}})
		// vertical timing pattern.
		positionAxesRange = append(positionAxesRange, &PositionAxesRange{From: &PositionAxes{X: 0, Y: p1}, To: &PositionAxes{X: 0, Y: moduleSize}})
	} else {
		// QRCode Timing Pattern
		// Horizontal timing pattern.
		positionAxesRange = append(positionAxesRange, &PositionAxesRange{From: &PositionAxes{X: p1, Y: p2}, To: &PositionAxes{X: p3, Y: p2}})
		// vertical timing pattern.
		positionAxesRange = append(positionAxesRange, &PositionAxesRange{From: &PositionAxes{X: p2, Y: p1}, To: &PositionAxes{X: p2, Y: p3}})
	}
	return positionAxesRange
}
