package repository

import (
	"AnaWarehouseService/model"
	"AnaWarehouseService/request"
	"database/sql"
	"log"
)

type WarehouseRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *WarehouseRepository {
	return &WarehouseRepository{db}
}

func (u *WarehouseRepository) StocksTransferRequest(param request.WarehouseTransferRequest) error {
	tx, err := u.db.Begin()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	var transferStockId model.ProductTransfer
	query := "WITH _param AS (SELECT\n                    $1::bigint AS product_id,\n                    $2::bigint AS from_warehouse_id,\n                    $3::bigint AS to_warehouse_id,\n                    $4::bigint AS quantity)\n\n   , _source_product_stock AS (SELECT\n                                   id,\n                                   product_id,\n                                   warehouse_id,\n                                   stock_quantity - (SELECT quantity FROM _param) AS stock_quantity,\n                                   CASE\n                                       WHEN stock_quantity - (SELECT quantity FROM _param) < 0 THEN\n                                           FALSE\n                                       ELSE TRUE END AS is_valid\n                               FROM product_stock\n                               WHERE\n                                   product_id = (SELECT product_id FROM _param) AND\n                                   warehouse_id = (SELECT from_warehouse_id FROM _param))\n\n   , _target_product_stock AS (SELECT\n                                   id,\n                                   product_id,\n                                   warehouse_id,\n                                   CASE\n                                       WHEN (SELECT is_valid FROM _source_product_stock) IS TRUE THEN\n                                           stock_quantity + (SELECT quantity FROM _param)\n                                       ELSE stock_quantity END AS stock_quantity,\n                                   (SELECT is_valid FROM _source_product_stock) AS is_valid\n                               FROM product_stock\n                               WHERE\n                                       (SELECT is_valid FROM _source_product_stock) IS TRUE AND\n                                       product_id = (SELECT product_id FROM _param) AND\n                                       warehouse_id = (SELECT to_warehouse_id FROM _param))\n\n   , _data_update AS (SELECT\n                          *\n                      FROM _source_product_stock\n                      UNION ALL\n                      SELECT\n                          *\n                      FROM _target_product_stock)\n\n   , _update_product_stock AS (\n    UPDATE product_stock\n        SET stock_quantity = _data_update.stock_quantity\n        FROM _data_update\n        WHERE\n                is_valid IS TRUE AND\n                _data_update.id = product_stock.id)\n\nINSERT\nINTO product_transfer (product_id, from_warehouse_id, to_warehouse_id, quantity)\nSELECT\n    product_id,\n    from_warehouse_id,\n    to_warehouse_id,\n    quantity\nFROM _param\nWHERE\n        (SELECT is_valid FROM _source_product_stock) IS TRUE\nRETURNING id"
	err = tx.QueryRow(query, param.ProductId, param.FromWarehouseId, param.ToWarehouseId, param.Quantity).Scan(&transferStockId.ID)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	} else {
		// Use the retrieved ID here
		log.Println("Inserted ID:", transferStockId)
	}

	tx.Commit()
	return nil
}

func (u *WarehouseRepository) UpdateStatusWarehouse(param request.WarehouseStatus) error {
	tx, err := u.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	queryUpdateStock := "UPDATE warehouse SET is_active = $2 WHERE id = $1"
	_, err = tx.Exec(queryUpdateStock, param.ID, param.IsActive)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	tx.Commit()
	return nil
}
