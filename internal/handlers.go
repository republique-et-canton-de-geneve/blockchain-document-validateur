package internal

import (
	"context"
	op "github.com/geneva_validateur/restapi/operations"
	models "github.com/geneva_validateur/models"

	middleware "github.com/go-openapi/runtime/middleware"

)

func MonitoringHandler(ctx context.Context, params op.MonitoringParams) middleware.Responder {
	nodeOk := GetNodeSignal(ctx)

	var sondeResp []*models.Sonde
	sondeResp_rcpt := models.Sonde{EthereumActive: nodeOk}
	sondeResp = append(sondeResp, &sondeResp_rcpt)
	return op.NewMonitoringOK().WithPayload(sondeResp)
}
