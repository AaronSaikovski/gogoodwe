package interfaces

import (
	"github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata/currentkpidata"
	inverteallpoint "github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata/inverterallpoint"
	"github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata/monitordetail"
	"github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata/monitorsummary"
	"github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata/plantdetail"
	plantchartdata "github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata/plantpowerchart"
	"github.com/AaronSaikovski/gogoodwe/internal/features/fetchdata/powerflow"
)

// SEMS Data Constraints
type SemsDataConstraint interface {
	inverteallpoint.InverterAllPoint | monitordetail.MonitorData | monitorsummary.DailySummaryData | plantdetail.PlantDetailByPowerstationId | plantchartdata.PlantPowerChart | powerflow.Powerflow | currentkpidata.KPIMonitorData
}
