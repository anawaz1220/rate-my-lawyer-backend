package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/lawyer"
	"github.com/mikestefanello/pagoda/pkg/services"
)

type DashboardAPI struct {
	services.Container
}

func (h *DashboardAPI) Init(c *services.Container) error {
	h.Container = *c
	return nil
}

func (h *DashboardAPI) Routes(g *echo.Group) {
	g.GET("/api/dashboard/top-lawyers", h.GetTopLawyers)
	g.GET("/api/dashboard/stats", h.GetStats)
}

type TopLawyerResponse struct {
	ID               int     `json:"id"`
	FullName         string  `json:"full_name"`
	AverageRating    float64 `json:"average_rating"`
	ReviewCount      int     `json:"review_count"`
	PracticingStatus string  `json:"practicing_status"`
}

func (h *DashboardAPI) GetTopLawyers(c echo.Context) error {
	ctx := c.Request().Context()

	topLawyers, err := h.Container.ORM.Lawyer.Query().
		Order(ent.Desc(lawyer.FieldAverageRating)).
		Limit(5).
		All(ctx)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch top lawyers"})
	}

	var response []TopLawyerResponse
	for _, l := range topLawyers {
		response = append(response, TopLawyerResponse{
			ID:               l.ID,
			FullName:         l.FullName,
			AverageRating:    l.AverageRating,
			ReviewCount:      l.ReviewCount,
			PracticingStatus: l.PracticingStatus,
		})
	}

	return c.JSON(http.StatusOK, response)
}

type StatsResponse struct {
	TotalLawyers       int `json:"total_lawyers"`
	TotalJurisdictions int `json:"total_jurisdictions"`
	TotalUsers         int `json:"total_users"`
}

func (h *DashboardAPI) GetStats(c echo.Context) error {
	ctx := c.Request().Context()

	lawyerCount, err := h.Container.ORM.Lawyer.Query().Count(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch total lawyers"})
	}

	jurisdictionCount, err := h.Container.ORM.Jurisdiction.Query().Count(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch total jurisdictions"})
	}

	userCount, err := h.Container.ORM.RMLUser.Query().Count(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch total users"})
	}

	stats := StatsResponse{
		TotalLawyers:       lawyerCount,
		TotalJurisdictions: jurisdictionCount,
		TotalUsers:         userCount,
	}

	return c.JSON(http.StatusOK, stats)
}

func init() {
	Register(new(DashboardAPI))
}
