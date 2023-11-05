package masking

import (
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
		//result := FindAndMaskLinks(testCase.maskedStr)
		result := ""
		assert.Equal(t, testCase.expected, result)
	}
}

/*
func TestService_findAndMaskLinks(t *testing.T) {
	type fields struct {
		prod producer
		pres presenter
	}
	type args struct {
		sourceStr string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				prod: tt.fields.prod,
				pres: tt.fields.pres,
			}
			assert.Equalf(t, tt.want, s.findAndMaskLinks(tt.args.sourceStr), "findAndMaskLinks(%v)", tt.args.sourceStr)
		})
	}
}
*/
