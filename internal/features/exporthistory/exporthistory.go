package exporthistory

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/AaronSaikovski/gogoodwe/internal/shared/apihelpers"
	"github.com/AaronSaikovski/gogoodwe/internal/shared/auth"
	"github.com/AaronSaikovski/gogoodwe/internal/shared/utils"
)

const (
	dateTimeFormat = "2006-01-02 15:04"
)

// defaultTargets returns the standard set of inverter metrics to query
func defaultTargets() []Target {
	return []Target{
		{TargetKey: "Vpv1", TargetIndex: 1},
		{TargetKey: "Vpv2", TargetIndex: 2},
		{TargetKey: "Ipv1", TargetIndex: 5},
		{TargetKey: "Ipv2", TargetIndex: 6},
		{TargetKey: "Vac1", TargetIndex: 9},
		{TargetKey: "Iac1", TargetIndex: 12},
		{TargetKey: "Fac1", TargetIndex: 15},
		{TargetKey: "Pac", TargetIndex: 18},
		{TargetKey: "WorkMode", TargetIndex: 19},
		{TargetKey: "Tempperature", TargetIndex: 20},
		{TargetKey: "ETotal", TargetIndex: 22},
		{TargetKey: "HTotal", TargetIndex: 23},
		{TargetKey: "Reserved5", TargetIndex: 36},
		{TargetKey: "PF", TargetIndex: 407},
	}
}

// targetKeyToIndex maps target key names to their API index values
var targetKeyToIndex = map[string]int{
	"Vpv1":         1,
	"Vpv2":         2,
	"Ipv1":         5,
	"Ipv2":         6,
	"Vac1":         9,
	"Iac1":         12,
	"Fac1":         15,
	"Pac":          18,
	"WorkMode":     19,
	"Tempperature": 20,
	"ETotal":       22,
	"HTotal":       23,
	"Reserved5":    36,
	"PF":           407,
}

// ParseTargets parses a comma-separated string of target keys into Target structs.
// If the input is empty, returns the default targets.
func ParseTargets(targetsStr string) ([]Target, error) {
	if targetsStr == "" {
		return defaultTargets(), nil
	}

	keys := strings.Split(targetsStr, ",")
	targets := make([]Target, 0, len(keys))
	for _, key := range keys {
		key = strings.TrimSpace(key)
		idx, ok := targetKeyToIndex[key]
		if !ok {
			return nil, fmt.Errorf("unknown target key '%s'. Valid keys: Vpv1, Vpv2, Ipv1, Ipv2, Vac1, Iac1, Fac1, Pac, WorkMode, Tempperature, ETotal, HTotal, Reserved5, PF", key)
		}
		targets = append(targets, Target{TargetKey: key, TargetIndex: idx})
	}
	return targets, nil
}

// BuildRequest constructs the HistoryDataRequest from parameters
func BuildRequest(loginInfo *auth.LoginInfo, qryTimeStart, qryTimeEnd string, targets []Target) (*HistoryDataRequest, error) {
	startTime, err := time.Parse(dateTimeFormat, qryTimeStart)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start time: %w", err)
	}
	endTime, err := time.Parse(dateTimeFormat, qryTimeEnd)
	if err != nil {
		return nil, fmt.Errorf("failed to parse end time: %w", err)
	}

	days := int(endTime.Sub(startTime).Hours()/24) + 1

	return &HistoryDataRequest{
		DataType:     0,
		TimesType:    1,
		QryTimeStart: qryTimeStart,
		QryTimeEnd:   qryTimeEnd,
		Times:        days,
		QryStatus:    0,
		PwsHistorys: []PwsHistory{
			{
				ID:        loginInfo.SemsLoginCredentials.PowerStationID,
				PwName:    loginInfo.SemsLoginCredentials.Account,
				Status:    1,
				PwAddress: "",
				Inverters: []Inverter{
					{
						SN:         loginInfo.SemsLoginCredentials.PowerStationID,
						Name:       loginInfo.SemsLoginCredentials.Account,
						ChangeNum:  0,
						ChangeType: 0,
						Status:     1,
					},
				},
			},
		},
		Targets: targets,
	}, nil
}

// FetchExportHistory calls the ExportExcelStationHistoryData API and prints the response.
func FetchExportHistory(ctx context.Context, loginInfo *auth.LoginInfo, qryTimeStart, qryTimeEnd string, targets []Target) error {
	request, err := BuildRequest(loginInfo, qryTimeStart, qryTimeEnd, targets)
	if err != nil {
		return err
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal history data request: %w", err)
	}

	var response HistoryDataResponse
	rawJSON, err := apihelpers.FetchHistoryExportData(ctx, loginInfo, requestBody, &response)
	if err != nil {
		return err
	}

	return utils.ProcessRawJSON(rawJSON)
}
