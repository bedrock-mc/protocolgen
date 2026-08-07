// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type RequestPermissions struct {
	TargetPlayerIdSRawID  int64
	PlayerPermissionLevel int32
	CustomPermissionFlags uint16
}

func (p *RequestPermissions) Encode(w Encoder) error {
	if err := w.Write("RequestPermissionsPacket.Target Player Id's Raw ID", Shape{Kind: "primitive", PrimitiveCode: "i64le"}, p.TargetPlayerIdSRawID); err != nil {
		return err
	}
	if err := w.Write("RequestPermissionsPacket.Player Permission Level", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"}, p.PlayerPermissionLevel); err != nil {
		return err
	}
	if err := w.Write("RequestPermissionsPacket.Custom Permission Flags", Shape{Kind: "primitive", PrimitiveCode: "u16le"}, p.CustomPermissionFlags); err != nil {
		return err
	}
	return nil
}

func DecodeRequestPermissions(r Decoder) (RequestPermissions, error) {
	var p RequestPermissions
	{
		raw, err := r.Read("RequestPermissionsPacket.Target Player Id's Raw ID", Shape{Kind: "primitive", PrimitiveCode: "i64le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int64)
		if !ok {
			return p, fmt.Errorf("field RequestPermissionsPacket.Target Player Id's Raw ID has unexpected decoded type %T", raw)
		}
		p.TargetPlayerIdSRawID = value
	}
	{
		raw, err := r.Read("RequestPermissionsPacket.Player Permission Level", Shape{Kind: "primitive", PrimitiveCode: "zigzag_i32"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(int32)
		if !ok {
			return p, fmt.Errorf("field RequestPermissionsPacket.Player Permission Level has unexpected decoded type %T", raw)
		}
		p.PlayerPermissionLevel = value
	}
	{
		raw, err := r.Read("RequestPermissionsPacket.Custom Permission Flags", Shape{Kind: "primitive", PrimitiveCode: "u16le"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint16)
		if !ok {
			return p, fmt.Errorf("field RequestPermissionsPacket.Custom Permission Flags has unexpected decoded type %T", raw)
		}
		p.CustomPermissionFlags = value
	}
	return p, nil
}
