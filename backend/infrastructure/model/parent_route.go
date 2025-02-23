package model

import (
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/model"
)

type ParentRoute struct {
	parent_route_id string
	parent_route_name string
}

func (pr *ParentRoute) ToParentRoute() *model.ParentRoute {
	return model.NewParentRoute(
		pr.parent_route_id,
		pr.parent_route_name,
	)
}
