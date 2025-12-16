package app

import (
	"fmt"
)

// ParseReportType converts a string report type to its integer constant.
func ParseReportType(reportTypeStr string) (int, error) {
	switch reportTypeStr {
	case "detail", "0":
		return Detail, nil
	case "summary", "1":
		return Summary, nil
	case "point", "2":
		return Point, nil
	case "plant", "3":
		return Plant, nil
	case "plantchart", "4":
		return PlantChart, nil
	case "powerflow", "5":
		return PowerFlow, nil
	case "kpidata", "6":
		return KPIData, nil
	default:
		return -1, fmt.Errorf("invalid report type '%s'. Valid options are: detail, summary, point, plant, plantchart, powerflow, kpidata", reportTypeStr)
	}
}
