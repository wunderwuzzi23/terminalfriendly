# terminalfriendly

`GetTerminalFriendlyString` returns escaped control characters in the provided string in caret notation.

The only characters that are kept are \r \t \n everything else is replace with 
caret notation and extended control characters 128-159 are shown as hex \x.

This is useful to safely print text from untrusted sources in Go.
