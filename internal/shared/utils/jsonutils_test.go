package utils

import (
	"testing"
)

func TestMarshalStructToJSON(t *testing.T) {
	type sample struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	tests := []struct {
		name    string
		input   interface{}
		want    string
		wantErr bool
	}{
		{"simple struct", sample{Name: "test", Value: 42}, `{"name":"test","value":42}`, false},
		{"empty struct", sample{}, `{"name":"","value":0}`, false},
		{"nil input", nil, "null", false},
		{"string input", "hello", `"hello"`, false},
		{"unmarshalable input", make(chan int), "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MarshalStructToJSON(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalStructToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && string(got) != tt.want {
				t.Errorf("MarshalStructToJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestUnmarshalDataToStruct(t *testing.T) {
	type sample struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	t.Run("valid JSON", func(t *testing.T) {
		var result sample
		err := UnmarshalDataToStruct([]byte(`{"name":"test","value":42}`), &result)
		if err != nil {
			t.Fatalf("UnmarshalDataToStruct() unexpected error: %v", err)
		}
		if result.Name != "test" || result.Value != 42 {
			t.Errorf("UnmarshalDataToStruct() = %+v, want {Name:test Value:42}", result)
		}
	})

	t.Run("invalid JSON", func(t *testing.T) {
		var result sample
		err := UnmarshalDataToStruct([]byte(`not json`), &result)
		if err == nil {
			t.Error("UnmarshalDataToStruct() expected error for invalid JSON")
		}
	})

	t.Run("empty JSON object", func(t *testing.T) {
		var result sample
		err := UnmarshalDataToStruct([]byte(`{}`), &result)
		if err != nil {
			t.Fatalf("UnmarshalDataToStruct() unexpected error: %v", err)
		}
		if result.Name != "" || result.Value != 0 {
			t.Errorf("UnmarshalDataToStruct() = %+v, want zero value", result)
		}
	})
}
