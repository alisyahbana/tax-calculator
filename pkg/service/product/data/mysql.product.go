package data

import (
	"database/sql"
	"github.com/alisyahbana/tax-calculator/pkg/common/database"
	"github.com/jmoiron/sqlx"
)

type MysqlProductData struct {
}

const (
	CreateProductQuery     = `INSERT INTO products (name, tax_code, price) VALUES ( ?, ?, ?);`
	GetProductIdByAllQuery = `SELECT id FROM products WHERE name = ? AND tax_code = ? AND price = ?;`
)

type MysqlProductStatement struct {
	CreateProductQuery     *sqlx.Stmt
	GetProductIdByAllQuery *sqlx.Stmt
}

var stmt MysqlProductStatement

func init() {
	stmt.CreateProductQuery = database.Prepare(database.GetDBMaster(), CreateProductQuery)
	stmt.GetProductIdByAllQuery = database.Prepare(database.GetDBMaster(), GetProductIdByAllQuery)
}

func (m MysqlProductData) CreateProduct(product Product) (uint64, error) {
	newProduct, err := stmt.CreateProductQuery.Exec(
		product.Name,
		product.TaxCode,
		product.Price,
	)

	if err != nil {
		return 0, err
	}

	newProductId, _ := newProduct.LastInsertId()

	return uint64(newProductId), nil
}

func (m MysqlProductData) GetProductId(product Product) (uint64, error) {
	var productId uint64
	err := stmt.GetProductIdByAllQuery.Get(
		&productId,
		product.Name,
		product.TaxCode,
		product.Price,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return productId, nil
}
