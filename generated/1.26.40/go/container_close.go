// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

import "fmt"

type ContainerClose struct {
	ContainerId          uint8
	ContainerType        uint8
	ServerInitiatedClose bool
}

func (p *ContainerClose) Encode(w Encoder) error {
	if err := w.Write("ContainerClosePacket.Container Id", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.ContainerId); err != nil {
		return err
	}
	if err := w.Write("ContainerClosePacket.Container Type", Shape{Kind: "primitive", PrimitiveCode: "u8"}, p.ContainerType); err != nil {
		return err
	}
	if err := w.Write("ContainerClosePacket.Server Initiated Close", Shape{Kind: "primitive", PrimitiveCode: "bool"}, p.ServerInitiatedClose); err != nil {
		return err
	}
	return nil
}

func DecodeContainerClose(r Decoder) (ContainerClose, error) {
	var p ContainerClose
	{
		raw, err := r.Read("ContainerClosePacket.Container Id", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field ContainerClosePacket.Container Id has unexpected decoded type %T", raw)
		}
		p.ContainerId = value
	}
	{
		raw, err := r.Read("ContainerClosePacket.Container Type", Shape{Kind: "primitive", PrimitiveCode: "u8"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(uint8)
		if !ok {
			return p, fmt.Errorf("field ContainerClosePacket.Container Type has unexpected decoded type %T", raw)
		}
		p.ContainerType = value
	}
	{
		raw, err := r.Read("ContainerClosePacket.Server Initiated Close", Shape{Kind: "primitive", PrimitiveCode: "bool"})
		if err != nil {
			return p, err
		}
		value, ok := raw.(bool)
		if !ok {
			return p, fmt.Errorf("field ContainerClosePacket.Server Initiated Close has unexpected decoded type %T", raw)
		}
		p.ServerInitiatedClose = value
	}
	return p, nil
}
