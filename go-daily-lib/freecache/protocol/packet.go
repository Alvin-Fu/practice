package protocol

import (
	"fmt"
	"encoding/binary"
)

type Packet struct {
	version [2]byte
	length  uint32
	pType   byte
	payload []byte
}

func (p *Packet) Version()[2]byte{
	return p.version
}

func (p *Packet) Length()uint32{
	return p.length
}

func (p *Packet) PType()byte{
	return p.pType
}

func (p *Packet) ParseHead(data []byte)error{
	if len(data) < HeadTotalLength{
		return fmt.Errorf("no enough data for packet head")
	}
	p.version[0], p.version[1] = data[0], data[1]
	p.length = binary.BigEndian.Uint32(data[2:6])
	p.pType = data[6]
	return nil
}

func (p *Packet) CopyPayload(data []byte){
	p.payload = data
}

func (p *Packet) Payload()[]byte{
	return p.payload
}

func (p *Packet) Fill(version [2]byte, pType byte, payload []byte)error{
	p.version = version
	p.length = uint32(HeadTotalLength + len(payload))
	p.pType = pType
	p.payload = payload
	return nil
}

func (p *Packet) Marshal()([]byte, error){
	data := make([]byte, p.Length())
	err := p.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *Packet) Unmarshal(data []byte)error{
	if len(data) < HeadTotalLength{
		return fmt.Errorf("input data size is to small")
	}
	p.version[0], p.version[1] = data[0], data[1]
	p.length = binary.BigEndian.Uint32(data[2:6])
	p.pType = data[6]
	if int(p.Length()) > len(data) {
		return fmt.Errorf("no enough payload data to unmarshal")
	}
	payload := make([]byte, p.length - HeadTotalLength)
	copy(payload, data[HeadTotalLength:p.length])
	p.payload = payload
	return nil
}

func (p *Packet) MarshalTo(data []byte)error{
	if len(data) <= 0 || len(data) < int(p.Length()){
		return fmt.Errorf("out of range")
	}
	data[0] = p.version[0]
	data[1] = p.version[1]
	binary.BigEndian.PutUint32(data[2:6], p.Length())
	data[6] = p.PType()
	copy(data[HeadTotalLength:], p.Payload())
	return nil
}
