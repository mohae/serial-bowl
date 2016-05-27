package tmsgp

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Message) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.ID, err = dc.ReadBytes(z.ID)
			if err != nil {
				return
			}
		case "dest_id":
			z.DestID, err = dc.ReadUint32()
			if err != nil {
				return
			}
		case "type":
			z.Type, err = dc.ReadInt8()
			if err != nil {
				return
			}
		case "kind":
			z.Kind, err = dc.ReadInt16()
			if err != nil {
				return
			}
		case "data":
			z.Data, err = dc.ReadBytes(z.Data)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Message) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "id"
	err = en.Append(0x85, 0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.ID)
	if err != nil {
		return
	}
	// write "dest_id"
	err = en.Append(0xa7, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteUint32(z.DestID)
	if err != nil {
		return
	}
	// write "type"
	err = en.Append(0xa4, 0x74, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt8(z.Type)
	if err != nil {
		return
	}
	// write "kind"
	err = en.Append(0xa4, 0x6b, 0x69, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt16(z.Kind)
	if err != nil {
		return
	}
	// write "data"
	err = en.Append(0xa4, 0x64, 0x61, 0x74, 0x61)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Data)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Message) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "id"
	o = append(o, 0x85, 0xa2, 0x69, 0x64)
	o = msgp.AppendBytes(o, z.ID)
	// string "dest_id"
	o = append(o, 0xa7, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64)
	o = msgp.AppendUint32(o, z.DestID)
	// string "type"
	o = append(o, 0xa4, 0x74, 0x79, 0x70, 0x65)
	o = msgp.AppendInt8(o, z.Type)
	// string "kind"
	o = append(o, 0xa4, 0x6b, 0x69, 0x6e, 0x64)
	o = msgp.AppendInt16(o, z.Kind)
	// string "data"
	o = append(o, 0xa4, 0x64, 0x61, 0x74, 0x61)
	o = msgp.AppendBytes(o, z.Data)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Message) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "id":
			z.ID, bts, err = msgp.ReadBytesBytes(bts, z.ID)
			if err != nil {
				return
			}
		case "dest_id":
			z.DestID, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				return
			}
		case "type":
			z.Type, bts, err = msgp.ReadInt8Bytes(bts)
			if err != nil {
				return
			}
		case "kind":
			z.Kind, bts, err = msgp.ReadInt16Bytes(bts)
			if err != nil {
				return
			}
		case "data":
			z.Data, bts, err = msgp.ReadBytesBytes(bts, z.Data)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *Message) Msgsize() (s int) {
	s = 1 + 3 + msgp.BytesPrefixSize + len(z.ID) + 8 + msgp.Uint32Size + 5 + msgp.Int8Size + 5 + msgp.Int16Size + 5 + msgp.BytesPrefixSize + len(z.Data)
	return
}
