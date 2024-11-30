package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/page"
	"github.com/mikestefanello/pagoda/pkg/services"
)

type MvpRoutePlanner struct {
	services.Container
}

func (h *MvpRoutePlanner) Init(c *services.Container) error {
	h.Container = *c
	return nil
}

func (h *MvpRoutePlanner) Routes(g *echo.Group) {
	g.GET("/route-planner", h.Get).Name = "route_planner"
}

func (h *MvpRoutePlanner) Get(c echo.Context) error {
	p := page.New(c)
	p.Name = "route_planner"
	p.Layout = "main"
	// p.Title = "Route Planner"

	return h.Container.TemplateRenderer.RenderPage(c, p)
}

func init() {
	Register(new(MvpRoutePlanner))
}
