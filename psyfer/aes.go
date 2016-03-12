package psyfer

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

var Sbox = [][]byte{
	{0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76},
	{0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0, 0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0},
	{0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc, 0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15},
	{0x04, 0xc7, 0x23, 0xc3, 0x18, 0x96, 0x05, 0x9a, 0x07, 0x12, 0x80, 0xe2, 0xeb, 0x27, 0xb2, 0x75},
	{0x09, 0x83, 0x2c, 0x1a, 0x1b, 0x6e, 0x5a, 0xa0, 0x52, 0x3b, 0xd6, 0xb3, 0x29, 0xe3, 0x2f, 0x84},
	{0x53, 0xd1, 0x00, 0xed, 0x20, 0xfc, 0xb1, 0x5b, 0x6a, 0xcb, 0xbe, 0x39, 0x4a, 0x4c, 0x58, 0xcf},
	{0xd0, 0xef, 0xaa, 0xfb, 0x43, 0x4d, 0x33, 0x85, 0x45, 0xf9, 0x02, 0x7f, 0x50, 0x3c, 0x9f, 0xa8},
	{0x51, 0xa3, 0x40, 0x8f, 0x92, 0x9d, 0x38, 0xf5, 0xbc, 0xb6, 0xda, 0x21, 0x10, 0xff, 0xf3, 0xd2},
	{0xcd, 0x0c, 0x13, 0xec, 0x5f, 0x97, 0x44, 0x17, 0xc4, 0xa7, 0x7e, 0x3d, 0x64, 0x5d, 0x19, 0x73},
	{0x60, 0x81, 0x4f, 0xdc, 0x22, 0x2a, 0x90, 0x88, 0x46, 0xee, 0xb8, 0x14, 0xde, 0x5e, 0x0b, 0xdb},
	{0xe0, 0x32, 0x3a, 0x0a, 0x49, 0x06, 0x24, 0x5c, 0xc2, 0xd3, 0xac, 0x62, 0x91, 0x95, 0xe4, 0x79},
	{0xe7, 0xc8, 0x37, 0x6d, 0x8d, 0xd5, 0x4e, 0xa9, 0x6c, 0x56, 0xf4, 0xea, 0x65, 0x7a, 0xae, 0x08},
	{0xba, 0x78, 0x25, 0x2e, 0x1c, 0xa6, 0xb4, 0xc6, 0xe8, 0xdd, 0x74, 0x1f, 0x4b, 0xbd, 0x8b, 0x8a},
	{0x70, 0x3e, 0xb5, 0x66, 0x48, 0x03, 0xf6, 0x0e, 0x61, 0x35, 0x57, 0xb9, 0x86, 0xc1, 0x1d, 0x9e},
	{0xe1, 0xf8, 0x98, 0x11, 0x69, 0xd9, 0x8e, 0x94, 0x9b, 0x1e, 0x87, 0xe9, 0xce, 0x55, 0x28, 0xdf},
	{0x8c, 0xa1, 0x89, 0x0d, 0xbf, 0xe6, 0x42, 0x68, 0x41, 0x99, 0x2d, 0x0f, 0xb0, 0x54, 0xbb, 0x16},
}

var InvSbox = [][]byte{
	{0x52, 0x09, 0x6a, 0xd5, 0x30, 0x36, 0xa5, 0x38, 0xbf, 0x40, 0xa3, 0x9e, 0x81, 0xf3, 0xd7, 0xfb},
	{0x7c, 0xe3, 0x39, 0x82, 0x9b, 0x2f, 0xff, 0x87, 0x34, 0x8e, 0x43, 0x44, 0xc4, 0xde, 0xe9, 0xcb},
	{0x54, 0x7b, 0x94, 0x32, 0xa6, 0xc2, 0x23, 0x3d, 0xee, 0x4c, 0x95, 0x0b, 0x42, 0xfa, 0xc3, 0x4e},
	{0x08, 0x2e, 0xa1, 0x66, 0x28, 0xd9, 0x24, 0xb2, 0x76, 0x5b, 0xa2, 0x49, 0x6d, 0x8b, 0xd1, 0x25},
	{0x72, 0xf8, 0xf6, 0x64, 0x86, 0x68, 0x98, 0x16, 0xd4, 0xa4, 0x5c, 0xcc, 0x5d, 0x65, 0xb6, 0x92},
	{0x6c, 0x70, 0x48, 0x50, 0xfd, 0xed, 0xb9, 0xda, 0x5e, 0x15, 0x46, 0x57, 0xa7, 0x8d, 0x9d, 0x84},
	{0x90, 0xd8, 0xab, 0x00, 0x8c, 0xbc, 0xd3, 0x0a, 0xf7, 0xe4, 0x58, 0x05, 0xb8, 0xb3, 0x45, 0x06},
	{0xd0, 0x2c, 0x1e, 0x8f, 0xca, 0x3f, 0x0f, 0x02, 0xc1, 0xaf, 0xbd, 0x03, 0x01, 0x13, 0x8a, 0x6b},
	{0x3a, 0x91, 0x11, 0x41, 0x4f, 0x67, 0xdc, 0xea, 0x97, 0xf2, 0xcf, 0xce, 0xf0, 0xb4, 0xe6, 0x73},
	{0x96, 0xac, 0x74, 0x22, 0xe7, 0xad, 0x35, 0x85, 0xe2, 0xf9, 0x37, 0xe8, 0x1c, 0x75, 0xdf, 0x6e},
	{0x47, 0xf1, 0x1a, 0x71, 0x1d, 0x29, 0xc5, 0x89, 0x6f, 0xb7, 0x62, 0x0e, 0xaa, 0x18, 0xbe, 0x1b},
	{0xfc, 0x56, 0x3e, 0x4b, 0xc6, 0xd2, 0x79, 0x20, 0x9a, 0xdb, 0xc0, 0xfe, 0x78, 0xcd, 0x5a, 0xf4},
	{0x1f, 0xdd, 0xa8, 0x33, 0x88, 0x07, 0xc7, 0x31, 0xb1, 0x12, 0x10, 0x59, 0x27, 0x80, 0xec, 0x5f},
	{0x60, 0x51, 0x7f, 0xa9, 0x19, 0xb5, 0x4a, 0x0d, 0x2d, 0xe5, 0x7a, 0x9f, 0x93, 0xc9, 0x9c, 0xef},
	{0xa0, 0xe0, 0x3b, 0x4d, 0xae, 0x2a, 0xf5, 0xb0, 0xc8, 0xeb, 0xbb, 0x3c, 0x83, 0x53, 0x99, 0x61},
	{0x17, 0x2b, 0x04, 0x7e, 0xba, 0x77, 0xd6, 0x26, 0xe1, 0x69, 0x14, 0x63, 0x55, 0x21, 0x0c, 0x7d},
}

var MM = []byte{
	2, 3, 1, 1,
	1, 2, 3, 1,
	1, 1, 2, 3,
	3, 1, 1, 2,
}

var iMM = []byte{
	14, 11, 13, 9,
	9, 14, 11, 13,
	13, 9, 14, 11,
	11, 13, 9, 14,
}

var Rcon = []byte{0x00, // Rcon[] is 1-based, so the first entry is just a place holder
	0x01, 0x02, 0x04, 0x08,
	0x10, 0x20, 0x40, 0x80,
	0x1B, 0x36, 0x6C, 0xD8,
	0xAB, 0x4D, 0x9A, 0x2F,
	0x5E, 0xBC, 0x63, 0xC6,
	0x97, 0x35, 0x6A, 0xD4,
	0xB3, 0x7D, 0xFA, 0xEF,
	0xC5, 0x91, 0x39, 0x72,
	0xE4, 0xD3, 0xBD, 0x61,
	0xC2, 0x9F, 0x25, 0x4A,
	0x94, 0x33, 0x66, 0xCC,
	0x83, 0x1D, 0x3A, 0x74,
	0xE8, 0xCB, 0x8D,
}

type Block []byte

var keyexpanded []Block
var key Block

func ToString(all []Block, keysize int, k Block, decrypt bool) string {
	final := ""
	for _, bl := range all {
		result := Block{}
		if decrypt {
			result = InvCipher(bl, keysize, k)
		} else {
			result = Cipher(bl, keysize, k)
		}
		final += string(result)
	}
	return final
}

func ToHex(all []Block, keysize int, k Block, decrypt bool) string {
	final := ""
	for _, bl := range all {
		result := Block{}
		if decrypt {
			result = InvCipher(bl, keysize, k)
		} else {
			result = Cipher(bl, keysize, k)
		}
		for i := 0; i < 16; i++ {
			final += fmt.Sprintf("0x%x ", result[i])
		}
	}
	return final
}

func BlockGen(arg string) []Block {
	all := []Block{}
	b := Block{}
	for i, char := range arg {
		value, err := strconv.Atoi(hex.EncodeToString([]byte(string(char))))
		if err != nil {
			log.Fatal(err)
		}
		if i%16 == 0 && i > 0 {
			all = append(all, b)
			b = b[:0]
		}
		b = append(b, byte(value))
		if i == len(arg)-1 {
			all = append(all, b)
		}
	}
	return all
}

func Cipher(cur Block, bit int, incomingKey Block) Block {
	if len(cur) != 16 {
		missing := 16 - len(cur)
		for i := 0; i < missing; i++ {
			cur = append(cur, 0x00)
		}
	}
	key = Block{}
	keyexpanded = []Block{}
	AssignKey(incomingKey)
	if bit == 128 {
		KeyExpansionBase(128)
		cur = AddRoundKey(cur, 0)
		for i := 0; i < 9; i++ {
			cur = SubBytes(cur)
			cur = ShiftRows(cur)
			cur = MixColumns(cur)
			cur = AddRoundKey(cur, i+1)
		}
		cur = SubBytes(cur)
		cur = ShiftRows(cur)
		cur = AddRoundKey(cur, 10)
		return cur
	}
	if bit == 192 {
		KeyExpansionBase(192)
		cur = AddRoundKey(cur, 0)
		for i := 0; i < 11; i++ {
			cur = SubBytes(cur)
			cur = ShiftRows(cur)
			cur = MixColumns(cur)
			cur = AddRoundKey(cur, i+1)
			if i == 0 {
			}
		}
		cur = SubBytes(cur)
		cur = ShiftRows(cur)
		cur = AddRoundKey(cur, 12)
		return cur
	}
	if bit == 256 {
		KeyExpansionBase(256)
		cur = AddRoundKey(cur, 0)
		for i := 0; i < 13; i++ {
			cur = SubBytes(cur)
			cur = ShiftRows(cur)
			cur = MixColumns(cur)
			cur = AddRoundKey(cur, i+1)
			if i == 0 {
			}
		}
		cur = SubBytes(cur)
		cur = ShiftRows(cur)
		cur = AddRoundKey(cur, 14)
		return cur
	}
	return cur
}

func InvCipher(cur Block, bit int, incomingKey Block) Block {
	if len(cur) != 16 {
		missing := 16 - len(cur)
		for i := 0; i < missing; i++ {
			cur = append(cur, 0x00)
		}
	}
	key = Block{}
	keyexpanded = []Block{}
	AssignKey(incomingKey)
	if bit == 128 {
		KeyExpansionBase(128)
		cur = AddRoundKey(cur, 10)
		for i := 9; i > 0; i-- {
			cur = InvShiftRows(cur)
			cur = InvSubBytes(cur)
			cur = AddRoundKey(cur, i)
			cur = InvMixColumns(cur)
		}
		cur = InvSubBytes(cur)
		cur = InvShiftRows(cur)
		cur = AddRoundKey(cur, 0)
		return cur
	}
	if bit == 192 {
		KeyExpansionBase(192)
		cur = AddRoundKey(cur, 12)
		for i := 11; i > 0; i-- {
			cur = InvShiftRows(cur)
			cur = InvSubBytes(cur)
			cur = AddRoundKey(cur, i)
			cur = InvMixColumns(cur)
		}
		cur = InvSubBytes(cur)
		cur = InvShiftRows(cur)
		cur = AddRoundKey(cur, 0)
		return cur
	}
	if bit == 256 {
		KeyExpansionBase(256)
		cur = AddRoundKey(cur, 14)
		for i := 13; i > 0; i-- {
			cur = InvShiftRows(cur)
			cur = InvSubBytes(cur)
			cur = AddRoundKey(cur, i)
			cur = InvMixColumns(cur)
		}
		cur = InvSubBytes(cur)
		cur = InvShiftRows(cur)
		cur = AddRoundKey(cur, 0)
		return cur
	}
	return cur
}

func AddRoundKey(cur Block, iteration int) Block {
	for i := 0; i < 16; i++ {
		cur[i] = cur[i] ^ keyexpanded[iteration][i]
	}
	return cur
}

func AssignKey(cur Block) {
	key = cur
}

func KeyExpansionBase(keysize int) {
	if keysize == 128 {
		keyexpanded = append(keyexpanded, key)
		for i := 0; i < 10; i++ {
			KeyExpansion(keyexpanded[i], i+1)
		}
	} else if keysize == 192 {
		keyexpanded = append(keyexpanded, key)
		for i := 0; i < 8; i++ {
			KeyExpansion192(keyexpanded[i], i+1)
		}
		temp := keyexpanded
		keyexpanded = []Block{}
		for i := 0; i < 9; i++ {
			if i == 8 {
				a := Block{
					temp[i][0], temp[i][1], temp[i][2], temp[i][3],
					temp[i][6], temp[i][7], temp[i][8], temp[i][9],
					temp[i][12], temp[i][13], temp[i][14], temp[i][15],
					temp[i][18], temp[i][19], temp[i][20], temp[i][21],
				}
				b := Block{
					temp[i][4], temp[i][5], 0, 0,
					temp[i][10], temp[i][11], 0, 0,
					temp[i][16], temp[i][17], 0, 0,
					temp[i][22], temp[i][23], 0, 0,
				}
				keyexpanded = append(keyexpanded, a)
				keyexpanded = append(keyexpanded, b)
			} else {
				a := Block{
					temp[i][0], temp[i][1], temp[i][2], temp[i][3],
					temp[i][6], temp[i][7], temp[i][8], temp[i][9],
					temp[i][12], temp[i][13], temp[i][14], temp[i][15],
					temp[i][18], temp[i][19], temp[i][20], temp[i][21],
				}
				b := Block{
					temp[i][4], temp[i][5], temp[i+1][0], temp[i+1][1],
					temp[i][10], temp[i][11], temp[i+1][6], temp[i+1][7],
					temp[i][16], temp[i][17], temp[i+1][12], temp[i+1][13],
					temp[i][22], temp[i][23], temp[i+1][18], temp[i+1][19],
				}
				c := Block{
					temp[i+1][2], temp[i+1][3], temp[i+1][4], temp[i+1][5],
					temp[i+1][8], temp[i+1][9], temp[i+1][10], temp[i+1][11],
					temp[i+1][14], temp[i+1][15], temp[i+1][16], temp[i+1][17],
					temp[i+1][20], temp[i+1][21], temp[i+1][22], temp[i+1][23],
				}
				keyexpanded = append(keyexpanded, a)
				keyexpanded = append(keyexpanded, b)
				keyexpanded = append(keyexpanded, c)
			}
			i++
		}
	} else if keysize == 256 {
		keyexpanded = append(keyexpanded, key)
		for i := 0; i < 8; i++ {
			KeyExpansion256(keyexpanded[i], i+1)
		}
		temp := keyexpanded
		keyexpanded = []Block{}
		for i := 0; i < 9; i++ {
			a := Block{
				temp[i][0], temp[i][1], temp[i][2], temp[i][3],
				temp[i][8], temp[i][9], temp[i][10], temp[i][11],
				temp[i][16], temp[i][17], temp[i][18], temp[i][19],
				temp[i][24], temp[i][25], temp[i][26], temp[i][27],
			}
			b := Block{
				temp[i][4], temp[i][5], temp[i][6], temp[i][7],
				temp[i][12], temp[i][13], temp[i][14], temp[i][15],
				temp[i][20], temp[i][21], temp[i][22], temp[i][23],
				temp[i][28], temp[i][29], temp[i][30], temp[i][31],
			}
			keyexpanded = append(keyexpanded, a)
			keyexpanded = append(keyexpanded, b)
		}
	}
}

func KeyExpansion(cur Block, iteration int) Block {
	var nb = Block{ //nb = nextBlock
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}
	rotword := []byte{cur[7], cur[11], cur[15], cur[3]}
	for i := 0; i < 4; i++ {
		upper, lower := SplitBytes(rotword[i])
		rotword[i] = Sbox[upper][lower]
	}
	rotword[0] = rotword[0] ^ cur[0] ^ Rcon[iteration]
	rotword[1] = rotword[1] ^ cur[4]
	rotword[2] = rotword[2] ^ cur[8]
	rotword[3] = rotword[3] ^ cur[12]

	nb[0], nb[4], nb[8], nb[12] = rotword[0], rotword[1], rotword[2], rotword[3]
	nb[1], nb[5], nb[9], nb[13] = nb[0]^cur[1], nb[4]^cur[5], nb[8]^cur[9], nb[12]^cur[13]
	nb[2], nb[6], nb[10], nb[14] = nb[1]^cur[2], nb[5]^cur[6], nb[9]^cur[10], nb[13]^cur[14]
	nb[3], nb[7], nb[11], nb[15] = nb[2]^cur[3], nb[6]^cur[7], nb[10]^cur[11], nb[14]^cur[15]
	keyexpanded = append(keyexpanded, nb)
	return nb
}

func KeyExpansion192(cur Block, iteration int) Block {
	var nb = Block{ //nb = nextBlock
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	rotword := []byte{cur[11], cur[17], cur[23], cur[5]}
	for i := 0; i < 4; i++ {
		upper, lower := SplitBytes(rotword[i])
		rotword[i] = Sbox[upper][lower]
	}
	rotword[0] = rotword[0] ^ cur[0] ^ Rcon[iteration]
	rotword[1] = rotword[1] ^ cur[6]
	rotword[2] = rotword[2] ^ cur[12]
	rotword[3] = rotword[3] ^ cur[18]

	nb[0], nb[6], nb[12], nb[18] = rotword[0], rotword[1], rotword[2], rotword[3]
	nb[1], nb[7], nb[13], nb[19] = nb[0]^cur[1], nb[6]^cur[7], nb[12]^cur[13], nb[18]^cur[19]
	nb[2], nb[8], nb[14], nb[20] = nb[1]^cur[2], nb[7]^cur[8], nb[13]^cur[14], nb[19]^cur[20]
	nb[3], nb[9], nb[15], nb[21] = nb[2]^cur[3], nb[8]^cur[9], nb[14]^cur[15], nb[20]^cur[21]
	nb[4], nb[10], nb[16], nb[22] = nb[3]^cur[4], nb[9]^cur[10], nb[15]^cur[16], nb[21]^cur[22]
	nb[5], nb[11], nb[17], nb[23] = nb[4]^cur[5], nb[10]^cur[11], nb[16]^cur[17], nb[22]^cur[23]
	keyexpanded = append(keyexpanded, nb)

	return nb
}

func KeyExpansion256(cur Block, iteration int) Block {
	var nb = Block{ //nb = nextBlock
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	rotword := []byte{cur[15], cur[23], cur[31], cur[7]}
	for i := 0; i < 4; i++ {
		upper, lower := SplitBytes(rotword[i])
		rotword[i] = Sbox[upper][lower]
	}
	rotword[0] = rotword[0] ^ cur[0] ^ Rcon[iteration]
	rotword[1] = rotword[1] ^ cur[8]
	rotword[2] = rotword[2] ^ cur[16]
	rotword[3] = rotword[3] ^ cur[24]

	nb[0], nb[8], nb[16], nb[24] = rotword[0], rotword[1], rotword[2], rotword[3]
	nb[1], nb[9], nb[17], nb[25] = nb[0]^cur[1], nb[8]^cur[9], nb[16]^cur[17], nb[24]^cur[25]
	nb[2], nb[10], nb[18], nb[26] = nb[1]^cur[2], nb[9]^cur[10], nb[17]^cur[18], nb[25]^cur[26]
	nb[3], nb[11], nb[19], nb[27] = nb[2]^cur[3], nb[10]^cur[11], nb[18]^cur[19], nb[26]^cur[27]
	sw := []byte{nb[3], nb[11], nb[19], nb[27]} //sw = subword
	for i := 0; i < 4; i++ {
		upper, lower := SplitBytes(sw[i])
		sw[i] = Sbox[upper][lower]
	}
	nb[4], nb[12], nb[20], nb[28] = sw[0]^cur[4], sw[1]^cur[12], sw[2]^cur[20], sw[3]^cur[28]
	nb[5], nb[13], nb[21], nb[29] = nb[4]^cur[5], nb[12]^cur[13], nb[20]^cur[21], nb[28]^cur[29]
	nb[6], nb[14], nb[22], nb[30] = nb[5]^cur[6], nb[13]^cur[14], nb[21]^cur[22], nb[29]^cur[30]
	nb[7], nb[15], nb[23], nb[31] = nb[6]^cur[7], nb[14]^cur[15], nb[22]^cur[23], nb[30]^cur[31]
	keyexpanded = append(keyexpanded, nb)

	return nb
}

func SplitBytes(b byte) (byte, byte) {
	return b >> 4, b & 0x0f
}

func SubBytes(cur Block) Block {
	for i := 0; i < 16; i++ {
		upper, lower := SplitBytes(cur[i])
		cur[i] = Sbox[upper][lower]
	}
	return cur
}

func InvSubBytes(cur Block) Block {
	for i := 0; i < 16; i++ {
		upper, lower := SplitBytes(cur[i])
		cur[i] = InvSbox[upper][lower]
	}
	return cur
}

func Xtime(cur byte) []byte {
	var bytes []byte
	bytes = append(bytes, cur)
	for i := 1; i < 8; i++ { // first iteration done outside of for-loop
		if (cur >> 7) == 1 {
			cur = cur << 1
			cur = cur ^ 0x1b
		} else {
			cur = cur << 1
		}
		bytes = append(bytes, cur)
	}
	return bytes
}

func FFmult(cur []byte, multiplier byte) byte {
	if multiplier == 1 {
		return cur[0]
	} else if multiplier == 2 {
		return cur[1]
	} else if multiplier == 3 {
		return cur[0] ^ cur[1]
	} else if multiplier == 9 {
		return cur[0] ^ cur[3]
	} else if multiplier == 11 {
		return cur[0] ^ cur[1] ^ cur[3]
	} else if multiplier == 13 {
		return cur[0] ^ cur[2] ^ cur[3]
	} else if multiplier == 14 {
		return cur[1] ^ cur[2] ^ cur[3]
	}
	return 0
}

func mixColumnsAssist(cur []byte) []byte {
	a1 := FFmult(Xtime(cur[0]), MM[0]) ^ FFmult(Xtime(cur[1]), MM[1]) ^ FFmult(Xtime(cur[2]), MM[2]) ^ FFmult(Xtime(cur[3]), MM[3])
	a2 := FFmult(Xtime(cur[0]), MM[4]) ^ FFmult(Xtime(cur[1]), MM[5]) ^ FFmult(Xtime(cur[2]), MM[6]) ^ FFmult(Xtime(cur[3]), MM[7])
	a3 := FFmult(Xtime(cur[0]), MM[8]) ^ FFmult(Xtime(cur[1]), MM[9]) ^ FFmult(Xtime(cur[2]), MM[10]) ^ FFmult(Xtime(cur[3]), MM[11])
	a4 := FFmult(Xtime(cur[0]), MM[12]) ^ FFmult(Xtime(cur[1]), MM[13]) ^ FFmult(Xtime(cur[2]), MM[14]) ^ FFmult(Xtime(cur[3]), MM[15])
	return []byte{a1, a2, a3, a4}
}

func MixColumns(cur Block) Block {
	col1 := []byte{cur[0], cur[4], cur[8], cur[12]}
	col2 := []byte{cur[1], cur[5], cur[9], cur[13]}
	col3 := []byte{cur[2], cur[6], cur[10], cur[14]}
	col4 := []byte{cur[3], cur[7], cur[11], cur[15]}
	col1 = mixColumnsAssist(col1)
	col2 = mixColumnsAssist(col2)
	col3 = mixColumnsAssist(col3)
	col4 = mixColumnsAssist(col4)
	cur = Block{
		col1[0], col2[0], col3[0], col4[0],
		col1[1], col2[1], col3[1], col4[1],
		col1[2], col2[2], col3[2], col4[2],
		col1[3], col2[3], col3[3], col4[3],
	}
	return cur
}

func InvMixColumns(cur Block) Block {
	col1 := []byte{cur[0], cur[4], cur[8], cur[12]}
	col2 := []byte{cur[1], cur[5], cur[9], cur[13]}
	col3 := []byte{cur[2], cur[6], cur[10], cur[14]}
	col4 := []byte{cur[3], cur[7], cur[11], cur[15]}
	col1 = InvMixColumnsAssist(col1)
	col2 = InvMixColumnsAssist(col2)
	col3 = InvMixColumnsAssist(col3)
	col4 = InvMixColumnsAssist(col4)
	cur = Block{
		col1[0], col2[0], col3[0], col4[0],
		col1[1], col2[1], col3[1], col4[1],
		col1[2], col2[2], col3[2], col4[2],
		col1[3], col2[3], col3[3], col4[3],
	}
	return cur
}

func InvMixColumnsAssist(cur []byte) []byte {
	a1 := FFmult(Xtime(cur[0]), iMM[0]) ^ FFmult(Xtime(cur[1]), iMM[1]) ^ FFmult(Xtime(cur[2]), iMM[2]) ^ FFmult(Xtime(cur[3]), iMM[3])
	a2 := FFmult(Xtime(cur[0]), iMM[4]) ^ FFmult(Xtime(cur[1]), iMM[5]) ^ FFmult(Xtime(cur[2]), iMM[6]) ^ FFmult(Xtime(cur[3]), iMM[7])
	a3 := FFmult(Xtime(cur[0]), iMM[8]) ^ FFmult(Xtime(cur[1]), iMM[9]) ^ FFmult(Xtime(cur[2]), iMM[10]) ^ FFmult(Xtime(cur[3]), iMM[11])
	a4 := FFmult(Xtime(cur[0]), iMM[12]) ^ FFmult(Xtime(cur[1]), iMM[13]) ^ FFmult(Xtime(cur[2]), iMM[14]) ^ FFmult(Xtime(cur[3]), iMM[15])
	return []byte{a1, a2, a3, a4}
}

func ShiftRows(cur Block) Block {
	cur[4], cur[5], cur[6], cur[7] = cur[5], cur[6], cur[7], cur[4]
	cur[8], cur[9], cur[10], cur[11] = cur[10], cur[11], cur[8], cur[9]
	cur[12], cur[13], cur[14], cur[15] = cur[15], cur[12], cur[13], cur[14]
	return cur
}

func InvShiftRows(cur Block) Block {
	cur[4], cur[5], cur[6], cur[7] = cur[7], cur[4], cur[5], cur[6]
	cur[8], cur[9], cur[10], cur[11] = cur[10], cur[11], cur[8], cur[9]
	cur[12], cur[13], cur[14], cur[15] = cur[13], cur[14], cur[15], cur[12]
	return cur
}
