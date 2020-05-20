package reverse

// String takes the specified string and reverses it.
func String(str string) string {

	// Convert the input string into slice of runes for processing.
	// A rune represent a code point in the UTF-8 character set.
	runes := []rune(str)
	// Create an index that will traverse the collection of
	// runes from the beginning to the end.
	var beg int

	// Create an index that will traverse the collection of
	// runes from the end to the beginning.
	end := len(runes) - 1
