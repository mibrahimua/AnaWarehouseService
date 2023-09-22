package request

type WarehouseTransferRequest struct {
	ProductId       int `json:"product_id"`
	FromWarehouseId int `json:"from_warehouse_id"`
	ToWarehouseId   int `json:"to_warehouse_id"`
	Quantity        int `json:"quantity"`
}

type WarehouseDetail struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Location       string `json:"location"`
	ProductStockId int    `json:"product_stock_id"`
	ProductId      int    `json:"product_id"`
	Quantity       int    `json:"quantity"`
}

type WarehouseStatus struct {
	ID       int  `json:"id"`
	IsActive bool `json:"is_active"`
}
