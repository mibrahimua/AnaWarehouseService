package service

import (
	"AnaWarehouseService/model"
	"AnaWarehouseService/repository"
	"AnaWarehouseService/request"
)

type WarehouseService struct {
	warehouseRepository *repository.WarehouseRepository
}

func NewProductService(warehouseRepository *repository.WarehouseRepository) *WarehouseService {
	return &WarehouseService{warehouseRepository}
}

func (us *WarehouseService) StocksTransferRequest(param request.WarehouseTransferRequest) ([]model.Product, error) {
	return us.warehouseRepository.StocksTransferRequest(param)
}

func (us *WarehouseService) UpdateStatusWarehouse(productName string) ([]model.Product, error) {
	return us.warehouseRepository.UpdateStatusWarehouse(productName)
}
