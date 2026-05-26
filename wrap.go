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

// DefaultColorizer is global colorizer that used for package root color
// methods.
var DefaultColorizer = New(WithColors(true), WithHyperlinks(true))

// Colorize wraps given value into Value with given colors. For example
//
//	var s = Colorize("some", BlueFg|GreenBg|BoldFm)
//
// returns a Value with blue foreground, green background and bold. Unlike
// functions like Red/BgBlue/Bold etc. This function clears all previous colors
// and formats. Thus
//
//	var s = Colorize(Red("some"), BgBlue)
//
// clears red color from value.
func Colorize(arg interface{}, color Color) Value { _ = "STUB: not implemented"; return *new(Value) }

// Reset wraps given argument returning Value without formats, colors and links.
func Reset(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Clear wraps given argument returning Value without formats and colors. But
// preserving links.
func Clear(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

//
// Formats
//

// Bold or increased intensity (1).
func Bold(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Faint decreases intensity (2). The Faint rejects the Bold.
func Faint(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// DoublyUnderline or Bold off, double-underline per ECMA-48 (21).
func DoublyUnderline(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Fraktur is rarely supported (20).
func Fraktur(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Italic is not widely supported, sometimes treated as inverse (3).
func Italic(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Underline (4).
func Underline(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// SlowBlink makes text blink less than 150 per minute (5).
func SlowBlink(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// RapidBlink makes text blink 150+ per minute. It is not widely supported (6).
func RapidBlink(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Blink is alias for the SlowBlink.
func Blink(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Reverse video, swap foreground and background colors (7).
func Reverse(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Inverse is alias for the Reverse
func Inverse(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Conceal hides text, preserving an ability to select the text and copy it. It
// is not widely supported (8).
func Conceal(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Hidden is alias for the Conceal
func Hidden(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// CrossedOut makes characters legible, but marked for deletion (9).
func CrossedOut(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// StrikeThrough is alias for the CrossedOut.
func StrikeThrough(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Framed (51).
func Framed(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Encircled (52).
func Encircled(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Overlined (53).
func Overlined(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

//
// Foreground colors
//
//

// Black foreground color (30)
func Black(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Red foreground color (31)
func Red(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Green foreground color (32)
func Green(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Yellow foreground color (33)
func Yellow(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Blue foreground color (34)
func Blue(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Magenta foreground color (35)
func Magenta(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Cyan foreground color (36)
func Cyan(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// White foreground color (37)
func White(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

//
// Bright foreground colors
//

// BrightBlack foreground color (90)
func BrightBlack(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightRed foreground color (91)
func BrightRed(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightGreen foreground color (92)
func BrightGreen(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightYellow foreground color (93)
func BrightYellow(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightBlue foreground color (94)
func BrightBlue(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightMagenta foreground color (95)
func BrightMagenta(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightCyan foreground color (96)
func BrightCyan(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BrightWhite foreground color (97)
func BrightWhite(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

//
// Other
//

// Index of pre-defined 8-bit foreground color from 0 to 255 (38;5;n).
//
//	  0-  7:  standard colors (as in ESC [ 30–37 m)
//	  8- 15:  high intensity colors (as in ESC [ 90–97 m)
//	 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
//	232-255:  grayscale from black to white in 24 steps
func Index(n ColorIndex, arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// Gray from 0 to 24.
func Gray(n GrayIndex, arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

//
// Background colors
//
//

// BgBlack background color (40)
func BgBlack(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgRed background color (41)
func BgRed(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgGreen background color (42)
func BgGreen(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgYellow background color (43)
func BgYellow(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBlue background color (44)
func BgBlue(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgMagenta background color (45)
func BgMagenta(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgCyan background color (46)
func BgCyan(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgWhite background color (47)
func BgWhite(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

//
// Bright background colors
//

// BgBrightBlack background color (100)
func BgBrightBlack(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBrightRed background color (101)
func BgBrightRed(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBrightGreen background color (102)
func BgBrightGreen(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBrightYellow background color (103)
func BgBrightYellow(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBrightBlue background color (104)
func BgBrightBlue(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBrightMagenta background color (105)
func BgBrightMagenta(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBrightCyan background color (106)
func BgBrightCyan(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgBrightWhite background color (107)
func BgBrightWhite(arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

//
// Other
//

// BgIndex of 8-bit pre-defined background color from 0 to 255 (48;5;n).
//
//	  0-  7:  standard colors (as in ESC [ 40–47 m)
//	  8- 15:  high intensity colors (as in ESC [100–107 m)
//	 16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
//	232-255:  grayscale from black to white in 24 steps
func BgIndex(n ColorIndex, arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// BgGray from 0 to 24.
func BgGray(n GrayIndex, arg interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

//
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
func Hyperlink(arg interface{}, target string, params ...HyperlinkParam) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

// HyperlinkTarget of the argument if it's a Value.
func HyperlinkTarget(arg interface{}) (target string) { _ = "STUB: not implemented"; return "" }

// HyperlinkParams of the argument if it's a Value.
func HyperlinkParams(arg interface{}) (params []HyperlinkParam) {
	_ = "STUB: not implemented"
	return nil
}

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
func Sprintf(format interface{}, args ...interface{}) string { _ = "STUB: not implemented"; return "" }
