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

import (
	"fmt"
)

// compile-time check
var (
	_ fmt.Stringer  = Value{}
	_ fmt.Formatter = Value{}
	_ Colored       = Value{}
)

func coloredFormat(color Color, s fmt.State, verb rune) string {
	_ = "STUB: not implemented"

	// it's enough for many cases (%-+020.10f)
	// %          - 1
	// availFlags - 3 (5)
	// width      - 2
	// prec       - 3 (.23)
	// verb       - 1
	// --------------
	//             10
	// +
	// \033[                            5
	// 0;1;3;4;5;7;8;9;20;21;51;52;53  30
	// 38;5;216                         8
	// 48;5;216                         8
	// m                                1
	// +
	// \033[0m                          7
	//
	// x2 (possible tail color)
	//
	// 10 + 59 * 2 = 128
	return ""
}

// just clear

type colorConfig uint64

const (
	colorPin      colorConfig = 1 << 32
	hyperlinksPin colorConfig = 1 << 33
)

func (cc colorConfig) colorsEnabled() bool { _ = "STUB: not implemented"; return false }

func (cc colorConfig) hyperlinksEnbaled() bool { _ = "STUB: not implemented"; return false }

func (cc colorConfig) color() Color { _ = "STUB: not implemented"; return *new(Color) }

// lower 32 bits only

// even if a color set

func (cc colorConfig) resetColor() colorConfig { _ = "STUB: not implemented"; return *new(colorConfig) }

// A Value represents any printable value
// with or without colors, formats and a link.
type Value struct {
	value     interface{} // value as is
	cc        colorConfig // color & config
	hyperlink *hyperlink  // hyperlink target and parameters
}

// String implements standard fmt.Stringer interface.
func (v Value) String() string { _ = "STUB: not implemented"; return "" }

// calculate length

// fill

// no links, only colors & formats

// no links, no colors, no formats, just the value

// Color returns colors and formats of the Value.
func (v Value) Color() Color {
	_ = "STUB: not implemented"
	return *

	// Reset colors, formats and links.
	new(Color)
}

func (v Value) Reset() Value { _ = "STUB: not implemented"; return *new(Value) }

// Clear colors and formats, preserving links.
func (v Value) Clear() Value { _ = "STUB: not implemented"; return *new(Value) }

// Value returns value's value (welcome to the tautology club)
func (v Value) Value() interface{} {
	_ = "STUB: not implemented"

	// Format implements standard fmt.Formatter interface.
	return nil
}

func (v Value) Format(s fmt.State, verb rune) { _ = "STUB: not implemented"; return }

// Formats
//
// Bold or increased intensity (1).
func (v Value) Bold() Value { _ = "STUB: not implemented"; return *new(Value) }

// Faint, decreased intensity, reset the Bold (2).
func (v Value) Faint() Value { _ = "STUB: not implemented"; return *new(Value) }

// DoublyUnderline or Bold off, double-underline per ECMA-48 (21). It depends.
func (v Value) DoublyUnderline() Value { _ = "STUB: not implemented"; return *new(Value) }

// Fraktur, rarely supported (20).
func (v Value) Fraktur() Value { _ = "STUB: not implemented"; return *new(Value) }

// Italic, not widely supported, sometimes treated as inverse (3).
func (v Value) Italic() Value { _ = "STUB: not implemented"; return *new(Value) }

// Underline (4).
func (v Value) Underline() Value { _ = "STUB: not implemented"; return *new(Value) }

// SlowBlink, blinking less than 150 per minute (5).
func (v Value) SlowBlink() Value { _ = "STUB: not implemented"; return *new(Value) }

// RapidBlink, blinking 150+ per minute, not widely supported (6).
func (v Value) RapidBlink() Value { _ = "STUB: not implemented"; return *new(Value) }

// Blink is alias for the SlowBlink.
func (v Value) Blink() Value {
	_ = "STUB: not implemented"
	return *

	// Reverse video, swap foreground and background colors (7).
	new(Value)
}

func (v Value) Reverse() Value { _ = "STUB: not implemented"; return *new(Value) }

// Inverse is alias for the Reverse.
func (v Value) Inverse() Value {
	_ = "STUB: not implemented"

	// Conceal, hidden, not widely supported (8).
	return *new(Value)
}

func (v Value) Conceal() Value { _ = "STUB: not implemented"; return *new(Value) }

// Hidden is alias for the Conceal.
func (v Value) Hidden() Value {
	_ = "STUB: not implemented"

	// CrossedOut, characters legible, but marked for deletion (9).
	return *new(Value)
}

func (v Value) CrossedOut() Value { _ = "STUB: not implemented"; return *new(Value) }

// StrikeThrough is alias for the CrossedOut.
func (v Value) StrikeThrough() Value {
	_ = "STUB: not implemented"
	return *

	// Framed (51).
	new(Value)
}

func (v Value) Framed() Value { _ = "STUB: not implemented"; return *new(Value) }

// Encircled (52).
func (v Value) Encircled() Value { _ = "STUB: not implemented"; return *new(Value) }

// Overlined (53).
func (v Value) Overlined() Value { _ = "STUB: not implemented"; return *new(Value) }

// Foreground colors.
//
// Black foreground color (30).
func (v Value) Black() Value { _ = "STUB: not implemented"; return *new(Value) }

// Red foreground color (31).
func (v Value) Red() Value { _ = "STUB: not implemented"; return *new(Value) }

// Green foreground color (32).
func (v Value) Green() Value { _ = "STUB: not implemented"; return *new(Value) }

// Yellow foreground color (33).
func (v Value) Yellow() Value { _ = "STUB: not implemented"; return *new(Value) }

// Blue foreground color (34).
func (v Value) Blue() Value { _ = "STUB: not implemented"; return *new(Value) }

// Magenta foreground color (35).
func (v Value) Magenta() Value { _ = "STUB: not implemented"; return *new(Value) }

// Cyan foreground color (36).
func (v Value) Cyan() Value { _ = "STUB: not implemented"; return *new(Value) }

// White foreground color (37).
func (v Value) White() Value { _ = "STUB: not implemented"; return *new(Value) }

// Bright foreground colors.
//
// BrightBlack foreground color (90).
func (v Value) BrightBlack() Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightRed foreground color (91).
func (v Value) BrightRed() Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightGreen foreground color (92).
func (v Value) BrightGreen() Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightYellow foreground color (93).
func (v Value) BrightYellow() Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightBlue foreground color (94).
func (v Value) BrightBlue() Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightMagenta foreground color (95).
func (v Value) BrightMagenta() Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightCyan foreground color (96).
func (v Value) BrightCyan() Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightWhite foreground color (97).
func (v Value) BrightWhite() Value { _ = "STUB: not implemented"; return *new(Value) }

// Other colors.
//
// Index of pre-defined 8-bit foreground color from 0 to 255 (38;5;n).
//
//	  0-  7:  standard colors (as in ESC [ 30–37 m)
//	  8- 15:  high intensity colors (as in ESC [ 90–97 m)
//	 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
//	232-255:  grayscale from black to white in 24 steps
func (v Value) Index(n ColorIndex) Value { _ = "STUB: not implemented"; return *new(Value) }

// Gray from 0 to 24.
func (v Value) Gray(n GrayIndex) Value { _ = "STUB: not implemented"; return *new(Value) }

// Background colors
//
// BgBlack background color (40).
func (v Value) BgBlack() Value { _ = "STUB: not implemented"; return *new(Value) }

// BgRed background color (41).
func (v Value) BgRed() Value { _ = "STUB: not implemented"; return *new(Value) }

// BgGreen background color (42).
func (v Value) BgGreen() Value { _ = "STUB: not implemented"; return *new(Value) }

// BgYellow background color (43).
func (v Value) BgYellow() Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBlue background color (44).
func (v Value) BgBlue() Value { _ = "STUB: not implemented"; return *new(Value) }

// BgMagenta background color (45).
func (v Value) BgMagenta() Value { _ = "STUB: not implemented"; return *new(Value) }

// BgCyan background color (46).
func (v Value) BgCyan() Value { _ = "STUB: not implemented"; return *new(Value) }

// BgWhite background color (47).
func (v Value) BgWhite() Value { _ = "STUB: not implemented"; return *new(Value) }

// Bright background colors.
//
// BgBrightBlack background color (100).
func (v Value) BgBrightBlack() Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBrightRed background color (101).
func (v Value) BgBrightRed() Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBrightGreen background color (102).
func (v Value) BgBrightGreen() Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBrightYellow background color (103).
func (v Value) BgBrightYellow() Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBrightBlue background color (104).
func (v Value) BgBrightBlue() Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBrightMagenta background color (105).
func (v Value) BgBrightMagenta() Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBrightCyan background color (106).
func (v Value) BgBrightCyan() Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBrightWhite background color (107).
func (v Value) BgBrightWhite() Value { _ = "STUB: not implemented"; return *new(Value) }

// Other background colors.
//
// BgIndex of 8-bit pre-defined background color from 0 to 255 (48;5;n).
//
//	  0-  7:  standard colors (as in ESC [ 40–47 m)
//	  8- 15:  high intensity colors (as in ESC [100–107 m)
//	 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
//	232-255:  grayscale from black to white in 24 steps
func (v Value) BgIndex(n ColorIndex) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgGray from 0 to 24.
func (v Value) BgGray(n GrayIndex) Value { _ = "STUB: not implemented"; return *new(Value) }

// Special colorization method.
//
// Colorize removes existing colors and formats of the argument and applies
// given.
func (v Value) Colorize(color Color) Value { _ = "STUB: not implemented"; return *new(Value) }

// Hyperlinks feature
//
// Hyperlink with given target and parameters. If hyperlinks feature is
// disabled, then the 'arg' argument dropped and the 'target' used instead,
// inheriting all colors and format from the 'arg' (if it's a Colored).
//
// See https://gist.github.com/egmontkob/eb114294efbcd5adb1944c9f3cb5feda
// for details about the hyperlinks feature.
//
// The Hyperlink doesn't escape the target and the parameters. They should be
// checked and escaped before. See HyperlinkEscape function.
//
// See also HyperlinkID function.
//
// For a simple example
//
//	val.Hyperlink("http://example.com")
//
// and an example with ID
//
//	val.Hyperlink("http://example.com", aurora.HyperlinkID("10"))
//
// Successive calls replace previously set target and parameters.
func (v Value) Hyperlink(target string, params ...HyperlinkParam) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

// drop value, use the target

// keep for the HyperlinkTarget method

// HyperlinkTarget if any.
func (v Value) HyperlinkTarget() (target string) { _ = "STUB: not implemented"; return "" }

// nothing

// HyperlinkParams if any.
func (v Value) HyperlinkParams() (params []HyperlinkParam) { _ = "STUB: not implemented"; return nil }

// nil
