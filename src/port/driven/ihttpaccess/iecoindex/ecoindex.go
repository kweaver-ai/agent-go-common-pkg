package iecoindex

import (
	"context"

	"github.com/kweaver-ai/agent-go-common-pkg/src/domain/valueobject/comvalobj"
	"github.com/kweaver-ai/agent-go-common-pkg/src/domain/valueobject/daconfvalobj/datasourcevalobj"
	"github.com/kweaver-ai/agent-go-common-pkg/src/drivenadapter/httpaccess/ecoindexhttp/ecoindexdto"
)

type IEcoIndex interface {
	AddBotSourceIndex(ctx context.Context, uniqueFlag string, sources []*datasourcevalobj.DocSourceField) (err error)

	RemoveSourceIndex(ctx context.Context, uniqueFlag string, sources []*datasourcevalobj.DocSourceField) (err error)

	// DeleteBotIndex(ctx context.Context, botID string) (err error)

	GetBotIndexStatus(ctx context.Context, uniqueFlag string, showFailInfos bool, failInfoPagination *comvalobj.Pagination) (info *ecoindexdto.BotIndexStatusInfo, err error)
}
