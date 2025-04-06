package input

import (
	"context"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
)

type ParentRoutesInputPort interface {
	GetAll(ctx context.Context) ([]entity.ParentRoute, error)
}

