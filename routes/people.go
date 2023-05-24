package routes

import (
	"dumbmerch/handlers"

	"github.com/labstack/echo/v4"
)

func PeopleRoutes(e *echo.Group) {
	e.GET("/peoples", handlers.FindPeoples)
	e.GET("/people/:id", handlers.GetPeople)
	e.POST("/people", handlers.AddPeople)
	e.DELETE("/people/:id", handlers.DeletePeople)
}
