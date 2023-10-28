package masking

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAndMaskLinks(t *testing.T) {
	testTable := []struct {
		maskedStr string
		expected  string
	}{
		{
			maskedStr: "Here's my spammy page: http://hehefouls.netHAHAHA see you.",
			expected:  "Here's my spammy page: http://******************* see you.",
		},
		{
			maskedStr: "Here's my spammy page: http://hehefouls.netHAHAHA see you. http://sdsd",
			expected:  "Here's my spammy page: http://******************* see you. http://****",
		},
		{
			maskedStr: "http://hehefouls.netHAHAHA see you. http://sdsd",
			expected:  "http://******************* see you. http://****",
		},
		{
			maskedStr: "htehefouls.netHAHAHA see you.",
			expected:  "htehefouls.netHAHAHA see you.",
		},
		{
			maskedStr: "Here's my spammy pagehttp://hehefouls.netHAHAHA see you. http://sdsd",
			expected:  "Here's my spammy pagehttp://hehefouls.netHAHAHA see you. http://****",
		},
		{
			maskedStr: "http://hehef",
			expected:  "http://*****",
		},
	}

	for _, testCase := range testTable {
		result := FindAndMaskLinks(testCase.maskedStr)
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrect result, expected \"%s\", result \"%s\"", testCase.expected, result))
	}
}
