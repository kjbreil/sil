package silread

import "testing"

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
