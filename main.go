package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type People struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	IsMarried bool   `isMarried:"isMarried"`
}

var Talent = []People{
	{
		Id:        "1",
		Name:      "Rizal Brutal",
		IsMarried: true,
	},
	{
		Id:        "2",
		Name:      "Gendi Brutal",
		IsMarried: false,
	},
}

func main() {
	e := echo.New()

	e.GET("/peoples", FindPeoples)
	e.GET("/people/:id", GetPeople)
	e.POST("/people", AddPeople)
	e.DELETE("/people/:id", DeletePeople)

	fmt.Println("Running on port 5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}

func FindPeoples(c echo.Context) error {
	c.Response().Header().Set("Content-type", "application/json")
	c.Response().WriteHeader(http.StatusOK)

	return json.NewEncoder(c.Response()).Encode(Talent)
}

func AddPeople(c echo.Context) error {
	var data People

	json.NewDecoder(c.Request().Body).Decode(&data)

	Talent = append(Talent, data)
	c.Response().Header().Set("Content-type", "application/json")
	c.Response().WriteHeader(http.StatusOK)

	return json.NewEncoder(c.Response()).Encode(Talent)
}

func GetPeople(c echo.Context) error {
	c.Response().Header().Set("Content-type", "application/json")
	id := c.Param("id")
	var Data People
	var cekId = false

	for _, talent := range Talent {
		if id == talent.Id {
			cekId = true
			Data = talent
		}
	}

	if !cekId {
		c.Response().WriteHeader(http.StatusNotFound)
		return json.NewEncoder(c.Response()).Encode("ID: " + id + " not found")
	}

	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(Data)
}

func DeletePeople(c echo.Context) error {
	id := c.Param("id")
	var cekId = false
	var index = 0

	for i, talent := range Talent {
		if id == talent.Id {
			cekId = true
			index = i
		}
	}

	if !cekId {
		c.Response().WriteHeader(http.StatusNotFound)
		return json.NewEncoder(c.Response()).Encode("ID: " + id + " not found")
	}

	Talent = append(Talent[:index], Talent[index+1:]...)

	c.Response().Header().Set("Content-type", "application/json")
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(Talent)
}
