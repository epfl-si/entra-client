package client

import "epfl-entra/internal/models"

type Service interface {
	CreateApplication(app *models.Application, options models.ClientOptions) (err error)
	DeleteApplication(id string, options models.ClientOptions) (err error)
	GetApplication(id string, options models.ClientOptions) (apps *models.Application, err error)
	GetApplications(options models.ClientOptions) (apps []*models.Application, nextURL string, err error)

	CreateGroup(app *models.Group, options models.ClientOptions) (err error)
	DeleteGroup(id string, options models.ClientOptions) (err error)
	GetGroup(id string, options models.ClientOptions) (groups *models.Group, err error)
	GetGroups(options models.ClientOptions) (groups []*models.Group, nextURL string, err error)

	CreateUser(app *models.User, options models.ClientOptions) (err error)
	DeleteUser(id string, options models.ClientOptions) (err error)
	GetUser(id string, options models.ClientOptions) (app *models.User, err error)
	GetUsers(options models.ClientOptions) (users []*models.User, nextURL string, err error)
	UpdateUser(app *models.User, options models.ClientOptions) (err error)
}
