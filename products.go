package baselinker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetProductsResponse struct {
	*BaseResponse
	Products []Product `json:"products"`
}

type GetProductsListParameters struct {
	StorageId          string  `json:"storage_id"`
	FilterCategoryId   string  `json:"filter_category_id,omitempty"`
	FilterSort         string  `json:"filter_sort,omitempty"`
	FilterId           string  `json:"filter_id,omitempty"`
	FilterEan          string  `json:"filter_ean,omitempty"`
	FilterSku          string  `json:"filter_sku,omitempty"`
	FilterName         string  `json:"filter_name,omitempty"`
	FilterPriceFrom    float32 `json:"filter_price_from,omitempty"`
	FilterPriceTo      float32 `json:"filter_price_to,omitempty"`
	FilterQuantityFrom int     `json:"filter_quantity_from,omitempty"`
	FilterQuantityTo   int     `json:"filter_quantity_to,omitempty"`
	FilterAvailable    int     `json:"filter_available,omitempty"`
	Page               int     `json:"page,omitempty"`
}

// Documentation: https://api.baselinker.com/index.php?method=getProductsList
func (baseLiner *BaseLinker) GetProductsList(parameters GetProductsListParameters) ([]Product, Error) {
	var (
		response GetProductsResponse
	)

	formData := baseLiner.createRequestForm("getProductsList", parameters)
	resp, err := http.PostForm(baseLiner.Url, formData)

	if nil != err {
		return response.Products, NewSimpleError(err)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)

	if nil != err {
		return response.Products, NewSimpleError(err)
	}

	if !response.IsSuccess() {
		return response.Products, response
	}

	return response.Products, nil
}

func (baseLiner *BaseLinker) GetProduct(storageId, productId string) (Product, Error) {
	products, err := baseLiner.GetProductsList(
		GetProductsListParameters{
			StorageId: storageId,
			FilterId:  productId,
		},
	)

	if nil != err {
		return Product{}, err
	}

	if len(products) == 0 {
		return Product{}, NewSimpleError(fmt.Errorf("Product with id: %s not found", productId))
	}

	return products[0], nil
}

type GetProductsDetailsParameters struct {
	StorageId   string   `json:"storage_id"`
	ProductsIds []string `json:"products"`
}

// Documentation: https://api.baselinker.com/index.php?method=getProductsData
func (baseLiner *BaseLinker) GetProdutsDetails(parameters GetProductsDetailsParameters) ([]Product, Error) {
	var (
		response GetProductsResponse
	)

	formData := baseLiner.createRequestForm("getProductsData", parameters)
	resp, err := http.PostForm(baseLiner.Url, formData)

	if nil != err {
		return response.Products, NewSimpleError(err)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)

	if nil != err {
		return response.Products, NewSimpleError(err)
	}

	if !response.IsSuccess() {
		return response.Products, response
	}

	return response.Products, nil
}

func (baseLiner *BaseLinker) GetProductDetails(storageId, productId string) (Product, Error) {
	products, err := baseLiner.GetProdutsDetails(
		GetProductsDetailsParameters{
			StorageId:   storageId,
			ProductsIds: []string{productId},
		},
	)

	if nil != err {
		return Product{}, err
	}

	if len(products) == 0 {
		return Product{}, NewSimpleError(fmt.Errorf("Product with id: %s not found", productId))
	}

	return products[0], nil

}
