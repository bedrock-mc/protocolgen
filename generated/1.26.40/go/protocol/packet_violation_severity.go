// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PacketViolationSeverity int32

const (
	PacketViolationSeverityUnknown               PacketViolationSeverity = -1
	PacketViolationSeverityWarning               PacketViolationSeverity = 0
	PacketViolationSeverityFinalWarning          PacketViolationSeverity = 1
	PacketViolationSeverityTerminatingConnection PacketViolationSeverity = 2
)
