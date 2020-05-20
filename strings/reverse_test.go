/*
	// This is the API you need to build for these tests. You will need to
	// change the import path in this test to point to your code.
	package reverse
	// String takes the specified string and reverses it.
	func String(str string) string
*/

package reverse_test

import (
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/strings/reverse"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestReverseString(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{"odd", "Hello World", "dlroW olleH"},
		{"even", "go", "og"},
		{"chinese", "汉字", "字汉"},

		// {"tworunes", "é́́", "é́́"}, -- Need to get this to work.
	}
	t.Log("Given the need to test reverse string functionality.")
	{
		for testID, test := range tt {
			tf := func(t *testing.T) {
				t.Logf("\tTest %d:\tWhen checking the word %q.", testID, test.input)
				{
