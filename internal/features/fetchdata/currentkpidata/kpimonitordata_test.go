package currentkpidata

import "testing"

func TestNewKPIMonitorData(t *testing.T) {
	data := NewKPIMonitorData()
	if data == nil {
		t.Fatal("NewKPIMonitorData() returned nil")
	}
}
