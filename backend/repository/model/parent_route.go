package model

import (
  "fmt"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

type ParentRoute struct {
	ParentRouteId   string
	ParentRouteName string
}

func (pr *ParentRoute) ToParentRoute() (*entity.ParentRoute, error) {
	name, err := vo.NewParentRouteName(pr.ParentRouteName)
	if err != nil {
		return nil, fmt.Errorf("failed to create ParentRouteName: %w", err)
	}

	return entity.NewParentRoute(
    pr.ParentRouteId,
    *name,
	), nil
}

