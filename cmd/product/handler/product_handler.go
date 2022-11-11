package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/storeapp/cmd/product/dto"
	"github.com/storeapp/cmd/product/service"
	"github.com/storeapp/pkg/err"
)

type Handlers struct {
	Service service.Service
}

func (d Handlers) AddNewProduct(w http.ResponseWriter, r *http.Request) {
	var addProductRequest dto.AddProductRequest
	json.NewDecoder(r.Body).Decode(&addProductRequest)
	res, err := d.Service.AddNewProduct(addProductRequest)
	createResponse(res, err, w)
}
func (d Handlers) GetProductList(w http.ResponseWriter, r *http.Request) {
	res, err := d.Service.GetProductList()
	createResponse(res, err, w)
}
func (d Handlers) GetProduct(w http.ResponseWriter, r *http.Request) {

}
func (d Handlers) GetCollectList(w http.ResponseWriter, r *http.Request) {
	pid := r.URL.Query().Get("product_id")
	fmt.Print(pid)
	res, e := d.Service.GetCollectList(pid)
	createResponse(res, e, w)

}
func (d Handlers) GetCollectionList(w http.ResponseWriter, r *http.Request) {
	pid := r.URL.Query().Get("product_id")
	fmt.Print(pid)
	res, e := d.Service.GetCollectionList(pid)
	createResponse(res, e, w)

}
func (d Handlers) GetProductVariant(w http.ResponseWriter, r *http.Request) {

}
func (d Handlers) GetProductImageList(w http.ResponseWriter, r *http.Request) {
	pid := r.URL.Query().Get("product_id")
	fmt.Print(pid)
	res, e := d.Service.GetProductImageList(pid)
	createResponse(res, e, w)
}
func (d Handlers) GetProductVariantList(w http.ResponseWriter, r *http.Request) {

}
func (d Handlers) AddProductImage(w http.ResponseWriter, r *http.Request) {

}
func (d Handlers) UpdateProductImage(w http.ResponseWriter, r *http.Request) {
	pid := r.URL.Query().Get("product_id")
	iid := r.URL.Query().Get("image_id")
	fmt.Print(pid)
	res, e := d.Service.UpdateProductImage(pid, iid)
	createResponse(res, e, w)
}

func createResponse(data interface{}, err *err.AppError, w http.ResponseWriter) {
	if err != nil {
		writeResponse(w, err.Messages.ResponseCode, err)
	} else {
		writeResponse(w, http.StatusOK, data)
	}
}
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
