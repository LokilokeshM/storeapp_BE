package domain

import (
	"encoding/json"
	"net/http"

	dt "github.com/storeapp/cmd/product/dto"
	"github.com/storeapp/pkg/constants"
	"github.com/storeapp/pkg/dto"
	"github.com/storeapp/pkg/err"
)

type HttpRepository struct {
	client *http.Client
}

func (h HttpRepository) GetProductList() (*dto.MasterStruct, *err.AppError) {
	res, e := dto.NewHttpRequest(
		h.client,
		http.MethodGet,
		constants.BaseUrl+"products.json",
		dto.HttpRequestDetails{})
	if e != nil {
		return nil, e
	}
	var productList dt.ProductList
	er := json.Unmarshal(res, &productList)
	if er != nil {
		return nil, err.NewBadRequest(er.Error())
	}
	return dto.CreateMasterStruct(productList), nil
}

func (h HttpRepository) AddNewProduct(req dt.AddProductRequest) (*dto.MasterStruct, *err.AppError) {

	buff, _ := json.Marshal(req)
	res, e := dto.NewHttpRequest(
		h.client,
		http.MethodPost,
		constants.BaseUrl+"products.json",
		dto.HttpRequestDetails{
			Body: buff,
		})
	if e != nil {
		return nil, e
	}
	var product dt.AddProductResponse
	er := json.Unmarshal(res, &product)
	if er != nil {
		return nil, err.NewBadRequest(er.Error())
	}
	return dto.CreateMasterStruct(product), nil
}
func (h HttpRepository) GetProduct() (*dto.MasterStruct, *err.AppError) {
	return nil, nil
}
func (h HttpRepository) GetCollectList(pid string) (*dto.MasterStruct, *err.AppError) {
	res, e := dto.NewHttpRequest(
		h.client,
		http.MethodGet,
		constants.BaseUrl+"collects.json?product_id="+pid,
		dto.HttpRequestDetails{})
	if e != nil {
		return nil, e
	}
	var productCollectsList dt.ProductCollectsList
	er := json.Unmarshal(res, &productCollectsList)
	if er != nil {
		return nil, err.NewBadRequest(er.Error())
	}
	return dto.CreateMasterStruct(productCollectsList), nil
}
func (h HttpRepository) GetCollectionList(pid string) (*dto.MasterStruct, *err.AppError) {
	res, e := dto.NewHttpRequest(
		h.client,
		http.MethodGet,
		constants.BaseUrl+"collections/"+pid+"/products.json",
		dto.HttpRequestDetails{})
	if e != nil {
		return nil, e
	}
	var productCollectionList dt.ProductCollectionList
	er := json.Unmarshal(res, &productCollectionList)
	if er != nil {
		return nil, err.NewBadRequest(er.Error())
	}
	return dto.CreateMasterStruct(productCollectionList), nil
}
func (h HttpRepository) GetProductVariant() (*dto.MasterStruct, *err.AppError) {
	return nil, nil
}
func (h HttpRepository) GetProductImageList(pid string) (*dto.MasterStruct, *err.AppError) {
	res, e := dto.NewHttpRequest(
		h.client,
		http.MethodGet,
		constants.BaseUrl+"products/"+pid+"/images.json?since_id=850703190",
		dto.HttpRequestDetails{})
	if e != nil {
		return nil, e
	}
	var productImageLIst dt.ProductImageLIst
	er := json.Unmarshal(res, &productImageLIst)
	if er != nil {
		return nil, err.NewBadRequest(er.Error())
	}
	return dto.CreateMasterStruct(productImageLIst), nil
}
func (h HttpRepository) UpdateProductImage(pid string, iid string) (*dto.MasterStruct, *err.AppError) {
	res, e := dto.NewHttpRequest(
		h.client,
		http.MethodGet,
		constants.BaseUrl+"products/"+pid+"/images.json?since_id=850703190",
		dto.HttpRequestDetails{})
	if e != nil {
		return nil, e
	}
	var productImage dt.ProductImage
	er := json.Unmarshal(res, &productImage)
	if er != nil {
		return nil, err.NewBadRequest(er.Error())
	}
	return dto.CreateMasterStruct(productImage), nil
}
func (h HttpRepository) GetProductVariantList() (*dto.MasterStruct, *err.AppError) {
	return nil, nil
}
func NewHttpRepository(client *http.Client) HttpRepository {
	return HttpRepository{client}
}
