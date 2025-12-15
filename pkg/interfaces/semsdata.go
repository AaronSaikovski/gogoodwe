package interfaces

import (
	"github.com/AaronSaikovski/gogoodwe/pkg/models/currentkpidata"
	inverteallpoint "github.com/AaronSaikovski/gogoodwe/pkg/models/inverterallpoint"
	"github.com/AaronSaikovski/gogoodwe/pkg/models/monitordetail"
	"github.com/AaronSaikovski/gogoodwe/pkg/models/monitorsummary"
	"github.com/AaronSaikovski/gogoodwe/pkg/models/plantdetail"
	plantchartdata "github.com/AaronSaikovski/gogoodwe/pkg/models/plantpowerchart"
	"github.com/AaronSaikovski/gogoodwe/pkg/models/powerflow"
)

// SEMS Data Constraints
type SemsDataConstraint interface {
	inverteallpoint.InverterAllPoint | monitordetail.MonitorData | monitorsummary.DailySummaryData | plantdetail.PlantDetailByPowerstationId | plantchartdata.PlantPowerChart | powerflow.Powerflow | currentkpidata.KPIMonitorData
}
