package bytestring

import (
	"encoding/hex"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
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
	return b.Base64Custom(base64.StdEncoding)
}

func (b Bytes) Base64URL() string {
	return b.Base64Custom(base64.URLEncoding.WithPadding(base64.NoPadding))
}

func (b Bytes) Base64Custom(encoding *base64.Encoding) string {
	return encoding.EncodeToString(b.data)
}

func (b Bytes) GoByteArray() string {
	return fmt.Sprintf("%v", b.data)
}

type Strings interface {
	toBytes(str string) ([]byte, error)
}

type NormalString struct {
}

func (n NormalString) toBytes(str string) ([]byte, error) {
	return *(*[]byte)(unsafe.Pointer(&str)), nil
}

type HexString struct {
}

func (h HexString) toBytes(str string) ([]byte, error) {
	return hex.DecodeString(str)
}

type Base64String struct {
	encoding *base64.Encoding
}

func (b Base64String) toBytes(str string) ([]byte, error) {
	return b.encoding.DecodeString(str)
}

type GoByteArrayString struct {
}

func (g GoByteArrayString) toBytes(str string) ([]byte, error) {
	strn := strings.Replace(str, "[", "", 1)
	strn = strings.Replace(strn, "]", "", 1)
	strs := strings.Split(strn, " ")
	bytes := make([]byte, len(strs), len(strs))
	for i, v := range strs {
		b, err := strconv.Atoi(v)
		if err != nil {
			return bytes, err
		}
		if b > 255 || b < 0 {
			return bytes, fmt.Errorf("bytestring: invalid value of %d at %d", b, i)
		}
		bytes[i] = byte(b)
	}
	return bytes, nil
}

var Normal = NormalString{}
var Hex = HexString{}
var Base64 = Base64String{base64.StdEncoding}
var Base64URL = Base64String{base64.URLEncoding.WithPadding(base64.NoPadding)}
var GoByteArray = GoByteArrayString{}

type String struct {
	data []byte
}

func NewBytesFrom(str string, typ Strings) (String, error) {
	d, err := typ.toBytes(str)
	return String{
		d,
	}, err
}

func (s String) ByteArray() []byte {
	return s.data
}
