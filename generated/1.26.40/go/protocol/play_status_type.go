// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

type PlayStatusType int32

const (
	PlayStatusTypeLoginSuccess                             PlayStatusType = 0
	PlayStatusTypeLoginFailedClientOld                     PlayStatusType = 1
	PlayStatusTypeLoginFailedServerOld                     PlayStatusType = 2
	PlayStatusTypePlayerSpawn                              PlayStatusType = 3
	PlayStatusTypeLoginFailedInvalidTenant                 PlayStatusType = 4
	PlayStatusTypeLoginFailedEditionMismatchEduToVanilla   PlayStatusType = 5
	PlayStatusTypeLoginFailedEditionMismatchVanillaToEdu   PlayStatusType = 6
	PlayStatusTypeLoginFailedServerFullSubClient           PlayStatusType = 7
	PlayStatusTypeLoginFailedEditorMismatchEditorToVanilla PlayStatusType = 8
	PlayStatusTypeLoginFailedEditorMismatchVanillaToEditor PlayStatusType = 9
)
