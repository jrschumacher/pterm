package pterm

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
)

var (
	// Sprint formats using the default formats for its operands and returns the resulting string.
	// Spaces are added between operands when neither is a string.
	Sprint = color.Sprint

	// Sprintf formats according to a format specifier and returns the resulting string.
	Sprintf = color.Sprintf

	// Println formats using the default formats for its operands and writes to standard output.
	// Spaces are always added between operands and a newline is appended.
	// It returns the number of bytes written and any write error encountered.
	Println = color.Println

	// Printf formats according to a format specifier and writes to standard output.
	// It returns the number of bytes written and any write error encountered.
	Printf = color.Printf

	// Print formats using the default formats for its operands and writes to standard output.
	// Spaces are added between operands when neither is a string.
	// It returns the number of bytes written and any write error encountered.
	Print = color.Print
)

// Foreground colors. basic foreground colors 30 - 37
const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
	// FgDefault revert default FG
	FgDefault Color = 39
)

// Extra foreground color 90 - 97(非标准)
const (
	FgDarkGray Color = iota + 90 // 亮黑（灰）
	FgLightRed
	FgLightGreen
	FgLightYellow
	FgLightBlue
	FgLightMagenta
	FgLightCyan
	FgLightWhite
	// FgGray is an alias of FgDarkGray
	FgGray Color = 90
)

// Background colors. basic background colors 40 - 47
const (
	BgBlack Color = iota + 40
	BgRed
	BgGreen
	BgYellow // BgBrown like yellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	// BgDefault reverts to the default background
	BgDefault Color = 49
)

// Extra background color 100 - 107(非标准)
const (
	BgDarkGray Color = iota + 100
	BgLightRed
	BgLightGreen
	BgLightYellow
	BgLightBlue
	BgLightMagenta
	BgLightCyan
	BgLightWhite
	// BgGray is an alias of BgDarkGray
	BgGray Color = 100
)

// Option settings
const (
	Reset Color = iota
	Bold
	Fuzzy
	Italic
	Underscore
	Blink
	FastBlink
	Reverse
	Concealed
	Strikethrough
)

var (
	// Red is an alias for FgRed.Sprint
	Red = FgRed.Sprint
	// Cyan is an alias for FgCyan.Sprint
	Cyan = FgCyan.Sprint
	// Gray is an alias for FgGray.Sprint
	Gray = FgGray.Sprint
	// Blue is an alias for FgBlue.Sprint
	Blue = FgBlue.Sprint
	// Black is an alias for FgBlack.Sprint
	Black = FgBlack.Sprint
	// Green is an alias for FgGreen.Sprint
	Green = FgGreen.Sprint
	// White is an alias for FgWhite.Sprint
	White = FgWhite.Sprint
	// Yellow is an alias for FgYellow.Sprint
	Yellow = FgYellow.Sprint
	// Magenta is an alias for FgMagenta.Sprint
	Magenta = FgMagenta.Sprint

	// Normal is an alias for FgDefault.Sprint
	Normal = FgDefault.Sprint

	// extra light

	// LightRed is a shortcut for FgLightRed.Sprint
	LightRed = FgLightRed.Sprint
	// LightCyan is a shortcut for FgLightCyan.Sprint
	LightCyan = FgLightCyan.Sprint
	// LightBlue is a shortcut for FgLightBlue.Sprint
	LightBlue = FgLightBlue.Sprint
	// LightGreen is a shortcut for FgLightGreen.Sprint
	LightGreen = FgLightGreen.Sprint
	// LightWhite is a shortcut for FgLightWhite.Sprint
	LightWhite = FgLightWhite.Sprint
	// LightYellow is a shortcut for FgLightYellow.Sprint
	LightYellow = FgLightYellow.Sprint
	// LightMagenta is a shortcut for FgLightMagenta.Sprint
	LightMagenta = FgLightMagenta.Sprint
)

// Color is a number which will be used to color strings in the terminal
type Color uint8

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
// Input will be colored with the parent Color
func (c Color) Sprintln(a ...interface{}) string {
	return color.RenderCode(color.Color(c).String(), a...) + "\n"
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
// Input will be colored with the parent Color
func (c Color) Sprint(a ...interface{}) string {
	return color.RenderCode(color.Color(c).String(), a...)
}

// Sprintf formats according to a format specifier and returns the resulting string.
// Input will be colored with the parent Color
func (c Color) Sprintf(format string, a ...interface{}) string {
	return color.RenderString(color.Color(c).String(), fmt.Sprintf(format, a...))
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
// Input will be colored with the parent Color
func (c Color) Println(a ...interface{}) {
	Print(c.Sprintln(a...))
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
// Input will be colored with the parent Color
func (c Color) Print(a ...interface{}) {
	Print(c.Sprint(a...))
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
// Input will be colored with the parent Color
func (c Color) Printf(format string, a ...interface{}) {
	Print(c.Sprintf(format, a...))
}

// String converts the color to a string. eg "35"
func (c Color) String() string {
	return fmt.Sprintf("%d", c)
}

// Style is a collection of colors.
// Can include foreground, background and styling (eg. Bold, Underscore, etc.) colors.
type Style []Color

// NewStyle returns a new Style.
// Accepts multiple colors.
func NewStyle(colors ...Color) Style {
	return colors
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
// Input will be colored with the parent Style
func (s Style) Sprint(a ...interface{}) string {
	return color.RenderCode(s.String(), a...)
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
// Input will be colored with the parent Style
func (s Style) Sprintln(a ...interface{}) string {
	return color.RenderCode(s.String(), a...) + "\n"
}

// Sprintf formats according to a format specifier and returns the resulting string.
// Input will be colored with the parent Style
func (s Style) Sprintf(format string, a ...interface{}) string {
	return color.RenderString(s.String(), fmt.Sprintf(format, a...))
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
// Input will be colored with the parent Style
func (s Style) Print(a ...interface{}) {
	Print(s.Sprint(a...))
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
// Input will be colored with the parent Style
func (s Style) Println(a ...interface{}) {
	Println(s.Sprint(a...))
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
// Input will be colored with the parent Style
func (s Style) Printf(format string, a ...interface{}) {
	Print(s.Sprintf(format, a...))
}

// Code convert to code string. returns like "32;45;3"
func (s Style) Code() string {
	return s.String()
}

// String convert to code string. returns like "32;45;3"
func (s Style) String() string {
	return colors2code(s...)
}

// Converts colors to code.
// Return format: "32;45;3"
func colors2code(colors ...Color) string {
	if len(colors) == 0 {
		return ""
	}

	var codes []string
	for _, c := range colors {
		codes = append(codes, c.String())
	}

	return strings.Join(codes, ";")
}
