package entity

import (
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

type ParentRoute struct {
	ParentRouteId   string
	ParentRouteName vo.ParentRouteName
}

func NewParentRoute(parentRouteId string, parentRouteName vo.ParentRouteName) *ParentRoute {
	return &ParentRoute{
		ParentRouteId:   parentRouteId,
		ParentRouteName: parentRouteName,
	}
}

