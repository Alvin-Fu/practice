// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cache.proto

/*
	Package PbCache is a generated protocol buffer package.

	It is generated from these files:
		cache.proto

	It has these top-level messages:
		CacheReq
		CacheResp
*/
package PbCache

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type OptionType int32

const (
	OptionType_Option_Type_Get OptionType = 1
	OptionType_Option_Type_Set OptionType = 2
	OptionType_Option_Type_Del OptionType = 3
)

var OptionType_name = map[int32]string{
	1: "Option_Type_Get",
	2: "Option_Type_Set",
	3: "Option_Type_Del",
}
var OptionType_value = map[string]int32{
	"Option_Type_Get": 1,
	"Option_Type_Set": 2,
	"Option_Type_Del": 3,
}

func (x OptionType) Enum() *OptionType {
	p := new(OptionType)
	*p = x
	return p
}
func (x OptionType) String() string {
	return proto.EnumName(OptionType_name, int32(x))
}
func (x *OptionType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(OptionType_value, data, "OptionType")
	if err != nil {
		return err
	}
	*x = OptionType(value)
	return nil
}
func (OptionType) EnumDescriptor() ([]byte, []int) { return fileDescriptorCache, []int{0} }

type CacheReq struct {
	OptionType       *OptionType `protobuf:"varint,1,opt,name=option_type,json=optionType,enum=PbCache.OptionType" json:"option_type,omitempty"`
	Key              *string     `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value            *string     `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	Expire           *int64      `protobuf:"varint,4,opt,name=expire" json:"expire,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CacheReq) Reset()                    { *m = CacheReq{} }
func (*CacheReq) ProtoMessage()               {}
func (*CacheReq) Descriptor() ([]byte, []int) { return fileDescriptorCache, []int{0} }

func (m *CacheReq) GetOptionType() OptionType {
	if m != nil && m.OptionType != nil {
		return *m.OptionType
	}
	return OptionType_Option_Type_Get
}

func (m *CacheReq) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *CacheReq) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func (m *CacheReq) GetExpire() int64 {
	if m != nil && m.Expire != nil {
		return *m.Expire
	}
	return 0
}

type CacheResp struct {
	ErrCode          *int32  `protobuf:"varint,1,opt,name=err_code,json=errCode" json:"err_code,omitempty"`
	ErrMsg           *string `protobuf:"bytes,2,opt,name=err_msg,json=errMsg" json:"err_msg,omitempty"`
	Rue              *string `protobuf:"bytes,3,opt,name=rue" json:"rue,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CacheResp) Reset()                    { *m = CacheResp{} }
func (*CacheResp) ProtoMessage()               {}
func (*CacheResp) Descriptor() ([]byte, []int) { return fileDescriptorCache, []int{1} }

func (m *CacheResp) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *CacheResp) GetErrMsg() string {
	if m != nil && m.ErrMsg != nil {
		return *m.ErrMsg
	}
	return ""
}

func (m *CacheResp) GetRue() string {
	if m != nil && m.Rue != nil {
		return *m.Rue
	}
	return ""
}

func init() {
	proto.RegisterType((*CacheReq)(nil), "PbCache.CacheReq")
	proto.RegisterType((*CacheResp)(nil), "PbCache.CacheResp")
	proto.RegisterEnum("PbCache.OptionType", OptionType_name, OptionType_value)
}
func (this *CacheReq) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*CacheReq)
	if !ok {
		that2, ok := that.(CacheReq)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *CacheReq")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *CacheReq but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *CacheReq but is not nil && this == nil")
	}
	if this.OptionType != nil && that1.OptionType != nil {
		if *this.OptionType != *that1.OptionType {
			return fmt.Errorf("OptionType this(%v) Not Equal that(%v)", *this.OptionType, *that1.OptionType)
		}
	} else if this.OptionType != nil {
		return fmt.Errorf("this.OptionType == nil && that.OptionType != nil")
	} else if that1.OptionType != nil {
		return fmt.Errorf("OptionType this(%v) Not Equal that(%v)", this.OptionType, that1.OptionType)
	}
	if this.Key != nil && that1.Key != nil {
		if *this.Key != *that1.Key {
			return fmt.Errorf("Key this(%v) Not Equal that(%v)", *this.Key, *that1.Key)
		}
	} else if this.Key != nil {
		return fmt.Errorf("this.Key == nil && that.Key != nil")
	} else if that1.Key != nil {
		return fmt.Errorf("Key this(%v) Not Equal that(%v)", this.Key, that1.Key)
	}
	if this.Value != nil && that1.Value != nil {
		if *this.Value != *that1.Value {
			return fmt.Errorf("Value this(%v) Not Equal that(%v)", *this.Value, *that1.Value)
		}
	} else if this.Value != nil {
		return fmt.Errorf("this.Value == nil && that.Value != nil")
	} else if that1.Value != nil {
		return fmt.Errorf("Value this(%v) Not Equal that(%v)", this.Value, that1.Value)
	}
	if this.Expire != nil && that1.Expire != nil {
		if *this.Expire != *that1.Expire {
			return fmt.Errorf("Expire this(%v) Not Equal that(%v)", *this.Expire, *that1.Expire)
		}
	} else if this.Expire != nil {
		return fmt.Errorf("this.Expire == nil && that.Expire != nil")
	} else if that1.Expire != nil {
		return fmt.Errorf("Expire this(%v) Not Equal that(%v)", this.Expire, that1.Expire)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *CacheReq) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CacheReq)
	if !ok {
		that2, ok := that.(CacheReq)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.OptionType != nil && that1.OptionType != nil {
		if *this.OptionType != *that1.OptionType {
			return false
		}
	} else if this.OptionType != nil {
		return false
	} else if that1.OptionType != nil {
		return false
	}
	if this.Key != nil && that1.Key != nil {
		if *this.Key != *that1.Key {
			return false
		}
	} else if this.Key != nil {
		return false
	} else if that1.Key != nil {
		return false
	}
	if this.Value != nil && that1.Value != nil {
		if *this.Value != *that1.Value {
			return false
		}
	} else if this.Value != nil {
		return false
	} else if that1.Value != nil {
		return false
	}
	if this.Expire != nil && that1.Expire != nil {
		if *this.Expire != *that1.Expire {
			return false
		}
	} else if this.Expire != nil {
		return false
	} else if that1.Expire != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CacheResp) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*CacheResp)
	if !ok {
		that2, ok := that.(CacheResp)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *CacheResp")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *CacheResp but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *CacheResp but is not nil && this == nil")
	}
	if this.ErrCode != nil && that1.ErrCode != nil {
		if *this.ErrCode != *that1.ErrCode {
			return fmt.Errorf("ErrCode this(%v) Not Equal that(%v)", *this.ErrCode, *that1.ErrCode)
		}
	} else if this.ErrCode != nil {
		return fmt.Errorf("this.ErrCode == nil && that.ErrCode != nil")
	} else if that1.ErrCode != nil {
		return fmt.Errorf("ErrCode this(%v) Not Equal that(%v)", this.ErrCode, that1.ErrCode)
	}
	if this.ErrMsg != nil && that1.ErrMsg != nil {
		if *this.ErrMsg != *that1.ErrMsg {
			return fmt.Errorf("ErrMsg this(%v) Not Equal that(%v)", *this.ErrMsg, *that1.ErrMsg)
		}
	} else if this.ErrMsg != nil {
		return fmt.Errorf("this.ErrMsg == nil && that.ErrMsg != nil")
	} else if that1.ErrMsg != nil {
		return fmt.Errorf("ErrMsg this(%v) Not Equal that(%v)", this.ErrMsg, that1.ErrMsg)
	}
	if this.Rue != nil && that1.Rue != nil {
		if *this.Rue != *that1.Rue {
			return fmt.Errorf("Rue this(%v) Not Equal that(%v)", *this.Rue, *that1.Rue)
		}
	} else if this.Rue != nil {
		return fmt.Errorf("this.Rue == nil && that.Rue != nil")
	} else if that1.Rue != nil {
		return fmt.Errorf("Rue this(%v) Not Equal that(%v)", this.Rue, that1.Rue)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *CacheResp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CacheResp)
	if !ok {
		that2, ok := that.(CacheResp)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ErrCode != nil && that1.ErrCode != nil {
		if *this.ErrCode != *that1.ErrCode {
			return false
		}
	} else if this.ErrCode != nil {
		return false
	} else if that1.ErrCode != nil {
		return false
	}
	if this.ErrMsg != nil && that1.ErrMsg != nil {
		if *this.ErrMsg != *that1.ErrMsg {
			return false
		}
	} else if this.ErrMsg != nil {
		return false
	} else if that1.ErrMsg != nil {
		return false
	}
	if this.Rue != nil && that1.Rue != nil {
		if *this.Rue != *that1.Rue {
			return false
		}
	} else if this.Rue != nil {
		return false
	} else if that1.Rue != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CacheReq) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&PbCache.CacheReq{")
	if this.OptionType != nil {
		s = append(s, "OptionType: "+valueToGoStringCache(this.OptionType, "OptionType")+",\n")
	}
	if this.Key != nil {
		s = append(s, "Key: "+valueToGoStringCache(this.Key, "string")+",\n")
	}
	if this.Value != nil {
		s = append(s, "Value: "+valueToGoStringCache(this.Value, "string")+",\n")
	}
	if this.Expire != nil {
		s = append(s, "Expire: "+valueToGoStringCache(this.Expire, "int64")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CacheResp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&PbCache.CacheResp{")
	if this.ErrCode != nil {
		s = append(s, "ErrCode: "+valueToGoStringCache(this.ErrCode, "int32")+",\n")
	}
	if this.ErrMsg != nil {
		s = append(s, "ErrMsg: "+valueToGoStringCache(this.ErrMsg, "string")+",\n")
	}
	if this.Rue != nil {
		s = append(s, "Rue: "+valueToGoStringCache(this.Rue, "string")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringCache(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *CacheReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CacheReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.OptionType != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCache(dAtA, i, uint64(*m.OptionType))
	}
	if m.Key != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCache(dAtA, i, uint64(len(*m.Key)))
		i += copy(dAtA[i:], *m.Key)
	}
	if m.Value != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCache(dAtA, i, uint64(len(*m.Value)))
		i += copy(dAtA[i:], *m.Value)
	}
	if m.Expire != nil {
		dAtA[i] = 0x20
		i++
		i = encodeVarintCache(dAtA, i, uint64(*m.Expire))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CacheResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CacheResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ErrCode != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCache(dAtA, i, uint64(*m.ErrCode))
	}
	if m.ErrMsg != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCache(dAtA, i, uint64(len(*m.ErrMsg)))
		i += copy(dAtA[i:], *m.ErrMsg)
	}
	if m.Rue != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCache(dAtA, i, uint64(len(*m.Rue)))
		i += copy(dAtA[i:], *m.Rue)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintCache(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedCacheReq(r randyCache, easy bool) *CacheReq {
	this := &CacheReq{}
	if r.Intn(10) != 0 {
		v1 := OptionType([]int32{1, 2, 3}[r.Intn(3)])
		this.OptionType = &v1
	}
	if r.Intn(10) != 0 {
		v2 := string(randStringCache(r))
		this.Key = &v2
	}
	if r.Intn(10) != 0 {
		v3 := string(randStringCache(r))
		this.Value = &v3
	}
	if r.Intn(10) != 0 {
		v4 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		this.Expire = &v4
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedCache(r, 5)
	}
	return this
}

func NewPopulatedCacheResp(r randyCache, easy bool) *CacheResp {
	this := &CacheResp{}
	if r.Intn(10) != 0 {
		v5 := int32(r.Int31())
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		this.ErrCode = &v5
	}
	if r.Intn(10) != 0 {
		v6 := string(randStringCache(r))
		this.ErrMsg = &v6
	}
	if r.Intn(10) != 0 {
		v7 := string(randStringCache(r))
		this.Rue = &v7
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedCache(r, 4)
	}
	return this
}

type randyCache interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneCache(r randyCache) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringCache(r randyCache) string {
	v8 := r.Intn(100)
	tmps := make([]rune, v8)
	for i := 0; i < v8; i++ {
		tmps[i] = randUTF8RuneCache(r)
	}
	return string(tmps)
}
func randUnrecognizedCache(r randyCache, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldCache(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldCache(dAtA []byte, r randyCache, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateCache(dAtA, uint64(key))
		v9 := r.Int63()
		if r.Intn(2) == 0 {
			v9 *= -1
		}
		dAtA = encodeVarintPopulateCache(dAtA, uint64(v9))
	case 1:
		dAtA = encodeVarintPopulateCache(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateCache(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateCache(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateCache(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateCache(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *CacheReq) Size() (n int) {
	var l int
	_ = l
	if m.OptionType != nil {
		n += 1 + sovCache(uint64(*m.OptionType))
	}
	if m.Key != nil {
		l = len(*m.Key)
		n += 1 + l + sovCache(uint64(l))
	}
	if m.Value != nil {
		l = len(*m.Value)
		n += 1 + l + sovCache(uint64(l))
	}
	if m.Expire != nil {
		n += 1 + sovCache(uint64(*m.Expire))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CacheResp) Size() (n int) {
	var l int
	_ = l
	if m.ErrCode != nil {
		n += 1 + sovCache(uint64(*m.ErrCode))
	}
	if m.ErrMsg != nil {
		l = len(*m.ErrMsg)
		n += 1 + l + sovCache(uint64(l))
	}
	if m.Rue != nil {
		l = len(*m.Rue)
		n += 1 + l + sovCache(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCache(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCache(x uint64) (n int) {
	return sovCache(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CacheReq) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CacheReq{`,
		`OptionType:` + valueToStringCache(this.OptionType) + `,`,
		`Key:` + valueToStringCache(this.Key) + `,`,
		`Value:` + valueToStringCache(this.Value) + `,`,
		`Expire:` + valueToStringCache(this.Expire) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CacheResp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CacheResp{`,
		`ErrCode:` + valueToStringCache(this.ErrCode) + `,`,
		`ErrMsg:` + valueToStringCache(this.ErrMsg) + `,`,
		`Rue:` + valueToStringCache(this.Rue) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCache(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CacheReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCache
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CacheReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CacheReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionType", wireType)
			}
			var v OptionType
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (OptionType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptionType = &v
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCache
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Key = &s
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCache
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Value = &s
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expire", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Expire = &v
		default:
			iNdEx = preIndex
			skippy, err := skipCache(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCache
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CacheResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCache
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CacheResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CacheResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrCode", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ErrCode = &v
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCache
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.ErrMsg = &s
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCache
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCache
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Rue = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCache(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCache
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCache(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCache
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCache
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCache
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCache
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCache
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCache(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCache = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCache   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cache.proto", fileDescriptorCache) }

var fileDescriptorCache = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8e, 0xbd, 0x4e, 0x02, 0x41,
	0x14, 0x85, 0xbd, 0xac, 0xfc, 0x5d, 0x12, 0xdd, 0x0c, 0x46, 0x57, 0x8b, 0x09, 0xa1, 0x22, 0xc6,
	0x2c, 0x89, 0xe1, 0x09, 0xc4, 0xc4, 0xc2, 0x18, 0x75, 0xb5, 0xdf, 0xc0, 0x72, 0x5d, 0x88, 0xc0,
	0xac, 0xc3, 0x40, 0xa4, 0xe3, 0x71, 0x7c, 0x04, 0x4b, 0x4b, 0x4b, 0x4b, 0x4b, 0x76, 0x9e, 0xc0,
	0xd2, 0xd2, 0xcc, 0xec, 0x06, 0x13, 0xba, 0xef, 0x3b, 0x33, 0x37, 0xe7, 0x60, 0x2d, 0xea, 0x45,
	0x43, 0xf2, 0x13, 0x29, 0x94, 0x60, 0xe5, 0xbb, 0x7e, 0xd7, 0xe8, 0x49, 0x67, 0x41, 0xd3, 0x81,
	0x90, 0xed, 0x78, 0xa4, 0x86, 0xf3, 0xbe, 0x1f, 0x89, 0x49, 0x3b, 0x16, 0xb1, 0x68, 0xdb, 0x6f,
	0xfd, 0xf9, 0x93, 0x35, 0x2b, 0x96, 0xb2, 0xf3, 0xe6, 0x0a, 0xb0, 0x62, 0xef, 0x03, 0x7a, 0x61,
	0x1d, 0xac, 0x89, 0x44, 0x8d, 0xc4, 0x34, 0x54, 0xcb, 0x84, 0x3c, 0x68, 0x40, 0x6b, 0xef, 0xbc,
	0xee, 0xe7, 0x0d, 0xfe, 0xad, 0x7d, 0x7b, 0x5c, 0x26, 0x14, 0xa0, 0xd8, 0x30, 0x73, 0xd1, 0x79,
	0xa6, 0xa5, 0x57, 0x68, 0x40, 0xab, 0x1a, 0x18, 0x64, 0x07, 0x58, 0x5c, 0xf4, 0xc6, 0x73, 0xf2,
	0x1c, 0x9b, 0x65, 0xc2, 0x0e, 0xb1, 0x44, 0xaf, 0xc9, 0x48, 0x92, 0xb7, 0xdb, 0x80, 0x96, 0x13,
	0xe4, 0xd6, 0xbc, 0xc7, 0x6a, 0xbe, 0x60, 0x96, 0xb0, 0x63, 0xac, 0x90, 0x94, 0x61, 0x24, 0x06,
	0x59, 0x7f, 0x31, 0x28, 0x93, 0x94, 0x5d, 0x31, 0x20, 0x76, 0x84, 0x06, 0xc3, 0xc9, 0x2c, 0xce,
	0xbb, 0x4a, 0x24, 0xe5, 0xcd, 0x2c, 0x36, 0x03, 0xe4, 0xa6, 0xcc, 0xe0, 0xe9, 0x35, 0xe2, 0xff,
	0x58, 0x56, 0xc7, 0xfd, 0xcc, 0x42, 0xa3, 0xe1, 0x15, 0x29, 0x17, 0xb6, 0xc3, 0x07, 0x52, 0x6e,
	0x61, 0x3b, 0xbc, 0xa4, 0xb1, 0xeb, 0x5c, 0x9c, 0x7d, 0xa7, 0x7c, 0x67, 0x9d, 0x72, 0xf8, 0x49,
	0x39, 0xfc, 0xa6, 0x1c, 0x56, 0x9a, 0xc3, 0x9b, 0xe6, 0xf0, 0xae, 0x39, 0x7c, 0x68, 0x0e, 0x9f,
	0x9a, 0xc3, 0x97, 0xe6, 0xb0, 0xd6, 0x1c, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x1d, 0xac,
	0xc0, 0x9d, 0x01, 0x00, 0x00,
}
