// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/constructs.proto

package pb

import (
	context "context"
	encoding_json "encoding/json"
	fmt "fmt"
	github_com_danielvladco_go_proto_gql_gqltypes "github.com/danielvladco/go-proto-gql/gqltypes"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	github_com_golang_protobuf_ptypes_any "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/empty"
	github_com_golang_protobuf_ptypes_empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConstructsGQLServer struct{ Service ConstructsServer }

func (s *ConstructsGQLServer) ConstructsScalars(ctx context.Context, in *Scalars) (*Scalars, error) {
	return s.Service.Scalars_(ctx, in)
}
func (s *ConstructsGQLServer) ConstructsRepeated(ctx context.Context, in *Repeated) (*Repeated, error) {
	return s.Service.Repeated_(ctx, in)
}
func (s *ConstructsGQLServer) ConstructsMaps(ctx context.Context, in *Maps) (*Maps, error) {
	return s.Service.Maps_(ctx, in)
}
func (s *ConstructsGQLServer) ConstructsAny(ctx context.Context, in *github_com_golang_protobuf_ptypes_any.Any) (*Any, error) {
	return s.Service.Any_(ctx, in)
}
func (s *ConstructsGQLServer) ConstructsEmpty(ctx context.Context) (*bool, error) {
	_, err := s.Service.Empty_(ctx, &github_com_golang_protobuf_ptypes_empty.Empty{})
	return nil, err
}
func (s *ConstructsGQLServer) ConstructsEmpty2(ctx context.Context) (*bool, error) {
	_, err := s.Service.Empty2_(ctx, &EmptyRecursive{})
	return nil, err
}
func (s *ConstructsGQLServer) ConstructsEmpty3(ctx context.Context) (*bool, error) {
	_, err := s.Service.Empty3_(ctx, &Empty3{})
	return nil, err
}
func (s *ConstructsGQLServer) ConstructsRef(ctx context.Context, in *Ref) (*Ref, error) {
	return s.Service.Ref_(ctx, in)
}
func (s *ConstructsGQLServer) ConstructsOneof(ctx context.Context, in *Oneof) (*Oneof, error) {
	return s.Service.Oneof_(ctx, in)
}
func (s *ConstructsGQLServer) ConstructsCallWithID(ctx context.Context) (*bool, error) {
	_, err := s.Service.CallWithId(ctx, &Empty{})
	return nil, err
}

func MarshalMaps_Int32Int32Entry(mp map[int32]int32) github_com_danielvladco_go_proto_gql_gqltypes.Marshaler {
	return github_com_danielvladco_go_proto_gql_gqltypes.WriterFunc(func(w io.Writer) {
		err := encoding_json.NewEncoder(w).Encode(mp)
		if err != nil {
			panic("stupid map")
		}
	})
}

func UnmarshalMaps_Int32Int32Entry(v interface{}) (mp map[int32]int32, err error) {
	switch vv := v.(type) {
	case []byte:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	case encoding_json.RawMessage:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	default:
		return nil, fmt.Errorf("%T is not map[int32]int32", v)
	}
}

func MarshalMaps_Int64Int64Entry(mp map[int64]int64) github_com_danielvladco_go_proto_gql_gqltypes.Marshaler {
	return github_com_danielvladco_go_proto_gql_gqltypes.WriterFunc(func(w io.Writer) {
		err := encoding_json.NewEncoder(w).Encode(mp)
		if err != nil {
			panic("stupid map")
		}
	})
}

func UnmarshalMaps_Int64Int64Entry(v interface{}) (mp map[int64]int64, err error) {
	switch vv := v.(type) {
	case []byte:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	case encoding_json.RawMessage:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	default:
		return nil, fmt.Errorf("%T is not map[int64]int64", v)
	}
}

func MarshalMaps_Uint32Uint32Entry(mp map[uint32]uint32) github_com_danielvladco_go_proto_gql_gqltypes.Marshaler {
	return github_com_danielvladco_go_proto_gql_gqltypes.WriterFunc(func(w io.Writer) {
		err := encoding_json.NewEncoder(w).Encode(mp)
		if err != nil {
			panic("stupid map")
		}
	})
}

func UnmarshalMaps_Uint32Uint32Entry(v interface{}) (mp map[uint32]uint32, err error) {
	switch vv := v.(type) {
	case []byte:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	case encoding_json.RawMessage:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	default:
		return nil, fmt.Errorf("%T is not map[uint32]uint32", v)
	}
}

func MarshalMaps_Uint64Uint64Entry(mp map[uint64]uint64) github_com_danielvladco_go_proto_gql_gqltypes.Marshaler {
	return github_com_danielvladco_go_proto_gql_gqltypes.WriterFunc(func(w io.Writer) {
		err := encoding_json.NewEncoder(w).Encode(mp)
		if err != nil {
			panic("stupid map")
		}
	})
}

func UnmarshalMaps_Uint64Uint64Entry(v interface{}) (mp map[uint64]uint64, err error) {
	switch vv := v.(type) {
	case []byte:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	case encoding_json.RawMessage:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	default:
		return nil, fmt.Errorf("%T is not map[uint64]uint64", v)
	}
}

func MarshalMaps_Sint32Sint32Entry(mp map[int32]int32) github_com_danielvladco_go_proto_gql_gqltypes.Marshaler {
	return github_com_danielvladco_go_proto_gql_gqltypes.WriterFunc(func(w io.Writer) {
		err := encoding_json.NewEncoder(w).Encode(mp)
		if err != nil {
			panic("stupid map")
		}
	})
}

func UnmarshalMaps_Sint32Sint32Entry(v interface{}) (mp map[int32]int32, err error) {
	switch vv := v.(type) {
	case []byte:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	case encoding_json.RawMessage:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	default:
		return nil, fmt.Errorf("%T is not map[int32]int32", v)
	}
}

func MarshalMaps_Sint64Sint64Entry(mp map[int64]int64) github_com_danielvladco_go_proto_gql_gqltypes.Marshaler {
	return github_com_danielvladco_go_proto_gql_gqltypes.WriterFunc(func(w io.Writer) {
		err := encoding_json.NewEncoder(w).Encode(mp)
		if err != nil {
			panic("stupid map")
		}
	})
}

func UnmarshalMaps_Sint64Sint64Entry(v interface{}) (mp map[int64]int64, err error) {
	switch vv := v.(type) {
	case []byte:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	case encoding_json.RawMessage:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	default:
		return nil, fmt.Errorf("%T is not map[int64]int64", v)
	}
}

func MarshalMaps_Fixed32Fixed32Entry(mp map[uint32]uint32) github_com_danielvladco_go_proto_gql_gqltypes.Marshaler {
	return github_com_danielvladco_go_proto_gql_gqltypes.WriterFunc(func(w io.Writer) {
		err := encoding_json.NewEncoder(w).Encode(mp)
		if err != nil {
			panic("stupid map")
		}
	})
}

func UnmarshalMaps_Fixed32Fixed32Entry(v interface{}) (mp map[uint32]uint32, err error) {
	switch vv := v.(type) {
	case []byte:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	case encoding_json.RawMessage:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	default:
		return nil, fmt.Errorf("%T is not map[uint32]uint32", v)
	}
}

func MarshalMaps_Fixed64Fixed64Entry(mp map[uint64]uint64) github_com_danielvladco_go_proto_gql_gqltypes.Marshaler {
	return github_com_danielvladco_go_proto_gql_gqltypes.WriterFunc(func(w io.Writer) {
		err := encoding_json.NewEncoder(w).Encode(mp)
		if err != nil {
			panic("stupid map")
		}
	})
}

func UnmarshalMaps_Fixed64Fixed64Entry(v interface{}) (mp map[uint64]uint64, err error) {
	switch vv := v.(type) {
	case []byte:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	case encoding_json.RawMessage:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	default:
		return nil, fmt.Errorf("%T is not map[uint64]uint64", v)
	}
}

func MarshalMaps_Sfixed32Sfixed32Entry(mp map[int32]int32) github_com_danielvladco_go_proto_gql_gqltypes.Marshaler {
	return github_com_danielvladco_go_proto_gql_gqltypes.WriterFunc(func(w io.Writer) {
		err := encoding_json.NewEncoder(w).Encode(mp)
		if err != nil {
			panic("stupid map")
		}
	})
}

func UnmarshalMaps_Sfixed32Sfixed32Entry(v interface{}) (mp map[int32]int32, err error) {
	switch vv := v.(type) {
	case []byte:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	case encoding_json.RawMessage:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	default:
		return nil, fmt.Errorf("%T is not map[int32]int32", v)
	}
}

func MarshalMaps_Sfixed64Sfixed64Entry(mp map[int64]int64) github_com_danielvladco_go_proto_gql_gqltypes.Marshaler {
	return github_com_danielvladco_go_proto_gql_gqltypes.WriterFunc(func(w io.Writer) {
		err := encoding_json.NewEncoder(w).Encode(mp)
		if err != nil {
			panic("stupid map")
		}
	})
}

func UnmarshalMaps_Sfixed64Sfixed64Entry(v interface{}) (mp map[int64]int64, err error) {
	switch vv := v.(type) {
	case []byte:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	case encoding_json.RawMessage:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	default:
		return nil, fmt.Errorf("%T is not map[int64]int64", v)
	}
}

func MarshalMaps_BoolBoolEntry(mp map[bool]bool) github_com_danielvladco_go_proto_gql_gqltypes.Marshaler {
	return github_com_danielvladco_go_proto_gql_gqltypes.WriterFunc(func(w io.Writer) {
		err := encoding_json.NewEncoder(w).Encode(mp)
		if err != nil {
			panic("stupid map")
		}
	})
}

func UnmarshalMaps_BoolBoolEntry(v interface{}) (mp map[bool]bool, err error) {
	switch vv := v.(type) {
	case []byte:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	case encoding_json.RawMessage:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	default:
		return nil, fmt.Errorf("%T is not map[bool]bool", v)
	}
}

func MarshalMaps_StringStringEntry(mp map[string]string) github_com_danielvladco_go_proto_gql_gqltypes.Marshaler {
	return github_com_danielvladco_go_proto_gql_gqltypes.WriterFunc(func(w io.Writer) {
		err := encoding_json.NewEncoder(w).Encode(mp)
		if err != nil {
			panic("stupid map")
		}
	})
}

func UnmarshalMaps_StringStringEntry(v interface{}) (mp map[string]string, err error) {
	switch vv := v.(type) {
	case []byte:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	case encoding_json.RawMessage:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	default:
		return nil, fmt.Errorf("%T is not map[string]string", v)
	}
}

func MarshalMaps_StringBytesEntry(mp map[string][]byte) github_com_danielvladco_go_proto_gql_gqltypes.Marshaler {
	return github_com_danielvladco_go_proto_gql_gqltypes.WriterFunc(func(w io.Writer) {
		err := encoding_json.NewEncoder(w).Encode(mp)
		if err != nil {
			panic("stupid map")
		}
	})
}

func UnmarshalMaps_StringBytesEntry(v interface{}) (mp map[string][]byte, err error) {
	switch vv := v.(type) {
	case []byte:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	case encoding_json.RawMessage:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	default:
		return nil, fmt.Errorf("%T is not map[string][]byte", v)
	}
}

func MarshalMaps_StringFloatEntry(mp map[string]float32) github_com_danielvladco_go_proto_gql_gqltypes.Marshaler {
	return github_com_danielvladco_go_proto_gql_gqltypes.WriterFunc(func(w io.Writer) {
		err := encoding_json.NewEncoder(w).Encode(mp)
		if err != nil {
			panic("stupid map")
		}
	})
}

func UnmarshalMaps_StringFloatEntry(v interface{}) (mp map[string]float32, err error) {
	switch vv := v.(type) {
	case []byte:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	case encoding_json.RawMessage:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	default:
		return nil, fmt.Errorf("%T is not map[string]float32", v)
	}
}

func MarshalMaps_StringDoubleEntry(mp map[string]float64) github_com_danielvladco_go_proto_gql_gqltypes.Marshaler {
	return github_com_danielvladco_go_proto_gql_gqltypes.WriterFunc(func(w io.Writer) {
		err := encoding_json.NewEncoder(w).Encode(mp)
		if err != nil {
			panic("stupid map")
		}
	})
}

func UnmarshalMaps_StringDoubleEntry(v interface{}) (mp map[string]float64, err error) {
	switch vv := v.(type) {
	case []byte:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	case encoding_json.RawMessage:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	default:
		return nil, fmt.Errorf("%T is not map[string]float64", v)
	}
}

func MarshalMaps_StringFooEntry(mp map[string]*Foo) github_com_danielvladco_go_proto_gql_gqltypes.Marshaler {
	return github_com_danielvladco_go_proto_gql_gqltypes.WriterFunc(func(w io.Writer) {
		err := encoding_json.NewEncoder(w).Encode(mp)
		if err != nil {
			panic("stupid map")
		}
	})
}

func UnmarshalMaps_StringFooEntry(v interface{}) (mp map[string]*Foo, err error) {
	switch vv := v.(type) {
	case []byte:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	case encoding_json.RawMessage:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	default:
		return nil, fmt.Errorf("%T is not map[string]*Foo", v)
	}
}

func MarshalMaps_StringBarEntry(mp map[string]Bar) github_com_danielvladco_go_proto_gql_gqltypes.Marshaler {
	return github_com_danielvladco_go_proto_gql_gqltypes.WriterFunc(func(w io.Writer) {
		err := encoding_json.NewEncoder(w).Encode(mp)
		if err != nil {
			panic("stupid map")
		}
	})
}

func UnmarshalMaps_StringBarEntry(v interface{}) (mp map[string]Bar, err error) {
	switch vv := v.(type) {
	case []byte:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	case encoding_json.RawMessage:
		err = encoding_json.Unmarshal(vv, &mp)
		return mp, err
	default:
		return nil, fmt.Errorf("%T is not map[string]Bar", v)
	}
}

type IsOneof_Oneof1 interface {
	isOneof_Oneof1()
}
type IsOneof_Oneof2 interface {
	isOneof_Oneof2()
}
type IsOneof_Oneof3 interface {
	isOneof_Oneof3()
}

func (c *Bar) UnmarshalGQL(v interface{}) error {
	code, ok := v.(string)
	if ok {
		*c = Bar(Bar_value[code])
		return nil
	}
	return fmt.Errorf("cannot unmarshal Bar enum")
}

func (c Bar) MarshalGQL(w io.Writer) {
	fmt.Fprintf(w, "%q", c.String())
}

func (c *Ref_Foo_En) UnmarshalGQL(v interface{}) error {
	code, ok := v.(string)
	if ok {
		*c = Ref_Foo_En(Ref_Foo_En_value[code])
		return nil
	}
	return fmt.Errorf("cannot unmarshal Ref_Foo_En enum")
}

func (c Ref_Foo_En) MarshalGQL(w io.Writer) {
	fmt.Fprintf(w, "%q", c.String())
}

func (c *Ref_Foo_Bar_En) UnmarshalGQL(v interface{}) error {
	code, ok := v.(string)
	if ok {
		*c = Ref_Foo_Bar_En(Ref_Foo_Bar_En_value[code])
		return nil
	}
	return fmt.Errorf("cannot unmarshal Ref_Foo_Bar_En enum")
}

func (c Ref_Foo_Bar_En) MarshalGQL(w io.Writer) {
	fmt.Fprintf(w, "%q", c.String())
}
