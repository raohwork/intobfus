/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package intobfus

import (
	"fmt"
	"math"
	"testing"
)

func ExampleObfuscator_obfuscate() {
	// obfuscates 0~255 (1~255 actually)
	max := uint64(255)
	// a number that GCD(max+1, key) = 1. prime number is suggested
	key := uint64(3 * 7 * 11)

	o := MustRestore(max, key)
	for i := uint64(1); i < 255; i++ {
		fmt.Printf("enc(%3d) = %3d\n", i, o.Obfuscate(i))
	}

	// output: enc(  1) = 113
	// enc(  2) = 206
	// enc(  3) = 181
	// enc(  4) = 156
	// enc(  5) = 131
	// enc(  6) = 106
	// enc(  7) =  81
	// enc(  8) =  56
	// enc(  9) =  31
	// enc( 10) =   6
	// enc( 11) = 237
	// enc( 12) = 212
	// enc( 13) = 187
	// enc( 14) = 162
	// enc( 15) = 137
	// enc( 16) = 112
	// enc( 17) =  87
	// enc( 18) =  62
	// enc( 19) =  37
	// enc( 20) =  12
	// enc( 21) = 243
	// enc( 22) = 218
	// enc( 23) = 193
	// enc( 24) = 168
	// enc( 25) = 143
	// enc( 26) = 118
	// enc( 27) =  93
	// enc( 28) =  68
	// enc( 29) =  43
	// enc( 30) =  18
	// enc( 31) = 249
	// enc( 32) = 224
	// enc( 33) = 199
	// enc( 34) = 174
	// enc( 35) = 149
	// enc( 36) = 124
	// enc( 37) =  99
	// enc( 38) =  74
	// enc( 39) =  49
	// enc( 40) =  24
	// enc( 41) = 255
	// enc( 42) = 230
	// enc( 43) = 205
	// enc( 44) = 180
	// enc( 45) = 155
	// enc( 46) = 130
	// enc( 47) = 105
	// enc( 48) =  80
	// enc( 49) =  55
	// enc( 50) =  30
	// enc( 51) =   5
	// enc( 52) = 236
	// enc( 53) = 211
	// enc( 54) = 186
	// enc( 55) = 161
	// enc( 56) = 136
	// enc( 57) = 111
	// enc( 58) =  86
	// enc( 59) =  61
	// enc( 60) =  36
	// enc( 61) =  11
	// enc( 62) = 242
	// enc( 63) = 217
	// enc( 64) = 192
	// enc( 65) = 167
	// enc( 66) = 142
	// enc( 67) = 117
	// enc( 68) =  92
	// enc( 69) =  67
	// enc( 70) =  42
	// enc( 71) =  17
	// enc( 72) = 248
	// enc( 73) = 223
	// enc( 74) = 198
	// enc( 75) = 173
	// enc( 76) = 148
	// enc( 77) = 123
	// enc( 78) =  98
	// enc( 79) =  73
	// enc( 80) =  48
	// enc( 81) =  23
	// enc( 82) = 254
	// enc( 83) = 229
	// enc( 84) = 204
	// enc( 85) = 179
	// enc( 86) = 154
	// enc( 87) = 129
	// enc( 88) = 104
	// enc( 89) =  79
	// enc( 90) =  54
	// enc( 91) =  29
	// enc( 92) =   4
	// enc( 93) = 235
	// enc( 94) = 210
	// enc( 95) = 185
	// enc( 96) = 160
	// enc( 97) = 135
	// enc( 98) = 110
	// enc( 99) =  85
	// enc(100) =  60
	// enc(101) =  35
	// enc(102) =  10
	// enc(103) = 241
	// enc(104) = 216
	// enc(105) = 191
	// enc(106) = 166
	// enc(107) = 141
	// enc(108) = 116
	// enc(109) =  91
	// enc(110) =  66
	// enc(111) =  41
	// enc(112) =  16
	// enc(113) = 231
	// enc(114) = 222
	// enc(115) = 197
	// enc(116) = 172
	// enc(117) = 147
	// enc(118) = 122
	// enc(119) =  97
	// enc(120) =  72
	// enc(121) =  47
	// enc(122) =  22
	// enc(123) = 253
	// enc(124) = 228
	// enc(125) = 203
	// enc(126) = 178
	// enc(127) = 153
	// enc(128) = 128
	// enc(129) = 103
	// enc(130) =  78
	// enc(131) =  53
	// enc(132) =  28
	// enc(133) =   3
	// enc(134) = 234
	// enc(135) = 209
	// enc(136) = 184
	// enc(137) = 159
	// enc(138) = 134
	// enc(139) = 109
	// enc(140) =  84
	// enc(141) =  59
	// enc(142) =  34
	// enc(143) =   9
	// enc(144) = 240
	// enc(145) = 215
	// enc(146) = 190
	// enc(147) = 165
	// enc(148) = 140
	// enc(149) = 115
	// enc(150) =  90
	// enc(151) =  65
	// enc(152) =  40
	// enc(153) =  15
	// enc(154) = 246
	// enc(155) = 221
	// enc(156) = 196
	// enc(157) = 171
	// enc(158) = 146
	// enc(159) = 121
	// enc(160) =  96
	// enc(161) =  71
	// enc(162) =  46
	// enc(163) =  21
	// enc(164) = 252
	// enc(165) = 227
	// enc(166) = 202
	// enc(167) = 177
	// enc(168) = 152
	// enc(169) = 127
	// enc(170) = 102
	// enc(171) =  77
	// enc(172) =  52
	// enc(173) =  27
	// enc(174) =   2
	// enc(175) = 233
	// enc(176) = 208
	// enc(177) = 183
	// enc(178) = 158
	// enc(179) = 133
	// enc(180) = 108
	// enc(181) =  83
	// enc(182) =  58
	// enc(183) =  33
	// enc(184) =   8
	// enc(185) = 239
	// enc(186) = 214
	// enc(187) = 189
	// enc(188) = 164
	// enc(189) = 139
	// enc(190) = 114
	// enc(191) =  89
	// enc(192) =  64
	// enc(193) =  39
	// enc(194) =  14
	// enc(195) = 245
	// enc(196) = 220
	// enc(197) = 195
	// enc(198) = 170
	// enc(199) = 145
	// enc(200) = 120
	// enc(201) =  95
	// enc(202) =  70
	// enc(203) =  45
	// enc(204) =  20
	// enc(205) = 251
	// enc(206) = 226
	// enc(207) = 201
	// enc(208) = 176
	// enc(209) = 151
	// enc(210) = 126
	// enc(211) = 101
	// enc(212) =  76
	// enc(213) =  51
	// enc(214) =  26
	// enc(215) =   1
	// enc(216) = 232
	// enc(217) = 207
	// enc(218) = 182
	// enc(219) = 157
	// enc(220) = 132
	// enc(221) = 107
	// enc(222) =  82
	// enc(223) =  57
	// enc(224) =  32
	// enc(225) =   7
	// enc(226) = 238
	// enc(227) = 213
	// enc(228) = 188
	// enc(229) = 163
	// enc(230) = 138
	// enc(231) = 247
	// enc(232) =  88
	// enc(233) =  63
	// enc(234) =  38
	// enc(235) =  13
	// enc(236) = 244
	// enc(237) = 219
	// enc(238) = 194
	// enc(239) = 169
	// enc(240) = 144
	// enc(241) = 119
	// enc(242) =  94
	// enc(243) =  69
	// enc(244) =  44
	// enc(245) =  19
	// enc(246) = 250
	// enc(247) = 225
	// enc(248) = 200
	// enc(249) = 175
	// enc(250) = 150
	// enc(251) = 125
	// enc(252) = 100
	// enc(253) =  75
	// enc(254) =  50

}

func TestGenKey(t *testing.T) {
	const max uint64 = math.MaxUint64
	key, err := GenKey(math.MaxUint64)
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}
	t.Logf("max: %d, key: %d", max, key)
	o := MustRestore(max, key)
	for i := uint64(1); i < 1000; i++ {
		t.Logf("enc(%d) = %d", i, o.Obfuscate(i))
	}
}

func TestRestoreOK(t *testing.T) {
	_, err := Restore(
		math.MaxUint64,
		17743098580093710193,
	)
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}
}

func TestRestoreFail(t *testing.T) {
	_, err := Restore(
		9999999999,
		5000000000,
	)
	if err == nil {
		t.Fatal("no error?!")
	}
	if _, ok := err.(ErrKey); !ok {
		t.Fatalf("unexpected error: %+v (%T)", err, err)
	}
}

func safeTest(o Obfuscator, m map[uint64]uint64, i uint64) func(t *testing.T) {
	return func(t *testing.T) {
		x := o.Obfuscate(i)
		y, err := o.Explain(x)
		if err != nil {
			t.Fatal("unexpected error: ", err)
		}

		if i != y {
			t.Fatalf(
				"cannot enc(%d) = %d but dec(enc(%d)) = %d",
				i, x, i, y,
			)
		}

		if m[x] > 0 {
			t.Fatalf(
				"found collision: enc(%d) = enc(%d) = %d",
				m[x], i, x,
			)
		}
		m[x] = i
	}
}

func safeBitTest(bits int) func(t *testing.T) {
	return func(t *testing.T) {
		max := uint64(1) << bits
		key, err := GenKey(max)
		if err != nil {
			t.Fatal("unexpected error: ", err)
		}
		o := MustRestore(max, key)
		m := map[uint64]uint64{}
		for j := uint64(1); j <= max; j++ {
			t.Run(fmt.Sprintf("#%d", j), safeTest(o, m, j))
		}
	}
}

func Test8bitEncodeDecode(t *testing.T) {
	safeBitTest(8)(t)
}

func Test16bitEncodeDecode(t *testing.T) {
	safeBitTest(16)(t)
}

func BenchmarkRetore(b *testing.B) {
	for x := 0; x < b.N; x++ {
		_, err := Restore(math.MaxUint64, 17743098580093710193)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	o, err := Restore(math.MaxUint64, 17743098580093710193)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for x := 0; x < b.N; x++ {
		o.Obfuscate(2)
	}
}

func BenchmarkDecode(b *testing.B) {
	o, err := Restore(math.MaxUint64, 17743098580093710193)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for x := 0; x < b.N; x++ {
		o.Explain(2)
	}
}
