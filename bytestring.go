package bytestring

import (
	"encoding/hex"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

type Option func([]byte) ([]byte, error)

func Type(typ Strings) Option {
	return func(bytes []byte) ([]byte, error) {
		return typ.toBytes(bytes)
	}
}

// Bytes is byte array wrapper.
type Bytes struct {
	data []byte
}

// NewBytes returns new Bytes including given byte array
func NewBytes(bytes []byte, options ...Option) (Bytes, error) {
	b := bytes
	for _, option := range options {
		bytes, err := option(b)
		if err != nil {
			return Bytes{ nil }, err
		}
		b = bytes
	}
	return Bytes{
		b,
	}, nil
}

// NewBytesFromstring returns new Bytes including given string.
func NewBytesFromString(str string, options ...Option) (Bytes, error) {
	b := *(*[]byte)(unsafe.Pointer(&str))
	return NewBytes(b, options...)
}

// ByteArray returns byte array in Bytes.
func (b Bytes) ByteArray() []byte {
	return b.data
}

// String returns ascii string.
func (b Bytes) String() string {
	return *(*string)(unsafe.Pointer(&(b.data)))
}

// HexString returns the hexdecimall encoding.
func (b Bytes) HexString() string {
	return hex.EncodeToString(b.data)
}

// Base64 returns the base64 encoded string.
func (b Bytes) Base64() string {
	return b.Base64Custom(base64.StdEncoding)
}

// Base64URL returns the base64url encoded string.
func (b Bytes) Base64URL() string {
	return b.Base64Custom(base64.URLEncoding.WithPadding(base64.NoPadding))
}

// Base64Custom returns the string encoded by encoding.
func (b Bytes) Base64Custom(encoding *base64.Encoding) string {
	return encoding.EncodeToString(b.data)
}

// GoByteArray returns the byte array printed by golang.
func (b Bytes) GoByteArray() string {
	return fmt.Sprintf("%v", b.data)
}

type Strings interface {
	toBytes(bytes []byte) ([]byte, error)
}

type AsciiString struct {
}

func (a AsciiString) toBytes(bytes []byte) ([]byte, error) {
	return bytes, nil
}

type HexString struct {
}

func (h HexString) toBytes(bytes []byte) ([]byte, error) {
	c := len(bytes)
	dst := make([]byte, c / 2)
	_, err := hex.Decode(dst, bytes)
	return dst, err
}

type Base64String struct {
	encoding *base64.Encoding
}

func (b Base64String) toBytes(bytes []byte) ([]byte, error) {
	len := len(bytes)
	for {
		if bytes[len - 1] != '=' {
			break
		}
		len = len - 1
	}
	dstLen := len * 6 / 8
	dst := make([]byte, dstLen)
	_, err := b.encoding.Decode(dst, bytes)
	return dst, err
}

type GoByteArrayString struct {
}

func (g GoByteArrayString) toBytes(src []byte) ([]byte, error) {
	str := *(*string)(unsafe.Pointer(&(src)))
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

var Ascii = AsciiString{}
var Hex = HexString{}
var Base64 = Base64String{base64.StdEncoding}
var Base64URL = Base64String{base64.URLEncoding.WithPadding(base64.NoPadding)}
var GoByteArray = GoByteArrayString{}

