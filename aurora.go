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

// Package aurora implements ANSI-colors
package aurora

type Aurora struct {
	conf Config
	cc   colorConfig
}

// New returns new colorizer by given Options.
func New(opts ...Option) (a *Aurora) { _ = "STUB: not implemented"; return nil }

// set defaults
// apply options
// keep the short hand

// Config of the colorizer. It returns copy of the configurations.
func (a *Aurora) Config() Config {
	_ = "STUB: not implemented"

	// Reset wraps given argument returning Value without formats, colors and links.
	return *new(Config)
}

func (a *Aurora) Reset(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Clear wraps given argument returning Value without formats and colors. But
// preserving links.
func (a *Aurora) Clear(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Formats
//
// Bold or increased intensity (1).
func (a *Aurora) Bold(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Faint, decreased intensity (2).
func (a *Aurora) Faint(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// DoublyUnderline or Bold off, double-underline per ECMA-48 (21).
func (a *Aurora) DoublyUnderline(arg interface{}) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

// Fraktur, rarely supported (20).
func (a *Aurora) Fraktur(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Italic, not widely supported, sometimes treated as inverse (3).
func (a *Aurora) Italic(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Underline (4).
func (a *Aurora) Underline(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// SlowBlink, blinking less than 150 per minute (5).
func (a *Aurora) SlowBlink(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// RapidBlink, blinking 150+ per minute, not widely supported (6).
func (a *Aurora) RapidBlink(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Blink is alias for the SlowBlink.
func (a *Aurora) Blink(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Reverse video, swap foreground and background colors (7).
func (a *Aurora) Reverse(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Inverse is alias for the Reverse
func (a *Aurora) Inverse(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Conceal, hidden, not widely supported (8).
func (a *Aurora) Conceal(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Hidden is alias for the Conceal.
func (a *Aurora) Hidden(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// CrossedOut, characters legible, but marked for deletion (9).
func (a *Aurora) CrossedOut(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// StrikeThrough is alias for the CrossedOut.
func (a *Aurora) StrikeThrough(arg interface{}) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

// Framed (51).
func (a *Aurora) Framed(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Encircled (52).
func (a *Aurora) Encircled(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Overlined (53).
func (a *Aurora) Overlined(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Foreground colors
//
// Black foreground color (30).
func (a *Aurora) Black(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Red foreground color (31).
func (a *Aurora) Red(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Green foreground color (32).
func (a *Aurora) Green(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Yellow foreground color (33).
func (a *Aurora) Yellow(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Blue foreground color (34).
func (a *Aurora) Blue(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Magenta foreground color (35).
func (a *Aurora) Magenta(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Cyan foreground color (36).
func (a *Aurora) Cyan(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// White foreground color (37).
func (a *Aurora) White(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Bright foreground colors.
//
// BrightBlack foreground color (90).
func (a *Aurora) BrightBlack(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightRed foreground color (91).
func (a *Aurora) BrightRed(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightGreen foreground color (92).
func (a *Aurora) BrightGreen(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightYellow foreground color (93).
func (a *Aurora) BrightYellow(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightBlue foreground color (94).
func (a *Aurora) BrightBlue(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightMagenta foreground color (95).
func (a *Aurora) BrightMagenta(arg interface{}) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

// BrightCyan foreground color (96).
func (a *Aurora) BrightCyan(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightWhite foreground color (97).
func (a *Aurora) BrightWhite(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Other colors.
//
// Index of pre-defined 8-bit foreground color from 0 to 255 (38;5;n).
//
//	  0-  7:  standard colors (as in ESC [ 30–37 m)
//	  8- 15:  high intensity colors (as in ESC [ 90–97 m)
//	 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
//	232-255:  grayscale from black to white in 24 steps
func (a *Aurora) Index(n ColorIndex, arg interface{}) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

// Gray from 0 to 23.
func (a *Aurora) Gray(n GrayIndex, arg interface{}) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

// Background colors.
//
// BgBlack background color (40).
func (a *Aurora) BgBlack(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgRed background color (41).
func (a *Aurora) BgRed(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgGreen background color (42).
func (a *Aurora) BgGreen(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgYellow background color (43).
func (a *Aurora) BgYellow(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBlue background color (44).
func (a *Aurora) BgBlue(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgMagenta background color (45).
func (a *Aurora) BgMagenta(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgCyan background color (46).
func (a *Aurora) BgCyan(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgWhite background color (47).
func (a *Aurora) BgWhite(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Bright background colors.
//
// BgBrightBlack background color (100).
func (a *Aurora) BgBrightBlack(arg interface{}) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

// BgBrightRed background color (101).
func (a *Aurora) BgBrightRed(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBrightGreen background color (102).
func (a *Aurora) BgBrightGreen(arg interface{}) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

// BgBrightYellow background color (103).
func (a *Aurora) BgBrightYellow(arg interface{}) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

// BgBrightBlue background color (104).
func (a *Aurora) BgBrightBlue(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBrightMagenta background color (105).
func (a *Aurora) BgBrightMagenta(arg interface{}) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

// BgBrightCyan background color (106).
func (a *Aurora) BgBrightCyan(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBrightWhite background color (107).
func (a *Aurora) BgBrightWhite(arg interface{}) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

// Other background colors.
//
// BgIndex of 8-bit pre-defined background color from 0 to 255 (48;5;n).
//
//	  0-  7:  standard colors (as in ESC [ 40–47 m)
//	  8- 15:  high intensity colors (as in ESC [100–107 m)
//	 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
//	232-255:  grayscale from black to white in 24 steps
func (a *Aurora) BgIndex(n ColorIndex, arg interface{}) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

// BgGray from 0 to 23.
func (a *Aurora) BgGray(n GrayIndex, arg interface{}) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

// Special color functions.
//
// Colorize removes existing colors and
// formats of the argument and applies given.
func (a *Aurora) Colorize(arg interface{}, color Color) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

// Hyperlinks feature
//
// Hyperlink with given target and parameters. If hyperlinks feature is
// disabled, then the 'arg' argument dropped and the 'target' used instead
// inheriting all colors and format from the 'arg' (if it's a Colored).
//
// See https://gist.github.com/egmontkob/eb114294efbcd5adb1944c9f3cb5feda
// for details about the hyperlinks feature.
//
// The Hyperlink doesn't escape the target and the params. They should be
// checked and escaped before.
//
// See also HyperlinkID function.
//
// For a simple example
//
//	au.Hyperlink("Example", "http://example.com")
//
// and an example with ID
//
//	au.Hyperlink("Example", "http://example.com", aurora.HyperlinkID("10"))
func (a *Aurora) Hyperlink(arg interface{}, target string,
	params ...HyperlinkParam) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

// HyperlinkTarget of the argument if it's a Value.
func (a *Aurora) HyperlinkTarget(arg interface{}) (target string) {
	_ = "STUB: not implemented"
	return ""
}

// no target

// HyperlinkParams of the argument if it's a Value.
func (a *Aurora) HyperlinkParams(arg interface{}) (params []HyperlinkParam) {
	_ = "STUB: not implemented"
	return nil
}

// no target

func (a *Aurora) transform(arg interface{}) (val Value, ok bool) {
	_ = "STUB: not implemented"
	return *new(Value), false
}

// Value{}, false

// if ai.cc.resetColor() == a.cc.resetColor() {
// 	return // don't replace, same configurations
// }

// transformed value, true

// Sprintf allows to use Value as format. For example
//
//	var v = Sprintf(Red("total: +3.5f points"), Blue(3.14))
//
// In this case "total:" and "points" will be red, but
// 3.14 will be blue. But, in another example
//
//	var v = Sprintf(Red("total: +3.5f points"), 3.14)
//
// full string will be red. And no way to clear 3.14 to default format and
// color.
//
// It applies own configurations to all given Values.
func (a *Aurora) Sprintf(format interface{}, args ...interface{}) string {
	_ = "STUB: not implemented"
	// // clear colors & links as configured by the a
	return ""
}
