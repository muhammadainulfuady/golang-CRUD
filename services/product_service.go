package services

import (
	"golang_standart_project/models"
	"golang_standart_project/repositories"
)

type ProductService struct {
	Repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductService {
	return &ProductService{Repo: repo}
}

func (s *ProductService) AddProduct(name, desc string, price float64, stock int) (int, error) {
	p := models.Product{
		ProductName:   name,
		Description:   desc,
		Price:         price,
		StockQuantity: stock,
	}
	return s.Repo.Insert(p)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.Repo.FindAll()
}
