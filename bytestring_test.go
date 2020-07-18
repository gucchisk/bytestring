package bytestring

import(
	"encoding/base64"
	"reflect"
	"testing"
)

func TestBytesToString(t *testing.T) {
	b := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}
	bytes := NewBytes(b)
	str := bytes.String()
	expectStr := "hello"
	if str != expectStr {
		t.Errorf("string not equal (result: %s, expect: %s)", str, expectStr)
	}
}

func TestBytesToHexString(t *testing.T) {
	b := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}
	bytes := NewBytes(b)
	str := bytes.HexString()
	expectStr := "68656c6c6f"
	if str != expectStr {
		t.Errorf("string not equal (result: %s, expect: %s)", str, expectStr)
	}
}

func TestBytesToBase64(t *testing.T) {
	b := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}
	bytes := NewBytes(b)
	str := bytes.Base64()
	expectStr := "aGVsbG8="
	if str != expectStr {
		t.Errorf("string not equal (result: %s, expect: %s)", str, expectStr)
	}

	b = []byte{0xfb, 0xff}
	bytes = NewBytes(b)
	str = bytes.Base64()
	expectStr = "+/8="
	if str != expectStr {
		t.Errorf("string not equal (result: %s, expect: %s)", str, expectStr)
	}
}

func TestBytesToBase64URL(t *testing.T) {
	b := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}
	bytes := NewBytes(b)
	str := bytes.Base64URL()
	expectStr := "aGVsbG8"
	if str != expectStr {
		t.Errorf("string not equal (result: %s, expect: %s)", str, expectStr)
	}

	b = []byte{0xfb, 0xff}
	bytes = NewBytes(b)
	str = bytes.Base64URL()
	expectStr = "-_8"
	if str != expectStr {
		t.Errorf("string not equal (result: %s, expect: %s)", str, expectStr)
	}
}

func TestBytesToGoByteArray(t *testing.T) {
	b := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}
	bytes := NewBytes(b)
	str := bytes.GoByteArray()
	expectStr := "[104 101 108 108 111]"
	if str != expectStr {
		t.Errorf("string not equal (result: %s, expect: %s)", str, expectStr)
	}
}


func TestStringToBytes(t *testing.T) {
	s := "hello"
	bytes, err := NewBytesFrom(s, Normal)
	if err != nil {
		t.Errorf("error is not nil: %s\n", err)
	}
	b := bytes.ByteArray()
	expectBytes := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}
	if !reflect.DeepEqual(b, expectBytes) {
		t.Errorf("bytes not equal (result: %d, expect: %d)", bytes, expectBytes)
	}
}

func TestHexStringToBytes(t *testing.T) {
	s := "68656c6c6f"
	bytes, err := NewBytesFrom(s, Hex)
	if err != nil {
		t.Errorf("error is not nil: %s\n", err)
	}
	b := bytes.ByteArray()
	expectBytes := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}
	if !reflect.DeepEqual(b, expectBytes) {
		t.Errorf("bytes not equal (result: %d, expect: %d)", bytes, expectBytes)
	}

	s = "abcdefgh"
	bytes, err = NewBytesFrom(s, Hex)
	if err == nil {
		t.Error("error is nil")
	}
	expectErrMsg := "encoding/hex: invalid byte: U+0067 'g'"
	if err.Error() != expectErrMsg {
		t.Errorf("error message not eaual (result: %s, expect: %s)\n", err, expectErrMsg)
	}

	s = "0abcdef"
	bytes, err = NewBytesFrom(s, Hex)
	if err == nil {
		t.Error("error is nil")
	}
	expectErrMsg = "encoding/hex: odd length hex string"
	if err.Error() != expectErrMsg {
		t.Errorf("error message not eaual (result: %s, expect: %s)\n", err, expectErrMsg)
	}
}

func TestBase64StringToBytes(t *testing.T) {
	s := "aGVsbG8="
	bytes, err := NewBytesFrom(s, Base64)
	if err != nil {
		t.Errorf("error is not nil: %s\n", err)
	}
	b := bytes.ByteArray()
	expectBytes := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}
	if !reflect.DeepEqual(b, expectBytes) {
		t.Errorf("bytes not equal (result: %d, expect: %d)", bytes, expectBytes)
	}

	s = "+/8="
	bytes, err = NewBytesFrom(s, Base64)
	if err != nil {
		t.Errorf("error is not nil: %s\n", err)
	}
	b = bytes.ByteArray()
	expectBytes = []byte{0xfb, 0xff}
	if !reflect.DeepEqual(b, expectBytes) {
		t.Errorf("bytes not equal (result: %d, expect: %d)", bytes, expectBytes)
	}
}

func TestBase64StringToBytesNoPadding(t *testing.T) {
	s := "aGVsbG8"
	bytes, err := NewBytesFrom(s, Base64String{base64.StdEncoding.WithPadding(base64.NoPadding)})
	b := bytes.ByteArray()
	if err != nil {
		t.Error("error is nil")
	}
	expectBytes := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}
	if !reflect.DeepEqual(b, expectBytes) {
		t.Errorf("bytes not equal (result: %d, expect: %d)", bytes, expectBytes)
	}
}

func TestBase64StringToBytesWithError(t *testing.T) {
	s := "abcdefg-"
	_, err := NewBytesFrom(s, Base64)
	if err == nil {
		t.Error("error is nil")
	}
	expectErrMsg := "illegal base64 data at input byte 7"
	if err.Error() != expectErrMsg {
		t.Errorf("error message not eaual (result: %s, expect: %s)\n", err, expectErrMsg)
	}
}


func TestGoByteArrayStringToBytes(t *testing.T) {
	s := "[104 101 108 108 111]"
	bytes, err := NewBytesFrom(s, GoByteArray)
	if err != nil {
		t.Errorf("error is not nil: %s\n", err)
	}
	b := bytes.ByteArray()
	expectBytes := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}
	if !reflect.DeepEqual(b, expectBytes) {
		t.Errorf("bytes not equal (result: %d, expect: %d)", bytes, expectBytes)
	}
}

func TestGoByteArrayStringToBytesWithError(t *testing.T) {
	s := "[104 101 108 108 111 256]"
	bytes, err := NewBytesFrom(s, GoByteArray)
	if err == nil {
		t.Error("error is nil")
	}
	_ = bytes.ByteArray()
	expectErrMsg := "bytestring: invalid value of 256 at 5"
	if err.Error() != expectErrMsg {
		t.Errorf("error message not eaual (result: %s, expect: %s)\n", err, expectErrMsg)
	}

	s = "[104 101 108 108 -1]"
	bytes, err = NewBytesFrom(s, GoByteArray)
	if err == nil {
		t.Error("error is nil")
	}
	_ = bytes.ByteArray()
	expectErrMsg = "bytestring: invalid value of -1 at 4"
	if err.Error() != expectErrMsg {
		t.Errorf("error message not eaual (result: %s, expect: %s)\n", err, expectErrMsg)
	}

	s = "hello"
	bytes, err = NewBytesFrom(s, GoByteArray)
	if err == nil {
		t.Error("error is nil")
	}
	_ = bytes.ByteArray()
	expectErrMsg = "strconv.Atoi: parsing \"hello\": invalid syntax"
	if err.Error() != expectErrMsg {
		t.Errorf("error message not eaual (result: %s, expect: %s)\n", err, expectErrMsg)
	}
}


func Test(t *testing.T) {
}
