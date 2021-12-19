// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package push_models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson7aa4b9ffDecodePatreonInternalMicroservicesPushPush(in *jlexer.Lexer, out *PostPush) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "post_id":
			out.PostId = int64(in.Int64())
		case "creator_id":
			out.CreatorId = int64(in.Int64())
		case "creator_nickname":
			out.CreatorNickname = string(in.String())
		case "post_title":
			out.PostTitle = string(in.String())
		case "creator_avatar":
			out.CreatorAvatar = string(in.String())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7aa4b9ffEncodePatreonInternalMicroservicesPushPush(out *jwriter.Writer, in PostPush) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"post_id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.PostId))
	}
	{
		const prefix string = ",\"creator_id\":"
		out.RawString(prefix)
		out.Int64(int64(in.CreatorId))
	}
	{
		const prefix string = ",\"creator_nickname\":"
		out.RawString(prefix)
		out.String(string(in.CreatorNickname))
	}
	{
		const prefix string = ",\"post_title\":"
		out.RawString(prefix)
		out.String(string(in.PostTitle))
	}
	{
		const prefix string = ",\"creator_avatar\":"
		out.RawString(prefix)
		out.String(string(in.CreatorAvatar))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostPush) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7aa4b9ffEncodePatreonInternalMicroservicesPushPush(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostPush) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7aa4b9ffEncodePatreonInternalMicroservicesPushPush(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostPush) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7aa4b9ffDecodePatreonInternalMicroservicesPushPush(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostPush) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7aa4b9ffDecodePatreonInternalMicroservicesPushPush(l, v)
}
func easyjson7aa4b9ffDecodePatreonInternalMicroservicesPushPush1(in *jlexer.Lexer, out *PaymentApplyPush) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "creator_id":
			out.CreatorId = int64(in.Int64())
		case "creator_nickname":
			out.CreatorNickname = string(in.String())
		case "creator_avatar":
			out.CreatorAvatar = string(in.String())
		case "awards_id":
			out.AwardsId = int64(in.Int64())
		case "awards_name":
			out.AwardsName = string(in.String())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7aa4b9ffEncodePatreonInternalMicroservicesPushPush1(out *jwriter.Writer, in PaymentApplyPush) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"creator_id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.CreatorId))
	}
	{
		const prefix string = ",\"creator_nickname\":"
		out.RawString(prefix)
		out.String(string(in.CreatorNickname))
	}
	{
		const prefix string = ",\"creator_avatar\":"
		out.RawString(prefix)
		out.String(string(in.CreatorAvatar))
	}
	{
		const prefix string = ",\"awards_id\":"
		out.RawString(prefix)
		out.Int64(int64(in.AwardsId))
	}
	{
		const prefix string = ",\"awards_name\":"
		out.RawString(prefix)
		out.String(string(in.AwardsName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PaymentApplyPush) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7aa4b9ffEncodePatreonInternalMicroservicesPushPush1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PaymentApplyPush) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7aa4b9ffEncodePatreonInternalMicroservicesPushPush1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PaymentApplyPush) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7aa4b9ffDecodePatreonInternalMicroservicesPushPush1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PaymentApplyPush) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7aa4b9ffDecodePatreonInternalMicroservicesPushPush1(l, v)
}
func easyjson7aa4b9ffDecodePatreonInternalMicroservicesPushPush2(in *jlexer.Lexer, out *CommentPush) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "creator_id":
			out.CreatorId = int64(in.Int64())
		case "comment_id":
			out.CommentId = int64(in.Int64())
		case "post_id":
			out.PostId = int64(in.Int64())
		case "author_id":
			out.AuthorId = int64(in.Int64())
		case "author_nickname":
			out.AuthorNickname = string(in.String())
		case "author_avatar":
			out.AuthorAvatar = string(in.String())
		case "post_title":
			out.PostTitle = string(in.String())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7aa4b9ffEncodePatreonInternalMicroservicesPushPush2(out *jwriter.Writer, in CommentPush) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"creator_id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.CreatorId))
	}
	{
		const prefix string = ",\"comment_id\":"
		out.RawString(prefix)
		out.Int64(int64(in.CommentId))
	}
	{
		const prefix string = ",\"post_id\":"
		out.RawString(prefix)
		out.Int64(int64(in.PostId))
	}
	{
		const prefix string = ",\"author_id\":"
		out.RawString(prefix)
		out.Int64(int64(in.AuthorId))
	}
	{
		const prefix string = ",\"author_nickname\":"
		out.RawString(prefix)
		out.String(string(in.AuthorNickname))
	}
	{
		const prefix string = ",\"author_avatar\":"
		out.RawString(prefix)
		out.String(string(in.AuthorAvatar))
	}
	{
		const prefix string = ",\"post_title\":"
		out.RawString(prefix)
		out.String(string(in.PostTitle))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CommentPush) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7aa4b9ffEncodePatreonInternalMicroservicesPushPush2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CommentPush) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7aa4b9ffEncodePatreonInternalMicroservicesPushPush2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CommentPush) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7aa4b9ffDecodePatreonInternalMicroservicesPushPush2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CommentPush) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7aa4b9ffDecodePatreonInternalMicroservicesPushPush2(l, v)
}
