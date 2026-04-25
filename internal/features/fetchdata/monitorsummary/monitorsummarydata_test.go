package monitorsummary

import "testing"

func TestNewDailySummaryData(t *testing.T) {
	data := NewDailySummaryData()
	if data == nil {
		t.Fatal("NewDailySummaryData() returned nil")
	}
}
