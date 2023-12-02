package masking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findAndMaskLinks(t *testing.T) {
	var prod producer
	var pres presenter
	s := &Service{prod, pres}

	const countG = 10

	type args struct {
		sourceStr   string
		chMaskStr   chan string
		chMaskedStr chan string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "middle link",
			args: args{
				sourceStr:   "Here's my spammy page: http://hehefouls.netHAHAHA see you.",
				chMaskStr:   make(chan string, countG),
				chMaskedStr: make(chan string, countG),
			},
			want: "Here's my spammy page: http://******************* see you.",
		},
		{
			name: "beggin, end link",
			args: args{
				sourceStr:   "http://hehefouls.netHAHAHA see you. http://sdsd",
				chMaskStr:   make(chan string, countG),
				chMaskedStr: make(chan string, countG),
			},
			want: "http://******************* see you. http://****",
		},
		{
			name: "no link",
			args: args{
				sourceStr:   "htehefouls.netHAHAHA see you.",
				chMaskStr:   make(chan string, countG),
				chMaskedStr: make(chan string, countG),
			},
			want: "htehefouls.netHAHAHA see you.",
		},
		{
			name: "middle fake, end link",
			args: args{
				sourceStr:   "Here's my spammy pagehttp://hehefouls.netHAHAHA see you. http://sdsd",
				chMaskStr:   make(chan string, countG),
				chMaskedStr: make(chan string, countG),
			},
			want: "Here's my spammy pagehttp://hehefouls.netHAHAHA see you. http://****",
		},
		{
			name: "only link",
			args: args{
				sourceStr:   "http://hehef",
				chMaskStr:   make(chan string, countG),
				chMaskedStr: make(chan string, countG),
			},
			want: "http://*****",
		},
		{
			name: "one char",
			args: args{
				sourceStr:   "k",
				chMaskStr:   make(chan string, countG),
				chMaskedStr: make(chan string, countG),
			},
			want: "k",
		},
		{
			name: "link like link",
			args: args{
				sourceStr:   "http://http://http:/",
				chMaskStr:   make(chan string, countG),
				chMaskedStr: make(chan string, countG),
			},
			want: "http://*************",
		},
		{
			name: "integers",
			args: args{
				sourceStr:   "84857786",
				chMaskStr:   make(chan string, countG),
				chMaskedStr: make(chan string, countG),
			},
			want: "84857786",
		},
		{
			name: "spaces",
			args: args{
				sourceStr:   "          ",
				chMaskStr:   make(chan string, countG),
				chMaskedStr: make(chan string, countG),
			},
			want: "          ",
		},
		{
			name: "strange symbols",
			args: args{
				sourceStr:   "[]]!!#$@Fr",
				chMaskStr:   make(chan string, countG),
				chMaskedStr: make(chan string, countG),
			},
			want: "[]]!!#$@Fr",
		},
		{
			name: "empty",
			args: args{
				sourceStr:   "",
				chMaskStr:   make(chan string, countG),
				chMaskedStr: make(chan string, countG),
			},
			want: "",
		},
		{
			name: "incomplete link",
			args: args{
				sourceStr:   "ttp://tqwe",
				chMaskStr:   make(chan string, countG),
				chMaskedStr: make(chan string, countG),
			},
			want: "ttp://tqwe",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.chMaskStr <- tt.args.sourceStr
			s.findAndMaskLinks(tt.args.chMaskStr, tt.args.chMaskedStr)
			result := <-tt.args.chMaskedStr
			assert.Equalf(t, tt.want, result, "findAndMaskLinks(%v)", tt.args.sourceStr)
		})
	}
}
