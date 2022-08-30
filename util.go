package gosphinx

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func writeFloat32ToBytes(bs []byte, f float32) []byte {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, f); err != nil {
		fmt.Println(err)
	}
	return append(bs, buf.Bytes()...)
}

func writeInt16ToBytes(bs []byte, i int) []byte {
	var byte2 = make([]byte, 2)
	binary.BigEndian.PutUint16(byte2, uint16(i))
	return append(bs, byte2...)
}

func writeInt32ToBytes(bs []byte, i int) []byte {
	var byte4 = make([]byte, 4)
	binary.BigEndian.PutUint32(byte4, uint32(i))
	return append(bs, byte4...)
}

func writeInt64ToBytes(bs []byte, ui uint64) []byte {
	var byte8 = make([]byte, 8)
	binary.BigEndian.PutUint64(byte8, ui)
	return append(bs, byte8...)
}

func writeLenStrToBytes(bs []byte, s string) []byte {
	var byte4 = make([]byte, 4)
	binary.BigEndian.PutUint32(byte4, uint32(len(s)))
	bs = append(bs, byte4...)
	return append(bs, []byte(s)...)
}

// For SetGeoAnchor()
func DegreeToRadian(degree float32) float32 {
	return degree * math.Pi / 180
}

type byteParser struct {
	stream []byte
	p      int
}

func (bp *byteParser) Int32() (i int) {
	i = int(binary.BigEndian.Uint32(bp.stream[bp.p : bp.p+4]))
	bp.p += 4
	return
}

func (bp *byteParser) Uint32() (i uint32) {
	i = binary.BigEndian.Uint32(bp.stream[bp.p : bp.p+4])
	bp.p += 4
	return
}

func (bp *byteParser) Uint64() (i uint64) {
	i = binary.BigEndian.Uint64(bp.stream[bp.p : bp.p+8])
	bp.p += 8
	return
}

func (bp *byteParser) Float32() (f float32, err error) {
	buf := bytes.NewBuffer(bp.stream[bp.p : bp.p+4])
	bp.p += 4
	if err := binary.Read(buf, binary.BigEndian, &f); err != nil {
		return 0, err
	}
	return f, nil
}

func (bp *byteParser) String() (s string) {
	s = ""
	if slen := bp.Int32(); slen > 0 {
		s = string(bp.stream[bp.p : bp.p+slen])
		bp.p += slen
	}
	return
}

// SetBit description
//
// createTime: 2022-08-26 12:20:48
//
// author: hailaz
func SetBit(in int, bit int, on bool) int {
	if on {
		return in | (1 << bit)
	}
	return in &^ (1 << bit)
}
