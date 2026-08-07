// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type PlayStatus struct {
	Status PlayStatusType
}

func (p *PlayStatus) Encode(w Encoder) error {
	if err := w.Write("PlayStatusPacket.Status", Shape{Kind: "enum", Semantic: "PlayStatus", TypeID: "enums/PlayStatus", PrimitiveCode: "i32be", Variants: []ShapeVariant{{Value: 0, Name: "LoginSuccess", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "LoginFailed_ClientOld", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "LoginFailed_ServerOld", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "PlayerSpawn", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "LoginFailed_InvalidTenant", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "LoginFailed_EditionMismatchEduToVanilla", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "LoginFailed_EditionMismatchVanillaToEdu", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "LoginFailed_ServerFullSubClient", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "LoginFailed_EditorMismatchEditorToVanilla", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "LoginFailed_EditorMismatchVanillaToEditor", Shape: Shape{Kind: "void"}}}}, p.Status); err != nil {
		return err
	}
	return nil
}

func DecodePlayStatus(r Decoder) (PlayStatus, error) {
	var p PlayStatus
	{
		raw, err := r.Read("PlayStatusPacket.Status", Shape{Kind: "enum", Semantic: "PlayStatus", TypeID: "enums/PlayStatus", PrimitiveCode: "i32be", Variants: []ShapeVariant{{Value: 0, Name: "LoginSuccess", Shape: Shape{Kind: "void"}}, {Value: 1, Name: "LoginFailed_ClientOld", Shape: Shape{Kind: "void"}}, {Value: 2, Name: "LoginFailed_ServerOld", Shape: Shape{Kind: "void"}}, {Value: 3, Name: "PlayerSpawn", Shape: Shape{Kind: "void"}}, {Value: 4, Name: "LoginFailed_InvalidTenant", Shape: Shape{Kind: "void"}}, {Value: 5, Name: "LoginFailed_EditionMismatchEduToVanilla", Shape: Shape{Kind: "void"}}, {Value: 6, Name: "LoginFailed_EditionMismatchVanillaToEdu", Shape: Shape{Kind: "void"}}, {Value: 7, Name: "LoginFailed_ServerFullSubClient", Shape: Shape{Kind: "void"}}, {Value: 8, Name: "LoginFailed_EditorMismatchEditorToVanilla", Shape: Shape{Kind: "void"}}, {Value: 9, Name: "LoginFailed_EditorMismatchVanillaToEditor", Shape: Shape{Kind: "void"}}}})
		if err != nil {
			return p, err
		}
		value, ok := raw.(PlayStatusType)
		if !ok {
			return p, fmt.Errorf("field PlayStatusPacket.Status has unexpected decoded type %T", raw)
		}
		p.Status = value
	}
	return p, nil
}
