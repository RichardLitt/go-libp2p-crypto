// Code generated by protoc-gen-gogo.
// source: crypto.proto
// DO NOT EDIT!

/*
Package crypto_pb is a generated protocol buffer package.

It is generated from these files:
	crypto.proto

It has these top-level messages:
	PublicKey
	PrivateKey
*/
package crypto_pb

import proto "gx/QmchriuyWMcqHha3dG86rQyxoswSUpmLUBjuJ2kotB65qR/gogo-protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type KeyType int32

const (
	KeyType_RSA KeyType = 0
)

var KeyType_name = map[int32]string{
	0: "RSA",
}
var KeyType_value = map[string]int32{
	"RSA": 0,
}

func (x KeyType) Enum() *KeyType {
	p := new(KeyType)
	*p = x
	return p
}
func (x KeyType) String() string {
	return proto.EnumName(KeyType_name, int32(x))
}
func (x *KeyType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(KeyType_value, data, "KeyType")
	if err != nil {
		return err
	}
	*x = KeyType(value)
	return nil
}

type PublicKey struct {
	Type             *KeyType `protobuf:"varint,1,req,enum=crypto.pb.KeyType" json:"Type,omitempty"`
	Data             []byte   `protobuf:"bytes,2,req" json:"Data,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *PublicKey) Reset()         { *m = PublicKey{} }
func (m *PublicKey) String() string { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()    {}

func (m *PublicKey) GetType() KeyType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return KeyType_RSA
}

func (m *PublicKey) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PrivateKey struct {
	Type             *KeyType `protobuf:"varint,1,req,enum=crypto.pb.KeyType" json:"Type,omitempty"`
	Data             []byte   `protobuf:"bytes,2,req" json:"Data,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *PrivateKey) Reset()         { *m = PrivateKey{} }
func (m *PrivateKey) String() string { return proto.CompactTextString(m) }
func (*PrivateKey) ProtoMessage()    {}

func (m *PrivateKey) GetType() KeyType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return KeyType_RSA
}

func (m *PrivateKey) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("crypto.pb.KeyType", KeyType_name, KeyType_value)
}
