package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Decoder struct {
	r   io.Reader
	buf []byte
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r, buf: make([]byte, defaultPacketSize)}
}

func (d *Decoder) Decode(p *Packet) error {
	_, err := io.ReadFull(d.r, d.buf[:PacketHeadLength])
	if err != nil {
		return err
	}
	err = p.ParseHead(d.buf[:PacketHeadLength])
	if err != nil {
		return err
	}
	p.payload = make([]byte, int(p.Length())-PacketHeadLength)
	_, err = io.ReadFull(d.r, p.Payload())
	if err != nil {
		return err
	}
	return nil
}

type Encoder struct {
	w   io.Writer
	buf []byte
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w, buf: make([]byte, defaultPacketSize)}
}

func (e *Encoder) Encode(p *Packet) error {
	if int(p.Length()) > len(e.buf) {
		e.buf = make([]byte, p.Length())
	}
	binary.BigEndian.PutUint32(e.buf[:4], p.Length())
	e.buf[4], e.buf[5] = p.version[0], p.version[1]
	e.buf[6] = byte(p.pType)

	copy(e.buf[PacketHeadLength:], p.Payload())
	n, err := e.w.Write(e.buf[0:p.Length()])
	if err != nil {
		return err
	}
	fmt.Println("write: ", n)
	return nil
}
