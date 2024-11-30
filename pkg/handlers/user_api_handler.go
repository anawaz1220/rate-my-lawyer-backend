package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/rmluser"
	"github.com/mikestefanello/pagoda/pkg/services"
)

type UserAPI struct {
	services.Container
}

func (h *UserAPI) Init(c *services.Container) error {
	h.Container = *c
	return nil
}

type TopUserResponse struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	ReviewCount int    `json:"review_count"`
}

func (h *UserAPI) GetTopUsers(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit <= 0 {
		limit = 10 // Default limit
	}

	ctx := c.Request().Context()

	users, err := h.Container.ORM.RMLUser.Query().
		Order(ent.Desc(rmluser.FieldReviewCount)).
		Limit(limit).
		All(ctx)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch top users"})
	}

	var response []TopUserResponse
	for _, u := range users {
		response = append(response, TopUserResponse{
			Name:        u.Name,
			Email:       u.Email,
			ReviewCount: u.ReviewCount,
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *UserAPI) Routes(g *echo.Group) {
	g.GET("/api/users/top", h.GetTopUsers)
}

func init() {
	Register(new(UserAPI))
}
