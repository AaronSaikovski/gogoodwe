package plantdetail

import "testing"

func TestNewGetPlantDetailByPowerstationId(t *testing.T) {
	data := NewGetPlantDetailByPowerstationId()
	if data == nil {
		t.Fatal("NewGetPlantDetailByPowerstationId() returned nil")
	}
}
