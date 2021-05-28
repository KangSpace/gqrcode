package model

import (
	"errors"
	"fmt"
	"github.com/gqrcode/core/cons"
	"github.com/gqrcode/core/logger"
	"strconv"
)

// Define Version here

type VersionId = int
type VersionName = string

const (
	// VERSION_ALL : All QRCode Model2 Versions value is (VersionId >> 40)
	VERSION_ALL VersionId = 0
	VERSION1 VersionId = 1
	VERSION2 VersionId = 2
	VERSION3 VersionId = 3
	VERSION4 VersionId = 4
	VERSION5 VersionId = 5
	VERSION6 VersionId = 6
	VERSION7 VersionId = 7
	VERSION8 VersionId = 8
	VERSION9 VersionId = 9
	VERSION10 VersionId = 10
	VERSION11 VersionId = 11
	VERSION12 VersionId = 12
	VERSION13 VersionId = 13
	VERSION14 VersionId = 14
	VERSION15 VersionId = 15
	VERSION16 VersionId = 16
	VERSION17 VersionId = 17
	VERSION18 VersionId = 18
	VERSION19 VersionId = 19
	VERSION20 VersionId = 20
	VERSION21 VersionId = 21
	VERSION22 VersionId = 22
	VERSION23 VersionId = 23
	VERSION24 VersionId = 24
	VERSION25 VersionId = 25
	VERSION26 VersionId = 26
	VERSION27 VersionId = 27
	VERSION28 VersionId = 28
	VERSION29 VersionId = 29
	VERSION30 VersionId = 30
	VERSION31 VersionId = 31
	VERSION32 VersionId = 32
	VERSION33 VersionId = 33
	VERSION34 VersionId = 34
	VERSION35 VersionId = 35
	VERSION36 VersionId = 36
	VERSION37 VersionId = 37
	VERSION38 VersionId = 38
	VERSION39 VersionId = 39
	VERSION40 VersionId = 40
	VERSION_M1 VersionId = -1
	VERSION_M2 VersionId = -2
	VERSION_M3 VersionId = -3
	VERSION_M4 VersionId = -4

)

// Version :QRCode Version
// Contain 2 subtypes:
// 1. QRCode Version: 1...40
// 2. Micro QRCode Version: M1...M4
//
type Version struct {
	Id VersionId `json:"version"`
	Name VersionName `json:"version"`
	finderPattern *FinderPattern
}


func (v *Version) GetFinderPattern() *FinderPattern{
	return v.finderPattern
}

var versionM1ToM4IdNameMap = map[VersionId]VersionName{
	-1:"M1",
	-2:"M2",
	-3:"M3",
	-4:"M4",
}

func NewVersion(id VersionId) *Version{
	var name VersionName
	if id<0{
		name = versionM1ToM4IdNameMap[id]
	}else{
		name = strconv.Itoa(id)
	}
	return &Version{id,name,NewFinderPattern(id)}
}

// VersionSymbolCharsAndInputDataCapacity :Page 41,Table 7-Number of symbol characters and input data capacity for QR Code
type VersionSymbolCharsAndInputDataCapacity struct {
	Version              VersionId
	ErrorCorrectionLevel cons.ErrorCorrectionLevel
	// this data codewords number for 7.4.10 Bit stream to codeword conversion, not for final Error Correction code words
	NumberOfDataCodewords int
	NumberOfDataBits      int
	// mode name: data capacity , struct: {Numeric:5}
	DataCapacity map[cons.ModeType]int
	ErrorCorrectionBlockCapacity *ErrorCorrectionBlockCapacity

}

// ErrorCorrectionBlockCapacity :Page 46,Table 9-Error correction characteristics for QR Code
// Total number of data codewords = NoECBlocksG1 * NoDataCodewordsPerBlockG1 +  NoECBlocksG2 * NoDataCodewordsPerBlockG2
type ErrorCorrectionBlockCapacity struct {
	ErrorCorrectionLevel cons.ErrorCorrectionLevel
	// number of error Correction Codewords Per Block
	NoECCodewordsPerBlock int
	// number of error correction blocks in group1
	NoECBlocksG1 int
	// number of data codewords per blocks in group1
	NoDataCodewordsPerBlockG1 int
	// number of error correction blocks in group2
	NoECBlocksG2 int
	// number of data codewords per blocks in group2
	NoDataCodewordsPerBlockG2 int
}

// GetTotalECBlocksCount :
func (ecbc *ErrorCorrectionBlockCapacity) GetTotalECBlocksCount() int{
	return ecbc.NoECBlocksG1 + ecbc.NoECBlocksG2
}

// VersionSymbolCharsAndInputDataCapacityMap : {VersionName:{ErrorCorrectionLevel:{VersionDataCapacity}}}
// Page 41,Table 7-Number of symbol characters and input data capacity for QR Code
// Page 46,Table 9-Error correction characteristics for QR Code
var VersionSymbolCharsAndInputDataCapacityMap = map[VersionId]map[cons.ErrorCorrectionLevel]*VersionSymbolCharsAndInputDataCapacity{
	VERSION_M1 	:{cons.NONE: {VERSION_M1, cons.NONE,3,20		,map[string]int{cons.NumericMode: 5},&ErrorCorrectionBlockCapacity{cons.NONE,2,1,3,0,0}}},
	VERSION_M2 	:{cons.L: {VERSION_M2	, cons.L,5		,40		,map[string]int{cons.NumericMode: 10 	, cons.AlphanumericMode:6},&ErrorCorrectionBlockCapacity{cons.L,5,1,5,0,0}}, cons.M:{VERSION_M2, cons.M,4,  32 , map[string]int{cons.NumericMode: 8  , cons.AlphanumericMode:5},&ErrorCorrectionBlockCapacity{cons.M,6,1,4,0,0}}},
	VERSION_M3 	:{cons.L: {VERSION_M3	, cons.L,11		,84		,map[string]int{cons.NumericMode: 23 	, cons.AlphanumericMode:14 	, cons.ByteMode:9	, cons.KanjiMode:6},&ErrorCorrectionBlockCapacity{cons.L,6,1,11,0,0}},	cons.M:{VERSION_M3, cons.M,9,  68 , map[string]int{cons.NumericMode: 18 , cons.AlphanumericMode:11 , cons.ByteMode:7,  cons.KanjiMode:4},&ErrorCorrectionBlockCapacity{cons.M,8,1,9,0,0}}},
	VERSION_M4 	:{cons.L: {VERSION_M4	, cons.L,16		,128	,map[string]int{cons.NumericMode: 35 	, cons.AlphanumericMode:21 	, cons.ByteMode:15	, cons.KanjiMode:9},&ErrorCorrectionBlockCapacity{cons.L,8,1,16,0,0}}, 	cons.M:{VERSION_M4, cons.M,14, 112, map[string]int{cons.NumericMode: 30 , cons.AlphanumericMode:18 , cons.ByteMode:13, cons.KanjiMode:8},&ErrorCorrectionBlockCapacity{cons.M,10,1,14,0,0}}, cons.Q:{VERSION_M4, cons.Q,10, 80 , map[string]int{cons.NumericMode: 21 , cons.AlphanumericMode:12 , cons.ByteMode:9,  cons.KanjiMode:5},&ErrorCorrectionBlockCapacity{cons.Q,14,1,10,0,0}}},
	VERSION1   	:{cons.L: {VERSION1		, cons.L,19		,152	,map[string]int{cons.NumericMode: 41 	, cons.AlphanumericMode:25 	, cons.ByteMode:17	, cons.KanjiMode:10},&ErrorCorrectionBlockCapacity{cons.L, 7, 1, 19, 0, 0}},	        cons.M:{VERSION1 , cons.M,16, 128, map[string]int{cons.NumericMode: 34 , cons.AlphanumericMode:20 , cons.ByteMode:14, cons.KanjiMode:8}		            ,&ErrorCorrectionBlockCapacity{cons.M, 10, 1, 16, 0, 0}}, cons.Q:{VERSION1, cons.Q,13, 104, map[string]int{cons.NumericMode: 27 , cons.AlphanumericMode:16 , cons.ByteMode: 11, cons.KanjiMode:7}		                ,&ErrorCorrectionBlockCapacity{cons.Q, 13, 1, 13, 0, 0}}, cons.H:{VERSION1, cons.H,9,  72 , map[string]int{cons.NumericMode: 17 , cons.AlphanumericMode:10 , cons.ByteMode:7,  cons.KanjiMode:4},&ErrorCorrectionBlockCapacity{cons.H, 17, 1, 9, 0, 0}}},
	VERSION2   	:{cons.L: {VERSION2		, cons.L,34		,272	,map[string]int{cons.NumericMode: 77 	, cons.AlphanumericMode:47 	, cons.ByteMode:32	, cons.KanjiMode:20},&ErrorCorrectionBlockCapacity{cons.L, 10, 1, 34, 0, 0}},	        cons.M:{VERSION2 , cons.M,28, 224, map[string]int{cons.NumericMode: 63 , cons.AlphanumericMode:38 , cons.ByteMode:26, cons.KanjiMode:16}		            ,&ErrorCorrectionBlockCapacity{cons.M, 16, 1, 28, 0, 0}}, cons.Q:{VERSION2, cons.Q,22, 176, map[string]int{cons.NumericMode: 48 , cons.AlphanumericMode:29 , cons.ByteMode: 20, cons.KanjiMode:12}		            ,&ErrorCorrectionBlockCapacity{cons.Q, 22, 1, 22, 0, 0}}, cons.H:{VERSION2, cons.H,16, 128, map[string]int{cons.NumericMode: 34 , cons.AlphanumericMode:20 , cons.ByteMode:14, cons.KanjiMode:8},&ErrorCorrectionBlockCapacity{cons.H, 28, 1, 16, 0, 0}}},
	VERSION3   	:{cons.L: {VERSION3		, cons.L,55		,440	,map[string]int{cons.NumericMode: 127	, cons.AlphanumericMode:77 	, cons.ByteMode:53	, cons.KanjiMode:32},&ErrorCorrectionBlockCapacity{cons.L, 15, 1, 55, 0, 0}},	        cons.M:{VERSION3 , cons.M,44, 352, map[string]int{cons.NumericMode: 101, cons.AlphanumericMode:61 , cons.ByteMode:42, cons.KanjiMode:26}		            ,&ErrorCorrectionBlockCapacity{cons.M, 26, 1, 44, 0, 0}}, cons.Q:{VERSION3, cons.Q,34, 272, map[string]int{cons.NumericMode: 77 , cons.AlphanumericMode:47 , cons.ByteMode: 32, cons.KanjiMode:20}		            ,&ErrorCorrectionBlockCapacity{cons.Q, 18, 2, 17, 0, 0}}, cons.H:{VERSION3, cons.H,26, 208, map[string]int{cons.NumericMode: 58 , cons.AlphanumericMode:35 , cons.ByteMode:24, cons.KanjiMode:15},&ErrorCorrectionBlockCapacity{cons.H, 22, 2, 13, 0, 0}}},
	VERSION4   	:{cons.L: {VERSION4		, cons.L,80		,640	,map[string]int{cons.NumericMode: 187	, cons.AlphanumericMode:114	, cons.ByteMode:78	, cons.KanjiMode:48},&ErrorCorrectionBlockCapacity{cons.L, 20, 1, 80, 0, 0}},	        cons.M:{VERSION4 , cons.M,64, 512, map[string]int{cons.NumericMode: 149, cons.AlphanumericMode:90 , cons.ByteMode:62, cons.KanjiMode:38}		            ,&ErrorCorrectionBlockCapacity{cons.M, 18, 2, 32, 0, 0}}, cons.Q:{VERSION4, cons.Q,48, 384, map[string]int{cons.NumericMode: 111, cons.AlphanumericMode:67 , cons.ByteMode: 46, cons.KanjiMode:28}		            ,&ErrorCorrectionBlockCapacity{cons.Q, 26, 2, 24, 0, 0}}, cons.H:{VERSION4, cons.H,36, 288, map[string]int{cons.NumericMode: 82 , cons.AlphanumericMode:50 , cons.ByteMode:34, cons.KanjiMode:21},&ErrorCorrectionBlockCapacity{cons.H, 16, 4, 9, 0, 0}}},
	VERSION5   	:{cons.L: {VERSION5		, cons.L,108	,864	,map[string]int{cons.NumericMode: 255	, cons.AlphanumericMode:154	, cons.ByteMode:106	, cons.KanjiMode:65},&ErrorCorrectionBlockCapacity{cons.L, 26, 1, 108, 0, 0}},	    cons.M:{VERSION5 , cons.M,86, 688, map[string]int{cons.NumericMode: 202, cons.AlphanumericMode:122, cons.ByteMode:84, cons.KanjiMode:52}		            ,&ErrorCorrectionBlockCapacity{cons.M, 24, 2, 43, 0, 0}}, cons.Q:{VERSION5, cons.Q,62, 496, map[string]int{cons.NumericMode: 144, cons.AlphanumericMode:87 , cons.ByteMode: 60, cons.KanjiMode:37}		            ,&ErrorCorrectionBlockCapacity{cons.Q, 18, 2, 15, 2, 16}}, cons.H:{VERSION5, cons.H,46, 368, map[string]int{cons.NumericMode: 106, cons.AlphanumericMode:64 , cons.ByteMode:44, cons.KanjiMode:27},&ErrorCorrectionBlockCapacity{cons.H, 22, 2, 11, 2, 12}}},
	VERSION6	:{cons.L: {VERSION6		, cons.L,136	,1088	,map[string]int{cons.NumericMode: 322	, cons.AlphanumericMode:195	, cons.ByteMode:134	, cons.KanjiMode:82},&ErrorCorrectionBlockCapacity{cons.L, 18, 2, 68, 0, 0}},	        cons.M:{VERSION6 , cons.M,108	,864	,map[string]int{cons.NumericMode: 255	, cons.AlphanumericMode:154	, cons.ByteMode:106	, cons.KanjiMode:65}		,&ErrorCorrectionBlockCapacity{cons.M, 16, 4, 27, 0, 0}}, cons.Q:{VERSION6	, cons.Q,76     ,608	,map[string]int{cons.NumericMode: 178	, cons.AlphanumericMode:108	, cons.ByteMode:74	, cons.KanjiMode:45}		,&ErrorCorrectionBlockCapacity{cons.Q, 24, 4, 19, 0, 0}}, cons.H:{VERSION6	, cons.H,60	    ,480	,map[string]int{cons.NumericMode: 139	, cons.AlphanumericMode:84	, cons.ByteMode:58	, cons.KanjiMode:36},&ErrorCorrectionBlockCapacity{cons.H, 28, 4, 15, 0, 0}}},
	VERSION7	:{cons.L: {VERSION7		, cons.L,156	,1248	,map[string]int{cons.NumericMode: 370	, cons.AlphanumericMode:224	, cons.ByteMode:154	, cons.KanjiMode:95},&ErrorCorrectionBlockCapacity{cons.L, 20, 2, 78, 0, 0}},	        cons.M:{VERSION7 , cons.M,124	,992	,map[string]int{cons.NumericMode: 293	, cons.AlphanumericMode:178	, cons.ByteMode:122	, cons.KanjiMode:75}		,&ErrorCorrectionBlockCapacity{cons.M, 18, 4, 31, 0, 0}}, cons.Q:{VERSION7	, cons.Q,88	    ,704	,map[string]int{cons.NumericMode: 207	, cons.AlphanumericMode:125	, cons.ByteMode:86	, cons.KanjiMode:53}		,&ErrorCorrectionBlockCapacity{cons.Q, 18, 2, 14, 4, 15}}, cons.H:{VERSION7	, cons.H,66	    ,528	,map[string]int{cons.NumericMode: 154	, cons.AlphanumericMode:93	, cons.ByteMode:64	, cons.KanjiMode:39},&ErrorCorrectionBlockCapacity{cons.H, 26, 4, 13, 1, 14}}},
	VERSION8	:{cons.L: {VERSION8		, cons.L,194	,1552	,map[string]int{cons.NumericMode: 461	, cons.AlphanumericMode:279	, cons.ByteMode:192	, cons.KanjiMode:118},&ErrorCorrectionBlockCapacity{cons.L, 24, 2, 97, 0, 0}},	    cons.M:{VERSION8 , cons.M,154	,1232	,map[string]int{cons.NumericMode: 365	, cons.AlphanumericMode:221	, cons.ByteMode:152	, cons.KanjiMode:93}		,&ErrorCorrectionBlockCapacity{cons.M, 22, 2, 38, 2, 39}}, cons.Q:{VERSION8	, cons.Q,110	,880	,map[string]int{cons.NumericMode: 259	, cons.AlphanumericMode:157	, cons.ByteMode:108	, cons.KanjiMode:66}		,&ErrorCorrectionBlockCapacity{cons.Q, 22, 4, 18, 2, 19}}, cons.H:{VERSION8   , cons.H,86	    ,688	,map[string]int{cons.NumericMode: 202	, cons.AlphanumericMode:122	, cons.ByteMode:84	, cons.KanjiMode:52},&ErrorCorrectionBlockCapacity{cons.H, 26, 4, 14, 2, 15}}},
	VERSION9	:{cons.L: {VERSION9		, cons.L,232	,1856	,map[string]int{cons.NumericMode: 552	, cons.AlphanumericMode:335	, cons.ByteMode:230	, cons.KanjiMode:141},&ErrorCorrectionBlockCapacity{cons.L, 30, 2, 116, 0, 0}},	    cons.M:{VERSION9 , cons.M,182	,1456	,map[string]int{cons.NumericMode: 432	, cons.AlphanumericMode:262	, cons.ByteMode:180	, cons.KanjiMode:111}		,&ErrorCorrectionBlockCapacity{cons.M, 22, 3, 36, 2, 37}}, cons.Q:{VERSION9	, cons.Q,132	,1056	,map[string]int{cons.NumericMode: 312	, cons.AlphanumericMode:189	, cons.ByteMode:130	, cons.KanjiMode:80}		,&ErrorCorrectionBlockCapacity{cons.Q, 20, 4, 16, 4, 17}}, cons.H:{VERSION9	, cons.H,100	,800	,map[string]int{cons.NumericMode: 235	, cons.AlphanumericMode:143	, cons.ByteMode:98	, cons.KanjiMode:60},&ErrorCorrectionBlockCapacity{cons.H, 24, 4, 12, 4, 13}}},
	VERSION10	:{cons.L: {VERSION10	, cons.L,274	,2192	,map[string]int{cons.NumericMode: 652	, cons.AlphanumericMode:395	, cons.ByteMode:271	, cons.KanjiMode:167},&ErrorCorrectionBlockCapacity{cons.L, 18, 2, 68, 2, 69}},	    cons.M:{VERSION10, cons.M,216	,1728	,map[string]int{cons.NumericMode: 513	, cons.AlphanumericMode:311	, cons.ByteMode:213	, cons.KanjiMode:131}		,&ErrorCorrectionBlockCapacity{cons.M, 26, 4, 43, 1, 44}}, cons.Q:{VERSION10	, cons.Q,154	,1232	,map[string]int{cons.NumericMode: 364	, cons.AlphanumericMode:221	, cons.ByteMode:151	, cons.KanjiMode:93}		,&ErrorCorrectionBlockCapacity{cons.Q, 24, 6, 19, 2, 20}}, cons.H:{VERSION10	, cons.H,122	,976	,map[string]int{cons.NumericMode: 288	, cons.AlphanumericMode:174	, cons.ByteMode:119	, cons.KanjiMode:74},&ErrorCorrectionBlockCapacity{cons.H, 28, 6, 15, 2, 16}}},
	VERSION11	:{cons.L: {VERSION11	, cons.L,324	,2592	,map[string]int{cons.NumericMode: 772	, cons.AlphanumericMode:468	, cons.ByteMode:321	, cons.KanjiMode:198},&ErrorCorrectionBlockCapacity{cons.L, 20, 4, 81, 0, 0}},	    cons.M:{VERSION11, cons.M,254	,2032	,map[string]int{cons.NumericMode: 604	, cons.AlphanumericMode:366	, cons.ByteMode:251	, cons.KanjiMode:155}		,&ErrorCorrectionBlockCapacity{cons.M, 30, 1, 50, 4, 51}}, cons.Q:{VERSION11	, cons.Q,180	,1440	,map[string]int{cons.NumericMode: 427	, cons.AlphanumericMode:259	, cons.ByteMode:177	, cons.KanjiMode:109}		,&ErrorCorrectionBlockCapacity{cons.Q, 28, 4, 22, 4, 23}}, cons.H:{VERSION11	, cons.H,140	,1120	,map[string]int{cons.NumericMode: 331	, cons.AlphanumericMode:200	, cons.ByteMode:137	, cons.KanjiMode:85},&ErrorCorrectionBlockCapacity{cons.H, 24, 3, 12, 8, 13}}},
	VERSION12	:{cons.L: {VERSION12	, cons.L,370	,2960	,map[string]int{cons.NumericMode: 883	, cons.AlphanumericMode:535	, cons.ByteMode:367	, cons.KanjiMode:226},&ErrorCorrectionBlockCapacity{cons.L, 24, 2, 92, 2, 93}},	    cons.M:{VERSION12, cons.M,290	,2320	,map[string]int{cons.NumericMode: 691	, cons.AlphanumericMode:419	, cons.ByteMode:287	, cons.KanjiMode:177}		,&ErrorCorrectionBlockCapacity{cons.M, 22, 6, 36, 2, 37}}, cons.Q:{VERSION12	, cons.Q,206	,1648	,map[string]int{cons.NumericMode: 489	, cons.AlphanumericMode:296	, cons.ByteMode:203	, cons.KanjiMode:125}		,&ErrorCorrectionBlockCapacity{cons.Q, 26, 4, 20, 6, 21}}, cons.H:{VERSION12	, cons.H,158	,1264	,map[string]int{cons.NumericMode: 374	, cons.AlphanumericMode:227	, cons.ByteMode:155	, cons.KanjiMode:96},&ErrorCorrectionBlockCapacity{cons.H, 28, 7, 14, 4, 15}}},
	VERSION13	:{cons.L: {VERSION13	, cons.L,428	,3424	,map[string]int{cons.NumericMode: 1022	, cons.AlphanumericMode:619	, cons.ByteMode:425	, cons.KanjiMode:262},&ErrorCorrectionBlockCapacity{cons.L, 26, 4, 107, 0, 0}},	    cons.M:{VERSION13, cons.M,334	,2672	,map[string]int{cons.NumericMode: 796	, cons.AlphanumericMode:483	, cons.ByteMode:331	, cons.KanjiMode:204}		,&ErrorCorrectionBlockCapacity{cons.M, 22, 8, 37, 1, 38}}, cons.Q:{VERSION13	, cons.Q,244	,1952	,map[string]int{cons.NumericMode: 580	, cons.AlphanumericMode:352	, cons.ByteMode:241	, cons.KanjiMode:149}		,&ErrorCorrectionBlockCapacity{cons.Q, 24, 8, 20, 4, 21}}, cons.H:{VERSION13	, cons.H,180	,1440	,map[string]int{cons.NumericMode: 427	, cons.AlphanumericMode:259	, cons.ByteMode:177	, cons.KanjiMode:109},&ErrorCorrectionBlockCapacity{cons.H, 22, 12, 11, 4, 12}}},
	VERSION14	:{cons.L: {VERSION14	, cons.L,461	,3688	,map[string]int{cons.NumericMode: 1101	, cons.AlphanumericMode:667	, cons.ByteMode:458	, cons.KanjiMode:282},&ErrorCorrectionBlockCapacity{cons.L, 30, 3, 115, 1, 116}},    cons.M:{VERSION14, cons.M,365	,2920	,map[string]int{cons.NumericMode: 871	, cons.AlphanumericMode:528	, cons.ByteMode:362	, cons.KanjiMode:223}		,&ErrorCorrectionBlockCapacity{cons.M, 24, 4, 40, 5, 41}}, cons.Q:{VERSION14	, cons.Q,261	,2088	,map[string]int{cons.NumericMode: 621	, cons.AlphanumericMode:376	, cons.ByteMode:258	, cons.KanjiMode:159}		,&ErrorCorrectionBlockCapacity{cons.Q, 20, 11, 16, 5, 17}}, cons.H:{VERSION14	, cons.H,197	,1576	,map[string]int{cons.NumericMode: 468	, cons.AlphanumericMode:283	, cons.ByteMode:194	, cons.KanjiMode:120},&ErrorCorrectionBlockCapacity{cons.H, 24, 11, 12, 5, 13}}},
	VERSION15	:{cons.L: {VERSION15	, cons.L,523	,4184	,map[string]int{cons.NumericMode: 1250	, cons.AlphanumericMode:758	, cons.ByteMode:520	, cons.KanjiMode:320},&ErrorCorrectionBlockCapacity{cons.L, 22, 5, 87, 1, 88}},	    cons.M:{VERSION15, cons.M,415	,3320	,map[string]int{cons.NumericMode: 991	, cons.AlphanumericMode:600	, cons.ByteMode:412	, cons.KanjiMode:254}		,&ErrorCorrectionBlockCapacity{cons.M, 24, 5, 41, 5, 42}}, cons.Q:{VERSION15	, cons.Q,295	,2360	,map[string]int{cons.NumericMode: 703	, cons.AlphanumericMode:426	, cons.ByteMode:292	, cons.KanjiMode:180}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 5, 24, 7, 25}}, cons.H:{VERSION15	, cons.H,223	,1784	,map[string]int{cons.NumericMode: 530	, cons.AlphanumericMode:321	, cons.ByteMode:220	, cons.KanjiMode:136},&ErrorCorrectionBlockCapacity{cons.H, 24, 11, 12, 7, 13}}},
	VERSION16	:{cons.L: {VERSION16	, cons.L,589	,4712	,map[string]int{cons.NumericMode: 1408	, cons.AlphanumericMode:854	, cons.ByteMode:586	, cons.KanjiMode:361},&ErrorCorrectionBlockCapacity{cons.L, 24, 5, 98, 1, 99}},	    cons.M:{VERSION16, cons.M,453	,3624	,map[string]int{cons.NumericMode: 1082	, cons.AlphanumericMode:656	, cons.ByteMode:450	, cons.KanjiMode:277}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 7, 45, 3, 46}}, cons.Q:{VERSION16	, cons.Q,325	,2600	,map[string]int{cons.NumericMode: 775	, cons.AlphanumericMode:470	, cons.ByteMode:322	, cons.KanjiMode:198}		,&ErrorCorrectionBlockCapacity{cons.Q, 24, 15, 19, 2, 20}}, cons.H:{VERSION16  , cons.H,253	,2024	,map[string]int{cons.NumericMode: 602	, cons.AlphanumericMode:365	, cons.ByteMode:250	, cons.KanjiMode:154},&ErrorCorrectionBlockCapacity{cons.H, 30, 3, 15, 13, 16}}},
	VERSION17	:{cons.L: {VERSION17	, cons.L,647	,5176	,map[string]int{cons.NumericMode: 1548	, cons.AlphanumericMode:938	, cons.ByteMode:644	, cons.KanjiMode:397},&ErrorCorrectionBlockCapacity{cons.L, 28, 1, 107, 5, 108}},    cons.M:{VERSION17, cons.M,507	,4056	,map[string]int{cons.NumericMode: 1212	, cons.AlphanumericMode:734	, cons.ByteMode:504	, cons.KanjiMode:310}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 10, 46, 1, 47}}, cons.Q:{VERSION17	, cons.Q,367	,2936	,map[string]int{cons.NumericMode: 876	, cons.AlphanumericMode:531	, cons.ByteMode:364	, cons.KanjiMode:224}		,&ErrorCorrectionBlockCapacity{cons.Q, 28, 1, 22, 15, 23}}, cons.H:{VERSION17	, cons.H,283	,2264	,map[string]int{cons.NumericMode: 674	, cons.AlphanumericMode:408	, cons.ByteMode:280	, cons.KanjiMode:173},&ErrorCorrectionBlockCapacity{cons.H, 28, 2, 14, 17, 15}}},
	VERSION18	:{cons.L: {VERSION18	, cons.L,721	,5768	,map[string]int{cons.NumericMode: 1725	, cons.AlphanumericMode:1046, cons.ByteMode:718	, cons.KanjiMode:442},&ErrorCorrectionBlockCapacity{cons.L, 30, 5, 120, 1, 121}},    cons.M:{VERSION18, cons.M,563	,4504	,map[string]int{cons.NumericMode: 1346	, cons.AlphanumericMode:816	, cons.ByteMode:560	, cons.KanjiMode:345}		,&ErrorCorrectionBlockCapacity{cons.M, 26, 9, 43, 4, 44}}, cons.Q:{VERSION18	, cons.Q,397	,3176	,map[string]int{cons.NumericMode: 948	, cons.AlphanumericMode:574	, cons.ByteMode:394	, cons.KanjiMode:243}		,&ErrorCorrectionBlockCapacity{cons.Q, 28, 17, 22, 1, 23}}, cons.H:{VERSION18	, cons.H,313	,2504	,map[string]int{cons.NumericMode: 746	, cons.AlphanumericMode:452	, cons.ByteMode:310	, cons.KanjiMode:191},&ErrorCorrectionBlockCapacity{cons.H, 28, 2, 14, 19, 15}}},
	VERSION19	:{cons.L: {VERSION19	, cons.L,795	,6360	,map[string]int{cons.NumericMode: 1903	, cons.AlphanumericMode:1153, cons.ByteMode:792	, cons.KanjiMode:488},&ErrorCorrectionBlockCapacity{cons.L, 28, 3, 113, 4, 114}},    cons.M:{VERSION19, cons.M,627	,5016	,map[string]int{cons.NumericMode: 1500	, cons.AlphanumericMode:909	, cons.ByteMode:624	, cons.KanjiMode:384}		,&ErrorCorrectionBlockCapacity{cons.M, 26, 3, 44, 11, 45}}, cons.Q:{VERSION19	, cons.Q,445	,3560	,map[string]int{cons.NumericMode: 1063	, cons.AlphanumericMode:644	, cons.ByteMode:442	, cons.KanjiMode:272}		,&ErrorCorrectionBlockCapacity{cons.Q, 26, 17, 21, 4, 22}}, cons.H:{VERSION19	, cons.H,341	,2728	,map[string]int{cons.NumericMode: 813	, cons.AlphanumericMode:493	, cons.ByteMode:338	, cons.KanjiMode:208},&ErrorCorrectionBlockCapacity{cons.H, 26, 9, 13, 16, 14}}},
	VERSION20	:{cons.L: {VERSION20	, cons.L,861	,6888	,map[string]int{cons.NumericMode: 2061	, cons.AlphanumericMode:1249, cons.ByteMode:858	, cons.KanjiMode:528},&ErrorCorrectionBlockCapacity{cons.L, 28, 3, 107, 5, 108}},    cons.M:{VERSION20, cons.M,669	,5352	,map[string]int{cons.NumericMode: 1600	, cons.AlphanumericMode:970	, cons.ByteMode:666	, cons.KanjiMode:410}		,&ErrorCorrectionBlockCapacity{cons.M, 26, 3, 41, 13, 42}}, cons.Q:{VERSION20	, cons.Q,485	,3880	,map[string]int{cons.NumericMode: 1159	, cons.AlphanumericMode:702	, cons.ByteMode:482	, cons.KanjiMode:297}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 15, 24, 5, 25}}, cons.H:{VERSION20	, cons.H,385	,3080	,map[string]int{cons.NumericMode: 919	, cons.AlphanumericMode:557	, cons.ByteMode:382	, cons.KanjiMode:235},&ErrorCorrectionBlockCapacity{cons.H, 28, 15, 15, 10, 16}}},
	VERSION21	:{cons.L: {VERSION21	, cons.L,932	,7456	,map[string]int{cons.NumericMode: 2232	, cons.AlphanumericMode:1352, cons.ByteMode:929	, cons.KanjiMode:572},&ErrorCorrectionBlockCapacity{cons.L, 28, 4, 116, 4, 117}},    cons.M:{VERSION21, cons.M,714	,5712	,map[string]int{cons.NumericMode: 1708	, cons.AlphanumericMode:1035, cons.ByteMode:711	, cons.KanjiMode:438}		,&ErrorCorrectionBlockCapacity{cons.M, 26, 17, 42, 0, 0}}, cons.Q:{VERSION21	, cons.Q,512	,4096	,map[string]int{cons.NumericMode: 1224	, cons.AlphanumericMode:742	, cons.ByteMode:509	, cons.KanjiMode:314}		,&ErrorCorrectionBlockCapacity{cons.Q, 28, 17, 22, 6, 23}}, cons.H:{VERSION21	, cons.H,406	,3248	,map[string]int{cons.NumericMode: 969	, cons.AlphanumericMode:587	, cons.ByteMode:403	, cons.KanjiMode:248},&ErrorCorrectionBlockCapacity{cons.H, 30, 19, 16, 6, 17}}},
	VERSION22	:{cons.L: {VERSION22	, cons.L,1006	,8048	,map[string]int{cons.NumericMode: 2409	, cons.AlphanumericMode:1460, cons.ByteMode:1003, cons.KanjiMode:618},&ErrorCorrectionBlockCapacity{cons.L, 28, 2, 111, 7, 112}},    cons.M:{VERSION22, cons.M,782	,6256	,map[string]int{cons.NumericMode: 1872	, cons.AlphanumericMode:1134, cons.ByteMode:779	, cons.KanjiMode:480}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 17, 46, 0, 0}}, cons.Q:{VERSION22	, cons.Q,568	,4544	,map[string]int{cons.NumericMode: 1358	, cons.AlphanumericMode:823	, cons.ByteMode:565	, cons.KanjiMode:348}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 7, 24, 16, 25}}, cons.H:{VERSION22	, cons.H,442	,3536	,map[string]int{cons.NumericMode: 1056	, cons.AlphanumericMode:640	, cons.ByteMode:439	, cons.KanjiMode:270},&ErrorCorrectionBlockCapacity{cons.H, 24, 34, 13, 0, 0}}},
	VERSION23	:{cons.L: {VERSION23	, cons.L,1094	,8752	,map[string]int{cons.NumericMode: 2620	, cons.AlphanumericMode:1588, cons.ByteMode:1091, cons.KanjiMode:672},&ErrorCorrectionBlockCapacity{cons.L, 30, 4, 121, 5, 122}},    cons.M:{VERSION23, cons.M,860	,6880	,map[string]int{cons.NumericMode: 2059	, cons.AlphanumericMode:1248, cons.ByteMode:857	, cons.KanjiMode:528}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 4, 47, 14, 48}}, cons.Q:{VERSION23	, cons.Q,614	,4912	,map[string]int{cons.NumericMode: 1468	, cons.AlphanumericMode:890	, cons.ByteMode:611	, cons.KanjiMode:376}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 11, 24, 14, 25}}, cons.H:{VERSION23	, cons.H,464	,3712	,map[string]int{cons.NumericMode: 1108	, cons.AlphanumericMode:672	, cons.ByteMode:461	, cons.KanjiMode:284},&ErrorCorrectionBlockCapacity{cons.H, 30, 16, 15, 14, 16}}},
	VERSION24	:{cons.L: {VERSION24	, cons.L,1174	,9392	,map[string]int{cons.NumericMode: 2812	, cons.AlphanumericMode:1704, cons.ByteMode:1171, cons.KanjiMode:721},&ErrorCorrectionBlockCapacity{cons.L, 30, 6, 117, 4, 118}},    cons.M:{VERSION24, cons.M,914	,7312	,map[string]int{cons.NumericMode: 2188	, cons.AlphanumericMode:1326, cons.ByteMode:911	, cons.KanjiMode:561}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 6, 45, 14, 46}}, cons.Q:{VERSION24	, cons.Q,664	,5312	,map[string]int{cons.NumericMode: 1588	, cons.AlphanumericMode:963	, cons.ByteMode:661	, cons.KanjiMode:407}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 11, 24, 16, 25}}, cons.H:{VERSION24	, cons.H,514	,4112	,map[string]int{cons.NumericMode: 1228	, cons.AlphanumericMode:744	, cons.ByteMode:511	, cons.KanjiMode:315},&ErrorCorrectionBlockCapacity{cons.H, 30, 30, 16, 2, 17}}},
	VERSION25	:{cons.L: {VERSION25	, cons.L,1276	,10208	,map[string]int{cons.NumericMode: 3057	, cons.AlphanumericMode:1853, cons.ByteMode:1273, cons.KanjiMode:784},&ErrorCorrectionBlockCapacity{cons.L, 26, 8, 106, 4, 107}},    cons.M:{VERSION25, cons.M,1000	,8000	,map[string]int{cons.NumericMode: 2395	, cons.AlphanumericMode:1451, cons.ByteMode:997	, cons.KanjiMode:614}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 8, 47, 13, 48}}, cons.Q:{VERSION25	, cons.Q,718	,5744	,map[string]int{cons.NumericMode: 1718	, cons.AlphanumericMode:1041, cons.ByteMode:715	, cons.KanjiMode:440}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 7, 24, 22, 25}}, cons.H:{VERSION25	, cons.H,538	,4304	,map[string]int{cons.NumericMode: 1286	, cons.AlphanumericMode:779	, cons.ByteMode:535	, cons.KanjiMode:330},&ErrorCorrectionBlockCapacity{cons.H, 30, 22, 15, 13, 16}}},
	VERSION26	:{cons.L: {VERSION26	, cons.L,1370	,10960	,map[string]int{cons.NumericMode: 3283	, cons.AlphanumericMode:1990, cons.ByteMode:1367, cons.KanjiMode:842},&ErrorCorrectionBlockCapacity{cons.L, 28, 10, 114, 2, 115}},	cons.M:{VERSION26, cons.M,1062	,8496	,map[string]int{cons.NumericMode: 2544	, cons.AlphanumericMode:1542, cons.ByteMode:1059, cons.KanjiMode:652}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 19, 46, 4, 47}}, cons.Q:{VERSION26	, cons.Q,754	,6032	,map[string]int{cons.NumericMode: 1804	, cons.AlphanumericMode:1094, cons.ByteMode:751	, cons.KanjiMode:462}		,&ErrorCorrectionBlockCapacity{cons.Q, 28, 28, 22, 6, 23}}, cons.H:{VERSION26	, cons.H,596	,4768	,map[string]int{cons.NumericMode: 1425	, cons.AlphanumericMode:864	, cons.ByteMode:593	, cons.KanjiMode:365},&ErrorCorrectionBlockCapacity{cons.H, 30, 33, 16, 4, 17}}},
	VERSION27	:{cons.L: {VERSION27	, cons.L,1468	,11744	,map[string]int{cons.NumericMode: 3517	, cons.AlphanumericMode:2132, cons.ByteMode:1465, cons.KanjiMode:902},&ErrorCorrectionBlockCapacity{cons.L, 30, 8, 122, 4, 123}},	cons.M:{VERSION27, cons.M,1128	,9024	,map[string]int{cons.NumericMode: 2701	, cons.AlphanumericMode:1637, cons.ByteMode:1125, cons.KanjiMode:692}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 22, 45, 3, 46}}, cons.Q:{VERSION27	, cons.Q,808	,6464	,map[string]int{cons.NumericMode: 1933	, cons.AlphanumericMode:1172, cons.ByteMode:805	, cons.KanjiMode:496}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 8, 23, 26, 24}}, cons.H:{VERSION27	, cons.H,628	,5024	,map[string]int{cons.NumericMode: 1501	, cons.AlphanumericMode:910	, cons.ByteMode:625	, cons.KanjiMode:385},&ErrorCorrectionBlockCapacity{cons.H, 30, 12, 15, 28, 16}}},
	VERSION28	:{cons.L: {VERSION28	, cons.L,1531	,12248	,map[string]int{cons.NumericMode: 3669	, cons.AlphanumericMode:2223, cons.ByteMode:1528, cons.KanjiMode:940},&ErrorCorrectionBlockCapacity{cons.L, 30, 3, 117, 10, 118}},	cons.M:{VERSION28, cons.M,1193	,9544	,map[string]int{cons.NumericMode: 2857	, cons.AlphanumericMode:1732, cons.ByteMode:1190, cons.KanjiMode:732}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 3, 45, 23, 46}}, cons.Q:{VERSION28	, cons.Q,871	,6968	,map[string]int{cons.NumericMode: 2085	, cons.AlphanumericMode:1263, cons.ByteMode:868	, cons.KanjiMode:534}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 4, 24, 31, 25}}, cons.H:{VERSION28	, cons.H,661	,5288	,map[string]int{cons.NumericMode: 1581	, cons.AlphanumericMode:958	, cons.ByteMode:658	, cons.KanjiMode:405},&ErrorCorrectionBlockCapacity{cons.H, 30, 11, 15, 31, 16}}},
	VERSION29	:{cons.L: {VERSION29	, cons.L,1631	,13048	,map[string]int{cons.NumericMode: 3909	, cons.AlphanumericMode:2369, cons.ByteMode:1628, cons.KanjiMode:1002},&ErrorCorrectionBlockCapacity{cons.L, 30, 7, 116, 7, 117}},   cons.M:{VERSION29, cons.M,1267	,10136	,map[string]int{cons.NumericMode: 3035	, cons.AlphanumericMode:1839, cons.ByteMode:1264, cons.KanjiMode:778}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 21, 45, 7, 46}}, cons.Q:{VERSION29	, cons.Q,911	,7288	,map[string]int{cons.NumericMode: 2181	, cons.AlphanumericMode:1322, cons.ByteMode:908	, cons.KanjiMode:559}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 1, 23, 37, 24}}, cons.H:{VERSION29	, cons.H,701	,5608	,map[string]int{cons.NumericMode: 1677	, cons.AlphanumericMode:1016, cons.ByteMode:698	, cons.KanjiMode:430},&ErrorCorrectionBlockCapacity{cons.H, 30, 19, 15, 26, 16}}},
	VERSION30	:{cons.L: {VERSION30	, cons.L,1735	,13880	,map[string]int{cons.NumericMode: 4158	, cons.AlphanumericMode:2520, cons.ByteMode:1732, cons.KanjiMode:1066},&ErrorCorrectionBlockCapacity{cons.L, 30, 5, 115, 10, 116}},  cons.M:{VERSION30, cons.M,1373	,10984	,map[string]int{cons.NumericMode: 3289	, cons.AlphanumericMode:1994, cons.ByteMode:1370, cons.KanjiMode:843}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 19, 47, 10, 48}}, cons.Q:{VERSION30	, cons.Q,985	,7880	,map[string]int{cons.NumericMode: 2358	, cons.AlphanumericMode:1429, cons.ByteMode:982	, cons.KanjiMode:604}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 15, 24, 25, 25}}, cons.H:{VERSION30	, cons.H,745	,5960	,map[string]int{cons.NumericMode: 1782	, cons.AlphanumericMode:1080, cons.ByteMode:742	, cons.KanjiMode:457},&ErrorCorrectionBlockCapacity{cons.H, 30, 23, 15, 25, 16}}},
	VERSION31	:{cons.L: {VERSION31	, cons.L,1843	,14744	,map[string]int{cons.NumericMode: 4417	, cons.AlphanumericMode:2677, cons.ByteMode:1840, cons.KanjiMode:1132},&ErrorCorrectionBlockCapacity{cons.L, 30, 13, 115, 3, 116}},  cons.M:{VERSION31, cons.M,1455	,11640	,map[string]int{cons.NumericMode: 3486	, cons.AlphanumericMode:2113, cons.ByteMode:1452, cons.KanjiMode:894}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 2, 46, 29, 47}}, cons.Q:{VERSION31	, cons.Q,1033	,8264	,map[string]int{cons.NumericMode: 2473	, cons.AlphanumericMode:1499, cons.ByteMode:1030, cons.KanjiMode:634}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 42, 24, 1, 25}}, cons.H:{VERSION31	, cons.H,793	,6344	,map[string]int{cons.NumericMode: 1897	, cons.AlphanumericMode:1150, cons.ByteMode:790	, cons.KanjiMode:486},&ErrorCorrectionBlockCapacity{cons.H, 30, 23, 15, 28, 16}}},
	VERSION32	:{cons.L: {VERSION32	, cons.L,1955	,15640	,map[string]int{cons.NumericMode: 4686	, cons.AlphanumericMode:2840, cons.ByteMode:1952, cons.KanjiMode:1201},&ErrorCorrectionBlockCapacity{cons.L, 30, 17, 115, 0, 0}},    cons.M:{VERSION32, cons.M,1541	,12328	,map[string]int{cons.NumericMode: 3693	, cons.AlphanumericMode:2238, cons.ByteMode:1538, cons.KanjiMode:947}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 10, 46, 23, 47}}, cons.Q:{VERSION32	, cons.Q,1115	,8920	,map[string]int{cons.NumericMode: 2670	, cons.AlphanumericMode:1618, cons.ByteMode:1112, cons.KanjiMode:684}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 10, 24, 35, 25}}, cons.H:{VERSION32	, cons.H,845	,6760	,map[string]int{cons.NumericMode: 2022	, cons.AlphanumericMode:1226, cons.ByteMode:842	, cons.KanjiMode:518},&ErrorCorrectionBlockCapacity{cons.H, 30, 19, 15, 35, 16}}},
	VERSION33	:{cons.L: {VERSION33	, cons.L,2071	,16568	,map[string]int{cons.NumericMode: 4965	, cons.AlphanumericMode:3009, cons.ByteMode:2068, cons.KanjiMode:1273},&ErrorCorrectionBlockCapacity{cons.L, 30, 17, 115, 1, 116}},  cons.M:{VERSION33, cons.M,1631	,13048	,map[string]int{cons.NumericMode: 3909	, cons.AlphanumericMode:2369, cons.ByteMode:1628, cons.KanjiMode:1002}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 14, 46, 21, 47}}, cons.Q:{VERSION33, cons.Q,1171	,9368	,map[string]int{cons.NumericMode: 2805	, cons.AlphanumericMode:1700, cons.ByteMode:1168, cons.KanjiMode:719}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 29, 24, 19, 25}}, cons.H:{VERSION33	, cons.H,901	,7208	,map[string]int{cons.NumericMode: 2157	, cons.AlphanumericMode:1307, cons.ByteMode:898	, cons.KanjiMode:553},&ErrorCorrectionBlockCapacity{cons.H, 30, 11, 15, 46, 16}}},
	VERSION34	:{cons.L: {VERSION34	, cons.L,2191	,17528	,map[string]int{cons.NumericMode: 5253	, cons.AlphanumericMode:3183, cons.ByteMode:2188, cons.KanjiMode:1347},&ErrorCorrectionBlockCapacity{cons.L, 30, 13, 115, 6, 116}},  cons.M:{VERSION34, cons.M,1725	,13800	,map[string]int{cons.NumericMode: 4134	, cons.AlphanumericMode:2506, cons.ByteMode:1722, cons.KanjiMode:1060}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 14, 46, 23, 47}}, cons.Q:{VERSION34, cons.Q,1231	,9848	,map[string]int{cons.NumericMode: 2949	, cons.AlphanumericMode:1787, cons.ByteMode:1228, cons.KanjiMode:756}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 44, 24, 7, 25}}, cons.H:{VERSION34	, cons.H,961	,7688	,map[string]int{cons.NumericMode: 2301	, cons.AlphanumericMode:1394, cons.ByteMode:958	, cons.KanjiMode:590},&ErrorCorrectionBlockCapacity{cons.H, 30, 59, 16, 1, 17}}},
	VERSION35	:{cons.L: {VERSION35	, cons.L,2306	,18448	,map[string]int{cons.NumericMode: 5529	, cons.AlphanumericMode:3351, cons.ByteMode:2303, cons.KanjiMode:1417},&ErrorCorrectionBlockCapacity{cons.L, 30, 12, 121, 7, 122}},  cons.M:{VERSION35, cons.M,1812	,14496	,map[string]int{cons.NumericMode: 4343	, cons.AlphanumericMode:2632, cons.ByteMode:1809, cons.KanjiMode:1113}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 12, 47, 26, 48}}, cons.Q:{VERSION35, cons.Q,1286	,10288	,map[string]int{cons.NumericMode: 3081	, cons.AlphanumericMode:1867, cons.ByteMode:1283, cons.KanjiMode:790}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 39, 24, 14, 25}}, cons.H:{VERSION35	, cons.H,986	,7888	,map[string]int{cons.NumericMode: 2361	, cons.AlphanumericMode:1431, cons.ByteMode:983	, cons.KanjiMode:605},&ErrorCorrectionBlockCapacity{cons.H, 30, 22, 15, 41, 16}}},
	VERSION36	:{cons.L: {VERSION36	, cons.L,2434	,19472	,map[string]int{cons.NumericMode: 5836	, cons.AlphanumericMode:3537, cons.ByteMode:2431, cons.KanjiMode:1496},&ErrorCorrectionBlockCapacity{cons.L, 30, 6, 121, 14, 122}},  cons.M:{VERSION36, cons.M,1914	,15312	,map[string]int{cons.NumericMode: 4588	, cons.AlphanumericMode:2780, cons.ByteMode:1911, cons.KanjiMode:1176}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 6, 47, 34, 48}}, cons.Q:{VERSION36, cons.Q,1354	,10832	,map[string]int{cons.NumericMode: 3244	, cons.AlphanumericMode:1966, cons.ByteMode:1351, cons.KanjiMode:832}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 46, 24, 10, 25}}, cons.H:{VERSION36	, cons.H,1054	,8432	,map[string]int{cons.NumericMode: 2524	, cons.AlphanumericMode:1530, cons.ByteMode:1051, cons.KanjiMode:647},&ErrorCorrectionBlockCapacity{cons.H, 30, 2, 15, 64, 16}}},
	VERSION37	:{cons.L: {VERSION37	, cons.L,2566	,20528	,map[string]int{cons.NumericMode: 6153	, cons.AlphanumericMode:3729, cons.ByteMode:2563, cons.KanjiMode:1577},&ErrorCorrectionBlockCapacity{cons.L, 30, 17, 122, 4, 123}},  cons.M:{VERSION37, cons.M,1992	,15936	,map[string]int{cons.NumericMode: 4775	, cons.AlphanumericMode:2894, cons.ByteMode:1989, cons.KanjiMode:1224}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 29, 46, 14, 47}}, cons.Q:{VERSION37, cons.Q,1426	,11408	,map[string]int{cons.NumericMode: 3417	, cons.AlphanumericMode:2071, cons.ByteMode:1423, cons.KanjiMode:876}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 49, 24, 10, 25}}, cons.H:{VERSION37	, cons.H,1096	,8768	,map[string]int{cons.NumericMode: 2625	, cons.AlphanumericMode:1591, cons.ByteMode:1093, cons.KanjiMode:673},&ErrorCorrectionBlockCapacity{cons.H, 30, 24, 15, 46, 16}}},
	VERSION38	:{cons.L: {VERSION38	, cons.L,2702	,21616	,map[string]int{cons.NumericMode: 6479	, cons.AlphanumericMode:3927, cons.ByteMode:2699, cons.KanjiMode:1661},&ErrorCorrectionBlockCapacity{cons.L, 30, 4, 122, 18, 123}},  cons.M:{VERSION38, cons.M,2102	,16816	,map[string]int{cons.NumericMode: 5039	, cons.AlphanumericMode:3054, cons.ByteMode:2099, cons.KanjiMode:1292}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 13, 46, 32, 47}}, cons.Q:{VERSION38, cons.Q,1502	,12016	,map[string]int{cons.NumericMode: 3599	, cons.AlphanumericMode:2181, cons.ByteMode:1499, cons.KanjiMode:923}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 48, 24, 14, 25}}, cons.H:{VERSION38	, cons.H,1142	,9136	,map[string]int{cons.NumericMode: 2735	, cons.AlphanumericMode:1658, cons.ByteMode:1139, cons.KanjiMode:701},&ErrorCorrectionBlockCapacity{cons.H, 30, 42, 15, 32, 16}}},
	VERSION39	:{cons.L: {VERSION39	, cons.L,2812	,22496	,map[string]int{cons.NumericMode: 6743	, cons.AlphanumericMode:4087, cons.ByteMode:2809, cons.KanjiMode:1729},&ErrorCorrectionBlockCapacity{cons.L, 30, 20, 117, 4, 118}},  cons.M:{VERSION39, cons.M,2216	,17728	,map[string]int{cons.NumericMode: 5313	, cons.AlphanumericMode:3220, cons.ByteMode:2213, cons.KanjiMode:1362}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 40, 47, 7, 48}}, cons.Q:{VERSION39, cons.Q,1582	,12656	,map[string]int{cons.NumericMode: 3791	, cons.AlphanumericMode:2298, cons.ByteMode:1579, cons.KanjiMode:972}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 43, 24, 22, 25}}, cons.H:{VERSION39	, cons.H,1222	,9776	,map[string]int{cons.NumericMode: 2927	, cons.AlphanumericMode:1774, cons.ByteMode:1219, cons.KanjiMode:750},&ErrorCorrectionBlockCapacity{cons.H, 30, 10, 15, 67, 16}}},
	VERSION40	:{cons.L: {VERSION40	, cons.L,2956	,23648	,map[string]int{cons.NumericMode: 7089	, cons.AlphanumericMode:4296, cons.ByteMode:2953, cons.KanjiMode:1817},&ErrorCorrectionBlockCapacity{cons.L, 30, 19, 118, 6, 119}},  cons.M:{VERSION40, cons.M,2334	,18672	,map[string]int{cons.NumericMode: 5596	, cons.AlphanumericMode:3391, cons.ByteMode:2331, cons.KanjiMode:1435}		,&ErrorCorrectionBlockCapacity{cons.M, 28, 18, 47, 31, 48}}, cons.Q:{VERSION40, cons.Q,1666	,13328	,map[string]int{cons.NumericMode: 3993	, cons.AlphanumericMode:2420, cons.ByteMode:1663, cons.KanjiMode:1024}		,&ErrorCorrectionBlockCapacity{cons.Q, 30, 34, 24, 34, 25}}, cons.H:{VERSION40	, cons.H,1276	,10208	,map[string]int{cons.NumericMode: 3057	, cons.AlphanumericMode:1852, cons.ByteMode:1273, cons.KanjiMode:784},&ErrorCorrectionBlockCapacity{cons.H, 30, 20, 15, 61, 16}}},
}

func init() {
	logger.Info("Print Table 7-Number of symbol characters and input data capacity for QR Code " +
		"and Table 9-Error correction characteristics for QR Code")
	logger.Info(VersionSymbolCharsAndInputDataCapacityMap)
}

// VersionFinalCodewordCapacity :Page 26,Table 1-Codeword capacity of all versions of QR Code
// All codewords are 8 bits in length,except in version M1 and M3 where the final data codeword is 4 bit in length
type VersionFinalCodewordCapacity struct {
	Version VersionId
	// A
	NumberOfSideModules int
	// B
	FunctionPatternModules int
	// C
	FormatAndVersionInformationModules int
	// D: Data modules except (c) (D=A² - B -C)
	DataModules int
	// E: Codewords
	DataCapacityCodewords int
	RemainderBits int
}

// VersionFinalCodewordCapacityMap : The final message codeword capacity map
// Page 26,Table 1-Codeword capacity of all versions of QR Code
// Remainder Bits in {0,3,4,7}
// FormatAndVersionInformationModules in {15,31,67}
//
var VersionFinalCodewordCapacityMap = map[VersionId]*VersionFinalCodewordCapacity{
	VERSION_M1	:{VERSION_M1,11, 20		,15	,36		,5		,0},
	VERSION_M2	:{VERSION_M2,13	,74		,15	,80		,10	,0},
	VERSION_M3	:{VERSION_M3,15	,78		,15	,132	,17	,0},
	VERSION_M4	:{VERSION_M4,17	,82		,15	,192	,24	,0},
	VERSION1	:{VERSION1	,21	,202	,31	,208	,26	,0},
	VERSION2	:{VERSION2	,25	,235	,31	,359	,44	,7},
	VERSION3	:{VERSION3	,29	,243	,31	,567	,70	,7},
	VERSION4	:{VERSION4	,33	,251	,31	,807	,100	,7},
	VERSION5	:{VERSION5	,37	,259	,31	,1079	,134	,7},
	VERSION6	:{VERSION6	,41	,267	,31	,1383	,172	,7},
	VERSION7	:{VERSION7	,45	,390	,67	,1568	,196	,0},
	VERSION8	:{VERSION8	,49	,398	,67	,1936	,242	,0},
	VERSION9	:{VERSION9	,53	,406	,67	,2336	,292	,0},
	VERSION10	:{VERSION10	,57	,414	,67	,2768	,346	,0},
	VERSION11	:{VERSION11	,61	,422	,67	,3232	,406	,0},
	VERSION12	:{VERSION12	,65	,430	,67	,3728	,466	,0},
	VERSION13	:{VERSION13	,69	,438	,67	,4256	,532	,0},
	VERSION14	:{VERSION14	,73	,611	,67	,4651	,581	,3},
	VERSION15	:{VERSION15	,77	,619	,67	,5243	,655	,3},
	VERSION16	:{VERSION16	,81	,627	,67	,5867	,733	,3},
	VERSION17	:{VERSION17	,85	,635	,67	,6523	,815	,3},
	VERSION18	:{VERSION18	,89	,643	,67	,7211	,901	,3},
	VERSION19	:{VERSION19	,93	,651	,67	,7931	,991	,3},
	VERSION20	:{VERSION20	,97	,659	,67	,8683	,1085	,3},
	VERSION21	:{VERSION21	,101	,882	,67	,9252	,1156	,4},
	VERSION22	:{VERSION22	,105	,890	,67	,10068	,1258	,4},
	VERSION23	:{VERSION23	,109	,898	,67	,10916	,1364	,4},
	VERSION24	:{VERSION24	,113	,906	,67	,11796	,1474	,4},
	VERSION25	:{VERSION25	,117	,914	,67	,12708	,1588	,4},
	VERSION26	:{VERSION26	,121	,922	,67	,13652	,1706	,4},
	VERSION27	:{VERSION27	,125	,930	,67	,14628	,1828	,4},
	VERSION28	:{VERSION28	,129	,1203	,67	,15371	,1921	,3},
	VERSION29	:{VERSION29	,133	,1211	,67	,16411	,2051	,3},
	VERSION30	:{VERSION30	,137	,1219	,67	,17483	,2185	,3},
	VERSION31	:{VERSION31	,141	,1227	,67	,18587	,2323	,3},
	VERSION32	:{VERSION32	,145	,1235	,67	,19723	,2465	,3},
	VERSION33	:{VERSION33	,149	,1243	,67	,20891	,2611	,3},
	VERSION34	:{VERSION34	,153	,1251	,67	,22091	,2761	,3},
	VERSION35	:{VERSION35	,157	,1574	,67	,23008	,2876	,0},
	VERSION36	:{VERSION36	,161	,1582	,67	,24272	,3034	,0},
	VERSION37	:{VERSION37	,165	,1590	,67	,25568	,3196	,0},
	VERSION38	:{VERSION38	,169	,1598	,67	,26896	,3362	,0},
	VERSION39	:{VERSION39	,173	,1606	,67	,28256	,3532	,0},
	VERSION40	:{VERSION40	,177	,1614	,67	,29648	,3706	,0},
}

func GetVersionByInputDataLength(format cons.Format,dataLen int,mode cons.ModeType,level cons.ErrorCorrectionLevel) (*Version, cons.ErrorCorrectionLevel){
	maxVId := VERSION40
	minVId := VERSION1
	if format == cons.MicroQrcode {
		minVId = VERSION_M4
		maxVId = VERSION1
	}
	for vId := minVId; vId<= maxVId ; vId++{
		if ecMap,ok:= VersionSymbolCharsAndInputDataCapacityMap[vId];ok{
			for ecLevel:= cons.H ;ecLevel > 0; ecLevel--  {
				if capacity,ok := ecMap[ecLevel]; ok{
					if size,ok:=capacity.DataCapacity[mode];(level == ecLevel) || (ok && size >= dataLen) {
						return NewVersion(vId),ecLevel
					}else{
						if vId == maxVId && ecLevel == cons.L && size < dataLen{
							panic(errors.New("input data is to long for Mode: "+mode))
						}
					}
				}
			}
		}
	}
	panic(errors.New("can not find DataCapacity by mode:"+mode))
}

// GetVersionSymbolCharsAndInputDataCapacity :Get number of symbol characters and input data capacity for QR Code
func (v *Version) GetVersionSymbolCharsAndInputDataCapacity(ecLevel cons.ErrorCorrectionLevel) *VersionSymbolCharsAndInputDataCapacity{
	versionDataCapacity,ok :=VersionSymbolCharsAndInputDataCapacityMap[v.Id]
	if !ok {
		err := fmt.Errorf("can not find version data capacity in VersionDataCapacityMap by version: %s",v.Name)
		logger.Error("GetVersionDataCapacity: find version data capacity by version error: "+err.Error())
		panic(err)
	}
	if dataCapacity, ok := versionDataCapacity[ecLevel]; ok {
		return dataCapacity
	}else{
		err := fmt.Errorf("can not find version data capacity in VersionDataCapacityMap by version: %s, error correction: %d",v.Name,ecLevel)
		logger.Error("GetVersionDataCapacity: find version data capacity by errorCorrectionLevel error: "+err.Error())
		panic(err)
	}
	return nil
}

// GetVersionFinalCodewordCapacity : Get version codeword capacity by version.
func (v *Version) GetVersionFinalCodewordCapacity() *VersionFinalCodewordCapacity{
	if capacity,ok:= VersionFinalCodewordCapacityMap[v.Id];ok{
		return capacity
	}else{
		err := fmt.Errorf("can not find version codeword  capacity in VersionCodewordCapacityMap by version: %s",v.Name)
		logger.Error("GetVersionDataCapacity: find version data capacity by errorCorrectionLevel error: "+err.Error())
		panic(err)
	}
	return nil
}

// GetModuleSize : get module size of version,e.g.: Version 1 return 21
// QRCode size:
// Micro QRCode symbols: 11 x 11 modules to 17 x 17 modules(Versions M1 to M4,Increasing in steps of two modules per side)
// QRCode symbols: 21 x 21 modules to 177 x 177 modules(Versions 1 to 40,Increasing in steps of four modules per side)
func (v *Version) GetModuleSize() int{
	size := 0
	if v.Id >0 {
		size = (v.Id - 1) * 4 + 21
	}else{
		size = (-v.Id - 1) * 2 + 11
	}
	return size
}

func (v *Version) GetDefaultPixelSize() int{
	return v.GetModuleSize() * cons.DefaultPixelSizePerModule
}
