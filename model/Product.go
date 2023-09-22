package model

type Product struct {
	ID            int     `json:"id"`
	Category      string  `json:"category"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	Theme         string  `json:"theme"`
	StockQuantity int     `json:"stock_quantity"`
}

type ProductStock struct {
	ID            int    `json:"id"`
	ProductId     string `json:"product_id"`
	WarehouseId   string `json:"warehouse_id"`
	StockQuantity string `json:"stock_quantity"`
}

type ProductTransfer struct {
	ID              int `json:"id"`
	ProductId       int `json:"product_id"`
	FromWarehouseId int `json:"from_warehouse_id"`
	ToWarehouseId   int `json:"to_warehouse_id"`
	Quantity        int `json:"quantity"`
}
