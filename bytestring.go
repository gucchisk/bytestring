package bytestring

import (
	"encoding/hex"
	"encoding/base64"
	"unsafe"
)

type Bytes struct {
	data []byte
}

func NewBytes(bytes []byte) Bytes {
	return Bytes{
		bytes,
	}
}

func (b Bytes) String() string {
	return *(*string)(unsafe.Pointer(&(b.data)))
}

func (b Bytes) HexString() string {
	return hex.EncodeToString(b.data)
}

func (b Bytes) Base64() string {
	return base64.StdEncoding.EncodeToString(b.data)
}

type StringType int
const (
	Normal = iota
	Hex
	Base64
)

type String struct {
	data []byte
}

func NewString(str string) String {
	return String{
		stringToBytes(str),
	}
}

func NewStringT(str string, typ StringType) (String, error) {
	var d []byte
	var err error = nil
	switch typ {
	case Hex:
		d, err = hexStringToBytes(str)
	case Base64:
		d, err = base64StringToBytes(str)
	default:
		d = stringToBytes(str)
	}
	return String{
		d,
	}, err
}

func (s String) Bytes() []byte {
	return s.data
}

func stringToBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

func hexStringToBytes(str string) ([]byte, error) {
	return hex.DecodeString(str)
}

func base64StringToBytes(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}
