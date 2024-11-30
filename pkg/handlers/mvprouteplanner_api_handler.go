package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/mvpmaterial"
	"github.com/mikestefanello/pagoda/ent/mvpplannedroute"
	"github.com/mikestefanello/pagoda/ent/mvproute"
	"github.com/mikestefanello/pagoda/ent/mvpstaff"
	"github.com/mikestefanello/pagoda/ent/mvptruck"
	"github.com/mikestefanello/pagoda/pkg/services"
)

type MvpRoutePlannerAPI struct {
	services.Container
}

func (h *MvpRoutePlannerAPI) Init(c *services.Container) error {
	h.Container = *c
	return nil
}

func (h *MvpRoutePlannerAPI) Routes(g *echo.Group) {
	g.GET("/api/calendar/current-month", h.GetCurrentMonthData)
	g.GET("/api/routes-by-day", h.GetRoutesByDay)
	g.GET("/api/available/trucks", h.GetAvailableTrucks)
	g.GET("/api/available/drivers", h.GetAvailableDrivers)
	g.GET("/api/available/loaders", h.GetAvailableLoaders)
	g.GET("/api/available/materials", h.GetAvailableMaterials)
	g.POST("/api/route-schedule", h.SubmitRouteSchedule)
	g.GET("/api/route-schedule", h.GetRouteScheduleDetails)
}

// func (h *MvpRoutePlannerAPI) RoutePlannerPage(c echo.Context) error {
// 	p := page.New(c)
// 	p.Layout = "main"
// 	p.Name = "route_planner"
// 	p.Title = "Route Planner"
// 	return h.Container.ORM.TemplateRenderer.RenderPage(c, p)
// }

type CalendarDay struct {
	Date   string `json:"date"`
	Status string `json:"status"`
}

type CalendarData struct {
	Month int           `json:"month"`
	Year  int           `json:"year"`
	Days  []CalendarDay `json:"days"`
}

func (h *MvpRoutePlannerAPI) GetCurrentMonthData(c echo.Context) error {
	now := time.Now()
	firstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	lastDay := firstDay.AddDate(0, 1, -1)

	ctx := c.Request().Context()

	// Fetch all planned routes for the current month
	plannedRoutes, err := h.Container.ORM.MvpPlannedRoute.Query().
		Where(mvpplannedroute.DateGTE(firstDay), mvpplannedroute.DateLTE(lastDay)).
		All(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch planned routes"})
	}

	// Create a map to store the status of each day
	dayStatus := make(map[string]string)

	// Initialize all days as "unplanned"
	for d := firstDay; !d.After(lastDay); d = d.AddDate(0, 0, 1) {
		dayStatus[d.Format("2006-01-02")] = "unplanned"
	}

	// Update status based on planned routes
	for _, route := range plannedRoutes {
		date := route.Date.Format("2006-01-02")
		if dayStatus[date] == "unplanned" {
			dayStatus[date] = "partially_planned"
		} else if dayStatus[date] == "partially_planned" {
			dayStatus[date] = "fully_planned"
		}
	}

	// Create the response data
	calendarData := CalendarData{
		Month: int(now.Month()),
		Year:  now.Year(),
		Days:  make([]CalendarDay, 0),
	}

	for d := firstDay; !d.After(lastDay); d = d.AddDate(0, 0, 1) {
		date := d.Format("2006-01-02")
		calendarData.Days = append(calendarData.Days, CalendarDay{
			Date:   date,
			Status: dayStatus[date],
		})
	}

	return c.JSON(http.StatusOK, calendarData)
}

func (h *MvpRoutePlannerAPI) GetRoutesByDay(c echo.Context) error {
	day := c.QueryParam("day")
	if day == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Day parameter is required"})
	}

	// Normalize the day input
	day = strings.Title(strings.ToLower(day))

	ctx := c.Request().Context()

	// Fetch routes for the given day of the week
	routes, err := h.Container.ORM.MvpRoute.Query().
		Where(mvproute.DayOfWeekEQ(day)).
		All(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch routes"})
	}

	// Fetch planned routes for today
	// today := time.Now().Format("2006-01-02")
	today := time.Now().Truncate(24 * time.Hour)
	plannedRoutes, err := h.Container.ORM.MvpPlannedRoute.Query().
		Where(mvpplannedroute.DateEQ(today)).
		All(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch planned routes"})
	}

	log.Printf("Found %d routes for day %s", len(routes), day)
	log.Printf("Found %d planned routes for today", len(plannedRoutes))

	// Create a map of planned routes for quick lookup
	plannedRouteMap := make(map[string]*ent.MvpPlannedRoute)
	for _, pr := range plannedRoutes {
		log.Printf("Planned route: %s, Status: %s", pr.RouteName, pr.Status)
		plannedRouteMap[pr.RouteName] = pr
	}

	// Prepare the response
	type RouteResponse struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		DayOfWeek string `json:"day_of_week"`
		Status    string `json:"status"`
	}

	response := make([]RouteResponse, len(routes))
	for i, route := range routes {
		status := "unplanned"
		if pr, exists := plannedRouteMap[route.Name]; exists {

			status = pr.Status

		}
		response[i] = RouteResponse{
			ID:        route.ID,
			Name:      route.Name,
			DayOfWeek: route.DayOfWeek,
			Status:    status,
		}
	}

	return c.JSON(http.StatusOK, response)
}

func (h *MvpRoutePlannerAPI) GetAvailableTrucks(c echo.Context) error {
	date, err := time.Parse("2006-01-02", c.QueryParam("date"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid date format"})
	}

	ctx := c.Request().Context()

	// Get all trucks
	allTrucks, err := h.Container.ORM.MvpTruck.Query().All(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch trucks"})
	}

	// Get trucks already assigned for the given date
	assignedTrucks, err := h.Container.ORM.MvpPlannedRoute.Query().
		Where(mvpplannedroute.DateEQ(date)).
		QueryTrucks().
		All(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch assigned trucks"})
	}

	// Create a map of assigned truck IDs for quick lookup
	assignedTruckMap := make(map[int]bool)
	for _, truck := range assignedTrucks {
		assignedTruckMap[truck.ID] = true
	}

	// Filter out assigned trucks
	availableTrucks := make([]*ent.MvpTruck, 0)
	for _, truck := range allTrucks {
		if !assignedTruckMap[truck.ID] {
			availableTrucks = append(availableTrucks, truck)
		}
	}

	return c.JSON(http.StatusOK, availableTrucks)
}

func (h *MvpRoutePlannerAPI) GetAvailableDrivers(c echo.Context) error {
	date, err := time.Parse("2006-01-02", c.QueryParam("date"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid date format"})
	}

	ctx := c.Request().Context()

	// Get all drivers
	allDrivers, err := h.Container.ORM.MvpStaff.Query().
		Where(mvpstaff.RoleEQ("DRIVER")).
		All(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch drivers"})
	}

	// Get drivers already assigned for the given date
	assignedDrivers, err := h.Container.ORM.MvpPlannedRoute.Query().
		Where(mvpplannedroute.DateEQ(date)).
		QueryDriver().
		All(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch assigned drivers"})
	}

	// Create a map of assigned driver IDs for quick lookup
	assignedDriverMap := make(map[int]bool)
	for _, driver := range assignedDrivers {
		assignedDriverMap[driver.ID] = true
	}

	// Filter out assigned drivers
	availableDrivers := make([]*ent.MvpStaff, 0)
	for _, driver := range allDrivers {
		if !assignedDriverMap[driver.ID] {
			availableDrivers = append(availableDrivers, driver)
		}
	}

	return c.JSON(http.StatusOK, availableDrivers)
}

func (h *MvpRoutePlannerAPI) GetAvailableLoaders(c echo.Context) error {
	date, err := time.Parse("2006-01-02", c.QueryParam("date"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid date format"})
	}

	ctx := c.Request().Context()

	// Get all loaders
	allLoaders, err := h.Container.ORM.MvpStaff.Query().
		Where(mvpstaff.RoleEQ("LOADER")).
		All(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch loaders"})
	}

	// Get loaders already assigned for the given date
	assignedLoaders, err := h.Container.ORM.MvpPlannedRoute.Query().
		Where(mvpplannedroute.DateEQ(date)).
		QueryLoaders().
		All(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch assigned loaders"})
	}

	// Create a map of assigned loader IDs for quick lookup
	assignedLoaderMap := make(map[int]bool)
	for _, loader := range assignedLoaders {
		assignedLoaderMap[loader.ID] = true
	}

	// Filter out assigned loaders
	availableLoaders := make([]*ent.MvpStaff, 0)
	for _, loader := range allLoaders {
		if !assignedLoaderMap[loader.ID] {
			availableLoaders = append(availableLoaders, loader)
		}
	}

	return c.JSON(http.StatusOK, availableLoaders)
}

func (h *MvpRoutePlannerAPI) GetAvailableMaterials(c echo.Context) error {
	ctx := c.Request().Context()

	// Get all materials
	allMaterials, err := h.Container.ORM.MvpMaterial.Query().All(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch materials"})
	}

	return c.JSON(http.StatusOK, allMaterials)
}

type RouteScheduleInput struct {
	Date        string `json:"date"`
	RouteName   string `json:"route_name"`
	TruckIDs    []int  `json:"truck_ids"`
	DriverID    int    `json:"driver_id"`
	LoaderIDs   []int  `json:"loader_ids"`
	MaterialIDs []int  `json:"material_ids"`
}

func (h *MvpRoutePlannerAPI) SubmitRouteSchedule(c echo.Context) error {
	var input RouteScheduleInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input: " + err.Error()})
	}

	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid date format: " + err.Error()})
	}

	ctx := c.Request().Context()

	// Start a transaction
	tx, err := h.Container.ORM.Tx(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to start transaction: " + err.Error()})
	}
	defer tx.Rollback()

	// Basic validation: check if resources exist
	if err := h.validateResources(ctx, tx, input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Resource validation failed: " + err.Error()})
	}

	// Create the planned route
	plannedRoute, err := tx.MvpPlannedRoute.Create().
		SetDate(date).
		SetRouteName(input.RouteName).
		SetStatus("planned").
		Save(ctx)

	if err != nil {
		log.Printf("Failed to create planned route: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create planned route: " + err.Error()})
	}

	// Add relationships
	err = plannedRoute.Update().
		AddTruckIDs(input.TruckIDs...).
		SetDriverID(input.DriverID).
		AddLoaderIDs(input.LoaderIDs...).
		AddMaterialIDs(input.MaterialIDs...).
		Exec(ctx)

	if err != nil {
		log.Printf("Failed to add relationships: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add relationships: " + err.Error()})
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to commit transaction: " + err.Error()})
	}

	return c.JSON(http.StatusCreated, plannedRoute)
}

func (h *MvpRoutePlannerAPI) validateResources(ctx context.Context, tx *ent.Tx, input RouteScheduleInput) error {
	// Check if trucks exist
	for _, truckID := range input.TruckIDs {
		exists, err := tx.MvpTruck.Query().Where(mvptruck.ID(truckID)).Exist(ctx)
		if err != nil {
			log.Printf("Error checking truck existence: %v", err)
			return fmt.Errorf("error checking truck existence: %v", err)
		}
		if !exists {
			log.Printf("Truck with ID %d does not exist", truckID)
			return fmt.Errorf("truck with ID %d does not exist", truckID)
		}
	}

	// Check if driver exists
	driverExists, err := tx.MvpStaff.Query().Where(mvpstaff.ID(input.DriverID), mvpstaff.RoleEQ("DRIVER")).Exist(ctx)
	if err != nil {
		log.Printf("Error checking driver existence: %v", err)
		return fmt.Errorf("error checking driver existence: %v", err)
	}
	if !driverExists {
		log.Printf("Driver with ID %d does not exist", input.DriverID)
		return fmt.Errorf("driver with ID %d does not exist", input.DriverID)
	}

	// Check if loaders exist
	for _, loaderID := range input.LoaderIDs {
		exists, err := tx.MvpStaff.Query().Where(mvpstaff.ID(loaderID), mvpstaff.RoleEQ("LOADER")).Exist(ctx)
		if err != nil {
			log.Printf("Error checking loader existence: %v", err)
			return fmt.Errorf("error checking loader existence: %v", err)
		}
		if !exists {
			log.Printf("Loader with ID %d does not exist", loaderID)
			return fmt.Errorf("loader with ID %d does not exist", loaderID)
		}
	}

	// Check if materials exist
	for _, materialID := range input.MaterialIDs {
		exists, err := tx.MvpMaterial.Query().Where(mvpmaterial.ID(materialID)).Exist(ctx)
		if err != nil {
			log.Printf("Error checking material existence: %v", err)
			return fmt.Errorf("error checking material existence: %v", err)
		}
		if !exists {
			log.Printf("Material with ID %d does not exist", materialID)
			return fmt.Errorf("material with ID %d does not exist", materialID)
		}
	}

	return nil
}

func (h *MvpRoutePlannerAPI) GetRouteScheduleDetails(c echo.Context) error {
	date, err := time.Parse("2006-01-02", c.QueryParam("date"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid date format"})
	}

	routeName := c.QueryParam("route_name")
	if routeName == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Route name is required"})
	}

	ctx := c.Request().Context()

	plannedRoute, err := h.Container.ORM.MvpPlannedRoute.Query().
		Where(mvpplannedroute.DateEQ(date), mvpplannedroute.RouteNameEQ(routeName)).
		WithTrucks().
		WithDriver().
		WithLoaders().
		WithMaterials().
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Route schedule not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch route schedule"})
	}

	return c.JSON(http.StatusOK, plannedRoute)
}

func init() {
	Register(new(MvpRoutePlannerAPI))
}
