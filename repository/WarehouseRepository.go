package repository

import (
	"AnaWarehouseService/model"
	"AnaWarehouseService/request"
	"database/sql"
)

type WarehouseRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *WarehouseRepository {
	return &WarehouseRepository{db}
}

func (u *WarehouseRepository) StocksTransferRequest(param request.WarehouseTransferRequest) ([]model.Product, error) {
	var products []model.Product

	query := "WITH _param AS (SELECT $1 AS product_name)\n   , _warehouse AS (SELECT\n                        *\n                    FROM warehouse\n                    WHERE\n                        is_active IS TRUE)\n\n   , _product_stock AS (SELECT\n                            stock.*\n                        FROM product_stock stock\n                        JOIN _warehouse warehouse ON warehouse.id = stock.warehouse_id)\n\n   , _product AS (SELECT\n                      product.*,\n                      stock.stock_quantity\n                  FROM _product_stock stock\n                  JOIN product ON product.id = stock.product_id\n                  WHERE\n                      product.name ILIKE '%' || (SELECT product_name FROM _param) || '%')\n\nSELECT\n    *\nFROM _product"
	rows, err := u.db.Query(query, param)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Category, &product.Name, &product.Description, &product.Price, &product.Theme, &product.StockQuantity); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (u *WarehouseRepository) UpdateStatusWarehouse(name string) ([]model.Product, error) {

}
