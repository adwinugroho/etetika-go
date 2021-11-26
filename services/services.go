package services

import (
	"github.com/adwinugroho/etetika-go/models"
)

type (
	getServicesDao interface {
		// Implements from models
		GetUserByEmail(email string) (*models.User, error)
	}

	GetServices struct {
		dao getServicesDao
	}
)

func NewServices(dao getServicesDao) *GetServices {
	return &GetServices{dao}
}

func (service *GetServices) GetDataUserByEmail(email string) interface{} {
	// var userCtx config.UserContext
	// var user *request.User = ctx.Value(userCtx).(*request.User)
	data, err := service.dao.GetUserByEmail(email)
	if err != nil {
		return map[string]interface{}{
			"errMessage": "error when get data",
		}
	}
	return map[string]interface{}{
		"data": data,
	}
}
