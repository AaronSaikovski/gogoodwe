package fetchdata

// Main package - This is the main program entry point
import (
	"github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata/common"
	"github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata/currentkpidata"
	"github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata/interfaces"
	inverteallpoint "github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata/inverterallpoint"
	"github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata/monitordetail"
	"github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata/monitorsummary"
	"github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata/plantdetail"
	plantchartdata "github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata/plantpowerchart"
	"github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata/powerflow"
)

// lookupMonitorData returns a PowerData object based on the given reportData int.
//
// Parameters:
// - reportData: a string representing the type of data to retrieve.
//
// Returns:
// - interfaces.PowerData: the PowerData object corresponding to the reportData.
func LookupMonitorData(reportData int) interfaces.PowerData {

	switch reportData {
	case common.Point:
		return inverteallpoint.NewInverterAllPoint()
	case common.Detail:
		return monitordetail.NewMonitorData()
	case common.Summary:
		return monitorsummary.NewDailySummaryData()
	case common.Plant:
		return plantdetail.NewGetPlantDetailByPowerstationId()
	case common.PlantChart:
		return plantchartdata.NewPlantPowerChart()
	case common.PowerFlow:
		return powerflow.NewPowerflow()
	case common.KPIData:
		return currentkpidata.NewKPIMonitorData()
	default:
		return monitordetail.NewMonitorData()
	}
}

/// InverterAllPoint represents the data structure for inverter all point data.
func LookupReportStruct(reportData int) interfaces.PowerData {
	switch reportData {
	case common.Point:
		return &inverteallpoint.InverterAllPoint{}
	case common.Detail:
		return &monitordetail.MonitorData{}
	case common.Summary:
		return &monitorsummary.DailySummaryData{}
	case common.Plant:
		return &plantdetail.PlantDetailByPowerstationId{}
	case common.PlantChart:
		return &plantchartdata.PlantPowerChart{}
	case common.PowerFlow:
		return &powerflow.Powerflow{}
	case common.KPIData:
		return &currentkpidata.KPIMonitorData{}
	default:
		return &monitordetail.MonitorData{}
	}
}
