// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type RequestType uint8

const (
	RequestTypeSetActions             RequestType = 0
	RequestTypeExecuteAction          RequestType = 1
	RequestTypeExecuteClosingCommands RequestType = 2
	RequestTypeSetName                RequestType = 3
	RequestTypeSetSkin                RequestType = 4
	RequestTypeSetInteractText        RequestType = 5
	RequestTypeExecuteOpeningCommands RequestType = 6
)
