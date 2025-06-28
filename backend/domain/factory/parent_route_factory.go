package factory

import (
	"fmt"

	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/entity"
	"github.com/tsuna-can/express-bus-time-table-app/backend/domain/vo"
)

// ParentRouteRawData represents raw data from database for parent route
type ParentRouteRawData struct {
	ParentRouteId   string
	ParentRouteName string
}

// ParentRouteFactory creates ParentRoute entity from raw data
type ParentRouteFactory interface {
	ReconstructFromRawData(data ParentRouteRawData) (*entity.ParentRoute, error)
	ReconstructManyFromRawData(data []ParentRouteRawData) ([]entity.ParentRoute, error)
}

type parentRouteFactory struct{}

// NewParentRouteFactory creates a new instance of ParentRouteFactory
func NewParentRouteFactory() ParentRouteFactory {
	return &parentRouteFactory{}
}

// ReconstructFromRawData reconstructs a ParentRoute entity from raw data
func (f *parentRouteFactory) ReconstructFromRawData(data ParentRouteRawData) (*entity.ParentRoute, error) {
	name, err := vo.NewParentRouteName(data.ParentRouteName)
	if err != nil {
		return nil, fmt.Errorf("failed to create ParentRouteName from raw data: %w", err)
	}

	return entity.NewParentRoute(data.ParentRouteId, *name), nil
}

// ReconstructManyFromRawData reconstructs multiple ParentRoute entities from raw data slice
func (f *parentRouteFactory) ReconstructManyFromRawData(data []ParentRouteRawData) ([]entity.ParentRoute, error) {
	parentRoutes := make([]entity.ParentRoute, 0, len(data))

	for i, rawData := range data {
		parentRoute, err := f.ReconstructFromRawData(rawData)
		if err != nil {
			return nil, fmt.Errorf("failed to reconstruct ParentRoute at index %d: %w", i, err)
		}
		parentRoutes = append(parentRoutes, *parentRoute)
	}

	return parentRoutes, nil
}
