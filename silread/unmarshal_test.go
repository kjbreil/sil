package silread

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"

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
	PtrString *string    `sil:"F01,nullable"`
	String    string     `sil:"F01,nullable"`
	PtrTime   *time.Time `sil:"F02,nullable"`
	Time      time.Time  `sil:"F02,nullable"`
}

func Test_unmarshalValues(t *testing.T) {
	fieldMap := []int{0, 1, 2, 3}

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
		{
			name: "DateTime in pointer and non pointer",
			args: args{
				input: []string{"", "", "2020072 12:50:29", "2020072 12:50:29"},
				result: testS{
					PtrTime: ptr(time.Date(2020, 3, 12, 12, 50, 29, 0, time.UTC)),
					Time:    time.Date(2020, 3, 12, 12, 50, 29, 0, time.UTC),
				},
				fieldMap: fieldMap,
			},
		},
		{
			name: "just julian date without time",
			args: args{
				input: []string{"", "", "2022208", "2022208"},
				result: testS{
					PtrTime: ptr(time.Date(2022, 7, 27, 0, 0, 0, 0, time.UTC)),
					Time:    time.Date(2022, 7, 27, 0, 0, 0, 0, time.UTC),
				},
				fieldMap: fieldMap,
			},
		},
	}
	for _, tt := range tests {
		testV := testS{}
		v := reflect.ValueOf(&testV).Elem()
		t.Run(tt.name, func(t *testing.T) {
			if err := unmarshalValues(tt.args.input, v, tt.args.fieldMap); (err != nil) != tt.wantErr {
				t.Errorf("unmarshalValues() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(testV, tt.args.result) {
				t.Errorf("unmarshalValues() = \n%v\n want \n%v", tt.args.result, testV)
			}
		})
	}
}

func Test_unmarshalValue(t *testing.T) {
	type args struct {
		t     any
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "string",
			args: args{
				t:     "",
				input: "TEST",
			},
		},
		{
			name: "nil string",
			args: args{
				t:     nilPtr[string](),
				input: "TEST",
			},
		},
		{
			name: "int",
			args: args{
				t:     int(0),
				input: "123",
			},
			want: 123,
		},
		{
			name: "default int",
			args: args{
				t:     ptr(0),
				input: "0",
			},
			want: 0,
		},
		{
			name: "unhandled type",
			args: args{
				t:     complex(3, -5),
				input: "123",
			},
			wantErr: true,
		},
		{
			name: "ptr string",
			args: args{
				t:     ptr(""),
				input: "TEST",
			},
		},
		{
			name: "date",
			args: args{
				t:     time.Time{},
				input: "2022102",
			},
			want: time.Date(2022, 4, 12, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "datetime",
			args: args{
				t:     time.Time{},
				input: "2022102 18:19:20",
			},
			want: time.Date(2022, 4, 12, 18, 19, 20, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := reflect.ValueOf(&tt.args.t).Elem()
			var ty reflect.Type
			if v.Elem().Kind() == reflect.Ptr {
				ty = v.Elem().Type().Elem()
			} else {
				ty = v.Elem().Type()
			}
			newValue := reflect.New(ty).Elem()
			err := unmarshalValue(newValue, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshalValue() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err != nil {
				return
			}
			v.Set(newValue)
			if tt.want != nil {
				if !reflect.DeepEqual(tt.args.t, tt.want) {
					t.Errorf("unmarshalValue() = %v, want %v", tt.args.t, tt.want)
				}
			} else if !reflect.DeepEqual(tt.args.t, tt.args.input) {
				t.Errorf("unmarshalValue() = %v, want %v", tt.args.t, tt.args.input)
			}
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}

func nilPtr[T any]() *T {
	var v *T
	return v
}
