package entity

type ParentRoute struct {
	ParentRouteId   string
	ParentRouteName string
}

func NewParentRoute(parentRouteId, parentRouteName string) *ParentRoute {
	return &ParentRoute{
		ParentRouteId:   parentRouteId,
		ParentRouteName: parentRouteName,
	}
}

