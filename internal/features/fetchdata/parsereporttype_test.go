package fetchdata

import (
	"testing"

	"github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata/common"
)

func TestParseReportType(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{"detail string", "detail", common.Detail, false},
		{"summary string", "summary", common.Summary, false},
		{"point string", "point", common.Point, false},
		{"plant string", "plant", common.Plant, false},
		{"plantchart string", "plantchart", common.PlantChart, false},
		{"powerflow string", "powerflow", common.PowerFlow, false},
		{"kpidata string", "kpidata", common.KPIData, false},
		{"detail numeric", "0", common.Detail, false},
		{"summary numeric", "1", common.Summary, false},
		{"point numeric", "2", common.Point, false},
		{"plant numeric", "3", common.Plant, false},
		{"plantchart numeric", "4", common.PlantChart, false},
		{"powerflow numeric", "5", common.PowerFlow, false},
		{"kpidata numeric", "6", common.KPIData, false},
		{"invalid string", "invalid", -1, true},
		{"empty string", "", -1, true},
		{"out of range number", "99", -1, true},
		{"negative number", "-1", -1, true},
		{"uppercase", "Detail", -1, true},
		{"with spaces", " detail", -1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseReportType(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseReportType(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseReportType(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
