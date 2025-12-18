package fetchdata

import (
	"fmt"

	"github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata/common"
)

// ParseReportType converts a string report type to its integer constant.
func ParseReportType(reportTypeStr string) (int, error) {
	switch reportTypeStr {
	case "detail", "0":
		return common.Detail, nil
	case "summary", "1":
		return common.Summary, nil
	case "point", "2":
		return common.Point, nil
	case "plant", "3":
		return common.Plant, nil
	case "plantchart", "4":
		return common.PlantChart, nil
	case "powerflow", "5":
		return common.PowerFlow, nil
	case "kpidata", "6":
		return common.KPIData, nil
	default:
		return -1, fmt.Errorf("invalid report type '%s'. Valid options are: detail, summary, point, plant, plantchart, powerflow, kpidata", reportTypeStr)
	}
}
