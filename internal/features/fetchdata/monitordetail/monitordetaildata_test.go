package monitordetail

import "testing"

func TestNewMonitorData(t *testing.T) {
	data := NewMonitorData()
	if data == nil {
		t.Fatal("NewMonitorData() returned nil")
	}
}
