package data

type Product struct {
	Name    string `json:"name"`
	TaxCode int    `json:"tax_code"`
	Price   uint64 `json:"price"`
}

type ProductData interface {
	CreateProduct(product Product) (uint64, error)
	GetProductId(product Product) (uint64, error)
}
