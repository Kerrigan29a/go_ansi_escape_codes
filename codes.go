package ansi_escape_codes

import (
	"strconv"
	"strings"
)

type ANSICode string
type CSISequence []ANSICode

const (
	ESC ANSICode = "\u001b"
	CSI ANSICode = ESC + "["
)

var (
	/* Decorations */
	Reset       CSISequence = CSISequence{"0", "m"}
	Bold        CSISequence = CSISequence{"1", "m"}
	Bright                  = Bold
	Underline   CSISequence = CSISequence{"4", "m"}
	Negative    CSISequence = CSISequence{"7", "m"}
	Reverse                 = Negative
	NoUnderline CSISequence = CSISequence{"24", "m"}
	NoNegative  CSISequence = CSISequence{"24", "m"}
	NoReverse               = NoNegative
	Positive                = NoNegative

	/* Foreground */
	FgBlack         CSISequence = CSISequence{"30", "m"}
	FgRed           CSISequence = CSISequence{"31", "m"}
	FgGreen         CSISequence = CSISequence{"32", "m"}
	FgYellow        CSISequence = CSISequence{"33", "m"}
	FgBlue          CSISequence = CSISequence{"34", "m"}
	FgMagenta       CSISequence = CSISequence{"35", "m"}
	FgCyan          CSISequence = CSISequence{"36", "m"}
	FgWhite         CSISequence = CSISequence{"37", "m"}
	FgExtendedColor CSISequence = CSISequence{"38", "m"}
	FgDefaultColor  CSISequence = CSISequence{"39", "m"}
	FgBrightBlack   CSISequence = CSISequence{"90", "m"}
	FgBrightRed     CSISequence = CSISequence{"91", "m"}
	FgBrightGreen   CSISequence = CSISequence{"92", "m"}
	FgBrightYellow  CSISequence = CSISequence{"93", "m"}
	FgBrightBlue    CSISequence = CSISequence{"94", "m"}
	FgBrightMagenta CSISequence = CSISequence{"95", "m"}
	FgBrightCyan    CSISequence = CSISequence{"96", "m"}
	FgBrightWhite   CSISequence = CSISequence{"97", "m"}

	/* Background */
	BgBlack         CSISequence = CSISequence{"40", "m"}
	BgRed           CSISequence = CSISequence{"41", "m"}
	BgGreen         CSISequence = CSISequence{"42", "m"}
	BgYellow        CSISequence = CSISequence{"43", "m"}
	BgBlue          CSISequence = CSISequence{"44", "m"}
	BgMagenta       CSISequence = CSISequence{"45", "m"}
	BgCyan          CSISequence = CSISequence{"46", "m"}
	BgWhite         CSISequence = CSISequence{"47", "m"}
	BgExtendedColor CSISequence = CSISequence{"48", "m"}
	BgDefaultColor  CSISequence = CSISequence{"49", "m"}
	BgBrightBlack   CSISequence = CSISequence{"100", "m"}
	BgBrightRed     CSISequence = CSISequence{"101", "m"}
	BgBrightGreen   CSISequence = CSISequence{"102", "m"}
	BgBrightYellow  CSISequence = CSISequence{"103", "m"}
	BgBrightBlue    CSISequence = CSISequence{"104", "m"}
	BgBrightMagenta CSISequence = CSISequence{"105", "m"}
	BgBrightCyan    CSISequence = CSISequence{"106", "m"}
	BgBrightWhite   CSISequence = CSISequence{"107", "m"}
)

func (s CSISequence) String() string {
	return string(CSI) + s.ParamsString()
}

func (s CSISequence) ParamsString() string {
	b := strings.Builder{}
	for _, code := range s {
		b.WriteString(string(code))
	}
	return b.String()
}

func (s1 CSISequence) Compose(s2 CSISequence) CSISequence {
	if s1[len(s1)-1] != "m" {
		panic("(Sequene 1) Invalid CSISequence: " + s1.ParamsString())
	}
	if s2[len(s2)-1] != "m" {
		panic("(Sequene 2) Invalid CSISequence: " + s2.ParamsString())
	}
	result := CSISequence(append([]ANSICode(nil), s1[:len(s1)-1]...))
	for _, code := range s2[:len(s2)-1] {
		result = append(result, ";")
		result = append(result, code)
	}
	result = append(result, "m")
	return result
}

func ExtendedColor(code uint8, background bool) CSISequence {
	head := FgExtendedColor
	if background {
		head = BgExtendedColor
	}
	return CSISequence{head[0], ";", "5", ";", ANSICode(strconv.FormatUint(uint64(code), 10)), "m"}
}

func RGBColor(r, g, b uint8, background bool) CSISequence {
	head := FgExtendedColor
	if background {
		head = BgExtendedColor
	}
	return CSISequence{head[0], ";", "2", ";",
		ANSICode(strconv.FormatUint(uint64(r), 10)), ";",
		ANSICode(strconv.FormatUint(uint64(g), 10)), ";",
		ANSICode(strconv.FormatUint(uint64(b), 10)), ";",
		"m"}
}
