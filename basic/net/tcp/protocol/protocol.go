package protocol

import (
	"encoding/binary"
	"fmt"
)

//协议格式，4字节长度，2个字节版本，一个字节类型，后面为包体
type PacketType byte

type PacketVersion [2]byte

var PacketVersionV1 = [2]byte{'V', '1'}

const (
	PacketHeadLength    = 7
	DefaultReadeBufSize = 2048
	DefaultWriteBufSize = 2048
	defaultPacketSize   = 256
)

var (
	PacketHeadLengthErr = fmt.Errorf("packet head length err")
	PacketLengthErr     = fmt.Errorf("packet length err")
	SliceRangeOut       = fmt.Errorf("slice out of range")
)

type Packet struct {
	length  uint32
	version [2]byte
	pType   PacketType
	payload []byte
}

func NewPacket(pVersion PacketVersion) *Packet {
	return &Packet{
		length:  0,
		version: pVersion,
		pType:   0,
		payload: nil,
	}
}

func (p *Packet) Version() PacketVersion {
	return p.version
}

func (p *Packet) Type() PacketType {
	return p.pType
}

func (p *Packet) Length() uint32 {
	return p.length
}

func (p *Packet) ParseHead(data []byte) error {
	if len(data) != PacketHeadLength {
		return PacketHeadLengthErr
	}
	p.length = binary.BigEndian.Uint32(data[:4])
	p.version[0], p.version[1] = data[4], data[5]
	p.pType = PacketType(data[6])
	return nil
}

func (p *Packet) CopyPayload(data []byte) error {
	p.payload = data
	return nil
}

func (p *Packet) Payload() []byte {
	return p.payload
}

func (p *Packet) Fill(version PacketVersion, pType PacketType, payload []byte) error {
	p.length = uint32(PacketHeadLength + len(payload))
	p.version = version
	p.pType = pType
	p.payload = payload
	return nil
}

func (p *Packet) Marshal() ([]byte, error) {
	v := make([]byte, p.Length())
	err := p.MarshalTo(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (p *Packet) Unmarshal(data []byte) error {
	if len(data) < PacketHeadLength {
		return PacketLengthErr
	}
	p.length = binary.BigEndian.Uint32(data[:4])
	p.version[0], p.version[1] = data[4], data[5]
	p.pType = PacketType(data[6])
	if len(data) != int(p.Length()) {
		return PacketLengthErr
	}
	payload := make([]byte, p.length-PacketHeadLength)
	copy(payload, data[PacketHeadLength:])
	p.payload = payload
	return nil
}

func (p *Packet) MarshalTo(v []byte) error {
	if len(v) <= 0 || len(v) != int(p.Length()) {
		return SliceRangeOut
	}
	binary.BigEndian.PutUint32(v[:4], p.Length())
	v[4], v[5] = p.version[0], p.version[1]
	v[6] = byte(p.pType)
	copy(v[PacketHeadLength:], p.Payload())
	return nil
}
