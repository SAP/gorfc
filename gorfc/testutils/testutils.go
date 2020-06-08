package testutils

import "crypto/rand"

var RFC_MATH = map[string]interface{}{

	"RFC_INT1": map[string]uint8{
		"MIN": 0,
		"MAX": 255,
	},
	"RFC_INT2": map[string]int16{
		"MIN": -32768,
		"MAX": 32767,
	},
	"RFC_INT4": map[string]int32{
		"MIN": -2147483648,
		"MAX": 2147483647,
	},
	"RFC_INT8": map[string]int64{
		"MIN": -9223372036854775808,
		"MAX": 9223372036854775807,
	},
	"FLOAT": map[string]interface{}{
		"NEG": map[string]string{
			"MIN": "-2.2250738585072014E-308",
			"MAX": "-1.7976931348623157E+308",
		},
		"POS": map[string]string{
			"MIN": "2.2250738585072014E-308",
			"MAX": "1.7976931348623157E+308",
		},
	},
	"DECF16": map[string]interface{}{
		"NEG": map[string]string{
			"MIN": "-1E-383",
			"MAX": "-9.999999999999999E+384",
		},
		"POS": map[string]string{
			"MIN": "1E-383",
			"MAX": "9.999999999999999E+384",
		},
	},
	"DECF34": map[string]interface{}{
		"NEG": map[string]string{
			"MIN": "-1E-6143",
			"MAX": "-9.999999999999999999999999999999999E+6144",
		},
		"POS": map[string]string{
			"MIN": "1E-6143",
			"MAX": "9.999999999999999999999999999999999E+6144",
		},
	},
	"DATE": map[string]string{
		"MIN": "00010101",
		"MAX": "99991231",
	},
	"TIME": map[string]string{
		"MIN": "000000",
		"MAX": "235959",
	},
	"UTCLONG": map[string]string{
		"MIN":     "0001-01-01T00:00:00.0000000",
		"MAX":     "9999-12-31T23:59:59.9999999",
		"INITIAL": "0000-00-00T00:00:00.0000000",
	},
}

func XBytes(length uint) []byte {
	var xbytes = make([]byte, length)
	rand.Read(xbytes)
	return xbytes
}
