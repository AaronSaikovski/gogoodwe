package plantchartdata

import "testing"

func TestNewPlantPowerChart(t *testing.T) {
	data := NewPlantPowerChart()
	if data == nil {
		t.Fatal("NewPlantPowerChart() returned nil")
	}
}
