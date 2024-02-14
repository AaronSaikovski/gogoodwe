package interfaces

import (
	"github.com/AaronSaikovski/gogoodwe/cmd/gogoodwe/types"
)

// Constraints for functions that return data from the API via marshalled structs
type ISemsDataConstraint interface {
	types.InverterData | types.DailySummaryData
}
