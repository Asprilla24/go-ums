package controllers

import (
	"github.com/Asprilla24/go-ums/app/models"
	"github.com/revel/revel"
)

type AppUsers struct {
	*revel.Controller
}

func (c AppUsers) GetUsers() revel.Result {
	response := Response{}
	users := []models.AppUsers{}

	result := DB.Find(&users)
	if result.Error != nil {
		response.Error = result.Error.Error()
	} else {
		response.Result = users
	}

	return c.RenderJSON(response)
}

func (c AppUsers) GetUser(id string) revel.Result {
	response := Response{}
	user := models.AppUsers{}

	result := DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		response.Error = result.Error.Error()
	} else {
		response.Result = user
	}

	return c.RenderJSON(response)
}
