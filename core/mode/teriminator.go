package mode

import (
	"fmt"
	"github.com/KangSpace/gqrcode/core/logger"
	"github.com/KangSpace/gqrcode/core/model"
)

// Define Terminator handle

type Terminal struct {
}

// VersionTerminatorMap : map of Version and Terminator bits
var VersionTerminatorMap = map[model.VersionId][]byte{
	model.VersionAll: {0, 0, 0, 0},
	model.VersionM1:  {0, 0, 0},
	model.VersionM2:  {0, 0, 0, 0, 0},
	model.VersionM3:  {0, 0, 0, 0, 0, 0, 0},
	model.VersionM4:  {0, 0, 0, 0, 0, 0, 0, 0, 0},
}

// GetTerminalBits :
// The end of data in the symbol is signalled by the Terminator sequence of 0 bits,
//	as defined in Table2, appended to the data bit stream following the final mode segment.
// The terminator shall be omitted if the data bit stream completely fills the capacity of the symbol,
//	or abbreviated if the remaining capacity of the symbol is less than the required bit length of Terminator.
func (m *AbstractMode) GetTerminalBits(version *model.Version) []byte {
	// version.Id = 0 to distinguish version all of QR Code, -1 to -4 for version M1 to M4 of Micro QRCode
	versionId := version.Id
	if !version.IsMicroQRCode() {
		versionId = model.VersionAll
	}
	bits, ok := VersionTerminatorMap[versionId]
	if !ok {
		err := fmt.Errorf("Can't found Version Terminator info,input version: %s ", version.Name)
		logger.Error(err.Error())
		panic(err)
	}
	return bits[:]
}
