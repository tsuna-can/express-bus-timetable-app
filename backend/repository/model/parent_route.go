package model

import (
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
)

type ParentRoute struct {
	parent_route_id string
	parent_route_name string
}

func (pr *ParentRoute) ToParentRoute() *entity.ParentRoute {
	return entity.NewParentRoute(
		pr.parent_route_id,
		pr.parent_route_name,
	)
}
