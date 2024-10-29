package pub

import (
	"reflect"
	"testing"
)

// go test -v -run TestIsEmpty
func TestIsEmpty(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args *args
		want bool
	}{
		{
			"nil",
			&args{nil},
			true,
		},
		{
			"int",
			&args{0},
			true,
		},
		{
			"float",
			&args{0.0},
			true,
		},
		{
			"bool",
			&args{false},
			true,
		},
		{
			"string",
			&args{""},
			true,
		},
		{
			"slice",
			&args{[]int{}},
			true,
		},
		{
			"map",
			&args{map[string]interface{}{}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.obj); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -v -run TestStructToMap
func TestStructToMap(t *testing.T) {
	type args struct {
		Name      string
		Age       int
		Man       bool
		Timestamp int64
	}
	tests := []struct {
		name string
		args *args
		want map[string]interface{}
	}{
		{
			"normal",
			&args{"tom", 18, true, 1600000000},
			map[string]interface{}{"Name": "tom", "Age": 18, "Man": true, "Timestamp": int64(1600000000)},
		},
		{
			"empty",
			&args{"", 0, false, 0},
			map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructToMap(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructToMap() = %+v, want %+v", got, tt.want)
				for k, v := range got {
					if v != tt.want[k] {
						t.Errorf("k = %+v, got=%+v != want %+v", k, v, tt.want[k])
					}
				}
			}
		})
	}
}
