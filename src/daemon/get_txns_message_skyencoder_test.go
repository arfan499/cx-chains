// Code generated by github.com/SkycoinProject/skyencoder. DO NOT EDIT.
package daemon

import (
	"bytes"
	"fmt"
	mathrand "math/rand"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/SkycoinProject/cx-chains/src/cipher/encoder"
	"github.com/SkycoinProject/encodertest"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func newEmptyGetTxnsMessageForEncodeTest() *GetTxnsMessage {
	var obj GetTxnsMessage
	return &obj
}

func newRandomGetTxnsMessageForEncodeTest(t *testing.T, rand *mathrand.Rand) *GetTxnsMessage {
	var obj GetTxnsMessage
	err := encodertest.PopulateRandom(&obj, rand, encodertest.PopulateRandomOptions{
		MaxRandLen: 4,
		MinRandLen: 1,
	})
	if err != nil {
		t.Fatalf("encodertest.PopulateRandom failed: %v", err)
	}
	return &obj
}

func newRandomZeroLenGetTxnsMessageForEncodeTest(t *testing.T, rand *mathrand.Rand) *GetTxnsMessage {
	var obj GetTxnsMessage
	err := encodertest.PopulateRandom(&obj, rand, encodertest.PopulateRandomOptions{
		MaxRandLen:    0,
		MinRandLen:    0,
		EmptySliceNil: false,
		EmptyMapNil:   false,
	})
	if err != nil {
		t.Fatalf("encodertest.PopulateRandom failed: %v", err)
	}
	return &obj
}

func newRandomZeroLenNilGetTxnsMessageForEncodeTest(t *testing.T, rand *mathrand.Rand) *GetTxnsMessage {
	var obj GetTxnsMessage
	err := encodertest.PopulateRandom(&obj, rand, encodertest.PopulateRandomOptions{
		MaxRandLen:    0,
		MinRandLen:    0,
		EmptySliceNil: true,
		EmptyMapNil:   true,
	})
	if err != nil {
		t.Fatalf("encodertest.PopulateRandom failed: %v", err)
	}
	return &obj
}

func testSkyencoderGetTxnsMessage(t *testing.T, obj *GetTxnsMessage) {
	isEncodableField := func(f reflect.StructField) bool {
		// Skip unexported fields
		if f.PkgPath != "" {
			return false
		}

		// Skip fields disabled with and enc:"- struct tag
		tag := f.Tag.Get("enc")
		return !strings.HasPrefix(tag, "-,") && tag != "-"
	}

	hasOmitEmptyField := func(obj interface{}) bool {
		v := reflect.ValueOf(obj)
		switch v.Kind() {
		case reflect.Ptr:
			v = v.Elem()
		}

		switch v.Kind() {
		case reflect.Struct:
			t := v.Type()
			n := v.NumField()
			f := t.Field(n - 1)
			tag := f.Tag.Get("enc")
			return isEncodableField(f) && strings.Contains(tag, ",omitempty")
		default:
			return false
		}
	}

	// returns the number of bytes encoded by an omitempty field on a given object
	omitEmptyLen := func(obj interface{}) uint64 {
		if !hasOmitEmptyField(obj) {
			return 0
		}

		v := reflect.ValueOf(obj)
		switch v.Kind() {
		case reflect.Ptr:
			v = v.Elem()
		}

		switch v.Kind() {
		case reflect.Struct:
			n := v.NumField()
			f := v.Field(n - 1)
			if f.Len() == 0 {
				return 0
			}
			return uint64(4 + f.Len())

		default:
			return 0
		}
	}

	// encodeSize

	n1 := encoder.Size(obj)
	n2 := encodeSizeGetTxnsMessage(obj)

	if uint64(n1) != n2 {
		t.Fatalf("encoder.Size() != encodeSizeGetTxnsMessage() (%d != %d)", n1, n2)
	}

	// Encode

	// encoder.Serialize
	data1 := encoder.Serialize(obj)

	// Encode
	data2, err := encodeGetTxnsMessage(obj)
	if err != nil {
		t.Fatalf("encodeGetTxnsMessage failed: %v", err)
	}
	if uint64(len(data2)) != n2 {
		t.Fatal("encodeGetTxnsMessage produced bytes of unexpected length")
	}
	if len(data1) != len(data2) {
		t.Fatalf("len(encoder.Serialize()) != len(encodeGetTxnsMessage()) (%d != %d)", len(data1), len(data2))
	}

	// EncodeToBuffer
	data3 := make([]byte, n2+5)
	if err := encodeGetTxnsMessageToBuffer(data3, obj); err != nil {
		t.Fatalf("encodeGetTxnsMessageToBuffer failed: %v", err)
	}

	if !bytes.Equal(data1, data2) {
		t.Fatal("encoder.Serialize() != encode[1]s()")
	}

	// Decode

	// encoder.DeserializeRaw
	var obj2 GetTxnsMessage
	if n, err := encoder.DeserializeRaw(data1, &obj2); err != nil {
		t.Fatalf("encoder.DeserializeRaw failed: %v", err)
	} else if n != uint64(len(data1)) {
		t.Fatalf("encoder.DeserializeRaw failed: %v", encoder.ErrRemainingBytes)
	}
	if !cmp.Equal(*obj, obj2, cmpopts.EquateEmpty(), encodertest.IgnoreAllUnexported()) {
		t.Fatal("encoder.DeserializeRaw result wrong")
	}

	// Decode
	var obj3 GetTxnsMessage
	if n, err := decodeGetTxnsMessage(data2, &obj3); err != nil {
		t.Fatalf("decodeGetTxnsMessage failed: %v", err)
	} else if n != uint64(len(data2)) {
		t.Fatalf("decodeGetTxnsMessage bytes read length should be %d, is %d", len(data2), n)
	}
	if !cmp.Equal(obj2, obj3, cmpopts.EquateEmpty(), encodertest.IgnoreAllUnexported()) {
		t.Fatal("encoder.DeserializeRaw() != decodeGetTxnsMessage()")
	}

	// Decode, excess buffer
	var obj4 GetTxnsMessage
	n, err := decodeGetTxnsMessage(data3, &obj4)
	if err != nil {
		t.Fatalf("decodeGetTxnsMessage failed: %v", err)
	}

	if hasOmitEmptyField(&obj4) && omitEmptyLen(&obj4) == 0 {
		// 4 bytes read for the omitEmpty length, which should be zero (see the 5 bytes added above)
		if n != n2+4 {
			t.Fatalf("decodeGetTxnsMessage bytes read length should be %d, is %d", n2+4, n)
		}
	} else {
		if n != n2 {
			t.Fatalf("decodeGetTxnsMessage bytes read length should be %d, is %d", n2, n)
		}
	}
	if !cmp.Equal(obj2, obj4, cmpopts.EquateEmpty(), encodertest.IgnoreAllUnexported()) {
		t.Fatal("encoder.DeserializeRaw() != decodeGetTxnsMessage()")
	}

	// DecodeExact
	var obj5 GetTxnsMessage
	if err := decodeGetTxnsMessageExact(data2, &obj5); err != nil {
		t.Fatalf("decodeGetTxnsMessage failed: %v", err)
	}
	if !cmp.Equal(obj2, obj5, cmpopts.EquateEmpty(), encodertest.IgnoreAllUnexported()) {
		t.Fatal("encoder.DeserializeRaw() != decodeGetTxnsMessage()")
	}

	// Check that the bytes read value is correct when providing an extended buffer
	if !hasOmitEmptyField(&obj3) || omitEmptyLen(&obj3) > 0 {
		padding := []byte{0xFF, 0xFE, 0xFD, 0xFC}
		data4 := append(data2[:], padding...)
		if n, err := decodeGetTxnsMessage(data4, &obj3); err != nil {
			t.Fatalf("decodeGetTxnsMessage failed: %v", err)
		} else if n != uint64(len(data2)) {
			t.Fatalf("decodeGetTxnsMessage bytes read length should be %d, is %d", len(data2), n)
		}
	}
}

func TestSkyencoderGetTxnsMessage(t *testing.T) {
	rand := mathrand.New(mathrand.NewSource(time.Now().Unix()))

	type testCase struct {
		name string
		obj  *GetTxnsMessage
	}

	cases := []testCase{
		{
			name: "empty object",
			obj:  newEmptyGetTxnsMessageForEncodeTest(),
		},
	}

	nRandom := 10

	for i := 0; i < nRandom; i++ {
		cases = append(cases, testCase{
			name: fmt.Sprintf("randomly populated object %d", i),
			obj:  newRandomGetTxnsMessageForEncodeTest(t, rand),
		})
		cases = append(cases, testCase{
			name: fmt.Sprintf("randomly populated object %d with zero length variable length contents", i),
			obj:  newRandomZeroLenGetTxnsMessageForEncodeTest(t, rand),
		})
		cases = append(cases, testCase{
			name: fmt.Sprintf("randomly populated object %d with zero length variable length contents set to nil", i),
			obj:  newRandomZeroLenNilGetTxnsMessageForEncodeTest(t, rand),
		})
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			testSkyencoderGetTxnsMessage(t, tc.obj)
		})
	}
}

func decodeGetTxnsMessageExpectError(t *testing.T, buf []byte, expectedErr error) {
	var obj GetTxnsMessage
	if _, err := decodeGetTxnsMessage(buf, &obj); err == nil {
		t.Fatal("decodeGetTxnsMessage: expected error, got nil")
	} else if err != expectedErr {
		t.Fatalf("decodeGetTxnsMessage: expected error %q, got %q", expectedErr, err)
	}
}

func decodeGetTxnsMessageExactExpectError(t *testing.T, buf []byte, expectedErr error) {
	var obj GetTxnsMessage
	if err := decodeGetTxnsMessageExact(buf, &obj); err == nil {
		t.Fatal("decodeGetTxnsMessageExact: expected error, got nil")
	} else if err != expectedErr {
		t.Fatalf("decodeGetTxnsMessageExact: expected error %q, got %q", expectedErr, err)
	}
}

func testSkyencoderGetTxnsMessageDecodeErrors(t *testing.T, k int, tag string, obj *GetTxnsMessage) {
	isEncodableField := func(f reflect.StructField) bool {
		// Skip unexported fields
		if f.PkgPath != "" {
			return false
		}

		// Skip fields disabled with and enc:"- struct tag
		tag := f.Tag.Get("enc")
		return !strings.HasPrefix(tag, "-,") && tag != "-"
	}

	numEncodableFields := func(obj interface{}) int {
		v := reflect.ValueOf(obj)
		switch v.Kind() {
		case reflect.Ptr:
			v = v.Elem()
		}

		switch v.Kind() {
		case reflect.Struct:
			t := v.Type()

			n := 0
			for i := 0; i < v.NumField(); i++ {
				f := t.Field(i)
				if !isEncodableField(f) {
					continue
				}
				n++
			}
			return n
		default:
			return 0
		}
	}

	hasOmitEmptyField := func(obj interface{}) bool {
		v := reflect.ValueOf(obj)
		switch v.Kind() {
		case reflect.Ptr:
			v = v.Elem()
		}

		switch v.Kind() {
		case reflect.Struct:
			t := v.Type()
			n := v.NumField()
			f := t.Field(n - 1)
			tag := f.Tag.Get("enc")
			return isEncodableField(f) && strings.Contains(tag, ",omitempty")
		default:
			return false
		}
	}

	// returns the number of bytes encoded by an omitempty field on a given object
	omitEmptyLen := func(obj interface{}) uint64 {
		if !hasOmitEmptyField(obj) {
			return 0
		}

		v := reflect.ValueOf(obj)
		switch v.Kind() {
		case reflect.Ptr:
			v = v.Elem()
		}

		switch v.Kind() {
		case reflect.Struct:
			n := v.NumField()
			f := v.Field(n - 1)
			if f.Len() == 0 {
				return 0
			}
			return uint64(4 + f.Len())

		default:
			return 0
		}
	}

	n := encodeSizeGetTxnsMessage(obj)
	buf, err := encodeGetTxnsMessage(obj)
	if err != nil {
		t.Fatalf("encodeGetTxnsMessage failed: %v", err)
	}

	// A nil buffer cannot decode, unless the object is a struct with a single omitempty field
	if hasOmitEmptyField(obj) && numEncodableFields(obj) > 1 {
		t.Run(fmt.Sprintf("%d %s buffer underflow nil", k, tag), func(t *testing.T) {
			decodeGetTxnsMessageExpectError(t, nil, encoder.ErrBufferUnderflow)
		})

		t.Run(fmt.Sprintf("%d %s exact buffer underflow nil", k, tag), func(t *testing.T) {
			decodeGetTxnsMessageExactExpectError(t, nil, encoder.ErrBufferUnderflow)
		})
	}

	// Test all possible truncations of the encoded byte array, but skip
	// a truncation that would be valid where omitempty is removed
	skipN := n - omitEmptyLen(obj)
	for i := uint64(0); i < n; i++ {
		if i == skipN {
			continue
		}

		t.Run(fmt.Sprintf("%d %s buffer underflow bytes=%d", k, tag, i), func(t *testing.T) {
			decodeGetTxnsMessageExpectError(t, buf[:i], encoder.ErrBufferUnderflow)
		})

		t.Run(fmt.Sprintf("%d %s exact buffer underflow bytes=%d", k, tag, i), func(t *testing.T) {
			decodeGetTxnsMessageExactExpectError(t, buf[:i], encoder.ErrBufferUnderflow)
		})
	}

	// Append 5 bytes for omit empty with a 0 length prefix, to cause an ErrRemainingBytes.
	// If only 1 byte is appended, the decoder will try to read the 4-byte length prefix,
	// and return an ErrBufferUnderflow instead
	if hasOmitEmptyField(obj) {
		buf = append(buf, []byte{0, 0, 0, 0, 0}...)
	} else {
		buf = append(buf, 0)
	}

	t.Run(fmt.Sprintf("%d %s exact buffer remaining bytes", k, tag), func(t *testing.T) {
		decodeGetTxnsMessageExactExpectError(t, buf, encoder.ErrRemainingBytes)
	})
}

func TestSkyencoderGetTxnsMessageDecodeErrors(t *testing.T) {
	rand := mathrand.New(mathrand.NewSource(time.Now().Unix()))
	n := 10

	for i := 0; i < n; i++ {
		emptyObj := newEmptyGetTxnsMessageForEncodeTest()
		fullObj := newRandomGetTxnsMessageForEncodeTest(t, rand)
		testSkyencoderGetTxnsMessageDecodeErrors(t, i, "empty", emptyObj)
		testSkyencoderGetTxnsMessageDecodeErrors(t, i, "full", fullObj)
	}
}
