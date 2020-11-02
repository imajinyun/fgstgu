package a08

import (
	"reflect"
	. "strconv"
	"testing"
)

type parseUint64Test struct {
	in  string
	out uint64
	err error
}

var parseUint64Tests = []parseUint64Test{
	{"", 0, ErrSyntax},
	{"0", 0, nil},
	{"1", 1, nil},
	{"12345", 12345, nil},
	{"012345", 12345, nil},
	{"12345x", 0, ErrSyntax},
	{"98765432100", 98765432100, nil},
	{"18446744073709551615", 1<<64 - 1, nil},
	{"18446744073709551616", 1<<64 - 1, ErrRange},
	{"18446744073709551620", 1<<64 - 1, ErrRange},
	{"1_2_3_4_5", 0, ErrSyntax}, // base=10 so no underscores allowed
	{"_12345", 0, ErrSyntax},
	{"1__2345", 0, ErrSyntax},
	{"12345_", 0, ErrSyntax},
}

type parseInt64Test struct {
	in  string
	out int64
	err error
}

var parseInt64Tests = []parseInt64Test{
	{"", 0, ErrSyntax},
	{"0", 0, nil},
	{"-0", 0, nil},
	{"1", 1, nil},
	{"-1", -1, nil},
	{"12345", 12345, nil},
	{"-12345", -12345, nil},
	{"012345", 12345, nil},
	{"-012345", -12345, nil},
	{"98765432100", 98765432100, nil},
	{"-98765432100", -98765432100, nil},
	{"9223372036854775807", 1<<63 - 1, nil},
	{"-9223372036854775807", -(1<<63 - 1), nil},
	{"9223372036854775808", 1<<63 - 1, ErrRange},
	{"-9223372036854775808", -1 << 63, nil},
	{"9223372036854775809", 1<<63 - 1, ErrRange},
	{"-9223372036854775809", -1 << 63, ErrRange},
	{"-1_2_3_4_5", 0, ErrSyntax}, // base=10 so no underscores allowed
	{"-_12345", 0, ErrSyntax},
	{"_12345", 0, ErrSyntax},
	{"1__2345", 0, ErrSyntax},
	{"12345_", 0, ErrSyntax},
}

type parseInt32Test struct {
	in  string
	out int32
	err error
}

var parseInt32Tests = []parseInt32Test{
	{"", 0, ErrSyntax},
	{"0", 0, nil},
	{"-0", 0, nil},
	{"1", 1, nil},
	{"-1", -1, nil},
	{"12345", 12345, nil},
	{"-12345", -12345, nil},
	{"012345", 12345, nil},
	{"-012345", -12345, nil},
	{"12345x", 0, ErrSyntax},
	{"-12345x", 0, ErrSyntax},
	{"987654321", 987654321, nil},
	{"-987654321", -987654321, nil},
	{"2147483647", 1<<31 - 1, nil},
	{"-2147483647", -(1<<31 - 1), nil},
	{"2147483648", 1<<31 - 1, ErrRange},
	{"-2147483648", -1 << 31, nil},
	{"2147483649", 1<<31 - 1, ErrRange},
	{"-2147483649", -1 << 31, ErrRange},
	{"-1_2_3_4_5", 0, ErrSyntax}, // base=10 so no underscores allowed
	{"-_12345", 0, ErrSyntax},
	{"_12345", 0, ErrSyntax},
	{"1__2345", 0, ErrSyntax},
	{"12345_", 0, ErrSyntax},
}

type parseUint32Test struct {
	in  string
	out uint32
	err error
}

var parseUint32Tests = []parseUint32Test{
	{"", 0, ErrSyntax},
	{"0", 0, nil},
	{"1", 1, nil},
	{"12345", 12345, nil},
	{"012345", 12345, nil},
	{"12345x", 0, ErrSyntax},
	{"987654321", 987654321, nil},
	{"4294967295", 1<<32 - 1, nil},
	{"4294967296", 1<<32 - 1, ErrRange},
	{"1_2_3_4_5", 0, ErrSyntax}, // base=10 so no underscores allowed
	{"_12345", 0, ErrSyntax},
	{"_12345", 0, ErrSyntax},
	{"1__2345", 0, ErrSyntax},
	{"12345_", 0, ErrSyntax},
}

func init() {
	for i := range parseUint64Tests {
		test := &parseUint64Tests[i]
		if test.err != nil {
			test.err = &NumError{"ParseUint", test.in, test.err}
		}
	}
	for i := range parseInt64Tests {
		test := &parseInt64Tests[i]
		if test.err != nil {
			test.err = &NumError{"ParseInt", test.in, test.err}
		}
	}
	for i := range parseUint32Tests {
		test := &parseUint32Tests[i]
		if test.err != nil {
			test.err = &NumError{"ParseUint", test.in, test.err}
		}
	}
	for i := range parseInt32Tests {
		test := &parseInt32Tests[i]
		if test.err != nil {
			test.err = &NumError{"ParseInt", test.in, test.err}
		}
	}
}

func TestParseUint32(t *testing.T) {
	for i := range parseUint32Tests {
		test := &parseUint32Tests[i]
		out, err := ParseUint(test.in, 10, 32)
		if uint64(test.out) != out || !reflect.DeepEqual(test.err, err) {
			t.Errorf("ParseUint(%q, 10, 32) = %v, %v want %v, %v",
				test.in, out, err, test.out, test.err)
		}
	}
}

func TestParseUint64(t *testing.T) {
	for i := range parseUint64Tests {
		test := &parseUint64Tests[i]
		out, err := ParseUint(test.in, 10, 64)
		if test.out != out || !reflect.DeepEqual(test.err, err) {
			t.Errorf("ParseUint(%q, 10, 64) = %v, %v want %v, %v",
				test.in, out, err, test.out, test.err)
		}
	}
}

func TestParseInt(t *testing.T) {
	switch IntSize {
	case 32:
		for i := range parseInt32Tests {
			test := &parseInt32Tests[i]
			out, err := ParseInt(test.in, 10, 0)
			if int64(test.out) != out || !reflect.DeepEqual(test.err, err) {
				t.Errorf("ParseInt(%q, 10, 0) = %v, %v want %v, %v",
					test.in, out, err, test.out, test.err)
			}
		}
	case 64:
		for i := range parseInt64Tests {
			test := &parseInt64Tests[i]
			out, err := ParseInt(test.in, 10, 0)
			if test.out != out || !reflect.DeepEqual(test.err, err) {
				t.Errorf("ParseInt(%q, 10, 0) = %v, %v want %v, %v",
					test.in, out, err, test.out, test.err)
			}
		}
	}
}

func TestAtoi(t *testing.T) {
	switch IntSize {
	case 32:
		for i := range parseInt32Tests {
			test := &parseInt32Tests[i]
			out, err := Atoi(test.in)
			var testErr error
			if test.err != nil {
				testErr = &NumError{"Atoi", test.in, test.err.(*NumError).Err}
			}
			if int(test.out) != out || !reflect.DeepEqual(testErr, err) {
				t.Errorf("Atoi(%q) = %v, %v want %v, %v",
					test.in, out, err, test.out, testErr)
			}
		}
	case 64:
		for i := range parseInt64Tests {
			test := &parseInt64Tests[i]
			out, err := Atoi(test.in)
			var testErr error
			if test.err != nil {
				testErr = &NumError{"Atoi", test.in, test.err.(*NumError).Err}
			}
			if test.out != int64(out) || !reflect.DeepEqual(testErr, err) {
				t.Errorf("Atoi(%q) = %v, %v want %v, %v",
					test.in, out, err, test.out, testErr)
			}
		}
	}
}
