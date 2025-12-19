package fetchdata

// import (
// 	"context"

// 	"github.com/AaronSaikovski/gogoodwe/internal/shared/apihelpers"
// 	"github.com/AaronSaikovski/gogoodwe/internal/shared/auth"
// )

// func callMonitorApi[T any](
// 	ctx context.Context,
// 	authLoginInfo *auth.LoginInfo,
// 	powerStationURL string,
// 	HTTPTimeout int,
// ) (*T, []byte, error) {

// 	var output T

// 	raw, err := apihelpers.FetchMonitorAPIData(
// 		ctx,
// 		authLoginInfo,
// 		powerStationURL,
// 		HTTPTimeout,
// 		&output,
// 	)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	return &output, raw, nil
// }

// func GetPowerData[T any](
// 	ctx context.Context,
// 	authLoginInfo *auth.LoginInfo,
// ) (*T, error) {

// 	data, _, err := callMonitorApi[T](ctx, authLoginInfo)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := utils.ProcessData(data); err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }
