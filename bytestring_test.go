package bytestring

import(
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
}


func TestStringToBytes(t *testing.T) {
	s := "hello"
	str := NewString(s)
	bytes := str.Bytes()
	expectBytes := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}
	if !reflect.DeepEqual(bytes, expectBytes) {
		t.Errorf("bytes not equal (result: %d, expect: %d)", bytes, expectBytes)
	}
}

func TestHexStringToBytes(t *testing.T) {
	s := "68656c6c6f"
	str, err := NewStringT(s, Hex)
	if err != nil {
		t.Errorf("error is not nil: %s\n", err)
	}
	bytes := str.Bytes()
	expectBytes := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}
	if !reflect.DeepEqual(bytes, expectBytes) {
		t.Errorf("bytes not equal (result: %d, expect: %d)", bytes, expectBytes)
	}

	s = "abcdefgh"
	str, err = NewStringT(s, Hex)
	if err == nil {
		t.Error("error is nil")
	}
	expectErrMsg := "encoding/hex: invalid byte: U+0067 'g'"
	if err.Error() != expectErrMsg {
		t.Errorf("error message not eaual (result: %s, expect: %s)\n", err, expectErrMsg)
	}

	s = "0abcdef"
	str, err = NewStringT(s, Hex)
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
	str, err := NewStringT(s, Base64)
	if err != nil {
		t.Errorf("error is not nil: %s\n", err)
	}
	bytes := str.Bytes()
	expectBytes := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f}
	if !reflect.DeepEqual(bytes, expectBytes) {
		t.Errorf("bytes not equal (result: %d, expect: %d)", bytes, expectBytes)
	}

	s = "aGVsbG8"
	str, err = NewStringT(s, Base64)
	if err == nil {
		t.Error("error is nil")
	}
	expectErrMsg := "illegal base64 data at input byte 4"
	if err.Error() != expectErrMsg {
		t.Errorf("error message not eaual (result: %s, expect: %s)\n", err, expectErrMsg)
	}

	s = "abcdefg-"
	str, err = NewStringT(s, Base64)
	if err == nil {
		t.Error("error is nil")
	}
	expectErrMsg = "illegal base64 data at input byte 7"
	if err.Error() != expectErrMsg {
		t.Errorf("error message not eaual (result: %s, expect: %s)\n", err, expectErrMsg)
	}

	// s = "0abcdef"
	// str, err = NewStringT(s, Hex)
	// if err == nil {
	// 	t.Error("error is nil")
	// }
	// expectErrMsg = "encoding/hex: odd length hex string"
	// if err.Error() != expectErrMsg {
	// 	t.Errorf("error message not eaual (result: %s, expect: %s)\n", err, expectErrMsg)
	// }
}
