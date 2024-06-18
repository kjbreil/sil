package silread

import (
	"reflect"
	"strings"
	"testing"
)

func Test_parseTable(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name       string
		args       args
		wantName   string
		wantAction string
	}{
		{
			name:       "SAL_HDR",
			args:       args{text: "SAL_HDR_LOAD"},
			wantName:   "SAL_HDR",
			wantAction: "LOAD",
		},
		{
			name:       "SCL_NUT_LOAD",
			args:       args{text: "SCL_NUT_LOAD"},
			wantName:   "SCL_NUT",
			wantAction: "LOAD",
		},
		{
			name:       "PRICE_LOAD",
			args:       args{text: "PRICE_LOAD"},
			wantName:   "PRICE",
			wantAction: "LOAD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, gotAction := parseTable(tt.args.text)
			if gotName != tt.wantName {
				t.Errorf("parseTable() gotName = %v, want %v", gotName, tt.wantName)
			}
			if gotAction != tt.wantAction {
				t.Errorf("parseTable() gotAction = %v, want %v", gotAction, tt.wantAction)
			}
		})
	}
}

func textToParsed(text string) (parsed, int, int) {
	p := newParser(strings.NewReader(text))
	columns := strings.Count(text, ",") + 1
	if strings.Contains(text, "),") {
		columns--
	}
	return p.parse(), 0, columns
}

func Test_readDataLine(t *testing.T) {
	// text := "(297,'PAL',118,'2021246 14:39:27','FOOD SAFETY: Keep refrigerated at or below 41F.','08670056',0,'','','',,'JACOBSONS MARKET GARLIC STICK PER LB','','','','','','','','999','','','','','',),"

	// text := "('PUB-3','PAL',,,1,'',,10,'pub_3.htm','PUB1',,,);\n"

	//
	tests := []struct {
		name    string
		line    string
		want    rowData
		wantErr bool
	}{
		{
			name:    "Empty Line",
			line:    "",
			want:    []string{},
			wantErr: true,
		},
		{
			name:    "Single Quote",
			line:    "('PUB-3','PAL',,,1,'',,10,'pub_3.htm','PUB1',,,);\n",
			want:    []string{"PUB-3", "PAL", "", "", "1", "''", "", "10", "pub_3.htm", "PUB1", "", "", ""},
			wantErr: false,
		},
		{
			name: "Double Space in Text",
			line: "(297,'PAL',118,'2021246 14:39:27','FOOD SAFETY: Keep refrigerated at or below 41F.  ','08670056',0,'','','',,'JACOBSONS MARKET GARLIC STICK PER LB','','','','','','','','999','','','','','',),\n",
			want: []string{"297", "PAL", "118", "2021246 14:39:27", "FOOD SAFETY: Keep refrigerated at or below 41F.  ", "08670056", "0", "''", "''", "''", "", "JACOBSONS MARKET GARLIC STICK PER LB", "''", "''", "''", "''", "''", "''", "''", "999", "''", "''", "''", "''", "''", ""},
		},
		{
			name: "Two single quotes in Text",
			line: "('POPN''CTN CDY ','PAL','9112','DIRECT',0,,,,,,,,1,,,1,1,'2020072 12:50:29',1,,,,'999',),\n",
			want: []string{"POPN''CTN CDY ", "PAL", "9112", "DIRECT", "0", "", "", "", "", "", "", "", "1", "", "", "1", "1", "2020072 12:50:29", "1", "", "", "", "999", ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, s, c := textToParsed(tt.line)
			got, got1, err := readDataLine(p, s, c)
			if (err != nil) != tt.wantErr {
				t.Errorf("readDataLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr {
				return
			}
			if got1 != len(p)-1 {
				t.Errorf("readDataLine() ending s = %d, want %d", got1, len(p)-1)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readDataLine() got = \n%v\n want \n%v\n", got.want(), tt.want.want())
			}

		})
	}
}
