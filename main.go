package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type User struct {
	Name string `json:"name" xml:"name" form:"name" query:"name" validate:"required"`
	Id   int    `json:"id" xml:"id" form:"id" query:"id"`
}

type Message struct {
	Text string `json:"text" xml:"text" form:"text" query:"text"`
}

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, []User{
			{Id: 1, Name: "john doe"},
			{Id: 2, Name: "anna boe"},
		})
	})

	e.GET("/new", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Message{Text: "welcome to the new page"})
	})

	e.GET("/nonexisting", func(c echo.Context) error {
		return c.String(http.StatusNotFound, "")
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusMovedPermanently, "")
		//return c.Redirect(http.StatusMovedPermanently, "/redirect-to")
	})

	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.String(http.StatusOK, fmt.Sprintf("%s has been added to the users list", u.Name))
	})

	e.Logger.Fatal(e.Start(":3001"))
}
