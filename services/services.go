package services

type (
	getServicesDao interface {
		// Implements from models
	}

	GetServices struct {
		dao getServicesDao
	}
)

func NewServices(dao getServicesDao) *GetServices {
	return &GetServices{dao}
}
