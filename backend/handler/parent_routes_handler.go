package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tsuna-can/express-bus-time-table-app/backend/handler/response"
	"github.com/tsuna-can/express-bus-time-table-app/backend/usecase"
)

type ParentRoutesHandler struct {
	getParentRoutesUsecase usecase.GetParentRoutesUsecase
}

func NewParentRoutesHandler(getParentRoutesUsecase usecase.GetParentRoutesUsecase) *ParentRoutesHandler {
	return &ParentRoutesHandler{
		getParentRoutesUsecase: getParentRoutesUsecase,
	}
}

// GetParentRoutes godoc
// @Summary Get all parent routes
// @Description Get all parent routes
// @Tags parent_routes
// @Accept json
// @Produce json
// @Success 200 {array} response.ParentRouteResponse
// @Router /parent-routes [get]
func (h ParentRoutesHandler) GetParentRoutes(e echo.Context) error {
	ctx := e.Request().Context()
	parentRoutes, err := h.getParentRoutesUsecase.GetAll(ctx)
	if err != nil {
		log.Printf("Error getting parent routes: %v", err)
		return e.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: "Internal server error"})
	}

	response := response.NewParentRoutesResponse(parentRoutes)
	return e.JSON(http.StatusOK, response)
}
