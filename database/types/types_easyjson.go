// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package database_types

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

func easyjson6601e8cdDecodeGithubComMachinefiW3bstreamWasmGolangSdkDatabaseTypes(in *jlexer.Lexer, out *param) {
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
		case "int32":
			out.Int32 = int32(in.Int32())
		case "int64":
			out.Int64 = int64(in.Int64())
		case "float32":
			out.Float32 = float32(in.Float32())
		case "float64":
			out.Float64 = float64(in.Float64())
		case "string":
			out.String = string(in.String())
		case "time":
			out.Time = string(in.String())
		case "bool":
			out.Bool = bool(in.Bool())
		case "bytes":
			out.Bytes = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComMachinefiW3bstreamWasmGolangSdkDatabaseTypes(out *jwriter.Writer, in param) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Int32 != 0 {
		const prefix string = ",\"int32\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.Int32))
	}
	if in.Int64 != 0 {
		const prefix string = ",\"int64\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Int64))
	}
	if in.Float32 != 0 {
		const prefix string = ",\"float32\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Float32))
	}
	if in.Float64 != 0 {
		const prefix string = ",\"float64\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Float64))
	}
	if in.String != "" {
		const prefix string = ",\"string\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.String))
	}
	if in.Time != "" {
		const prefix string = ",\"time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Time))
	}
	if in.Bool {
		const prefix string = ",\"bool\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Bool))
	}
	if in.Bytes != "" {
		const prefix string = ",\"bytes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Bytes))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v param) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComMachinefiW3bstreamWasmGolangSdkDatabaseTypes(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v param) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComMachinefiW3bstreamWasmGolangSdkDatabaseTypes(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *param) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComMachinefiW3bstreamWasmGolangSdkDatabaseTypes(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *param) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComMachinefiW3bstreamWasmGolangSdkDatabaseTypes(l, v)
}
func easyjson6601e8cdDecodeGithubComMachinefiW3bstreamWasmGolangSdkDatabaseTypes1(in *jlexer.Lexer, out *dBQuery) {
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
		case "statement":
			out.Statement = string(in.String())
		case "params":
			if in.IsNull() {
				in.Skip()
				out.Params = nil
			} else {
				in.Delim('[')
				if out.Params == nil {
					if !in.IsDelim(']') {
						out.Params = make([]*param, 0, 8)
					} else {
						out.Params = []*param{}
					}
				} else {
					out.Params = (out.Params)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *param
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(param)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Params = append(out.Params, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComMachinefiW3bstreamWasmGolangSdkDatabaseTypes1(out *jwriter.Writer, in dBQuery) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"statement\":"
		out.RawString(prefix[1:])
		out.String(string(in.Statement))
	}
	{
		const prefix string = ",\"params\":"
		out.RawString(prefix)
		if in.Params == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Params {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v dBQuery) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComMachinefiW3bstreamWasmGolangSdkDatabaseTypes1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v dBQuery) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComMachinefiW3bstreamWasmGolangSdkDatabaseTypes1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *dBQuery) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComMachinefiW3bstreamWasmGolangSdkDatabaseTypes1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *dBQuery) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComMachinefiW3bstreamWasmGolangSdkDatabaseTypes1(l, v)
}
