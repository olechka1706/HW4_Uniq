package uniq

import (
	"bytes"
	"testing"
)

var testOk = []struct {
	in  string
	opt Options
	out string
}{
	{
		in: "one\none\nOne\ntwo\n\n\n\n",
		opt: Options{
			Count:      false,
			Repeated:   false,
			Unique:     false,
			SkipFields: 0,
			SkipChars:  0,
			IgnoreCase: false,
		},
		out: "one\nOne\ntwo\n\n",
	},
	{
		in:  "",
		opt: Options{},
		out: "",
	},
	{
		in: "a\na\nb\nb\nc\n",
		opt: Options{
			Count: true,
		},
		out: "2 a\n2 b\n1 c\n",
	},
	{
		in: "a\na\nb\nc\nc\nd\n",
		opt: Options{
			Repeated: true,
		},
		out: "a\nc\n",
	},
	{
		in: "A\na\nB\nb\nA\n",
		opt: Options{
			IgnoreCase: true,
			Count:      true,
		},
		out: "2 A\n2 B\n1 A\n",
	},
	{
		in: "A\na\nB\nb\nC\n",
		opt: Options{
			Repeated:   true,
			IgnoreCase: true,
		},
		out: "A\nB\n",
	},
	{
		in: "I love apple\nWe love apple\nI love melon\nb\nb\n",
		opt: Options{
			SkipFields: 1,
			Repeated:   true,
		},
		out: "I love apple\nb\n",
	},
	{
		in: "I love apple\nWe love apple\nI love melon\nb\nb\n",
		opt: Options{
			SkipChars: 100,
			Repeated:  true,
		},
		out: "I love apple\n",
	},
}

func TestOK(t *testing.T) {
	for _, test := range testOk {
		t.Run(test.in, func(t *testing.T) {
			input := bytes.NewBufferString(test.in)
			output := bytes.NewBuffer(nil)
			Uniq(input, output, test.opt)
			result := output.String()
			if result != test.out {
				t.Errorf("test OK failed, got %q, want %q", result, test.out)
			}
		})
	}
}


