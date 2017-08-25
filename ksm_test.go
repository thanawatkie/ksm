package ksm

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/easonlin404/ksm/d"
	"github.com/stretchr/testify/assert"
)

type spcTest struct {
	filePath      string
	outFilePath   string
	playload_size int
	iv            []byte
	encrypted_key []byte

	ttls []TLLVBlock
}

var pub = `-----BEGIN CERTIFICATE-----
MIIDfTCCAmWgAwIBAgIIboBT3GOPJ50wDQYJKoZIhvcNAQEFBQAwfTELMAkGA1UE
BhMCVVMxEzARBgNVBAoMCkFwcGxlIEluYy4xJjAkBgNVBAsMHUFwcGxlIENlcnRp
ZmljYXRpb24gQXV0aG9yaXR5MTEwLwYDVQQDDChEUk0gVGVjaG5vbG9naWVzIENl
cnRpZmljYXRpb24gQXV0aG9yaXR5MB4XDTExMTAxODAxNTcyMloXDTEzMTAxNzAx
NTcyMlowRjERMA8GA1UEAwwIUGFydG5lcjIxETAPBgNVBAsMCFBhcnRuZXIyMREw
DwYDVQQKDAhQYXJ0bmVyMjELMAkGA1UEBhMCVVMwgZ8wDQYJKoZIhvcNAQEBBQAD
gY0AMIGJAoGBALReAQ24va6MquxUkOyrVLE0vjc3rv3a16qndKKKGL6afpkN19xc
/cWw9A2W0FCSJYgkY+iyhGPAO4BLWe0QSonJz08GdeEMS2wmj87h8PLe6Yyu8Ida
3hH+snc7hv2bxX5AI72ETSQWlElky3tHLCYV2tqbTW4BGQZvvE4LfM+tAgMBAAGj
gbswgbgwJwYLKoZIhvdjZAYNAQMEGAGAgEeXuoURG4c6qSNQztlZmgq9dM3kTzAv
BgsqhkiG92NkBg0BBAQgAaWxaRPd6O3itrSL3iqhd3fcpUMMhDQTIebXMN1IfmQw
HQYDVR0OBBYEFDdUHOfoNQC1nqz9IzDvC/WJR1ssMAwGA1UdEwEB/wQCMAAwHwYD
VR0jBBgwFoAU6rShbWWjpF5JZST6HCRnrVoa0DMwDgYDVR0PAQH/BAQDAgUgMA0G
CSqGSIb3DQEBBQUAA4IBAQB4gFunl0sKeqGza5fdDd9Dj0O+rutFPqIFFLY60Qgl
jQdkzaHegMBqoON3I2KWRxgOeaewArmlgZjK8LoTv++HALB1Thf7N9AulyWVCg7J
i/hFKhTNpbNWBXSkKYn1QpcnohAnjLsrNED7R0b4A7z1yBhUjU96uRsKU+Dd6St9
XMlvvK49iSWNadfz7IictPrOjvHj4hRzepE43U5unevsth2FXu553LMCZw7gy4h9
IMYU4NZSWhf5z+wYpjtzYxdoqynjvihqFdGqYDC2drzpLLhaCXZhZUq2D1mXoQaY
6URsYkp6FRwIAx++KnIwE7Q3kK6s+5sRpKK4zZ0y0O9Z
-----END CERTIFICATE-----`

var pri = `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQC0XgENuL2ujKrsVJDsq1SxNL43N6792teqp3Siihi+mn6ZDdfc
XP3FsPQNltBQkiWIJGPosoRjwDuAS1ntEEqJyc9PBnXhDEtsJo/O4fDy3umMrvCH
Wt4R/rJ3O4b9m8V+QCO9hE0kFpRJZMt7RywmFdram01uARkGb7xOC3zPrQIDAQAB
AoGBAIO+vkpFjNd4jEi/pHQa2WvuuJogpENsnGdclYc8E8L1mk81m1ys1/iUvk9G
v7Z6acu9uPR5oNYzzcJyR6cvZSFxtGIZnWNdDOAB71b+YqMvj3lr6MgUdMUgUfxZ
EDXLEhIoVzyQWIt+f6hjSG/hzyw+Jglo4ogCWPsV3S6UG2WBAkEA5HPddGIUa34k
2/EGQqyCAo4VYlCUdCFTp9+eFIUedequgsSIZhgblT+FSvMPYARuG/ywLoOivRy1
dFl0dIB1sQJBAModyMskK0r312kro+URq8VxlwwY0fv2rF1aS0/clQUw5OH/OxEn
Dgz3l3PNTXDCcQDh9wyEZV0SgIp7SYCDrL0CQEo8HEolVN1ZMEEIITCpPdX2tZws
8xCJg9WZJJUmbK+EgxCbLHeAffYRng6szOI2jlEp21ZCEC/DlHMqXl09IQECQGSn
EoC/oWOzKy4v0m3YL/+iwsL+dUwSGuJefhTmV7v/DmzRixvOpDum7WB5BDC8VERJ
Q5uTL1t7RFIydXcvm80CQH/E17mWT66PPeqloAfSH/5tJyak2gagkuFnMh779JRF
rl5YIIiAh+q5DkcjWw6eni5O4+UuwXRp29vZaxmDlIE=
-----END RSA PRIVATE KEY-----`

var spcContainerTests = []spcTest{
	{"testdata/FPS/spc1.bin",
		"testdata/FPS/ckc1.bin",
		2688,
		[]byte{0x5d, 0x16, 0x44, 0xea, 0xec, 0x11, 0xf9, 0x83, 0x14, 0x75, 0x41, 0xe4, 0x6e, 0xeb, 0x27, 0x74},
		[]byte{0x92, 0x66, 0x48, 0xb9, 0x86, 0x1e, 0xc0, 0x47, 0x1b, 0xa2, 0x17, 0x58, 0x85, 0x1c, 0x3d, 0xda, 0x31, 0xc9, 0x3b, 0x1d, 0xd6, 0x1, 0xaa, 0x4e, 0xad, 0x44, 0x15, 0xa2, 0x7, 0x59, 0xaa, 0xb9, 0xa6, 0xd8, 0x9f, 0x55, 0x13, 0x85, 0x85, 0x6e, 0x73, 0x57, 0x17, 0x29, 0xdf, 0x2f, 0x1d, 0x46, 0xd2, 0x5c, 0x13, 0xda, 0x2a, 0xd7, 0x5d, 0x0, 0xfd, 0x34, 0x13, 0xeb, 0xd9, 0x6c, 0xa4, 0x7d, 0x2, 0x95, 0x5c, 0x56, 0x9f, 0x7f, 0xab, 0x40, 0xf1, 0xa7, 0xfb, 0x23, 0x41, 0x41, 0x67, 0xa6, 0x53, 0xea, 0xbd, 0xf1, 0xad, 0x28, 0x3d, 0xf5, 0xe0, 0x7e, 0x7c, 0xf4, 0xaa, 0x2f, 0xba, 0xc6, 0x4f, 0x1d, 0x46, 0xf, 0xdf, 0x9a, 0x21, 0xee, 0xb2, 0x7a, 0x7f, 0x60, 0x72, 0x78, 0x53, 0xa4, 0x14, 0xc1, 0xc4, 0x50, 0xc5, 0x25, 0xe8, 0xda, 0xb6, 0xa3, 0xf1, 0x3c, 0xfa, 0x57, 0x17, 0x1a},

		[]TLLVBlock{
			{
				Tag:         Tag_SessionKey_R1_integrity,
				BlockLength: 0x40,
				ValueLength: 0x10,
				Value:       []byte{0x54, 0xa1, 0x6b, 0xe0, 0x13, 0x7e, 0xf2, 0x59, 0xab, 0x3e, 0x4f, 0xc7, 0x96, 0x90, 0x82, 0x5f},
			},
			{
				Tag:         Tag_SessionKey_R1,
				BlockLength: 0x100,
				ValueLength: 0x70,
				Value:       []byte{0x4f, 0x45, 0xd8, 0x5c, 0xe2, 0x62, 0x73, 0x10, 0x1a, 0x97, 0xf3, 0x30, 0x81, 0xc1, 0xd0, 0x4a, 0x93, 0xb2, 0xdd, 0x3, 0x55, 0xe3, 0x63, 0x72, 0x9d, 0x92, 0xa4, 0x5a, 0x45, 0xce, 0x8d, 0x25, 0x8b, 0xc, 0x8, 0xaa, 0x65, 0x1c, 0x9, 0x64, 0x97, 0x6b, 0xf0, 0x94, 0x4d, 0x28, 0x25, 0xf3, 0xac, 0x8d, 0xde, 0x7e, 0xd2, 0x31, 0x4f, 0xa0, 0xef, 0x3f, 0xb4, 0x5b, 0x97, 0xa2, 0x26, 0xe8, 0xc5, 0x36, 0x6d, 0xef, 0xe5, 0xf1, 0xe1, 0x2b, 0xd7, 0xb7, 0x21, 0x98, 0xa4, 0xa8, 0xf2, 0x65, 0x3a, 0xe, 0xf0, 0xde, 0x8c, 0x37, 0xa4, 0x7c, 0x3c, 0x40, 0xf0, 0x12, 0xe1, 0x5c, 0x8b, 0x59, 0x3d, 0xf1, 0x2d, 0x4b, 0x1, 0x60, 0x3a, 0x97, 0x35, 0x7e, 0x6a, 0xe0, 0xa1, 0x1c, 0xa3, 0xe3},
			},
			{
				Tag:         Tag_AntiReplaySeed,
				BlockLength: 0xd0,
				ValueLength: 0x10,
				Value:       []byte{0xf3, 0xc6, 0x9d, 0x1e, 0x8c, 0xc4, 0x27, 0x5a, 0x6d, 0x32, 0x86, 0xd3, 0x32, 0x61, 0x3e, 0x13},
			},
			{
				Tag:         Tag_R2,
				BlockLength: 0xb0,
				ValueLength: 0x15,
				Value:       []byte{0x11, 0xf7, 0xbe, 0x61, 0x2c, 0xa9, 0x5e, 0xf5, 0xe0, 0x7, 0xce, 0x51, 0x89, 0x6a, 0xe4, 0x50, 0x2c, 0xa3, 0xd8, 0x80, 0x1b},
			},
			{
				Tag:         Tag_AssetID,
				BlockLength: 0x80,
				ValueLength: 0x12,
				Value:       []byte{0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			},
			{
				Tag:         Tag_TransactionID,
				BlockLength: 0x70,
				ValueLength: 0x8,
				Value:       []byte{0x14, 0x73, 0xe5, 0xcc, 0x53, 0xe1, 0xe5, 0xd6},
			},
			{
				Tag:         Tag_ProtocolVersionUsed,
				BlockLength: 0xc0,
				ValueLength: 0x4,
				Value:       []byte{0x0, 0x0, 0x0, 0x1},
			},
			{
				Tag:         Tag_ProtocolVersionsSupported,
				BlockLength: 0xa0,
				ValueLength: 0x4,
				Value:       []byte{0x0, 0x0, 0x0, 0x1},
			},
			{
				Tag:         Tag_ReturnRequest,
				BlockLength: 0x60,
				ValueLength: 0x38,
				Value:       []byte{0x1b, 0xf7, 0xf5, 0x3f, 0x5d, 0x5d, 0x5a, 0x1f, 0x47, 0xaa, 0x7a, 0xd3, 0x44, 0x5, 0x77, 0xde, 0xf9, 0x11, 0xf0, 0x4d, 0xa5, 0x4b, 0xf5, 0x99, 0xba, 0x8, 0xcc, 0x74, 0xda, 0xc9, 0x17, 0x6d, 0x13, 0xd, 0x99, 0x4c, 0xb8, 0x94, 0xb9, 0xe3, 0x66, 0xc8, 0x23, 0xf3, 0x79, 0xb8, 0x7b, 0xb5, 0x18, 0xd4, 0x2c, 0x5f, 0x8e, 0x54, 0x5a, 0x4b},
			},
		},
	},
	{"testdata/FPS/spc2.bin",
		"testdata/FPS/ckc2.bin",
		1984,
		[]byte{0x88, 0x88, 0xc5, 0xa9, 0x42, 0x13, 0x3, 0x23, 0xaa, 0x4c, 0x9f, 0x1d, 0xa7, 0x7c, 0x88, 0x88},
		[]byte{0x1f, 0x60, 0x34, 0x70, 0x48, 0x9d, 0xde, 0x55, 0xfb, 0xd2, 0x17, 0x90, 0x66, 0x29, 0x5, 0xf9, 0xfc, 0x7c, 0xf0, 0x41, 0x94, 0xf6, 0x49, 0xa7, 0x55, 0x1e, 0x19, 0xd7, 0xf2, 0x14, 0x5b, 0xfe, 0xcb, 0x8a, 0x7, 0xf9, 0x41, 0xcc, 0xb8, 0x9a, 0xce, 0xc0, 0xdc, 0x7c, 0xf, 0x8, 0x2c, 0x8c, 0xd5, 0x55, 0x60, 0xad, 0xb5, 0x6a, 0xae, 0x1c, 0xd9, 0x6e, 0xb2, 0x44, 0x43, 0x41, 0xeb, 0xbe, 0x84, 0x51, 0x8a, 0x93, 0x19, 0x30, 0x55, 0x6a, 0x88, 0x1f, 0xe2, 0x50, 0xf5, 0x7d, 0x21, 0x97, 0xea, 0x23, 0x9a, 0x79, 0xa8, 0x3c, 0xb4, 0x7d, 0xc8, 0x46, 0x1e, 0x4b, 0xe5, 0xeb, 0xee, 0xc7, 0x4b, 0x4e, 0x40, 0x79, 0xa9, 0x77, 0x29, 0xd9, 0xa5, 0x6d, 0xcd, 0x9f, 0x20, 0x54, 0xc3, 0x3f, 0xe0, 0xac, 0xeb, 0x61, 0x77, 0xea, 0x9e, 0x79, 0xd, 0x1a, 0xed, 0x62, 0x87, 0xbf, 0x8b, 0xb},

		[]TLLVBlock{
			{
				Tag:         Tag_SessionKey_R1_integrity,
				BlockLength: 0x90,
				ValueLength: 0x10,
				Value:       []byte{0x11, 0x66, 0x4f, 0xe9, 0x6e, 0x8c, 0xb8, 0x36, 0xaa, 0x2e, 0xb0, 0x64, 0xe8, 0x65, 0x66, 0x11},
			},
			{
				Tag:         Tag_SessionKey_R1,
				BlockLength: 0x90,
				ValueLength: 0x70,
				Value:       []byte{0xaa, 0xaa, 0xbb, 0xbb, 0xcc, 0xcc, 0x82, 0x30, 0xe8, 0x9a, 0xce, 0xfa, 0xdd, 0xdd, 0xaa, 0xaa, 0x79, 0x2e, 0xf7, 0x42, 0xf, 0x47, 0x9b, 0x80, 0x56, 0x7e, 0x4f, 0x8b, 0xe8, 0x3e, 0xac, 0x8b, 0xfe, 0xb, 0x75, 0x51, 0x4a, 0x7c, 0xda, 0x90, 0xe8, 0xef, 0x55, 0xac, 0xf4, 0xf, 0x3c, 0x59, 0xcd, 0xfc, 0x7d, 0xdd, 0x34, 0xcb, 0x92, 0x29, 0x73, 0x3d, 0x3, 0xcd, 0x59, 0xbd, 0x1a, 0x6, 0xba, 0x21, 0xe9, 0x3e, 0x2a, 0xe, 0xf4, 0x50, 0x25, 0xb1, 0x14, 0x7c, 0xb5, 0x3, 0x44, 0x96, 0x5e, 0xa0, 0x66, 0x18, 0x35, 0xf, 0xb4, 0x5b, 0x4a, 0xb6, 0xf7, 0xb5, 0xb6, 0xd1, 0x6c, 0x3f, 0x9a, 0x3d, 0x9f, 0x7c, 0x38, 0xa4, 0xaa, 0x97, 0xea, 0x32, 0x66, 0x3d, 0x6d, 0x4a, 0xac, 0x11},
			},
			{
				Tag:         Tag_AntiReplaySeed,
				BlockLength: 0xa0,
				ValueLength: 0x10,
				Value:       []byte{0x2d, 0x9e, 0x8f, 0x28, 0x97, 0xc0, 0xc7, 0xa7, 0x17, 0xe0, 0x95, 0x5d, 0x3c, 0xea, 0xfa, 0x2d},
			},
			{
				Tag:         Tag_R2,
				BlockLength: 0xb0,
				ValueLength: 0x15,
				Value:       []byte{0x11, 0xf7, 0xbe, 0x61, 0x2c, 0xa9, 0x5e, 0xf5, 0xe0, 0x7, 0xce, 0x51, 0x89, 0x6a, 0xe4, 0x50, 0x2c, 0xa3, 0xd8, 0x80, 0x1b},
			},
			{
				Tag:         Tag_AssetID,
				BlockLength: 0x50,
				ValueLength: 0x3,
				Value:       []byte{0x6f, 0x6e, 0x65},
			},
			{
				Tag:         Tag_TransactionID,
				BlockLength: 0x80,
				ValueLength: 0x8,
				Value:       []byte{0xde, 0xad, 0xc0, 0xde, 0xde, 0xad, 0xc0, 0xde},
			},
			{
				Tag:         Tag_ProtocolVersionUsed,
				BlockLength: 0x80,
				ValueLength: 0x4,
				Value:       []byte{0x0, 0x0, 0x0, 0x1},
			},
			{
				Tag:         Tag_ProtocolVersionsSupported,
				BlockLength: 0x10,
				ValueLength: 0x4,
				Value:       []byte{0x0, 0x0, 0x0, 0x1},
			},
			{
				Tag:         Tag_ReturnRequest,
				BlockLength: 0x50,
				ValueLength: 0x18,
				Value:       []byte{0x1b, 0xf7, 0xf5, 0x3f, 0x5d, 0x5d, 0x5a, 0x1f, 0x47, 0xaa, 0x7a, 0xd3, 0x44, 0x5, 0x77, 0xde, 0xab, 0xb0, 0x25, 0x6a, 0x31, 0x84, 0x39, 0x74},
			},
		},
	},
	{"testdata/FPS/spc3.bin",
		"testdata/FPS/ckc3.bin",
		3952,
		[]byte{0x5d, 0x16, 0x44, 0xea, 0xec, 0x11, 0xf9, 0x83, 0x14, 0x75, 0x41, 0xe4, 0x6e, 0xeb, 0x27, 0x74},
		[]byte{0x60, 0x23, 0x52, 0xc2, 0xe6, 0xee, 0xfc, 0xa, 0x14, 0x86, 0x87, 0xf, 0x48, 0xe7, 0xa4, 0xe6, 0x44, 0x33, 0x52, 0x8f, 0x2a, 0x92, 0xe6, 0xc8, 0xe0, 0x60, 0xaf, 0x12, 0x91, 0x70, 0x86, 0x36, 0x84, 0xfd, 0x4a, 0x6a, 0x1b, 0xf4, 0xbe, 0xa7, 0x9f, 0xd9, 0x1b, 0xb2, 0xb0, 0x7a, 0xcc, 0xb4, 0x8b, 0xad, 0xa4, 0xdd, 0x87, 0x86, 0x23, 0x32, 0xc7, 0x5b, 0x8, 0xd9, 0x6f, 0x2e, 0x50, 0x1b, 0x8d, 0x23, 0xe2, 0x78, 0x55, 0x6e, 0x8d, 0xd5, 0xc1, 0x6a, 0x3e, 0xfe, 0x75, 0xad, 0x77, 0x56, 0x7a, 0xb, 0x52, 0x9e, 0xe2, 0x73, 0x34, 0x54, 0x58, 0xbd, 0x39, 0x20, 0xb8, 0xaa, 0x2d, 0x54, 0xc0, 0x50, 0xb0, 0x3c, 0x63, 0x0, 0x26, 0xc8, 0x27, 0x19, 0xf8, 0xe8, 0x63, 0xba, 0xbd, 0x4a, 0xc5, 0xab, 0x10, 0x9a, 0xf, 0xe2, 0x85, 0xfb, 0xee, 0xf, 0x26, 0xe0, 0x37, 0x48, 0x45, 0xa0},

		[]TLLVBlock{
			{
				Tag:         Tag_SessionKey_R1_integrity,
				BlockLength: 0x90,
				ValueLength: 0x10,
				Value:       []byte{0x54, 0xa1, 0x6b, 0xe0, 0x13, 0x7e, 0xf2, 0x59, 0xab, 0x3e, 0x4f, 0xc7, 0x96, 0x90, 0x82, 0x5f},
			},
			{
				Tag:         Tag_SessionKey_R1,
				BlockLength: 0xf0,
				ValueLength: 0x70,
				Value:       []byte{0x32, 0xc0, 0x82, 0xc9, 0xe2, 0x62, 0x73, 0x10, 0x1a, 0x97, 0xf3, 0x30, 0x81, 0xc1, 0xd0, 0x4a, 0xde, 0x58, 0xda, 0x53, 0xd7, 0xee, 0x17, 0x3e, 0xfd, 0xf0, 0xd1, 0xd3, 0xf0, 0x51, 0x4, 0xa9, 0xcb, 0x3a, 0x21, 0x3c, 0x7a, 0xc5, 0xab, 0x2b, 0x4c, 0x45, 0x2b, 0x95, 0x77, 0x79, 0x9f, 0x38, 0x6, 0xa0, 0x96, 0x47, 0xa5, 0xe5, 0xd2, 0x16, 0x74, 0x61, 0x73, 0x33, 0xe3, 0x2e, 0xc1, 0x20, 0x14, 0x74, 0x3c, 0x4e, 0x16, 0xb6, 0xd9, 0x6a, 0x61, 0xfc, 0xa5, 0x6b, 0xcb, 0x69, 0x6c, 0x75, 0xe9, 0x23, 0x75, 0x95, 0x35, 0xb3, 0x51, 0xb8, 0xc6, 0x27, 0xa2, 0x9c, 0x76, 0x69, 0xc8, 0xbf, 0x88, 0xc1, 0x40, 0x86, 0x4, 0x56, 0xef, 0x89, 0xf2, 0x8c, 0xe, 0xc6, 0xc9, 0x5a, 0xa6, 0xc3},
			},
			{
				Tag:         Tag_AntiReplaySeed,
				BlockLength: 0x20,
				ValueLength: 0x10,
				Value:       []byte{0xe0, 0xb4, 0xf2, 0x62, 0xf7, 0x6b, 0x54, 0x6d, 0x85, 0xf6, 0x22, 0xa6, 0xff, 0x48, 0x6c, 0xab},
			},
			{
				Tag:         Tag_R2,
				BlockLength: 0xc0,
				ValueLength: 0x15,
				Value:       []byte{0x11, 0xf7, 0xbe, 0x61, 0x2c, 0xa9, 0x5e, 0xf5, 0xe0, 0x7, 0xce, 0x51, 0x89, 0x6a, 0xe4, 0x50, 0x2c, 0xa3, 0xd8, 0x80, 0x1b},
			},
			{
				Tag:         Tag_AssetID,
				BlockLength: 0x70,
				ValueLength: 0x12,
				Value:       []byte{0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			},
			{
				Tag:         Tag_TransactionID,
				BlockLength: 0x80,
				ValueLength: 0x8,
				Value:       []byte{0x14, 0x73, 0xe5, 0xcc, 0x53, 0xe1, 0xe5, 0xd6},
			},
			{
				Tag:         Tag_ProtocolVersionUsed,
				BlockLength: 0x90,
				ValueLength: 0x4,
				Value:       []byte{0x0, 0x0, 0x0, 0x1},
			},
			{
				Tag:         Tag_ProtocolVersionsSupported,
				BlockLength: 0x70,
				ValueLength: 0x4,
				Value:       []byte{0x0, 0x0, 0x0, 0x1},
			},
			{
				Tag:         Tag_ReturnRequest,
				BlockLength: 0xb0,
				ValueLength: 0x38,
				Value:       []byte{0x1b, 0xf7, 0xf5, 0x3f, 0x5d, 0x5d, 0x5a, 0x1f, 0x47, 0xaa, 0x7a, 0xd3, 0x44, 0x5, 0x77, 0xde, 0x93, 0xc8, 0x23, 0x2c, 0x79, 0xb8, 0x1a, 0xb5, 0xce, 0xe6, 0x71, 0xf4, 0x89, 0xb4, 0xa, 0x6b, 0xf1, 0xd1, 0xeb, 0xe8, 0x98, 0xc0, 0x97, 0xd7, 0xbf, 0x91, 0xbc, 0x75, 0xab, 0xba, 0x19, 0xf7, 0x12, 0x47, 0x41, 0x85, 0x46, 0x71, 0x2c, 0x7d},
			},
		},
	},
	{"testdata/FPS-lease/spc1.bin",
		"testdata/FPS-lease/ckc1.bin",
		2032,
		[]byte{0xaf, 0xbf, 0xde, 0x35, 0xd2, 0x72, 0xa6, 0xb7, 0xc4, 0xd3, 0x85, 0xe2, 0xd4, 0xde, 0x2a, 0xb6},
		[]byte{0x8, 0x17, 0xc3, 0xe6, 0x18, 0xb2, 0x68, 0xab, 0x1a, 0xac, 0x5, 0xd3, 0xe, 0x91, 0x4b, 0xd2, 0xa8, 0x32, 0x61, 0x43, 0xe4, 0x0, 0xdd, 0xce, 0xa0, 0x63, 0x7b, 0xa6, 0xa1, 0xd3, 0xe4, 0x9a, 0x99, 0x6, 0x36, 0x1d, 0x79, 0x14, 0x34, 0x87, 0xee, 0xc4, 0x17, 0xe2, 0xec, 0x8c, 0xf7, 0xe3, 0x7a, 0x4b, 0xa4, 0xba, 0x3e, 0x81, 0xd7, 0xbe, 0xd, 0xf, 0xf1, 0x1b, 0xcb, 0x9, 0x74, 0xab, 0xb5, 0xb7, 0x40, 0x35, 0xc5, 0xff, 0xc7, 0x97, 0xe7, 0xb2, 0xc9, 0xf5, 0x1d, 0xf9, 0x69, 0x4a, 0xd7, 0x40, 0x30, 0x37, 0xd6, 0x64, 0x73, 0x8, 0xa7, 0xbf, 0x1, 0xc7, 0xde, 0xf7, 0xb4, 0x3f, 0x57, 0x17, 0xbf, 0x94, 0x8b, 0x24, 0xe3, 0x74, 0x52, 0x9d, 0xb2, 0xc6, 0x53, 0xbf, 0x98, 0x17, 0x2e, 0xff, 0x7b, 0xad, 0x4c, 0x82, 0x4d, 0xe5, 0xdc, 0x59, 0xc2, 0xba, 0x43, 0xe3, 0xf4, 0xd3},

		[]TLLVBlock{
			{
				Tag:         Tag_SessionKey_R1_integrity,
				BlockLength: 0xb0,
				ValueLength: 0x10,
				Value:       []byte{0x82, 0x9b, 0x46, 0x34, 0x85, 0x5d, 0x65, 0x7c, 0xfe, 0xac, 0x62, 0x7e, 0x2d, 0x6d, 0x63, 0x5a},
			},
			{
				Tag:         Tag_SessionKey_R1,
				BlockLength: 0xe0,
				ValueLength: 0x70,
				Value:       []byte{0xaa, 0xaa, 0xbb, 0xbb, 0xcc, 0xcc, 0x82, 0x30, 0xe8, 0x9a, 0xce, 0xfa, 0xdd, 0xdd, 0xaa, 0xaa, 0x79, 0x2e, 0xf7, 0x42, 0xf, 0x47, 0x9b, 0x80, 0x56, 0x7e, 0x4f, 0x8b, 0xe8, 0x3e, 0xac, 0x8b, 0xfe, 0xb, 0x75, 0x51, 0x4a, 0x7c, 0xda, 0x90, 0xe8, 0xef, 0x55, 0xac, 0xf4, 0xf, 0x3c, 0x59, 0x29, 0x2f, 0xd5, 0x0, 0xca, 0x4b, 0x3b, 0xdb, 0x48, 0x33, 0xd2, 0x45, 0x87, 0xaa, 0x87, 0x35, 0x8b, 0xc6, 0xe5, 0x97, 0xdb, 0x6c, 0x3c, 0xd1, 0xa4, 0xb5, 0xd0, 0x63, 0x8d, 0x99, 0xcf, 0xeb, 0x1a, 0x4d, 0x10, 0xa3, 0xf0, 0x90, 0x9d, 0x24, 0x2b, 0x2d, 0x8d, 0xf4, 0x5d, 0xd8, 0xe, 0x67, 0x51, 0x44, 0x6, 0xa9, 0x28, 0x3c, 0xd3, 0x68, 0x71, 0x82, 0xb4, 0xcd, 0x39, 0xa2, 0x75, 0xe},
			},
			{
				Tag:         Tag_AntiReplaySeed,
				BlockLength: 0x30,
				ValueLength: 0x10,
				Value:       []byte{0x2d, 0x9e, 0x8f, 0x28, 0x97, 0xc0, 0xc7, 0xa7, 0x17, 0xe0, 0x95, 0x5d, 0x3c, 0xea, 0xfa, 0x2d},
			},
			{
				Tag:         Tag_R2,
				BlockLength: 0x30,
				ValueLength: 0x15,
				Value:       []byte{0x11, 0xf7, 0xbe, 0x61, 0x2c, 0xa9, 0x5e, 0xf5, 0xe0, 0x7, 0xce, 0x51, 0x89, 0x6a, 0xe4, 0x50, 0x2c, 0xa3, 0xd8, 0x80, 0x1b},
			},
			{
				Tag:         Tag_AssetID,
				BlockLength: 0xa0,
				ValueLength: 0x3,
				Value:       []byte{0x6f, 0x6e, 0x65},
			},
			{
				Tag:         Tag_TransactionID,
				BlockLength: 0xb0,
				ValueLength: 0x8,
				Value:       []byte{0x13, 0x6b, 0x1a, 0x97, 0xa3, 0xd6, 0xb0, 0xfa},
			},
			{
				Tag:         Tag_ProtocolVersionUsed,
				BlockLength: 0x60,
				ValueLength: 0x4,
				Value:       []byte{0x0, 0x0, 0x0, 0x1},
			},
			{
				Tag:         Tag_ProtocolVersionsSupported,
				BlockLength: 0x30,
				ValueLength: 0x8,
				Value:       []byte{0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x2},
			},
			{
				Tag:         Tag_ReturnRequest,
				BlockLength: 0x40,
				ValueLength: 0x18,
				Value:       []byte{0x1b, 0xf7, 0xf5, 0x3f, 0x5d, 0x5d, 0x5a, 0x1f, 0x47, 0xaa, 0x7a, 0xd3, 0x44, 0x5, 0x77, 0xde, 0xc3, 0x92, 0xec, 0x58, 0x91, 0xef, 0xa7, 0xd0},
			},
		},
	},
}

func TestGenCKC(t *testing.T) {
	k := &Ksm{
		Pub:       pub,
		Pri:       pri,
		Rck:       RandomContentKey{},
		DFunction: d.AppleD{},
		Ask:       []byte{},
	}
	for _, test := range spcContainerTests {
		spcMessage := readBin(test.filePath)

		ckc, err := k.GenCKC(spcMessage)
		assert.NoError(t, err)

		assert.NoError(t, ioutil.WriteFile(test.outFilePath, ckc, 0777))
	}
}

func TestDebugCKC(t *testing.T) {
	ckcMessage := readBin("testdata/FPS/ckc1.bin")
	DebugCKC(ckcMessage)
}

func TestParseSPCV1(t *testing.T) {
	for _, test := range spcContainerTests {
		spcMessage := readBin(test.filePath)

		spcContainer, err := ParseSPCV1(spcMessage, pub, pri)
		assert.NoError(t, err)

		assert.Equal(t, test.playload_size, len(spcContainer.SPCPlayload))
		assert.Equal(t, test.encrypted_key, spcContainer.EncryptedAesKey)
		assert.Equal(t, test.iv, spcContainer.AesKeyIV)

		for _, tllv := range test.ttls {
			actualTtlv := spcContainer.TTLVS[tllv.Tag]
			assert.NotNil(t, actualTtlv)
			assert.Equal(t, tllv, actualTtlv)
		}
	}
}

func readBin(filePath string) []byte {
	f, err := os.Open(filePath)
	defer f.Close()
	checkErr(err)

	spcMessage, err := ioutil.ReadAll(f)
	checkErr(err)

	return spcMessage

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
