package fetchdata

import (
	"fmt"
	"testing"

	"github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata/common"
)

func TestLookupMonitorData(t *testing.T) {
	tests := []struct {
		name       string
		reportType int
		wantNil    bool
	}{
		{"detail", common.Detail, false},
		{"summary", common.Summary, false},
		{"point", common.Point, false},
		{"plant", common.Plant, false},
		{"plantchart", common.PlantChart, false},
		{"powerflow", common.PowerFlow, false},
		{"kpidata", common.KPIData, false},
		{"default/invalid returns monitordetail", 99, false},
		{"negative returns monitordetail", -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LookupMonitorData(tt.reportType)
			if (result == nil) != tt.wantNil {
				t.Errorf("LookupMonitorData(%d) = %v, wantNil %v", tt.reportType, result, tt.wantNil)
			}
		})
	}
}

func TestLookupMonitorDataReturnsDifferentTypes(t *testing.T) {
	// Verify different report types return different concrete types
	detail := LookupMonitorData(common.Detail)
	summary := LookupMonitorData(common.Summary)
	point := LookupMonitorData(common.Point)

	detailType := typeString(detail)
	summaryType := typeString(summary)
	pointType := typeString(point)

	if detailType == summaryType {
		t.Errorf("Detail and Summary should return different types, both got %s", detailType)
	}
	if detailType == pointType {
		t.Errorf("Detail and Point should return different types, both got %s", detailType)
	}
}

func typeString(v any) string {
	return fmt.Sprintf("%T", v)
}
