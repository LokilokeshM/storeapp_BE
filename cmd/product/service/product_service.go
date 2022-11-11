package service

import (
	"github.com/storeapp/cmd/product/domain"
	d "github.com/storeapp/cmd/product/dto"
	"github.com/storeapp/pkg/dto"
	"github.com/storeapp/pkg/err"
)

type Service interface {
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

func (d DefaultService) GetProductList() (*dto.MasterStruct, *err.AppError) {
	return d.Repo.GetProductList()
}

func (d DefaultService) AddNewProduct(req d.AddProductRequest) (*dto.MasterStruct, *err.AppError) {
	return d.Repo.AddNewProduct(req)
}

func (d DefaultService) GetProduct() (*dto.MasterStruct, *err.AppError) {
	return d.Repo.GetProduct()
}
func (d DefaultService) GetCollectList(pid string) (*dto.MasterStruct, *err.AppError) {
	return d.Repo.GetCollectList(pid)
}
func (d DefaultService) GetCollectionList(pid string) (*dto.MasterStruct, *err.AppError) {
	return d.Repo.GetCollectionList(pid)
}
func (d DefaultService) GetProductVariant() (*dto.MasterStruct, *err.AppError) {
	return d.Repo.GetProductVariant()
}
func (d DefaultService) GetProductImageList(pid string) (*dto.MasterStruct, *err.AppError) {
	return d.Repo.GetProductImageList(pid)
}
func (d DefaultService) GetProductVariantList() (*dto.MasterStruct, *err.AppError) {
	return d.Repo.GetProductVariantList()
}
func (d DefaultService) UpdateProductImage(pid string, iid string) (*dto.MasterStruct, *err.AppError) {
	return d.Repo.UpdateProductImage(pid, iid)
}

type DefaultService struct {
	Repo domain.Repository
}

func NewService(Repo domain.Repository) DefaultService {
	return DefaultService{Repo}
}
