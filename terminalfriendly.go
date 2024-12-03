package terminalfriendly

import (
	"fmt"
	"strings"
)

/*
 * terminalfriendly package
 * Author: Johann Rehberger (wunderwuzzi23)
 * Blog: https://embracethered.com
 *
 * GetTerminalFriendlyString returns escaped control characters in the provided string in 
 * caret notation. The only characters that are kept are \r \t \n everything else is replace
 * with caret notation ^ and extended control characters 128-159 are shown as hex \x.
 * This is useful to safely print text from untrusted sources in Go.
 *
 * License: MIT
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 */

// IsControlCharacter determines if a character is a control character
func IsControlCharacter(char rune) bool {
	return (char >= 0 && char <= 31 && char != 9 && char != 10 && char != 13) ||
		char == 127 ||
		(char >= 128 && char <= 159)
}

// ToCaretNotation converts a control character to its caret notation
func ToCaretNotation(char rune) string {
	charCode := uint32(char)
	switch {
	case charCode <= 31:
		return string([]rune{'^', rune(charCode + 64)})
	case charCode == 127:
		return "^?"
	default:
		return fmt.Sprintf("\\x%02x", charCode)
	}
}

// GetTerminalFriendlyString converts a string to a terminal-friendly representation
func GetTerminalFriendlyString(text string) string {
	var result strings.Builder
	for _, char := range text {
		if IsControlCharacter(char) {
			result.WriteString(ToCaretNotation(char))
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}
