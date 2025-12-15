package app

// Main package - This is the main program entry point
import (
	"github.com/AaronSaikovski/gogoodwe/pkg/interfaces"
	"github.com/AaronSaikovski/gogoodwe/pkg/models/currentkpidata"
	inverteallpoint "github.com/AaronSaikovski/gogoodwe/pkg/models/inverterallpoint"
	"github.com/AaronSaikovski/gogoodwe/pkg/models/monitordetail"
	"github.com/AaronSaikovski/gogoodwe/pkg/models/monitorsummary"
	"github.com/AaronSaikovski/gogoodwe/pkg/models/plantdetail"
	plantchartdata "github.com/AaronSaikovski/gogoodwe/pkg/models/plantpowerchart"
	"github.com/AaronSaikovski/gogoodwe/pkg/models/powerflow"
)

// lookupMonitorData returns a PowerData object based on the given reportData int.
//
// Parameters:
// - reportData: a string representing the type of data to retrieve.
//
// Returns:
// - interfaces.PowerData: the PowerData object corresponding to the reportData.
func lookupMonitorData(reportData int) interfaces.PowerData {

	switch reportData {

	case Point:
		return inverteallpoint.NewInverterAllPoint()
	case Detail:
		return monitordetail.NewMonitorData()
	case Summary:
		return monitorsummary.NewDailySummaryData()
	case Plant:
		return plantdetail.NewGetPlantDetailByPowerstationId()
	case PlantChart:
		return plantchartdata.NewPlantPowerChart()
	case PowerFlow:
		return powerflow.NewPowerflow()
	case KPIData:
		return currentkpidata.NewKPIMonitorData()
	default:
		return monitordetail.NewMonitorData()
	}
}
