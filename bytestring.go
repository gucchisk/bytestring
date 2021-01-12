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

func SetEncoding(format Encoding) Option {
	return func(bytes []byte) ([]byte, error) {
		return format.read(bytes)
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
	return Ascii.toString(b.data)
}

func (b Bytes) toString(enc Encoding) string {
	return enc.toString(b.data)
}

// HexString returns the hexdecimall encoding.
func (b Bytes) HexString() string {
	return Hex.toString(b.data)
}

// Base64 returns the base64 encoded string.
func (b Bytes) Base64() string {
	return Base64.toString(b.data)
}

// Base64URL returns the base64url encoded string.
func (b Bytes) Base64URL() string {
	return Base64URL.toString(b.data)
}

// Base64Custom returns the string encoded by encoding.
func (b Bytes) Base64Custom(encoding Encoding) string {
	return encoding.toString(b.data)
}

// GoByteArray returns the byte array printed by golang.
func (b Bytes) GoByteArray() string {
	return GoByteArray.toString(b.data)
}

type Encoding interface {
	read(src []byte) ([]byte, error)
	write(src []byte) []byte
	toString(src []byte) string
}

type AsciiEncoding struct {
}

func (a AsciiEncoding) read(src []byte) ([]byte, error) {
	return src, nil
}

func (a AsciiEncoding) write(src []byte) []byte {
	return src
}

func (a AsciiEncoding) toString(src []byte) string {
	b := a.write(src)
	return *(*string)(unsafe.Pointer(&b))
}

type HexEncoding struct {
}

func (h HexEncoding) read(src []byte) ([]byte, error) {
	c := len(src)
	dst := make([]byte, c / 2)
	_, err := hex.Decode(dst, src)
	return dst, err
}

func (h HexEncoding) write(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}

func (h HexEncoding) toString(src []byte) string {
	b := h.write(src)
	return *(*string)(unsafe.Pointer(&b))
}

type Base64Encoding struct {
	encoding *base64.Encoding
}

func (b Base64Encoding) read(src []byte) ([]byte, error) {
	// len := len(src)
	// for {
	// 	if src[len - 1] != '=' {
	// 		break
	// 	}
	// 	len = len - 1
	// }
	// dstLen := len * 6 / 8
	dst := make([]byte, b.encoding.DecodedLen(len(src)))
	l, err := b.encoding.Decode(dst, src)
	return dst[0:l], err
}

func (b Base64Encoding) write(src []byte) []byte {
	dstLen := b.encoding.EncodedLen(len(src))
	dst := make([]byte, dstLen)
	b.encoding.Encode(dst, src)
	return dst
}

func (b Base64Encoding) toString(src []byte) string {
	bytes := b.write(src)
	return *(*string)(unsafe.Pointer(&bytes))
}

type GoByteArrayEncoding struct {
}

func (g GoByteArrayEncoding) read(src []byte) ([]byte, error) {
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

func (g GoByteArrayEncoding) write(src []byte) []byte {
	s := g.toString(src)
	return *(*[]byte)(unsafe.Pointer(&s))
}

func (g GoByteArrayEncoding) toString(src []byte) string {
	return fmt.Sprintf("%v", src)
}

var Ascii = AsciiEncoding{}
var Hex = HexEncoding{}
var Base64 = Base64Encoding{base64.StdEncoding}
var Base64URL = Base64Encoding{base64.URLEncoding.WithPadding(base64.NoPadding)}
var GoByteArray = GoByteArrayEncoding{}

