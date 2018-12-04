package product

import (
	"github.com/alisyahbana/tax-calculator/pkg/constants"
	"github.com/alisyahbana/tax-calculator/pkg/service/product/data"
)

type ProductService struct {
	data data.MysqlProductData
}

func New() ProductService {
	return ProductService{
		data: data.MysqlProductData{},
	}
}

type ProductInput struct {
	Name    string `json:"name"`
	TaxCode int    `json:"tax_code"`
	Price   uint64 `json:"price"`
}

type Item struct {
	Name       string  `json:"name"`
	TaxCode    int     `json:"tax_code"`
	Type       string  `json:"type"`
	Refundable string  `json:"refundable"`
	Price      uint64  `json:"price"`
	Tax        float64 `json:"tax"`
	Amount     float64 `json:"amount"`
}

type BillingOutput struct {
	Items         []Item  `json:"items"`
	PriceSubtotal uint64  `json:"price_subtotal"`
	TaxSubtotal   float64 `json:"tax_subtotal"`
	GrandTotal    float64 `json:"grand_total"`
}

func (s ProductService) CreateProduct(productInput ProductInput) (uint64, error) {
	payloadInput := data.Product{
		Name:    productInput.Name,
		TaxCode: productInput.TaxCode,
		Price:   uint64(productInput.Price),
	}

	productIdExist, err := s.data.GetProductId(payloadInput)
	if err != nil {
		return 0, err
	}

	if productIdExist != 0 {
		return 0, nil
	}

	productId, err := s.data.CreateProduct(payloadInput)
	if err != nil {
		return 0, err
	}

	return productId, nil
}

func (s ProductService) GenerateBilling(products []ProductInput) (BillingOutput, error) {
	var billing BillingOutput
	var items []Item
	var taxSubTotal float64
	var priceSubTotal uint64

	for _, product := range products {
		var typeValue, refundable string
		var tax float64

		switch product.TaxCode {
		case 1:
			typeValue = constants.FNB
			refundable = "YES"
			tax = float64(product.Price * 10 / 100)

		case 2:
			typeValue = constants.TOBACCO
			refundable = "NO"
			tax = float64(10 + (product.Price * 2 / 100))
		case 3:
			typeValue = constants.ENTERTAINMENT
			refundable = "NO"
			if product.Price > 0 && product.Price < 100 {
				tax = 0
			} else if product.Price >= 100 {
				tax = 0.01 * float64(product.Price-100)
			}
		}

		item := Item{
			Name:       product.Name,
			TaxCode:    product.TaxCode,
			Type:       typeValue,
			Refundable: refundable,
			Price:      product.Price,
			Tax:        tax,
			Amount:     float64(product.Price) + tax,
		}

		items = append(items, item)
		taxSubTotal += tax
		priceSubTotal += product.Price
	}

	billing.Items = items
	billing.TaxSubtotal = taxSubTotal
	billing.PriceSubtotal = priceSubTotal
	billing.GrandTotal = taxSubTotal + float64(priceSubTotal)

	return billing, nil
}
