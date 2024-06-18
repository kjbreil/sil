package silread

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func textToScanner(text string) *scanner {
	return newScanner(bufio.NewReader(strings.NewReader(text)))
}

func Test_scanner_scan(t *testing.T) {
	type fields struct {
		r *bufio.Reader
	}
	tests := []struct {
		name string
		text string
		want *part
	}{
		{
			name: "newline detected as whitespace",
			text: "\n",
			want: &part{tok: CRLF, val: "\r\n"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := textToScanner(tt.text)
			if got := s.scan(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scan() = %v, want %v", got, tt.want)
			}
		})
	}
}
