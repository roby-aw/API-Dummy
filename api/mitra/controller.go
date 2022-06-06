package mitra

import (
	mitraBusiness "api-redeem-point/business/mitra"
)

type Controller struct {
	service mitraBusiness.Service
}

func NewController(service mitraBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}
