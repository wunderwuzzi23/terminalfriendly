package terminalfriendly

import (
	"testing"
)

/*
 * terminalfriendly package tests
 * Author: Johann Rehberger (wunderwuzzi23)
 * Blog: https://embracethered.com
 *
 * This test file is designed to validate the functionality of the terminalfriendly package,
 * ensuring proper conversion of control characters and handling of terminal-friendly strings.
 */

func TestGetTerminalFriendlyString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		// Original test cases
		{
			input:    "Hello\x00World\x1F\x7F",
			expected: "Hello^@World^_^?",
		},
		{
			input:    "\x1b[5;31mEmbrace the Red\x1B[0m",
			expected: "^[[5;31mEmbrace the Red^[[0m",
		},
		{
			input:    "\x03Control\x04Codes",
			expected: "^CControl^DCodes",
		},
		{
			input:    "Normal Text",
			expected: "Normal Text",
		},
		{
			input:    "\x01Start\x02End\x1A",
			expected: "^AStart^BEnd^Z",
		},
		{
			input:    "Line\rBreak\nTab\t",
			expected: "Line\rBreak\nTab\t",
		},
		{
			input:    "hello\x00world\x1B\x1C\x1D\x1E\x1F\x7F",
			expected: "hello^@world^[^\\^]^^^_^?",
		},
		{
			input:    "Text\r\n\x1B[3Jnormal\x01\x02\x7f",
			expected: "Text\r\n^[[3Jnormal^A^B^?",
		},
		{
			input:    "lowercase\x03test\x04\x7F",
			expected: "lowercase^Ctest^D^?",
		},
		{
			input:    "\x00\x01\x1b\x1C\x1d\x1E\x1f\x7F",
			expected: "^@^A^[^\\^]^^^_^?",
		},
		{
			input:    "just normal text",
			expected: "just normal text",
		},
		{
			input:    " ",
			expected: " ",
		},
	}

	for _, tc := range testCases {
		result := GetTerminalFriendlyString(tc.input)
		if result != tc.expected {
			t.Errorf("Input: %q\nExpected: %q\nGot: %q", tc.input, tc.expected, result)
		}
	}
}
