package internal

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"net/textproto"

	models "github.com/Magicking/rc-ge-ch-pdf/models"
	op "github.com/Magicking/rc-ge-ch-pdf/restapi/operations"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
)

type nopCloser struct {
	*bytes.Reader
}

func (nopCloser) Close() error { return nil }

func GetreceiptHandler(ctx context.Context, params op.GetreceiptParams) middleware.Responder {
	rcpt, ok, err := GetReceiptByHash(ctx, params.Hash)
	if err != nil {
		err_str := fmt.Sprintf("Failed to call %s: %v", "GetReceiptByHash", err)
		log.Printf(err_str)
		return op.NewGetreceiptDefault(500).WithPayload(&models.Error{Message: &err_str})
	}
	if !ok {
		err_str := fmt.Sprintf("No receipt found for %s", params.Hash)
		log.Printf(err_str)
		return op.NewGetreceiptDefault(500).WithPayload(&models.Error{Message: &err_str})
	}
	mimeHdr := textproto.MIMEHeader{"Content-Type": {"application/json"}}
	reader := bytes.NewReader(rcpt.JSONData)
	buf := nopCloser{reader}
	filename := fmt.Sprintf("%s.json", rcpt.Filename)
	fh := multipart.FileHeader{Filename: filename, Header: mimeHdr}
	return op.NewGetreceiptOK().WithPayload(runtime.File{Data: buf, Header: &fh})
}

func ListtimestampedHandler(ctx context.Context, params op.ListtimestampedParams) middleware.Responder {
	rcpts, err := GetAllReceipts(ctx)
	if err != nil {
		err_str := fmt.Sprintf("Failed to call %s: %v", "GetAllReceipts", err)
		log.Printf(err_str)
		return op.NewListtimestampedDefault(500).WithPayload(&models.Error{Message: &err_str})
	}
	var ret []*models.ReceiptFile
	for _, rcpt := range rcpts {
		ret_rcpt := models.ReceiptFile{Filename: rcpt.Filename,
			Hash:              rcpt.TargetHash,
			Horodatingaddress: "0xTODO",
			Transactionhash:   rcpt.TransactionHash,
		}
		ret = append(ret, &ret_rcpt)
	}
	return op.NewListtimestampedOK().WithPayload(ret)
}
