package idatahubacc

import (
	"context"

	"github.com/kweaver-ai/agent-go-common-pkg/src/drivenadapter/httpaccess/datahubcentralhttp/datahubcentraldto"
)

type IDataHubCentral interface {
	CreateDataset(ctx context.Context, req *datahubcentraldto.CreateDatasetsReq) (datasetID string, err error)
}
