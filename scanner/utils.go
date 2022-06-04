package scanner

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\n' || ch == '\r' || ch == '\t'
}

func isAsciiLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
