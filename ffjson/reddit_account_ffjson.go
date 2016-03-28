// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: reddit_account.go
// DO NOT EDIT!

package ffjson

import (
	"bytes"
	"errors"
	"fmt"

	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *AccountData) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *AccountData) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"comment_karma":`)
	fflib.FormatBits2(buf, uint64(mj.CommentKarma), 10, mj.CommentKarma < 0)
	if mj.HasMail {
		buf.WriteString(`,"has_mail":true`)
	} else {
		buf.WriteString(`,"has_mail":false`)
	}
	if mj.HasModMail {
		buf.WriteString(`,"has_mod_mail":true`)
	} else {
		buf.WriteString(`,"has_mod_mail":false`)
	}
	buf.WriteString(`,"id":`)
	fflib.WriteJsonString(buf, string(mj.ID))
	buf.WriteString(`,"inbox_count":`)
	fflib.FormatBits2(buf, uint64(mj.InboxCount), 10, mj.InboxCount < 0)
	if mj.IsFriend {
		buf.WriteString(`,"is_friend":true`)
	} else {
		buf.WriteString(`,"is_friend":false`)
	}
	if mj.IsGold {
		buf.WriteString(`,"is_gold":true`)
	} else {
		buf.WriteString(`,"is_gold":false`)
	}
	buf.WriteString(`,"link_karma":`)
	fflib.FormatBits2(buf, uint64(mj.LinkKarma), 10, mj.LinkKarma < 0)
	buf.WriteString(`,"mod_hash":`)
	fflib.WriteJsonString(buf, string(mj.ModHash))
	buf.WriteString(`,"name":`)
	fflib.WriteJsonString(buf, string(mj.Name))
	if mj.Over18 {
		buf.WriteString(`,"over_18":true`)
	} else {
		buf.WriteString(`,"over_18":false`)
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_AccountDatabase = iota
	ffj_t_AccountDatano_such_key

	ffj_t_AccountData_CommentKarma

	ffj_t_AccountData_HasMail

	ffj_t_AccountData_HasModMail

	ffj_t_AccountData_ID

	ffj_t_AccountData_InboxCount

	ffj_t_AccountData_IsFriend

	ffj_t_AccountData_IsGold

	ffj_t_AccountData_LinkKarma

	ffj_t_AccountData_ModHash

	ffj_t_AccountData_Name

	ffj_t_AccountData_Over18
)

var ffj_key_AccountData_CommentKarma = []byte("comment_karma")

var ffj_key_AccountData_HasMail = []byte("has_mail")

var ffj_key_AccountData_HasModMail = []byte("has_mod_mail")

var ffj_key_AccountData_ID = []byte("id")

var ffj_key_AccountData_InboxCount = []byte("inbox_count")

var ffj_key_AccountData_IsFriend = []byte("is_friend")

var ffj_key_AccountData_IsGold = []byte("is_gold")

var ffj_key_AccountData_LinkKarma = []byte("link_karma")

var ffj_key_AccountData_ModHash = []byte("mod_hash")

var ffj_key_AccountData_Name = []byte("name")

var ffj_key_AccountData_Over18 = []byte("over_18")

func (uj *AccountData) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *AccountData) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_AccountDatabase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_AccountDatano_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffj_key_AccountData_CommentKarma, kn) {
						currentKey = ffj_t_AccountData_CommentKarma
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'h':

					if bytes.Equal(ffj_key_AccountData_HasMail, kn) {
						currentKey = ffj_t_AccountData_HasMail
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_AccountData_HasModMail, kn) {
						currentKey = ffj_t_AccountData_HasModMail
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffj_key_AccountData_ID, kn) {
						currentKey = ffj_t_AccountData_ID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_AccountData_InboxCount, kn) {
						currentKey = ffj_t_AccountData_InboxCount
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_AccountData_IsFriend, kn) {
						currentKey = ffj_t_AccountData_IsFriend
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_AccountData_IsGold, kn) {
						currentKey = ffj_t_AccountData_IsGold
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffj_key_AccountData_LinkKarma, kn) {
						currentKey = ffj_t_AccountData_LinkKarma
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffj_key_AccountData_ModHash, kn) {
						currentKey = ffj_t_AccountData_ModHash
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffj_key_AccountData_Name, kn) {
						currentKey = ffj_t_AccountData_Name
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffj_key_AccountData_Over18, kn) {
						currentKey = ffj_t_AccountData_Over18
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.AsciiEqualFold(ffj_key_AccountData_Over18, kn) {
					currentKey = ffj_t_AccountData_Over18
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_AccountData_Name, kn) {
					currentKey = ffj_t_AccountData_Name
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_AccountData_ModHash, kn) {
					currentKey = ffj_t_AccountData_ModHash
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_AccountData_LinkKarma, kn) {
					currentKey = ffj_t_AccountData_LinkKarma
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_AccountData_IsGold, kn) {
					currentKey = ffj_t_AccountData_IsGold
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_AccountData_IsFriend, kn) {
					currentKey = ffj_t_AccountData_IsFriend
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_AccountData_InboxCount, kn) {
					currentKey = ffj_t_AccountData_InboxCount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_AccountData_ID, kn) {
					currentKey = ffj_t_AccountData_ID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_AccountData_HasModMail, kn) {
					currentKey = ffj_t_AccountData_HasModMail
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_AccountData_HasMail, kn) {
					currentKey = ffj_t_AccountData_HasMail
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_AccountData_CommentKarma, kn) {
					currentKey = ffj_t_AccountData_CommentKarma
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_AccountDatano_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_AccountData_CommentKarma:
					goto handle_CommentKarma

				case ffj_t_AccountData_HasMail:
					goto handle_HasMail

				case ffj_t_AccountData_HasModMail:
					goto handle_HasModMail

				case ffj_t_AccountData_ID:
					goto handle_ID

				case ffj_t_AccountData_InboxCount:
					goto handle_InboxCount

				case ffj_t_AccountData_IsFriend:
					goto handle_IsFriend

				case ffj_t_AccountData_IsGold:
					goto handle_IsGold

				case ffj_t_AccountData_LinkKarma:
					goto handle_LinkKarma

				case ffj_t_AccountData_ModHash:
					goto handle_ModHash

				case ffj_t_AccountData_Name:
					goto handle_Name

				case ffj_t_AccountData_Over18:
					goto handle_Over18

				case ffj_t_AccountDatano_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_CommentKarma:

	/* handler: uj.CommentKarma type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.CommentKarma = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_HasMail:

	/* handler: uj.HasMail type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.HasMail = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.HasMail = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_HasModMail:

	/* handler: uj.HasModMail type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.HasModMail = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.HasModMail = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ID:

	/* handler: uj.ID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.ID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_InboxCount:

	/* handler: uj.InboxCount type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.InboxCount = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_IsFriend:

	/* handler: uj.IsFriend type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.IsFriend = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.IsFriend = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_IsGold:

	/* handler: uj.IsGold type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.IsGold = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.IsGold = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LinkKarma:

	/* handler: uj.LinkKarma type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.LinkKarma = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ModHash:

	/* handler: uj.ModHash type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.ModHash = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Name:

	/* handler: uj.Name type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Name = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Over18:

	/* handler: uj.Over18 type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.Over18 = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.Over18 = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}

func (mj *ShRedditAccount) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *ShRedditAccount) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"id":`)
	fflib.WriteJsonString(buf, string(mj.ID))
	buf.WriteString(`,"name":`)
	fflib.WriteJsonString(buf, string(mj.Name))
	buf.WriteString(`,"kind":`)
	fflib.WriteJsonString(buf, string(mj.Kind))
	buf.WriteString(`,"data":`)

	{

		err = mj.Data.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_ShRedditAccountbase = iota
	ffj_t_ShRedditAccountno_such_key

	ffj_t_ShRedditAccount_ID

	ffj_t_ShRedditAccount_Name

	ffj_t_ShRedditAccount_Kind

	ffj_t_ShRedditAccount_Data
)

var ffj_key_ShRedditAccount_ID = []byte("id")

var ffj_key_ShRedditAccount_Name = []byte("name")

var ffj_key_ShRedditAccount_Kind = []byte("kind")

var ffj_key_ShRedditAccount_Data = []byte("data")

func (uj *ShRedditAccount) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *ShRedditAccount) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_ShRedditAccountbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_ShRedditAccountno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'd':

					if bytes.Equal(ffj_key_ShRedditAccount_Data, kn) {
						currentKey = ffj_t_ShRedditAccount_Data
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffj_key_ShRedditAccount_ID, kn) {
						currentKey = ffj_t_ShRedditAccount_ID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'k':

					if bytes.Equal(ffj_key_ShRedditAccount_Kind, kn) {
						currentKey = ffj_t_ShRedditAccount_Kind
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffj_key_ShRedditAccount_Name, kn) {
						currentKey = ffj_t_ShRedditAccount_Name
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_ShRedditAccount_Data, kn) {
					currentKey = ffj_t_ShRedditAccount_Data
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_ShRedditAccount_Kind, kn) {
					currentKey = ffj_t_ShRedditAccount_Kind
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_ShRedditAccount_Name, kn) {
					currentKey = ffj_t_ShRedditAccount_Name
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_ShRedditAccount_ID, kn) {
					currentKey = ffj_t_ShRedditAccount_ID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_ShRedditAccountno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_ShRedditAccount_ID:
					goto handle_ID

				case ffj_t_ShRedditAccount_Name:
					goto handle_Name

				case ffj_t_ShRedditAccount_Kind:
					goto handle_Kind

				case ffj_t_ShRedditAccount_Data:
					goto handle_Data

				case ffj_t_ShRedditAccountno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_ID:

	/* handler: uj.ID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.ID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Name:

	/* handler: uj.Name type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Name = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Kind:

	/* handler: uj.Kind type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Kind = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Data:

	/* handler: uj.Data type=shared.AccountData kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			state = fflib.FFParse_after_value
			goto mainparse
		}

		err = uj.Data.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}
