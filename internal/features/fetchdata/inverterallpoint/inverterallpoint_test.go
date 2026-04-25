package inverteallpoint

import "testing"

func TestNewInverterAllPoint(t *testing.T) {
	data := NewInverterAllPoint()
	if data == nil {
		t.Fatal("NewInverterAllPoint() returned nil")
	}
}
