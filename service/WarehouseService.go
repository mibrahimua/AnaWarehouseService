package service

import (
	"AnaWarehouseService/repository"
	"AnaWarehouseService/request"
)

type WarehouseService struct {
	warehouseRepository *repository.WarehouseRepository
}

func NewProductService(warehouseRepository *repository.WarehouseRepository) *WarehouseService {
	return &WarehouseService{warehouseRepository}
}

func (us *WarehouseService) StocksTransferRequest(param request.WarehouseTransferRequest) error {
	return us.warehouseRepository.StocksTransferRequest(param)
}

func (us *WarehouseService) UpdateStatusWarehouse(param request.WarehouseStatus) error {
	return us.warehouseRepository.UpdateStatusWarehouse(param)
}
