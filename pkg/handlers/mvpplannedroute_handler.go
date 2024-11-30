package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/services"
)

type MvpPlannedRoute struct {
	services.Container
}

func (h *MvpPlannedRoute) Init(c *services.Container) error {
	h.Container = *c
	return nil
}

func (h *MvpPlannedRoute) Routes(g *echo.Group) {
	g.POST("/api/planned-routes", h.Create).Name = "api_create_planned_route"
	g.PUT("/api/planned-routes/:id", h.Update).Name = "api_update_planned_route"
}

func (h *MvpPlannedRoute) Create(c echo.Context) error {
	var input struct {
		Date        string `json:"date"`
		RouteName   string `json:"routeName"`
		TruckIDs    []int  `json:"truckIDs"`
		DriverID    int    `json:"driverID"`
		LoaderIDs   []int  `json:"loaderIDs"`
		MaterialIDs []int  `json:"materialIDs"`
	}

	if err := c.Bind(&input); err != nil {
		return err
	}

	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()

	plannedRoute, err := h.ORM.MvpPlannedRoute.Create().
		SetDate(date).
		SetRouteName(input.RouteName).
		SetStatus("planned").
		Save(ctx)

	if err != nil {
		return err
	}

	// Add relationships
	err = plannedRoute.Update().
		AddTruckIDs(input.TruckIDs...).
		SetDriverID(input.DriverID).
		AddLoaderIDs(input.LoaderIDs...).
		AddMaterialIDs(input.MaterialIDs...).
		Exec(ctx)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, plannedRoute)
}

func (h *MvpPlannedRoute) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	var input struct {
		TruckIDs    []int `json:"truckIDs"`
		DriverID    int   `json:"driverID"`
		LoaderIDs   []int `json:"loaderIDs"`
		MaterialIDs []int `json:"materialIDs"`
	}

	if err := c.Bind(&input); err != nil {
		return err
	}

	ctx := c.Request().Context()

	plannedRoute, err := h.ORM.MvpPlannedRoute.UpdateOneID(id).
		SetStatus("edited").
		ClearTrucks().
		ClearDriver().
		ClearLoaders().
		ClearMaterials().
		Save(ctx)

	if err != nil {
		return err
	}

	err = plannedRoute.Update().
		AddTruckIDs(input.TruckIDs...).
		SetDriverID(input.DriverID).
		AddLoaderIDs(input.LoaderIDs...).
		AddMaterialIDs(input.MaterialIDs...).
		Exec(ctx)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, plannedRoute)
}

func init() {
	Register(new(MvpPlannedRoute))
}
