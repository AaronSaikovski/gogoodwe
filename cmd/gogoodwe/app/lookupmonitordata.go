/*
MIT License

# Copyright (c) 2024 Aaron Saikovski

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package app

// Main package - This is the main program entry point
import (
	"github.com/AaronSaikovski/gogoodwe/pkg/interfaces"
	inverteallpoint "github.com/AaronSaikovski/gogoodwe/pkg/inverterallpoint"
	"github.com/AaronSaikovski/gogoodwe/pkg/monitordetail"
	"github.com/AaronSaikovski/gogoodwe/pkg/monitorsummary"
	"github.com/AaronSaikovski/gogoodwe/pkg/plantdetail"
	plantchartdata "github.com/AaronSaikovski/gogoodwe/pkg/plantpowerchart"
	"github.com/AaronSaikovski/gogoodwe/pkg/powerflow"
)

// lookupMonitorData returns a PowerData object based on the given reportData string.
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
	default:
		return monitordetail.NewMonitorData()
	}
}
