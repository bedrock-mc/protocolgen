// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type SubChunkRequestResult uint8

const (
	SubChunkRequestResultSuccess               SubChunkRequestResult = 1
	SubChunkRequestResultLevelChunkDoesntExist SubChunkRequestResult = 2
	SubChunkRequestResultWrongDimension        SubChunkRequestResult = 3
	SubChunkRequestResultPlayerDoesntExist     SubChunkRequestResult = 4
	SubChunkRequestResultIndexOutOfBounds      SubChunkRequestResult = 5
	SubChunkRequestResultSuccessAllAir         SubChunkRequestResult = 6
)
