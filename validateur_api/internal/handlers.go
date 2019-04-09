package internal

import (
	"context"
	op "github.com/geneva_validateur/restapi/operations"
	models "github.com/geneva_validateur/models"

	middleware "github.com/go-openapi/runtime/middleware"

)

func MonitoringHandler(ctx context.Context, params op.MonitoringParams) middleware.Responder {
	// Check up Node signal and return false if ctx cannot connect
	nodeOk := GetNodeSignal(ctx)

	var sonde []*models.Sonde

	sonde_res := models.Sonde{EthereumActive: nodeOk}
	sonde = append(sonde, &sonde_res)

	return op.NewMonitoringOK().WithPayload(sonde)
}
