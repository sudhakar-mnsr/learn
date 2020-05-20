package reverse

// String takes the specified string and reverses it.
func String(str string) string {

	// Convert the input string into slice of runes for processing.
	// A rune represent a code point in the UTF-8 character set.
	runes := []rune(str)
