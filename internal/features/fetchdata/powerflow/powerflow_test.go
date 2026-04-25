package powerflow

import "testing"

func TestNewPowerflow(t *testing.T) {
	data := NewPowerflow()
	if data == nil {
		t.Fatal("NewPowerflow() returned nil")
	}
}
