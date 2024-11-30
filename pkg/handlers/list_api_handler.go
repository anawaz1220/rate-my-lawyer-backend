package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/mikestefanello/pagoda/ent/jurisdiction"
	"github.com/mikestefanello/pagoda/ent/lawyer"
	"github.com/mikestefanello/pagoda/pkg/services"
)

type ListAPI struct {
	services.Container
}

func (h *ListAPI) Init(c *services.Container) error {
	h.Container = *c
	return nil
}

func (h *ListAPI) Routes(g *echo.Group) {
	g.GET("/api/lists", h.GetLists)
}

func (h *ListAPI) GetLists(c echo.Context) error {
	ctx := c.Request().Context()

	// Fetch unique full names
	fullNames, err := h.Container.ORM.Lawyer.Query().
		GroupBy(lawyer.FieldFullName).
		Strings(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch full names"})
	}

	// Fetch unique jurisdictions
	jurisdictions, err := h.Container.ORM.Jurisdiction.Query().
		GroupBy(jurisdiction.FieldName).
		Strings(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch jurisdictions"})
	}

	// Fetch unique practicing statuses
	practicingStatuses, err := h.Container.ORM.Lawyer.Query().
		GroupBy(lawyer.FieldPracticingStatus).
		Strings(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch practicing statuses"})
	}

	response := map[string]interface{}{
		"full_names":          fullNames,
		"jurisdictions":       jurisdictions,
		"practicing_statuses": practicingStatuses,
	}

	return c.JSON(http.StatusOK, response)
}
func init() {
	Register(new(ListAPI))
}
