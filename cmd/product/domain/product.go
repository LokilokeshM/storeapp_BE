package domain

import (
	d "github.com/storeapp/cmd/product/dto"
	"github.com/storeapp/pkg/dto"
	"github.com/storeapp/pkg/err"
)

type Repository interface {
	GetProduct() (*dto.MasterStruct, *err.AppError)
	AddNewProduct(req d.AddProductRequest) (*dto.MasterStruct, *err.AppError)
	GetProductList() (*dto.MasterStruct, *err.AppError)
	GetCollectList(pid string) (*dto.MasterStruct, *err.AppError)
	GetCollectionList(pid string) (*dto.MasterStruct, *err.AppError)
	GetProductImageList(pid string) (*dto.MasterStruct, *err.AppError)
	GetProductVariant() (*dto.MasterStruct, *err.AppError)
	GetProductVariantList() (*dto.MasterStruct, *err.AppError)
	UpdateProductImage(pid string, iid string) (*dto.MasterStruct, *err.AppError)
}
