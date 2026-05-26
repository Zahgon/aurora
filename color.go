//
// Copyright (c) 2016-2022 The Aurora Authors. All rights reserved.
// This program is free software. It comes without any warranty,
// to the extent permitted by applicable law. You can redistribute
// it and/or modify it under the terms of the Unlicense. See LICENSE
// file for more details or see below.
//

//
// This is free and unencumbered software released into the public domain.
//
// Anyone is free to copy, modify, publish, use, compile, sell, or
// distribute this software, either in source code form or as a compiled
// binary, for any purpose, commercial or non-commercial, and by any
// means.
//
// In jurisdictions that recognize copyright laws, the author or authors
// of this software dedicate any and all copyright interest in the
// software to the public domain. We make this dedication for the benefit
// of the public at large and to the detriment of our heirs and
// successors. We intend this dedication to be an overt act of
// relinquishment in perpetuity of all present and future rights to this
// software under copyright law.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
// OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.
//
// For more information, please refer to <http://unlicense.org/>
//

package aurora

// A Color type is a color. It can contain
// one background color, one foreground color
// and a format, including ideogram related
// formats.
type Color uint

/*

	Developer note.

	The uint type is architecture depended and can be
	represented as int32 or int64.

	Thus, we can use 32-bits only to be fast and
	cross-platform.

	All supported formats requires 14 bits. It is
	first 14 bits.

	A foreground color requires 8 bit + 1 bit (presence flag).
	And the same for background color.

	The Color representations

	[ bg 8 bit ] [fg 8 bit ] [ fg/bg 2 bits ] [ fm 14 bits ]

	https://play.golang.org/p/fq2zcNstFoF

*/

// Special formats
const (
	BoldFm       Color = 1 << iota // 1
	FaintFm                        // 2
	ItalicFm                       // 3
	UnderlineFm                    // 4
	SlowBlinkFm                    // 5
	RapidBlinkFm                   // 6
	ReverseFm                      // 7
	ConcealFm                      // 8
	CrossedOutFm                   // 9

	FrakturFm         // 20
	DoublyUnderlineFm // 21 or bold off for some systems

	FramedFm    // 51
	EncircledFm // 52
	OverlinedFm // 53

	InverseFm       = ReverseFm    // alias to ReverseFm
	BlinkFm         = SlowBlinkFm  // alias to SlowBlinkFm
	HiddenFm        = ConcealFm    // alias to ConcealFm
	StrikeThroughFm = CrossedOutFm // alias to CrossedOutFm

	maskFm = BoldFm | FaintFm |
		ItalicFm | UnderlineFm |
		SlowBlinkFm | RapidBlinkFm |
		ReverseFm |
		ConcealFm | CrossedOutFm |

		FrakturFm | DoublyUnderlineFm |

		FramedFm | EncircledFm | OverlinedFm

	flagFg Color = 1 << 14 // presence flag (14th bit)
	flagBg Color = 1 << 15 // presence flag (15th bit)

	shiftFg = 16 // shift for foreground (starting from 16th bit)
	shiftBg = 24 // shift for background (starting from 24th bit)
)

// Foreground colors and related formats
const (

	// 8 bits

	// [  0;   7] - 30-37
	// [  8;  15] - 90-97 bright
	// [ 16; 231] - RGB
	// [232; 255] - grayscale

	BlackFg   Color = (iota << shiftFg) | flagFg // 30, 90
	RedFg                                        // 31, 91
	GreenFg                                      // 32, 92
	YellowFg                                     // 33, 93
	BlueFg                                       // 34, 94
	MagentaFg                                    // 35, 95
	CyanFg                                       // 36, 96
	WhiteFg                                      // 37, 97

	BrightFg Color = ((1 << 3) << shiftFg) | flagFg // -> 90

	// the BrightFg itself doesn't represent
	// a color, thus it has not flagFg

	// 5 bits

	//
	maskFg = (0xff << shiftFg) | flagFg
)

// Background colors and related formats
const (

	// 8 bits

	// [  0;   7] - 40-47
	// [  8;  15] - 100-107 bright
	// [ 16; 231] - RGB
	// [232; 255] - grayscale

	BlackBg   Color = (iota << shiftBg) | flagBg // 40, 100
	RedBg                                        // 41, 101
	GreenBg                                      // 42, 102
	YellowBg                                     // 43, 103
	BlueBg                                       // 44, 104
	MagentaBg                                    // 45, 105
	CyanBg                                       // 46, 106
	WhiteBg                                      // 47, 107

	BrightBg Color = ((1 << 3) << shiftBg) | flagBg // -> 100

	// the BrightBg itself doesn't represent
	// a color, thus it has not flagBg

	// 5 bits

	//
	maskBg = (0xff << shiftBg) | flagBg
)

const (
	availFlags = "-+# 0"
	esc        = "\033["
	clear      = esc + "0m"
)

// Nos returns string like 1;7;31;45. It
// may be an empty string for empty color.
// If the zero is true, then the string
// is prepended with 0;
func (c Color) Nos(zero bool) string { _ = "STUB: not implemented"; return "" }

func appendCond(bs []byte, cond, semi bool, vals ...byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// if the semi is true, then prepend with semicolon
func appendSemi(bs []byte, semi bool, vals ...byte) []byte { _ = "STUB: not implemented"; return nil }

func itoa(t byte) string { _ = "STUB: not implemented"; return "" }

func (c Color) appendFg(bs []byte, zero bool) []byte { _ = "STUB: not implemented"; return nil }

// 0- 7 :  30-37
// 8-15 :  90-97
// > 15 : 38;5;val

// '3' and the value itself

// '9' and the value itself
// clear bright flag

func (c Color) appendBg(bs []byte, zero bool) []byte { _ = "STUB: not implemented"; return nil }

// 0- 7 :  40- 47
// 8-15 : 100-107
// > 15 : 48;5;val

// '3' and the value itself

// '1', '0' and the value itself
// clear bright flag

func (c Color) appendFm9(bs []byte, zero bool) []byte { _ = "STUB: not implemented"; return nil }

// don't combine slow and rapid blink using only
// on of them, preferring slow blink

// including 1-2

// append 1;3;38;5;216 like string that represents ANSI
// color of the Color; the zero argument requires
// appending of '0' before to reset previous format
// and colors
func (c Color) appendNos(bs []byte, zero bool) []byte { _ = "STUB: not implemented"; return nil }

// reset previous

// formats
//

// 1-2

// don't combine bold and faint using only on of them, preferring bold

// 3-9

// 20-21

// 50-53

// foreground

// background

// ColorIndex is index of pre-defined 8-bit foreground or
// background colors from 0 to 255 (38;5;n).
//
//	  0-  7:  standard colors (as in ESC [ 30–37 m)
//	  8- 15:  high intensity colors (as in ESC [ 90–97 m)
//	 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
//	232-255:  grayscale from black to white in 24 steps
type ColorIndex uint8

// GrayIndex from 0 to 23.
type GrayIndex uint8

// The Colored interface represents a value with a Color.
type Colored interface {
	Color() Color // color of the value
}

// Reset returns Color without a color and formats.
func (c Color) Reset() Color {
	_ = "STUB: not implemented"

	// Formats
	return *new(Color)
}

// Bold or increased intensity (1).
func (c Color) Bold() Color { _ = "STUB: not implemented"; return *new(Color) }

// Faint, decreased intensity (2).
func (c Color) Faint() Color { _ = "STUB: not implemented"; return *new(Color) }

// DoublyUnderline or Bold off, double-underline
// per ECMA-48 (21).
func (c Color) DoublyUnderline() Color { _ = "STUB: not implemented"; return *new(Color) }

// Fraktur, rarely supported (20).
func (c Color) Fraktur() Color {
	_ = "STUB: not implemented"
	return *

	// Italic, not widely supported, sometimes
	// treated as inverse (3).
	new(Color)
}

func (c Color) Italic() Color {
	_ = "STUB: not implemented"
	return *

	// Underline (4).
	new(Color)
}

func (c Color) Underline() Color { _ = "STUB: not implemented"; return *new(Color) }

// SlowBlink, blinking less than 150
// per minute (5).
func (c Color) SlowBlink() Color { _ = "STUB: not implemented"; return *new(Color) }

// RapidBlink, blinking 150+ per minute,
// not widely supported (6).
func (c Color) RapidBlink() Color { _ = "STUB: not implemented"; return *new(Color) }

// Blink is alias for the SlowBlink.
func (c Color) Blink() Color {
	_ = "STUB: not implemented"
	return *

	// Reverse video, swap foreground and
	// background colors (7).
	new(Color)
}

func (c Color) Reverse() Color {
	_ = "STUB: not implemented"
	return *

	// Inverse is alias for the Reverse
	new(Color)
}

func (c Color) Inverse() Color {
	_ = "STUB: not implemented"

	// Conceal, hidden, not widely supported (8).
	return *new(Color)
}

func (c Color) Conceal() Color {
	_ = "STUB: not implemented"
	return *

	// Hidden is alias for the Conceal
	new(Color)
}

func (c Color) Hidden() Color {
	_ = "STUB: not implemented"

	// CrossedOut, characters legible, but
	// marked for deletion (9).
	return *new(Color)
}

func (c Color) CrossedOut() Color {
	_ = "STUB: not implemented"
	return *

	// StrikeThrough is alias for the CrossedOut.
	new(Color)
}

func (c Color) StrikeThrough() Color {
	_ = "STUB: not implemented"
	return *

	// Framed (51).
	new(Color)
}

func (c Color) Framed() Color {
	_ = "STUB: not implemented"
	return *

	// Encircled (52).
	new(Color)
}

func (c Color) Encircled() Color {
	_ = "STUB: not implemented"
	return *

	// Overlined (53).
	new(Color)
}

func (c Color) Overlined() Color {
	_ = "STUB: not implemented"
	return *

	// Foreground colors
	//
	// Black foreground color (30)
	new(Color)
}

func (c Color) Black() Color { _ = "STUB: not implemented"; return *new(Color) }

// Red foreground color (31)
func (c Color) Red() Color { _ = "STUB: not implemented"; return *new(Color) }

// Green foreground color (32)
func (c Color) Green() Color { _ = "STUB: not implemented"; return *new(Color) }

// Yellow foreground color (33)
func (c Color) Yellow() Color { _ = "STUB: not implemented"; return *new(Color) }

// Blue foreground color (34)
func (c Color) Blue() Color { _ = "STUB: not implemented"; return *new(Color) }

// Magenta foreground color (35)
func (c Color) Magenta() Color { _ = "STUB: not implemented"; return *new(Color) }

// Cyan foreground color (36)
func (c Color) Cyan() Color { _ = "STUB: not implemented"; return *new(Color) }

// White foreground color (37)
func (c Color) White() Color { _ = "STUB: not implemented"; return *new(Color) }

// Bright foreground colors
//
// BrightBlack foreground color (90)
func (c Color) BrightBlack() Color { _ = "STUB: not implemented"; return *new(Color) }

// BrightRed foreground color (91)
func (c Color) BrightRed() Color { _ = "STUB: not implemented"; return *new(Color) }

// BrightGreen foreground color (92)
func (c Color) BrightGreen() Color { _ = "STUB: not implemented"; return *new(Color) }

// BrightYellow foreground color (93)
func (c Color) BrightYellow() Color { _ = "STUB: not implemented"; return *new(Color) }

// BrightBlue foreground color (94)
func (c Color) BrightBlue() Color { _ = "STUB: not implemented"; return *new(Color) }

// BrightMagenta foreground color (95)
func (c Color) BrightMagenta() Color { _ = "STUB: not implemented"; return *new(Color) }

// BrightCyan foreground color (96)
func (c Color) BrightCyan() Color { _ = "STUB: not implemented"; return *new(Color) }

// BrightWhite foreground color (97)
func (c Color) BrightWhite() Color { _ = "STUB: not implemented"; return *new(Color) }

// Other
//
// Index of pre-defined 8-bit foreground color
// from 0 to 255 (38;5;n).
//
//	  0-  7:  standard colors (as in ESC [ 30–37 m)
//	  8- 15:  high intensity colors (as in ESC [ 90–97 m)
//	 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
//	232-255:  grayscale from black to white in 24 steps
func (c Color) Index(ci ColorIndex) Color { _ = "STUB: not implemented"; return *new(Color) }

// Gray from 0 to 23.
func (c Color) Gray(n GrayIndex) Color { _ = "STUB: not implemented"; return *new(Color) }

// Background colors
//
// BgBlack background color (40)
func (c Color) BgBlack() Color { _ = "STUB: not implemented"; return *new(Color) }

// BgRed background color (41)
func (c Color) BgRed() Color { _ = "STUB: not implemented"; return *new(Color) }

// BgGreen background color (42)
func (c Color) BgGreen() Color { _ = "STUB: not implemented"; return *new(Color) }

// BgYellow background color (43)
func (c Color) BgYellow() Color { _ = "STUB: not implemented"; return *new(Color) }

// BgBlue background color (44)
func (c Color) BgBlue() Color { _ = "STUB: not implemented"; return *new(Color) }

// BgMagenta background color (45)
func (c Color) BgMagenta() Color { _ = "STUB: not implemented"; return *new(Color) }

// BgCyan background color (46)
func (c Color) BgCyan() Color { _ = "STUB: not implemented"; return *new(Color) }

// BgWhite background color (47)
func (c Color) BgWhite() Color { _ = "STUB: not implemented"; return *new(Color) }

// Bright background colors
//
// BgBrightBlack background color (100)
func (c Color) BgBrightBlack() Color { _ = "STUB: not implemented"; return *new(Color) }

// BgBrightRed background color (101)
func (c Color) BgBrightRed() Color { _ = "STUB: not implemented"; return *new(Color) }

// BgBrightGreen background color (102)
func (c Color) BgBrightGreen() Color { _ = "STUB: not implemented"; return *new(Color) }

// BgBrightYellow background color (103)
func (c Color) BgBrightYellow() Color { _ = "STUB: not implemented"; return *new(Color) }

// BgBrightBlue background color (104)
func (c Color) BgBrightBlue() Color { _ = "STUB: not implemented"; return *new(Color) }

// BgBrightMagenta background color (105)
func (c Color) BgBrightMagenta() Color { _ = "STUB: not implemented"; return *new(Color) }

// BgBrightCyan background color (106)
func (c Color) BgBrightCyan() Color { _ = "STUB: not implemented"; return *new(Color) }

// BgBrightWhite background color (107)
func (c Color) BgBrightWhite() Color { _ = "STUB: not implemented"; return *new(Color) }

// Other
//
// BgIndex of 8-bit pre-defined background color
// from 0 to 255 (48;5;n).
//
//	  0-  7:  standard colors (as in ESC [ 40–47 m)
//	  8- 15:  high intensity colors (as in ESC [100–107 m)
//	 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
//	232-255:  grayscale from black to white in 24 steps
func (c Color) BgIndex(n ColorIndex) Color { _ = "STUB: not implemented"; return *new(Color) }

// BgGray from 0 to 23.
func (c Color) BgGray(n GrayIndex) Color { _ = "STUB: not implemented"; return *new(Color) }
