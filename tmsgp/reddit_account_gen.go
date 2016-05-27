package tmsgp

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *AccountData) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "comment_karma":
			z.CommentKarma, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "has_mail":
			z.HasMail, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "has_mod_mail":
			z.HasModMail, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "id":
			z.ID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "inbox_count":
			z.InboxCount, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "is_friend":
			z.IsFriend, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "is_gold":
			z.IsGold, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "link_karma":
			z.LinkKarma, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "mod_hash":
			z.ModHash, err = dc.ReadString()
			if err != nil {
				return
			}
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "over_18":
			z.Over18, err = dc.ReadBool()
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
func (z *AccountData) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 11
	// write "comment_karma"
	err = en.Append(0x8b, 0xad, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x61, 0x72, 0x6d, 0x61)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.CommentKarma)
	if err != nil {
		return
	}
	// write "has_mail"
	err = en.Append(0xa8, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x61, 0x69, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.HasMail)
	if err != nil {
		return
	}
	// write "has_mod_mail"
	err = en.Append(0xac, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x5f, 0x6d, 0x61, 0x69, 0x6c)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.HasModMail)
	if err != nil {
		return
	}
	// write "id"
	err = en.Append(0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ID)
	if err != nil {
		return
	}
	// write "inbox_count"
	err = en.Append(0xab, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.InboxCount)
	if err != nil {
		return
	}
	// write "is_friend"
	err = en.Append(0xa9, 0x69, 0x73, 0x5f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.IsFriend)
	if err != nil {
		return
	}
	// write "is_gold"
	err = en.Append(0xa7, 0x69, 0x73, 0x5f, 0x67, 0x6f, 0x6c, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.IsGold)
	if err != nil {
		return
	}
	// write "link_karma"
	err = en.Append(0xaa, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x6b, 0x61, 0x72, 0x6d, 0x61)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.LinkKarma)
	if err != nil {
		return
	}
	// write "mod_hash"
	err = en.Append(0xa8, 0x6d, 0x6f, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ModHash)
	if err != nil {
		return
	}
	// write "name"
	err = en.Append(0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "over_18"
	err = en.Append(0xa7, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x31, 0x38)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Over18)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *AccountData) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 11
	// string "comment_karma"
	o = append(o, 0x8b, 0xad, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x61, 0x72, 0x6d, 0x61)
	o = msgp.AppendInt64(o, z.CommentKarma)
	// string "has_mail"
	o = append(o, 0xa8, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x61, 0x69, 0x6c)
	o = msgp.AppendBool(o, z.HasMail)
	// string "has_mod_mail"
	o = append(o, 0xac, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x5f, 0x6d, 0x61, 0x69, 0x6c)
	o = msgp.AppendBool(o, z.HasModMail)
	// string "id"
	o = append(o, 0xa2, 0x69, 0x64)
	o = msgp.AppendString(o, z.ID)
	// string "inbox_count"
	o = append(o, 0xab, 0x69, 0x6e, 0x62, 0x6f, 0x78, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendInt64(o, z.InboxCount)
	// string "is_friend"
	o = append(o, 0xa9, 0x69, 0x73, 0x5f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64)
	o = msgp.AppendBool(o, z.IsFriend)
	// string "is_gold"
	o = append(o, 0xa7, 0x69, 0x73, 0x5f, 0x67, 0x6f, 0x6c, 0x64)
	o = msgp.AppendBool(o, z.IsGold)
	// string "link_karma"
	o = append(o, 0xaa, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x6b, 0x61, 0x72, 0x6d, 0x61)
	o = msgp.AppendInt64(o, z.LinkKarma)
	// string "mod_hash"
	o = append(o, 0xa8, 0x6d, 0x6f, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68)
	o = msgp.AppendString(o, z.ModHash)
	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "over_18"
	o = append(o, 0xa7, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x31, 0x38)
	o = msgp.AppendBool(o, z.Over18)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AccountData) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "comment_karma":
			z.CommentKarma, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "has_mail":
			z.HasMail, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "has_mod_mail":
			z.HasModMail, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "id":
			z.ID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "inbox_count":
			z.InboxCount, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "is_friend":
			z.IsFriend, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "is_gold":
			z.IsGold, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "link_karma":
			z.LinkKarma, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "mod_hash":
			z.ModHash, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "over_18":
			z.Over18, bts, err = msgp.ReadBoolBytes(bts)
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

func (z *AccountData) Msgsize() (s int) {
	s = 1 + 14 + msgp.Int64Size + 9 + msgp.BoolSize + 13 + msgp.BoolSize + 3 + msgp.StringPrefixSize + len(z.ID) + 12 + msgp.Int64Size + 10 + msgp.BoolSize + 8 + msgp.BoolSize + 11 + msgp.Int64Size + 9 + msgp.StringPrefixSize + len(z.ModHash) + 5 + msgp.StringPrefixSize + len(z.Name) + 8 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *RedditAccount) DecodeMsg(dc *msgp.Reader) (err error) {
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
			z.ID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "kind":
			z.Kind, err = dc.ReadString()
			if err != nil {
				return
			}
		case "data":
			err = z.Data.DecodeMsg(dc)
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
func (z *RedditAccount) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "id"
	err = en.Append(0x84, 0xa2, 0x69, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ID)
	if err != nil {
		return
	}
	// write "name"
	err = en.Append(0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "kind"
	err = en.Append(0xa4, 0x6b, 0x69, 0x6e, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Kind)
	if err != nil {
		return
	}
	// write "data"
	err = en.Append(0xa4, 0x64, 0x61, 0x74, 0x61)
	if err != nil {
		return err
	}
	err = z.Data.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *RedditAccount) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "id"
	o = append(o, 0x84, 0xa2, 0x69, 0x64)
	o = msgp.AppendString(o, z.ID)
	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "kind"
	o = append(o, 0xa4, 0x6b, 0x69, 0x6e, 0x64)
	o = msgp.AppendString(o, z.Kind)
	// string "data"
	o = append(o, 0xa4, 0x64, 0x61, 0x74, 0x61)
	o, err = z.Data.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RedditAccount) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
			z.ID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "kind":
			z.Kind, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "data":
			bts, err = z.Data.UnmarshalMsg(bts)
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

func (z *RedditAccount) Msgsize() (s int) {
	s = 1 + 3 + msgp.StringPrefixSize + len(z.ID) + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.StringPrefixSize + len(z.Kind) + 5 + z.Data.Msgsize()
	return
}
