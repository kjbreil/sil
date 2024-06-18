package silread

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/kjbreil/sil/tables"
)

// func TestUnmarshal(t *testing.T) {
// 	var obj tables.OBJ
//
// 	b, _ := ioutil.ReadFile("./examples/single.sil")
//
// 	s, err := Unmarshal(b, &obj)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	_, err = s.Marshal(true)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// }
//
// func TestUnmarshalHeaders(t *testing.T) {
// 	var cll tables.CLL
//
// 	b, _ := ioutil.ReadFile("./examples/with_header.sil")
//
// 	s, err := Unmarshal(b, &cll)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	_, err = s.Marshal(true)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

func TestUnmarshalLoad(t *testing.T) {
	var dss []tables.DSS

	b, _ := ioutil.ReadFile("./examples/dss.sql")

	err := Unmarshal(b, &dss)
	if err != nil {
		t.Fatal(err)
	}
	if len(dss) != 4 {
		t.Fatalf("expected dss length 4, got %d", len(dss))
	}
}

func TestUnmarshalLocLoad(t *testing.T) {
	var dss []*tables.DSS

	b, _ := os.ReadFile("./examples/Loc_Load.sql")

	err := Unmarshal(b, &dss)
	if err != nil {
		t.Fatal(err)
	}
}

type testS struct {
	PtrString *string `sil:"F01,nullable"`
	String    string  `sil:"F01,nullable"`
}

func Test_unmarshalValue(t *testing.T) {
	fieldMap := []int{0, 1}

	type args struct {
		input    []string
		result   testS
		fieldMap []int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "nil pointer data in non-pointer",
			args: args{
				input: []string{"", "TEST"},
				result: testS{
					PtrString: nil,
					String:    "TEST",
				},
				fieldMap: fieldMap,
			},
		},
		{
			name: "data in pointer blank non-pointer",
			args: args{
				input: []string{"TEST", ""},
				result: testS{
					PtrString: ptr("TEST"),
					String:    "",
				},
				fieldMap: fieldMap,
			},
		},
		{
			name: "double single quote in pointer and non pointer",
			args: args{
				input: []string{"''", "''"},
				result: testS{
					PtrString: ptr("''"),
					String:    "''",
				},
				fieldMap: fieldMap,
			},
		},
	}
	for _, tt := range tests {
		testV := testS{}
		v := reflect.ValueOf(&testV).Elem()
		t.Run(tt.name, func(t *testing.T) {
			if err := unmarshalValue(tt.args.input, v, tt.args.fieldMap); (err != nil) != tt.wantErr {
				t.Errorf("unmarshalValue() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(testV, tt.args.result) {
				t.Errorf("unmarshalValue() = %v, want %v", testV, tt.args.result)
			}
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}
