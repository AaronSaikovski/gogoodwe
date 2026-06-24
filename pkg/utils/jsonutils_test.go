package utils_test

import (
	"encoding/json"
	"testing"

	"github.com/AaronSaikovski/gogoodwe/pkg/utils"
)

func TestMarshalStructToJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		wantErr bool
	}{
		{name: "simple struct", input: struct {
			Name string
			Age  int
		}{"Alice", 30}, wantErr: false},
		{name: "map", input: map[string]int{"a": 1, "b": 2}, wantErr: false},
		{name: "slice", input: []int{1, 2, 3}, wantErr: false},
		{name: "empty struct", input: struct{}{}, wantErr: false},
		{name: "nil interface", input: nil, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := utils.MarshalStructToJSON(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalStructToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				var result any
				if err := json.Unmarshal(got, &result); err != nil {
					t.Errorf("MarshalStructToJSON() returned invalid JSON: %v", err)
				}
			}
		})
	}
}

func TestUnmarshalDataToStruct(t *testing.T) {
	tests := []struct {
		name      string
		jsonData  []byte
		target    any
		wantErr   bool
		checkFunc func(t *testing.T, target any)
	}{
		{
			name:     "valid struct unmarshal",
			jsonData: []byte(`{"Name":"Bob","Age":25}`),
			target: &struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
			}{},
			checkFunc: func(t *testing.T, target any) {
				s := target.(*struct {
					Name string `json:"name"`
					Age  int    `json:"age"`
				})
				if s.Name != "Bob" || s.Age != 25 {
					t.Errorf("got %+v, want {Name:Bob Age:25}", s)
				}
			},
		},
		{
			name:     "valid map unmarshal",
			jsonData: []byte(`{"key":"value","num":42}`),
			target:   &map[string]any{},
			checkFunc: func(t *testing.T, target any) {
				m := *target.(*map[string]any)
				if m["key"] != "value" || m["num"].(float64) != 42 {
					t.Errorf("got %v, want {key:value num:42}", m)
				}
			},
		},
		{
			name:     "invalid JSON",
			jsonData: []byte(`{invalid}`),
			target:   &map[string]any{},
			wantErr:  true,
		},
		{
			name:     "empty JSON",
			jsonData: []byte(``),
			target:   &map[string]any{},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := utils.UnmarshalDataToStruct(tt.jsonData, tt.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalDataToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.checkFunc != nil {
				tt.checkFunc(t, tt.target)
			}
		})
	}
}
